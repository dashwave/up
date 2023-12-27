package cmd

import (
	"os"

	"github.com/dashwave/up/internal/service"
	"github.com/spf13/cobra"
)

var downCmd = &cobra.Command{
	Use:   "down",
	Short: "stops the specified deployment",
	Long:  "stops the specified deployment",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()
		cleanupCompleted := make(chan bool)
		go service.CleanupDockerContainers(ctx, cleanupCompleted)
		<-cleanupCompleted
		os.Exit(0)
	},
}
