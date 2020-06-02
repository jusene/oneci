package cmd

import (
	"github.com/spf13/cobra"
	"strings"
	"zjhw.com/oneci/config"
	"zjhw.com/oneci/utils"
)

func checkFlag(flag interface{}, err error, param string) {
	if err != nil {
		panic(err)
	}

	if err == nil && flag == "" {
		panic("require param " + param)
	}
}

var dockerPreCmd = &cobra.Command{
	Use:   "prepare",
	Short: "docker build prepare",
	Long:  "docker build prepare",
	Run: func(cmd *cobra.Command, args []string) {
		project, err := cmd.Flags().GetString("project")
		checkFlag(project, err, "project")

		version, err := cmd.Flags().GetString("version")
		checkFlag(version, err, "version")

		apps, err := cmd.Flags().GetString("application")
		checkFlag(apps, err, "application")

		arch, err := cmd.Flags().GetString("arch")
		checkFlag(arch, err, "arch")

		tier, err := cmd.Flags().GetString("tier")
		checkFlag(tier, err, "tier")

		env, err := cmd.Flags().GetString("env")
		checkFlag(env, err, "env")

		ty, err := cmd.Flags().GetString("type")
		checkFlag(env, err, "type")

		if tier == "backend" {
			for _, app := range strings.Split(apps, ",") {
				utils.PreJavaDocker(app, project, version, env, arch, ty, config.JavaPre.Dockerfile, config.JavaPre.Entrypoint)
			}
		} else if tier == "front" {
			utils.PreJavaScriptDocker()
		}
	},
}

func init() {
	dockerCmd.AddCommand(dockerPreCmd)
}
