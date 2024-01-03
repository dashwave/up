package deploy

import (
	"context"
	"fmt"

	"github.com/dashwave/up/internal/service"
)

func Deploy(ctx context.Context, configFile string) error {
	if err := service.InitConfigFromYaml(ctx, configFile); err != nil {
		return err
	}
	fmt.Println(service.UpConfigs)
	for _, n := range service.UpConfigs.Networks {
		if err := n.Create(ctx); err != nil {
			fmt.Printf("error creating network %s: %v\n", n.Name, err)
			// return err
		}
	}
	for _, v := range service.UpConfigs.Volumes {
		if err := service.CreateDockerVolume(ctx, v); err != nil {
			fmt.Printf("error creating volume %s: %v\n", v, err)
			// return err
		}
	}
	for _, d := range service.UpConfigs.Deploy {
		var s service.Service
		serviceFound := false
		for _, service := range service.UpConfigs.Services {
			if service.Name == d {
				s = service
				serviceFound = true
				break
			}
		}
		if !serviceFound {
			return fmt.Errorf("service %s not found", d)
		}
		if err := s.Validate(ctx); err != nil {
			return err
		}
		if err := s.Deploy(ctx); err != nil {
			fmt.Println("error deploying service", err)
		}
	}
	return nil
}
