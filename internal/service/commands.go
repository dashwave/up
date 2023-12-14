package service

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os/exec"
	"strings"
)

type CommandConfig struct {
	Command    string
	EnvVars    []string
	WorkingDir string
}

func (cmdConfig *CommandConfig) Run(ctx context.Context) error {
	fmt.Printf("Running command: %s\n", cmdConfig.Command)

	// Split the command into base command and arguments
	parts := strings.Fields(cmdConfig.Command)
	name := parts[0]
	args := parts[1:]

	// Create the command
	cmd := exec.Command(name, args...)
	cmd.Env = append(cmd.Env, cmdConfig.EnvVars...)
	cmd.Dir = cmdConfig.WorkingDir

	// Get the Stdout and Stderr pipe
	stdoutPipe, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	stderrPipe, err := cmd.StderrPipe()
	if err != nil {
		return err
	}

	// Start the command
	if err := cmd.Start(); err != nil {
		return err
	}

	// Create a channel to signal when the command is done
	done := make(chan struct{})

	// Stream the output
	go streamOutput(&stdoutPipe, "STDOUT", done)
	go streamOutput(&stderrPipe, "STDERR", done)

	// Wait for the command to finish
	if err := cmd.Wait(); err != nil {
		return err
	}

	// Wait for output to finish streaming
	<-done
	<-done

	close(done)
	return nil
}

func streamOutput(pipeReader *io.ReadCloser, outputLabel string, done chan struct{}) {
	scanner := bufio.NewScanner(*pipeReader)
	for scanner.Scan() {
		fmt.Printf("%s\n", scanner.Text())
	}
	done <- struct{}{}
}
