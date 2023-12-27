package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/dashwave/up/internal/deploy"
	"github.com/dashwave/up/internal/service"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "up",
	Short: "runs the specified deployment",
	Long:  "runs the specified deployment",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()

		cleanupCompleted := make(chan bool)
		go service.CleanupDockerContainers(ctx, cleanupCompleted)
		<-cleanupCompleted

		if err := deploy.Deploy(ctx, deploymentFile); err != nil {
			fmt.Printf("error deploying: %v\n", err)
		}
		if background {
			os.Exit(0)
		}
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c
		go service.CleanupDockerContainers(ctx, cleanupCompleted)
		<-cleanupCompleted
		os.Exit(0)
	},
}
