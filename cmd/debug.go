// Copyright © 2019 Coleman Word
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

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// debugCmd represents the debug command
var debugCmd = &cobra.Command{
	Use:   "debug",
	Short: "debug current config file",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("------------------------------------------------")
		fmt.Println("DEBUG")
		fmt.Println("key| type| description|")
		fmt.Println("------------------------------------------------")
		fmt.Printf("|send|int| sending this many texts to each target:%v\n", viper.GetInt("send"))
		fmt.Printf("|from|[]string| randomly sending texts from these numbers:%v\n", viper.GetStringSlice("from"))
		fmt.Printf("|to|[]string| sending texts to these targets:%v\n", viper.GetStringSlice("to"))
	},
}

func init() {
	rootCmd.AddCommand(debugCmd)
}
