package service

import (
	"context"
	"fmt"
	"strings"
)

func getDockerEnvConfigs(ctx context.Context, env []string) ([]string, error) {
	dockerEnvs := make([]string, 0)
	for _, e := range env {
		for _, envVar := range UpConfigs.Env {
			envSplit := strings.Split(string(envVar), "=")
			if len(envSplit) < 2 {
				fmt.Printf("invalid env variable %s\n", envVar)
			}
			if e == envSplit[0] {
				dockerEnvs = append(dockerEnvs, string(envVar))
			}
		}
	}
	return dockerEnvs, nil
}
