package service

import "context"

func getDockerEnvConfigs(ctx context.Context, env []string) ([]string, error) {
	dockerEnvs := make([]string, 0)
	for _, e := range env {
		for _, envVar := range UpConfigs.Env {
			if e == envVar.Name {
				dockerEnvs = append(dockerEnvs, envVar.Name+"="+envVar.Value)
			}
		}
	}
	return dockerEnvs, nil
}
