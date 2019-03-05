package main

import (
	"testing"

	"github.com/spf13/viper"
)

func TestReadConfigFromFile(t *testing.T) {
	cfgFile = "./testdata/alp.yml"
	initConfig()
	var hostValue = viper.GetString("host")
	if hostValue != "alp_host" {
		t.Errorf("default config should have been used; given was '%s'", hostValue)
	}
}

// func TestReadConfigFromEnv(t *testing.T) {
// 	os.Setenv("ALP_HOST", "unicorn")
// 	initConfig()
// 	var hostValue = viper.Get("host")
// 	os.Setenv("ALP_HOST", "")
// 	if hostValue != "unicorn" {
// 		t.Errorf("unicorn should have been set; given was '%s'", hostValue)
// 	}
// }
//
// func TestReadConfigFromFileHostSsh(t *testing.T) {
// 	cfgFile = "./testdata/alp-ssh.json"
// 	initConfig()
// 	var hostValue = viper.Get("host")
// 	if hostValue != "https://my_host" {
// 		t.Errorf("host should have been read from config; given was '%s'", hostValue)
// 	}
// }
//
// func TestRtransHostnameCreation(t *testing.T) {
// 	hostConcat := createHostUrl("https://my_world", 1234)
//
// 	if hostConcat != "https://my_world:1234" {
// 		t.Errorf("host should have been https://my_world:1234; given was '%s'", hostConcat)
// 	}
// }
