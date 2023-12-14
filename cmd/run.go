package cmd

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/dashwave/up/internal/deploy"
	"github.com/dashwave/up/internal/service"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "runs the specified deployment",
	Long:  "runs the specified deployment",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()
		ctx, cancel := context.WithCancel(ctx)
		if err := deploy.Deploy(ctx, deploymentFile); err != nil {
			panic(err)
		}
		cleanupCompleted := make(chan bool)
		go service.CleanupDockerContainers(ctx, cleanupCompleted)
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c
		cancel()
		<-cleanupCompleted
		os.Exit(0)
	},
}
