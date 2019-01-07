package config

//init from db..?
type SystemConfig struct {
	Source struct {
		Path string
	}
	Project struct {
		Name     string
		RootPath string
	}
}
