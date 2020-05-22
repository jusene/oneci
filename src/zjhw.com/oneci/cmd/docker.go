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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if project, err := cmd.Flags().GetString("project"); err == nil && project != "" {
			fmt.Println("Project Name:", project)
		}
		if version, err := cmd.Flags().GetString("version"); err == nil && version != "" {
			fmt.Println("Project Version:", version)
		}

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
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dockerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
