package service

type Service struct {
	Name    string   `yaml:"name"`
	Source  string   `yaml:"source"`
	Image   string   `yaml:"image,omitempty"`
	Path    string   `yaml:"path,omitempty"`
	Ports   []string `yaml:"ports"`
	Volumes []string `yaml:"volumes"`
	Env     []string `yaml:"env"`
	Pre     []string `yaml:"pre,omitempty"`
	Exec    []string `yaml:"exec,omitempty"`
}

type Command struct {
	Command string   `yaml:"command"`
	Path    string   `yaml:"path"`
	Steps   []string `yaml:"steps"`
}

type EnvVar struct {
	Name  string `yaml:"name"`
	Value string `yaml:"value"`
}

type dockerService struct {
	Name    string
	Image   string
	Ports   []string
	Volumes []string
	Env     []string
	Pre     []string
}

type Config struct {
	Services []Service `yaml:"services"`
	Commands []Command `yaml:"commands"`
	Env      []EnvVar  `yaml:"env"`
}

type localService struct {
	Name string
	Path string
	Exec []string
	Pre  []string
	Env  []string
}
