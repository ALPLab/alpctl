package main

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	"os"
	"testing"
)

func TestReadConfigFromFile(t *testing.T) {
	cfgFile = "./testdata/alp.json"
	initConfig()
	var hostValue = viper.GetString("host")
	if hostValue != "json_host" {
		t.Errorf("The default config should have been used. Read unexpected %s", hostValue)
	}
}

func TestReadConfigFromEnv(t *testing.T) {
	os.Setenv("ALP_HOST", "zauberelfenfee")
	initConfig()
	var hostValue = viper.Get("host")
	if hostValue != "zauberelfenfee" {
		t.Errorf("zauberelfenfee should have been set. was %s", hostValue)
	}
}

func TestReadConfigFromCmdParams(t *testing.T) {
	os.Args = []string{"rtrans", "--host=param_host"}
	main()
	var hostValue = viper.Get("host")
	if hostValue != "param_host" {
		t.Errorf("param_host should have been read from config. Value read %s", hostValue)
	}
}

func TestReadDefaultCertDirectory(t *testing.T) {

	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	main()
	var certificatePath = viper.GetString("certificate")
	if certificatePath != home+"/.alp/ca-chain.cert.pem" {
		t.Errorf("certificate chain default %s ain't as it should be", certificatePath)
	}
}
