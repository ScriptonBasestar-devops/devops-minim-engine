package config

func systemConfigSample() SystemConfig {
	sc := SystemConfig{}
	sc.Source.Path = "/tmp/test1/devops/source"
	sc.Project.Name = "devops-MinG"
	sc.Project.RootPath = "/tmp/test1/devops/application"
	return sc
}
