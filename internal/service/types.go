package service

type Service struct {
	Name          string   `yaml:"name"`
	ContainerName string   `yaml:"container_name,omitempty"`
	Source        string   `yaml:"source"`
	Image         string   `yaml:"image,omitempty"`
	Path          string   `yaml:"path,omitempty"`
	Ports         []string `yaml:"ports"`
	Volumes       []string `yaml:"volumes"`
	Env           []string `yaml:"env"`
	Pre           []string `yaml:"pre,omitempty"`
	Exec          []string `yaml:"exec,omitempty"`
	Networks      []string `yaml:"networks,omitempty"`
}

type Command struct {
	Command string   `yaml:"command"`
	Path    string   `yaml:"path"`
	Steps   []string `yaml:"steps"`
}

type DockerNetwork struct {
	Name   string `yaml:"name"`
	Driver string `yaml:"driver"`
}

type EnvVar string

type dockerService struct {
	Name          string
	Image         string
	ContainerName string
	Ports         []string
	Volumes       []string
	Env           []string
	Pre           []string
	Networks      []string
}

type Config struct {
	Services []Service        `yaml:"services"`
	Commands []Command        `yaml:"commands"`
	Env      []EnvVar         `yaml:"env"`
	Networks []*DockerNetwork `yaml:"networks"`
	Deploy   []string         `yaml:"deploy"`
}

type localService struct {
	Name string
	Path string
	Exec []string
	Pre  []string
	Env  []string
}
