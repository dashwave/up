package service

import (
	"context"
	"fmt"
)

func (l *localService) validate(ctx context.Context) error {
	return nil
}

func (l *localService) deployLocal(ctx context.Context) error {
	fmt.Println("deploying local service", l.Name)
	envs, err := getDockerEnvConfigs(ctx, l.Env)
	if err != nil {
		return err
	}
	for _, e := range l.Exec {
		cmd := CommandConfig{
			Command:    e,
			EnvVars:    envs,
			WorkingDir: l.Path,
		}
		if err := cmd.Run(ctx); err != nil {
			fmt.Println("error running command", err)
		}
	}
	return nil
}
