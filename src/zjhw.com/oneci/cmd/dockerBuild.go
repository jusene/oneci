package cmd

import (
	"github.com/spf13/cobra"
	"strings"
	"zjhw.com/oneci/utils"
)

var dockerBuildCmd = &cobra.Command{
	Use:   "build",
	Short: "docker build",
	Long:  "docker build",
	Run: func(cmd *cobra.Command, args []string) {
		project, err := cmd.Flags().GetString("project")
		checkFlag(project, err, "project")

		version, err := cmd.Flags().GetString("version")
		checkFlag(version, err, "version")

		apps, err := cmd.Flags().GetString("application")
		checkFlag(apps, err, "application")

		timestamp, err := cmd.Flags().GetInt64("timestamp")
		checkFlag(timestamp, err, "timestamp")

		env, err := cmd.Flags().GetString("env")
		checkFlag(env, err, "env")

		tier, err := cmd.Flags().GetString("tier")
		checkFlag(tier, err, "tier")

		ty, err := cmd.Flags().GetString("type")
		checkFlag(env, err, "type")

		for _, app := range strings.Split(apps, ",") {
			utils.BuildDocker(app, version, project, env, ty, timestamp)
		}
	},
}

func init() {
	dockerCmd.AddCommand(dockerBuildCmd)

	dockerBuildCmd.Flags().Int64P("timestamp", "s", 0, "A timestamp tag of project")
}
