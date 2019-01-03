package config_test

import (
	"config"
	"fmt"
	"testing"
)

func YamlConfigSample() config.YamlConfig {
	y := (&config.YamlConfig{}).ReadConfig("MinG.yaml")
	return *y
}

func ExampleConfig_ReadConfig() {
	//c := config.YamlConfig{}
	//c.
}

func TestConfig_ReadConfig(t *testing.T) {
	y := YamlConfigSample()
	fmt.Println("===============================")
	fmt.Println(y.Vcs.Name)
	fmt.Println(y.Vcs.Branch)
	fmt.Println(y.Build.RootPath)
	fmt.Println(y.Deploy.RootPath)
	fmt.Println(y)
}
