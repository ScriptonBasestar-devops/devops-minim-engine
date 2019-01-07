package vcs

import (
	"config"
	"fmt"
	"path"
	"testing"
)

var workRoot = "/tmp/test-devosMinG"
var branch = "master"

func TestClone(t *testing.T) {
	//clone
	result := clone(path.Join(workRoot, "testvcs"), "https://github.com/cemacs-tv/devops-MinG.git", branch)
	fmt.Println(result)

	//work work work

	//fetch clean reset
	result = latest(path.Join(workRoot, "testvcs"), branch)
	fmt.Println(result)
}

func yamlConfigSample() config.YamlConfig {
	y := (&config.YamlConfig{}).ReadConfig("MinG.yaml")
	return *y
}
func systemConfigSample() config.SystemConfig {
	sc := config.SystemConfig{}
	sc.Source.Path = "/tmp/test1/devops/source"
	sc.Project.Name = "devops-MinG"
	sc.Project.RootPath = "/tmp/test1/devops/application"
	return sc
}

func TestPrepareVCS(t *testing.T) {
	y := yamlConfigSample()
	s := systemConfigSample()
	PrepareVCS(&s, &y)
}
