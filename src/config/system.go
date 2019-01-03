package config

type SystemConfig struct {
	Source struct {
		Path string `yaml:"path"`
	} `yaml:"source"`
	Project struct {
		Name string `yaml: "name"`
		RootPath string `yaml: "root_path"`
	} `yaml: "project"`
}
