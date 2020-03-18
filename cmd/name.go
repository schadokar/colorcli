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
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/spf13/cobra"
)

// nameCmd represents the name command
var nameCmd = &cobra.Command{
	Use:   "name",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		hexToName(args)
	},
}

func hexToName(args []string) {
	var hexMap map[string]string

	// read the color.min.json
	content, err := ioutil.ReadFile("colornames.min.json")

	if err != nil {
		fmt.Printf("Error while reading the file %v", err)
	}
	_ = json.Unmarshal(content, &hexMap)

	name, ok := hexMap[args[0]]
	if ok {
		fmt.Printf("Name: %s, Hex: %s\n", name, args[0])
	} else {

		fmt.Println("Color name not found")
	}
}

func init() {
	rootCmd.AddCommand(nameCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// nameCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// nameCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
