package main

import (
	"github.com/spf13/viper"
	"os"
	"testing"
)

func TestReadConfigFromFile(t *testing.T) {
	cfgFile = "./testdata/alp.json"
	initConfig()
	var hostValue = viper.GetString("host")
	if hostValue != "json_host" {
		t.Errorf("The default config should have been used")
	}
}

func TestReadConfigFromEnv(t *testing.T) {
	os.Setenv("ALP_HOST", "zauberelfenfee")
	initConfig()
	var hostValue = viper.Get("host")
	if hostValue != "zauberelfenfee" {
		t.Errorf("zauberelfenfee should have been set")
	}
}

func TestReadConfigFromCmdParams(t *testing.T) {
	os.Args = []string{"rtrans", "--host=param_host"}
	main()
	var hostValue = viper.Get("host")
	if hostValue != "param_host" {
		t.Errorf("param_host should have been red from config")
	}
}
