/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/phenixblue/hellogo-cli/pkg/common"
	"github.com/spf13/cobra"
)

// command1Cmd represents the command1 command
var command1Cmd = &cobra.Command{
	Use:   "command1",
	Short: "Command 1 does a thing",
	Long: `This is a really long,, multi-line description of
the thing that Command 1 does`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("command1 called")

		currDir, err := common.GetPath()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		files, err := common.GetFileList(currDir)
		if err != nil {
			fmt.Printf("Error listing files: %v\n", err)
		}

		for _, file := range files {
			fmt.Println(file.Name())
		}
	},
}

func init() {
	rootCmd.AddCommand(command1Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// command1Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// command1Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
