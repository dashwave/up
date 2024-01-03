package service

import (
	"context"
	"fmt"
	"os"
)

func (s *Service) Validate(ctx context.Context) error {
	if s.Name == "" {
		return fmt.Errorf("service name cannot be empty")
	}
	// add more validation
	return nil
}
func (s *Service) Deploy(ctx context.Context) error {
	switch s.Source {
	case "docker":
		if s.Auth && s.AuthConfig == nil {
			s.AuthConfig = &AuthConfig{
				Username: os.Getenv("DOCKER_USERNAME"),
				Password: os.Getenv("DOCKER_PASSWORD"),
			}
		}
		d := &dockerService{
			Name:          s.Name,
			ContainerName: s.ContainerName,
			Image:         s.Image,
			Ports:         s.Ports,
			Env:           s.Env,
			Pre:           s.Pre,
			Volumes:       s.Volumes,
			Networks:      s.Networks,
			Auth:          s.Auth,
			EnvFiles:      s.EnvFiles,
			AuthConfig:    s.AuthConfig,
		}
		if err := d.validate(ctx); err != nil {
			return err
		}
		return d.deployDocker(ctx)
	case "local":
		l := &localService{
			Name: s.Name,
			Pre:  s.Pre,
			Exec: s.Exec,
			Env:  s.Env,
			Path: s.Path,
		}
		if err := l.validate(ctx); err != nil {
			return err
		}
		go l.deployLocal(ctx)
	default:
		return fmt.Errorf("unknown source: %s", s.Source)
	}
	return nil
}
