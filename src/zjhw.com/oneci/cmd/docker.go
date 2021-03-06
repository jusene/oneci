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

// dockerCmd represents the docker command
var dockerCmd = &cobra.Command{
	Use:   "docker",
	Short: "docker command",
	Long:  `docker command`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Build Docker")
	},
}

func init() {
	rootCmd.AddCommand(dockerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dockerCmd.PersistentFlags().String("foo", "", "A help for foo")
	dockerCmd.PersistentFlags().StringP("version", "v", "", "A version of project")
	dockerCmd.PersistentFlags().StringP("project", "p", "", "A name of project")
	dockerCmd.PersistentFlags().StringP("application", "a", "", "A list of project")
	dockerCmd.PersistentFlags().StringP("arch", "r", "", "A arch name of project(arm64/amd64)")
	dockerCmd.PersistentFlags().StringP("tier", "t", "", "A tier name of project(front/backend)")
	dockerCmd.PersistentFlags().StringP("env", "e", "", "A env name of project")
	dockerCmd.PersistentFlags().StringP("type", "y", "", "A type of project")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dockerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
