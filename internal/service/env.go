package service

import (
	"bufio"
	"context"
	"os"
	"path"
	"strings"
)

func parseEnvFiles(ctx context.Context, files []string) ([]string, error) {
	envs := make([]string, 0)
	for _, f := range files {
		filePath := path.Join(os.Getenv("PWD"), f)
		file, err := os.Open(filePath)
		if err != nil {
			return nil, err
		}
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			if strings.Contains(line, "=") {
				envs = append(envs, line)
			}
		}
	}
	return envs, nil
}
