package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var deploymentFile string

var rootCmd = &cobra.Command{
	Use:   "up",
	Short: "up is a deployment tool for dashwave infrastructure.",
	Long:  "up is a deployment tool for dashwave infrastructure. Users can define their infrastructure in a single YAML file and deploy it to a machine with a single command. It supports both docker and local deployments.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// This function executes the root command for the tool.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	runCmd.Flags().StringVarP(&deploymentFile, "file", "f", "up.yaml", "deployment file to use")
	rootCmd.AddCommand(runCmd)
}
