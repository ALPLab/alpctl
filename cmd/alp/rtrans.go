// Copyright © 2018 ALP.Lab GmbH <office@alp-lab.at>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	rtrans "github.com/ALPLab/protorepo-infra-radar-transform-go"

	proto "github.com/gogo/protobuf/proto"
	cobra "github.com/spf13/cobra"
	viper "github.com/spf13/viper"
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
	credentials "google.golang.org/grpc/credentials"
)

var (
	Radar     string
	Car       string
	Plaintext bool
)

var rtransCmd = &cobra.Command{
	Use:        "rtrans",
	Short:      "Roadside and Sensor Data Transform",
	Long:       `Transforms roadside data and tracking data of ego car into ego car's relative coordinate system.`,
	SuggestFor: []string{"transform", "radar", "radartransform"},
	RunE: func(cmd *cobra.Command, args []string) error {
		Host := viper.GetString("host")
		Port := viper.GetInt("port")
		certFile := viper.GetString("certificate")
		outDir := viper.GetString("output")
		return radarTransform(context.Background(), Host, Port, certFile, outDir, &Radar, &Car, &Plaintext)
	},
}

func init() {
	rtransCmd.Flags().StringVarP(&Radar, "radar", "r", "", "file with radar data (required)")
	rtransCmd.Flags().StringVarP(&Car, "car", "c", "", "file with ego car's sensor data (required)")
	rtransCmd.Flags().BoolVar(&Plaintext, "plaintext", false, "sends data as plaintext if flag is set (insecure!)")
	rtransCmd.MarkFlagRequired("radar")
	rtransCmd.MarkFlagRequired("car")

	rootCmd.AddCommand(rtransCmd)
}

func radarTransform(ctx context.Context, host string, port int, certFile string,
	outDir string, radar *string, car *string, plaintext *bool) error {

	// read JSON file with radarData data
	radarData, err := ioutil.ReadFile(*radar)
	if err != nil {
		return fmt.Errorf("radar file could not be read: '%s'", err)
	}

	// read GPX file with ego carData's GPS tracking data
	carData, err := ioutil.ReadFile(*car)
	if err != nil {
		return fmt.Errorf("car file could not be read: '%s'", err)
	}

	// print to console
	fmt.Printf("\nConnection secure: %t\n", !*plaintext)
	server := createHostURL(host, port)
	fmt.Printf("Connecting to: '%s'\n", server)

	// check if correct certFile was provided
	if !path.IsAbs(certFile) {
		err := fmt.Errorf("Error: '%s' is not absolute; please provide absolute path", certFile)
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var creds credentials.TransportCredentials
	var conn *grpc.ClientConn

	// dial to server with gRPC
	if *plaintext == false {
		// connect with TLS certificate
		creds, err = credentials.NewClientTLSFromFile(certFile, "")
		if err != nil {
			return fmt.Errorf("could not fetch certificate: '%s'", err)
		}
		conn, err = grpc.Dial(server, grpc.WithTransportCredentials(creds))
		if err != nil {
			return fmt.Errorf("did not connect: '%s'", err)
		}
	} else {
		// connect without authentication
		conn, err = grpc.Dial(server, grpc.WithInsecure())
		if err != nil {
			fmt.Print(conn)
			return fmt.Errorf("did not connect: '%s'", err)
		}
	}

	// create new client
	client := rtrans.NewInfraRadarPositionTransformClient(conn)
	if err != nil {
		return fmt.Errorf("client could not be created: '%v'", err)
	}

	// transform timeout 30 seconds
	ctxTimeout, cancel := context.WithTimeout(ctx, 30*time.Second)
	// call transform
	log.Println("Sending transform request and waiting for response...")
	response, err := client.Transform(ctxTimeout, &rtrans.TransformRequest{Car: carData, Radar: radarData})
	if err != nil {
		return fmt.Errorf("transform failed: '%s'", err)
	}
	defer cancel()
	defer conn.Close()

	// encode response in JSON
	out, err := proto.Marshal(response)
	if err != nil {
		return fmt.Errorf("failed to encode response: '%v'", err)
	}

	// get current time in RFC3339 format and strip milliseconds
	now := time.Now().Format(time.RFC3339)
	datetime := strings.Replace(strings.Split(now, "+")[0], ":", "-", -1)

	// name output file according to convention and create file
	outPath := filepath.Join(outDir, "car_rsd_sensor_"+datetime+".osi3.pb")
	file, err := os.Create(outPath)
	if err != nil {
		return fmt.Errorf("could not create file: '%v'", err)
	}
	defer file.Close()

	// create a buffered writer from the file and write bytes to buffer
	bufferedWriter := bufio.NewWriter(file)
	_, err = bufferedWriter.Write(out)
	if err != nil {
		return fmt.Errorf("could not write to file: '%v'", err)
	}

	log.Printf("Response written to: '%v'", outPath)

	return nil
}

func createHostURL(host string, port int) string {
	server := net.JoinHostPort(host, strconv.Itoa(port))
	return server
}
