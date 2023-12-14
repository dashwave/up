package service

import (
	"context"
	"fmt"
	"strings"

	"github.com/docker/go-connections/nat"
)

func parsePortConfigs(ctx context.Context, portMappings []string) (nat.PortSet, nat.PortMap, error) {
	exposedPorts := nat.PortSet{}
	portBindings := nat.PortMap{}

	for _, portMapping := range portMappings {
		parts := strings.Split(portMapping, ":")
		if len(parts) != 2 {
			return nil, nil, fmt.Errorf("invalid port mapping format: %s", portMapping)
		}

		containerPort := parts[0] + "/tcp" // Assuming TCP. Change to "/udp" if using UDP ports
		hostPort := parts[1]

		// Set up exposed ports
		exposedPorts[nat.Port(containerPort)] = struct{}{}

		// Set up port bindings
		portBindings[nat.Port(containerPort)] = []nat.PortBinding{
			{
				HostIP:   "0.0.0.0",
				HostPort: hostPort,
			},
		}
	}

	return exposedPorts, portBindings, nil
}
