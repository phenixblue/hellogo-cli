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

	"github.com/spf13/cobra"
)

var (
	name string
)

// command2Cmd represents the command2 command
var command2Cmd = &cobra.Command{
	Use:   "command2",
	Short: "Command 2 has flags",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("command2 called")

		nameToggle, _ := cmd.Flags().GetBool("toggle-name")

		if nameToggle {
			fmt.Printf("\nI hope you are having a great day!\n")
		} else {
			fmt.Printf("\nHello %v, I hope you are having a great day!\n", name)
		}
	},
}

func init() {
	rootCmd.AddCommand(command2Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// command2Cmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// command2Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// This flag stores its value in a named variable
	command2Cmd.Flags().StringVarP(&name, "name", "n", "Stranger", "A name to reference for a custom greeting")
	// This flags value can be looked up with cmd.Flags().GetBool("toggle-name")
	command2Cmd.Flags().BoolP("toggle-name", "t", false, "Toggle name specific greeting")
}
