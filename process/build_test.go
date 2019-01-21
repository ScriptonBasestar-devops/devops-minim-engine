package process_test

import (
	"config"
	"fmt"
	"prepare"
	"testing"
)

func YamlConfigSample() config.YamlConfig {
	y := (&config.YamlConfig{}).ReadConfig("MinG.yaml")
	return *y
}

func TestBuild(t *testing.T) {
	y := YamlConfigSample()
	process.Build(&y, ".")
}

func TestConfig_ReadConfig(t *testing.T) {
	y := YamlConfigSample()
	fmt.Println("=============0 cache 0=============")
	fmt.Println("============= group =============")
	for idx, item := range y.Cache.Group {
		fmt.Println(idx, item.Key, item.Value)
	}
	fmt.Println("============= user =============")
	for idx, item := range y.Cache.User {
		fmt.Println(idx, item.Key, item.Value)
	}
	fmt.Println("============= project =============")
	for idx, item := range y.Cache.Project {
		fmt.Println(idx, item.Key, item.Value)
	}
	fmt.Println("============= temp =============")
	for idx, item := range y.Cache.Temp {
		fmt.Println(idx, item.Key, item.Value)
	}
	fmt.Println("=============0 build 0=============")
	fmt.Println("Dockerfile", y.Build.Dockerfile)
	fmt.Println("Arg", y.Build.Arg)
	fmt.Println("ImageStatus", y.Build.ImageStatus)
	fmt.Println("Env", y.Build.Env)
	fmt.Println("Volume", y.Build.Volume)
	fmt.Println("Exec", y.Build.Exec)
	fmt.Println("Extract", y.Build.Extract)

	fmt.Println("=============0 deploy 0=============")
	fmt.Println("Dockerfile", y.Deploy.Dockerfile)
	fmt.Println("Arg", y.Deploy.Arg)
	fmt.Println("Inject", y.Deploy.Inject)
	fmt.Println("Entrypoint", y.Deploy.Entrypoint)
	fmt.Println("PushTo", y.Deploy.PushTo)
}
