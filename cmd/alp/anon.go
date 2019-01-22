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
	"fmt"

	"github.com/spf13/cobra"
)

var Video string

// anonCmd represents the anon command
var anonCmd = &cobra.Command{
	Use:   "anon",
	Short: "Anonymize Test Drive Video (not yet available)",
	Long: `
A service to anonymize videos to make the recorded data available for
further testing and development while guaranteeing to protect recorded
personal information of other traffic participants and the like by
removing it as reliable as technically feasible. `,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("anon called")
	},
}

func init() {
	rootCmd.AddCommand(anonCmd)
	anonCmd.Flags().StringVarP(&Video, "video", "v", "", "video to be anonymized")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// anonCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// anonCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
