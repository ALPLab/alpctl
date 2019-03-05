package main

import (
	"os"
	"testing"

	"github.com/spf13/viper"
)

func TestReadConfigFromCmdParams(t *testing.T) {
	os.Args = []string{"rtrans", "--host=param_host"}
	main()
	var hostValue = viper.Get("host")
	if hostValue != "param_host" {
		t.Errorf("host should have been read from config; given was '%s'", hostValue)
	}
}

// func TestReadDefaultCertDirectory(t *testing.T) {
// 	home, err := homedir.Dir()
// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}
// 	main()
// 	var certificatePath = viper.GetString("certificate")
// 	if certificatePath != home+"/.alp/ca-chain.cert.pem" {
// 		t.Errorf("certificate chain default '%s' ain't as it should be", certificatePath)
// 	}
// }
