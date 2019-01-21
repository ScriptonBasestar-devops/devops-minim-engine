package config

import (
	"fmt"
	"github.com/cemacs/devops-engine/util"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"strings"
)

type CopySource struct {
	Source string `yaml:"source"`
	From   string `yaml:"from"`
	To     string `yaml:"to"`
}

type BoxItems map[string]interface{}
type BoxImage struct {
	Dockerimage string `yaml:"dockerimage"`
	Lifecycle   string `yaml:"lifecycle"`
}
type BoxFile struct {
	Dockerfile string `yaml:"dockerfile"`
	Context    string `yaml:"context"`
	Add        []struct {
		CopyIn CopySource `yaml:"copy_in"`
	} `yaml:"add"`
	Args      []string `yaml:"args"`
	Lifecycle string   `yaml:"lifecycle"`
}
type BoxRepo struct {
	Dockerrepo string `yaml:"dockerrepo"`
	Path       string `yaml:"path"`
	Branch     string `yaml:"branch"`
	Revision   string `yaml:"rev"`
	Lifecycle  string `yaml:"lifecycle"`
}

func ParseBoxRepo(bytes []byte) interface{} {
	var c interface{}
	switch {
	case strings.Contains(string(bytes), "dockerfile:"):
		c = &BoxFile{}
	case strings.Contains(string(bytes), "dockerimage:"):
		c = &BoxImage{}
	case strings.Contains(string(bytes), "dockerrepo:"):
		c = &BoxRepo{}
	default:
		panic(fmt.Sprintf("설정오류: box 타입에 없는 값이 입력됨 - %s", string(bytes)))
	}
	err := yaml.Unmarshal(bytes, c)
	util.OMG(err)
	return c
}

type ExecItems []interface{}
type ExecCopyIn struct {
	CopyIn CopySource `yaml:"copy_in"`
}
type ExecCopyOut struct {
	CopyOut CopySource `yaml:"copy_out"`
}
type ExecCommand struct {
	Cmd []string `yaml:"command"`
}
type ExecSnapOut struct {
	SnapOut struct {
		File map[string]string
	} `yaml:"snap_out"`
}

func ParseExecItem(bytes []byte) interface{} {
	var c interface{}
	switch {
	case strings.Contains(string(bytes), "copy_in:"):
		c = &ExecCopyIn{}
	case strings.Contains(string(bytes), "command:"):
		c = &ExecCommand{}
	case strings.Contains(string(bytes), "copy_out"):
		c = &ExecCopyOut{}
	case strings.Contains(string(bytes), "snap_out:"):
		c = &ExecSnapOut{}
	default:
		panic(fmt.Sprintf("설정오류: exec 타입에 없는 값이 입력됨 - %s", string(bytes)))
	}
	err := yaml.Unmarshal(bytes, c)
	util.OMG(err)
	return c
}

type PackageItems map[string]interface{}
type PackageImage struct {
	Dockerimage string            `yaml:"dockerimage"`
	Args        []string          `yaml:"args"`
	SnapIn      map[string]string `yaml:"snap_in"`
	Entrypoint  []string          `yaml:"entrypoint"`
	PushTo      []struct {
		Target string `yaml:"target"`
		Url    string `yaml:"url"`
		Name   string `yaml:"name"`
	} `yaml:"push_to"`
}
type PackageFile struct {
	CompressName string            `yaml:"compress_name"`
	CompressType string            `yaml:"compress_type"`
	SnapIn       map[string]string `yaml:"snap_in"`
	Add          []struct {
		CopyIn CopySource `yaml:"copy_in"`
	} `yaml:"add"`
}

func ParsePackagecItem(bytes []byte) interface{} {
	var c interface{}
	switch {
	case strings.Contains(string(bytes), "dockerimage:"):
		c = &PackageImage{}
	case strings.Contains(string(bytes), "compress_type:"):
		c = &PackageFile{}
	default:
		panic(fmt.Sprintf("설정오류: exec 타입에 없는 값이 입력됨 - %s", string(bytes)))
	}
	err := yaml.Unmarshal(bytes, c)
	util.OMG(err)
	return c
}

type YamlConfig struct {
	Meta struct {
		ProjectName    string `yaml:"project_name"`
		ProjectVersion string `yaml:"project_version"`
	} `yaml:"meta"`
	Cache struct {
		Group   map[string]string `yaml:"grp"`
		User    map[string]string `yaml:"usr"`
		Project map[string]string `yaml:"prj"`
		Temp    map[string]string `yaml:"tmp"`
		//Group   yaml.MapSlice `yaml:"grp"`
		//User    yaml.MapSlice `yaml:"usr"`
		//Project yaml.MapSlice `yaml:"prj"`
		//Temp    yaml.MapSlice `yaml:"tmp"`
	} `yaml:"cache"`
	Box BoxItems `yaml:"box"`
	//Box   map[string]interface{} `yaml:"box"`
	Build struct {
		Dockerimage string   `yaml:"dockerimage"`
		WorkRoot    string   `yaml:"work_root"`
		Envs        []string `yaml:"envs"`
		Volumes     []string `yaml:"volumes"`
		Exec        ExecItems
	} `yaml:"build"`
	Package PackageItems `yaml:"package"`
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
