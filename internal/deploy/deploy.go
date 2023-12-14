package deploy

import (
	"context"

	"github.com/dashwave/up/internal/service"
)

func Deploy(ctx context.Context, configFile string) error {
	if err := service.InitConfigFromYaml(ctx, configFile); err != nil {
		return err
	}
	for _, s := range service.UpConfigs.Services {
		if err := s.Validate(ctx); err != nil {
			return err
		}
		go func(s service.Service) {
			s.Deploy(ctx)
		}(s)
	}
	return nil
}
