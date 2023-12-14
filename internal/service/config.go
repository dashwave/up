package service

import (
	"context"
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

var UpConfigs *Config

func InitConfigFromYaml(ctx context.Context, filepath string) error {
	configData, err := os.ReadFile(filepath)
	if err != nil {
		return fmt.Errorf("error reading config file: %v", err)
	}
	if err := yaml.Unmarshal(configData, &UpConfigs); err != nil {
		return fmt.Errorf("error unmarshalling config file: %v", err)
	}
	return nil
}
