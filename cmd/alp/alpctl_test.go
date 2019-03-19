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
	"os"
	"testing"

	"github.com/spf13/viper"
)

func TestReadConfigFromFile(t *testing.T) {
	cfgFile = "./testdata/alp.yml"
	initConfig()
	Host := viper.GetString("host")
	if Host != "alp_host" {
		t.Errorf("default config should have been used; given was '%s'", Host)
	}
}

func TestReadConfigFromEnv(t *testing.T) {
	os.Setenv("ALP_HOST", "unicorn")
	initConfig()
	var Host = viper.Get("host")
	os.Setenv("ALP_HOST", "")
	if Host != "unicorn" {
		t.Errorf("unicorn should have been set; given was '%s'", Host)
	}
}

func TestReadConfigFromFileHostSsh(t *testing.T) {
	cfgFile = "./testdata/alp-ssh.yml"
	initConfig()
	var Host = viper.Get("host")
	if Host != "alp_ssh_host" {
		t.Errorf("host should have been read from config; given was '%s'", Host)
	}
}

func TestRtransHostnameCreation(t *testing.T) {
	hostConcat := createHostURL("radar-transform.example.io", 666)
	if hostConcat != "radar-transform.example.io:666" {
		t.Errorf("host should have been 'radar-transform.example.io:666'; given was '%s'", hostConcat)
	}
	hostConcat = createHostURL("239.21.12.3", 666)
	if hostConcat != "239.21.12.3:666" {
		t.Errorf("host should have been '239.21.12.3:666'; given was '%s'", hostConcat)
	}
	hostConcat = createHostURL("FE80:0000:0000:0000:0202:B3FF:FE1E:8329", 666)
	if hostConcat != "[FE80:0000:0000:0000:0202:B3FF:FE1E:8329]:666" {
		t.Errorf("host should have been '[FE80:0000:0000:0000:0202:B3FF:FE1E:8329]:666'; given was '%s'", hostConcat)
	}
}

func TestReadConfigFromCmdParams(t *testing.T) {
	os.Args = []string{"rtrans", "--host=param_host"}
	main()
	Host := viper.Get("host")
	if Host != "param_host" {
		t.Errorf("host should have been read from command parameter; given was '%s'", Host)
	}
}
