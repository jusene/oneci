package cmd

import (
	"github.com/spf13/cobra"
	"strings"
	"zjhw.com/oneci/config"
	"zjhw.com/oneci/utils"
)

func checkFlag(flag interface{}, err error) {
	if err != nil {
		panic(err)
	}

	if err == nil && flag == "" {
		panic(flag)
	}
}

var dockerPreCmd = &cobra.Command{
	Use: "prepare",
	Short: "docker build prepare",
	Long: "docker build prepare",
	Run: func(cmd *cobra.Command, args []string) {
		project, err := cmd.Flags().GetString("project")
		checkFlag(project, err)

		version, err := cmd.Flags().GetString("version")
		checkFlag(version, err)

		apps, err := cmd.Flags().GetString("application")
		checkFlag(apps, err)

		arch, err := cmd.Flags().GetString("arch")
		checkFlag(arch, err)

		tier, err := cmd.Flags().GetString("tier")
		checkFlag(tier, err)
		if tier == "backend" {
			for _, app := range strings.Split(apps, ",") {
				utils.PreJavaDocker(app, project, version, strings.Join([]string{config.JavaPre.Dockerfile, arch}, "/"),
					config.JavaPre.Entrypoint)
			}
		} else if tier == "front" {
			utils.PreJavaScriptDocker()
		}
	},
}

func init() {
	dockerCmd.AddCommand(dockerPreCmd)

	dockerPreCmd.Flags().StringP("timestamp", "s", "", "A tag of docker")
}