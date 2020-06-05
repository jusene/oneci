package cmd

import (
	"github.com/spf13/cobra"
	"strings"
	"zjhw.com/oneci/config"
	"zjhw.com/oneci/utils"
)

var kubeDeployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "kube deploy",
	Long:  "kube deploy",
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

		arch, err := cmd.Flags().GetString("arch")
		checkFlag(ty, err, "arch")

		if tier == "backend" {
			for _, app := range strings.Split(apps, ",") {
				utils.DeployJavaKube(config.Conf, app, version, project, env, ty, arch, timestamp)
			}
		} else if tier == "front" {
			utils.DeployJavaScriptKube(config.Conf, apps, version, project, env, ty, arch, timestamp)
		}
	},
}

func init() {
	kubeCmd.AddCommand(kubeDeployCmd)

	kubeDeployCmd.Flags().Int64P("timestamp", "s", 0, "A timestamp tag of project")
}
