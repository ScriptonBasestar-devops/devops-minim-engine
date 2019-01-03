package config

import (
	"gopkg.in/validator.v2"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"util"
)

//TODO 나중에 enum 형태로 수정
//type vcsType int
//const (
//	git vcsType = 0
//	hg
//	svn
//)
//
//var VcsType = [...]string{
//	"git",
//	"hg",
//	"svn",
//}

//TODO more validation
type YamlConfig struct {
	Vcs struct {
		//Name    vcsType  `yaml:"name"`
		Name    string   `yaml:"name" validate:"nonzero,regexp=(git|hg|svn)"`
		Repo    string   `yaml:"repo"`
		Branch  string   `yaml:"branch" validate:"nonzero"`
		Command []string `yaml:"command"`
	} `yaml:"vcs"`
	Build struct {
		RootPath    string   `yaml:"root_path" validate:"nonzero"`
		Dockerbuild string   `yaml:"dockerbuild"`
		Dockerimage string   `yaml:"dockerimage"`
		CopyTo      []string `yaml:"copy_to"`
		VolumeTo    []string `yaml:"volume_to"`
		ExtractFrom []string `yaml:"extract_from"`
		Test        []string `yaml:"test"`
		Script      []string `yaml:"script"`
	} `yaml:"build"`
	Deploy struct {
		RootPath    string   `yaml:"root_path" validate:"nonzero"`
		Dockerbuild string   `yaml:"dockerbuild"`
		Dockerimage string   `yaml:"dockerimage"`
		CopyTo      []string `yaml:"copy_to"`
		Test        []string `yaml:"test"`
		Script      []string `yaml:"script"`
	} `yaml:"deploy"`
}

func (c *YamlConfig) ReadConfig(filepath string) *YamlConfig {
	data, err := ioutil.ReadFile(filepath)
	util.OMG(err)
	err = yaml.Unmarshal([]byte(data), c)
	util.OMG(err)
	c.validate()
	return c
}

func (c *YamlConfig) validate() {
	err := validator.Validate(*c)
	panic(err)
}
