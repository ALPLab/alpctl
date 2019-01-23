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
		t.Errorf("The default config should have been used. Read unexpected %s", hostValue)
	}
}

func TestReadConfigFromEnv(t *testing.T) {
	os.Setenv("ALP_HOST", "zauberelfenfee")
	initConfig()
	var hostValue = viper.Get("host")
	os.Setenv("ALP_HOST", "")
	if hostValue != "zauberelfenfee" {
		t.Errorf("zauberelfenfee should have been set. was %s", hostValue)
	}
}

func TestReadConfigFromFileHostSsh(t *testing.T) {
	cfgFile = "./testdata/alp-ssh.json"
	initConfig()
	var hostValue = viper.Get("host")
	if hostValue != "https://my_host" {
		t.Errorf("host should have been read from config. Value read %s", hostValue)
	}
}

func TestRtransHostnameCreation(t *testing.T) {
	hostConcat := createHostUrl("https://my_world", 1234)

	if hostConcat != "https://my_world:1234" {
		t.Errorf("host should have been https://my_world:1234. Value read %s", hostConcat)
	}
}
