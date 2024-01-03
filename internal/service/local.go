package service

import (
	"context"
	"fmt"
	"os"
	"path"
)

func (l *localService) validate(ctx context.Context) error {
	return nil
}

func (l *localService) deployLocal(ctx context.Context) error {
	fmt.Println("deploying local service", l.Name)
	for _, e := range l.Exec {
		path := path.Join(os.Getenv("PWD"), l.Path)
		fmt.Println("running command", e, "in path", path)
		cmd := CommandConfig{
			Command:    e,
			EnvVars:    l.Env,
			WorkingDir: path,
		}
		if err := cmd.Run(ctx); err != nil {
			fmt.Println("error running command", err)
		}
	}
	return nil
}
