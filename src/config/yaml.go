package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"util"
)

type YamlConfig struct {
	Meta struct {
		ProjectName    string `yaml:"project_name"`
		ProjectVersion string `yaml:"project_version"`
	} `yaml:"meta"`
	Cache struct {
		Group   yaml.MapSlice `yaml:"grp"`
		User    yaml.MapSlice `yaml:"usr"`
		Project yaml.MapSlice `yaml:"prj"`
		Temp    yaml.MapSlice `yaml:"tmp"`
	} `yaml:"cache"`
	Build struct {
		Dockerfile  string   `yaml:"dockerfile"`
		Arg         []string `yaml:"arg"`
		ImageStatus string   `yaml:"image_status"`
		Env         []string `yaml:"env"`
		Volume      []string `yaml:"volume"`
		Exec        []struct {
			Act   string   `yaml:"act"`
			Value []string `yaml:"value"`
		} `yaml:"exec"`
		Extract yaml.MapSlice `yaml:"extract"`
	} `yaml:"build"`
	Deploy struct {
		Dockerfile string        `yaml:"dockerfile"`
		Arg        []string      `yaml:"arg"`
		Inject     yaml.MapSlice `yaml:"inject"`
		Entrypoint []string      `yaml:"entrypoint"`
		PushTo     []struct {
			Target string `yaml:"target"`
			Url    string `yaml:"url"`
			Name   string `yaml:"name"`
		} `yaml:"push_to"`
	} `yaml:"deploy"`
}

func (c *YamlConfig) ReadConfig(filepath string) *YamlConfig {
	data, err := ioutil.ReadFile(filepath)
	util.OMG(err)
	err = yaml.Unmarshal([]byte(data), c)
	util.OMG(err)
	//c.validate()
	return c
}

//func (c *YamlConfig) validate() {
//	err := validator.Validate(*c)
//	panic(err)
//}
