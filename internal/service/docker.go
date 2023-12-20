package service

import (
	"context"
	"fmt"
	"log"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
)

func newDockerClient(ctx context.Context) (*client.Client, error) {
	return client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
}

func (d *dockerService) getContainerConfigs(ctx context.Context) (*container.Config, error) {
	exposedPorts, _, err := parsePortConfigs(ctx, d.Ports)
	if err != nil {
		return nil, fmt.Errorf("error parsing port configs: %v", err)
	}
	dockerEnvs, err := getDockerEnvConfigs(ctx, d.Env)
	if err != nil {
		return nil, fmt.Errorf("error getting docker env configs: %v", err)
	}
	return &container.Config{
		Image:        d.Image,
		Env:          dockerEnvs,
		ExposedPorts: exposedPorts,
		Labels: map[string]string{
			"starter": "up",
		},
	}, nil
}

func (d *dockerService) getHostConfigs(ctx context.Context) (*container.HostConfig, error) {
	_, portBindings, err := parsePortConfigs(ctx, d.Ports)
	if err != nil {
		return nil, fmt.Errorf("error parsing port configs: %v", err)
	}
	return &container.HostConfig{
		Binds:        d.Volumes,
		PortBindings: portBindings,
	}, nil
}

func (d *dockerService) validate(ctx context.Context) error {
	return nil
}

func (d *dockerService) deployDocker(ctx context.Context) error {
	client, err := newDockerClient(ctx)
	if err != nil {
		return fmt.Errorf("error while creating docker client: %v", err)
	}
	containerConfigs, err := d.getContainerConfigs(ctx)
	if err != nil {
		return fmt.Errorf("error while getting container configs: %v", err)
	}
	hostConfigs, err := d.getHostConfigs(ctx)
	if err != nil {
		return fmt.Errorf("error while getting host configs: %v", err)
	}
	fmt.Println("Creating container...")
	if d.ContainerName == "" {
		d.ContainerName = d.Name
	}
	endpointConfigs := map[string]*network.EndpointSettings{}
	for _, p := range d.Networks {
		endpointConfigs[p] = &network.EndpointSettings{}
	}
	resp, err := client.ContainerCreate(ctx, containerConfigs, hostConfigs, &network.NetworkingConfig{
		EndpointsConfig: endpointConfigs,
	}, nil, d.ContainerName)
	if err != nil {
		return fmt.Errorf("error while creating container: %v", err)
	}
	fmt.Println("Container created with ID:", resp.ID)
	if err := client.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		return fmt.Errorf("error while starting container: %v", err)
	}
	fmt.Println("Container started with ID:", resp.ID)
	return nil
}

func CleanupDockerContainers(ctx context.Context, completed chan bool) {
	<-ctx.Done() // Wait for the context to be canceled

	ctx = context.Background()
	cli, err := newDockerClient(ctx)
	if err != nil {
		panic(fmt.Errorf("error while cleaning up docker containers: %v", err))
	}
	filter := filters.NewArgs()
	filter.Add("label", "starter=up")
	// List all containers
	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{Filters: filter, All: true})
	if err != nil {
		log.Fatalf("Error listing containers: %v", err)
	}

	// Kill and remove each container
	for _, container := range containers {
		// Kill the container
		if err := cli.ContainerKill(ctx, container.ID, "SIGKILL"); err != nil {
			log.Printf("Error killing container %s: %v", container.ID, err)
			// continue
		}

		// Remove the container
		if err := cli.ContainerRemove(ctx, container.ID, types.ContainerRemoveOptions{}); err != nil {
			log.Printf("Error removing container %s: %v", container.ID, err)
		}
	}

	fmt.Println("All containers have been killed and removed")
	completed <- true
}

func (d *DockerNetwork) Create(ctx context.Context) error {
	cli, err := newDockerClient(ctx)
	if err != nil {
		return fmt.Errorf("error while creating docker client: %v", err)
	}
	_, err = cli.NetworkCreate(ctx, d.Name, types.NetworkCreate{
		Driver:         d.Driver,
		Attachable:     true,
		CheckDuplicate: true,
		Labels: map[string]string{
			"starter": "up",
		},
	})
	if err != nil {
		return fmt.Errorf("error while creating network: %v", err)
	}
	return nil
}
