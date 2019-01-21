package unit_t_test

import (
	"fmt"
	"github.com/cemacs/devops-engine/config"
	"github.com/cemacs/devops-engine/util"
	"gopkg.in/yaml.v2"
	"testing"
)

func YamlConfigSample() config.YamlConfig {
	y := (&config.YamlConfig{}).ReadConfig("MinG.yaml")
	return *y
}

func TestConfig_ReadConfig(t *testing.T) {
	y := YamlConfigSample()
	fmt.Println("=============0 cache 0=============")
	fmt.Println("============= group =============")
	for key, val := range y.Cache.Group {
		fmt.Println(key, val)
	}
	fmt.Println("============= user =============")
	for key, val := range y.Cache.User {
		fmt.Println(key, val)
	}
	fmt.Println("============= project =============")
	for key, val := range y.Cache.Project {
		fmt.Println(key, val)
	}
	fmt.Println("============= temp =============")
	for key, val := range y.Cache.Temp {
		fmt.Println(key, val)
	}
	fmt.Println("=============0 box 0=============")
	fmt.Println("=============1 box  gradle5 1=============")
	bytes, err := yaml.Marshal(y.Box["gradle5"])
	util.OMG(err)

	boxGradle5i := config.ParseBoxRepo(bytes)
	fmt.Println(boxGradle5i)
	boxGradle5 := boxGradle5i.(*config.BoxFile)
	fmt.Println(boxGradle5.Dockerfile)
	fmt.Println(boxGradle5.Context)
	fmt.Println(boxGradle5.Args)
	fmt.Println(boxGradle5.Add)
	fmt.Println(boxGradle5.Lifecycle)

	fmt.Println("=============1 box springboot2 1=============")
	bytes, err = yaml.Marshal(y.Box["springboot2"])
	util.OMG(err)

	boxSpringboot2i := config.ParseBoxRepo(bytes)
	fmt.Println(boxSpringboot2i)
	boxSpringboot2 := boxSpringboot2i.(*config.BoxRepo)
	fmt.Println(boxSpringboot2.Dockerrepo)
	fmt.Println(boxSpringboot2.Revision)
	fmt.Println(boxSpringboot2.Path)
	fmt.Println(boxSpringboot2.Branch)
	fmt.Println(boxSpringboot2.Lifecycle)

	fmt.Println("=============1 box maven 1=============")
	bytes, err = yaml.Marshal(y.Box["maven"])
	util.OMG(err)

	boxMaveni := config.ParseBoxRepo(bytes)
	fmt.Println(boxMaveni)
	boxMaven := boxMaveni.(*config.BoxImage)
	fmt.Println(boxMaven.Dockerimage)
	fmt.Println(boxMaven.Lifecycle)

	fmt.Println("=============0 build 0=============")
	fmt.Println("Dockerimage", y.Build.Dockerimage)
	fmt.Println("WorkRoot", y.Build.WorkRoot)
	fmt.Println("Envs", y.Build.Envs)
	fmt.Println("Volumes", y.Build.Volumes)

	fmt.Println("=============1 build exec 1=============")
	fmt.Println(len(y.Build.Exec))
	bytes, err = yaml.Marshal(y.Build.Exec)
	util.OMG(err)

	for _, item := range y.Build.Exec {
		bytes, err := yaml.Marshal(item)
		util.OMG(err)
		execItem := config.ParseExecItem(bytes)
		fmt.Println(execItem)
	}

	fmt.Println("=============0 deploy 0=============")
	fmt.Println("Package len", len(y.Package))

	fmt.Println("=============1 deploy kos_sample_image 1=============")
	bytes, err = yaml.Marshal(y.Package["kos_sample_image"])
	util.OMG(err)
	packageImagei := config.ParsePackagecItem(bytes)
	packageImage := packageImagei.(*config.PackageImage)
	fmt.Println(packageImage.Dockerimage)
	fmt.Println(packageImage.Args)
	fmt.Println(packageImage.Entrypoint)
	fmt.Println(packageImage.SnapIn)
	fmt.Println(packageImage.PushTo)

	fmt.Println("=============1 deploy kos_sample_tar 1=============")
	bytes, err = yaml.Marshal(y.Package["kos_sample_tar"])
	util.OMG(err)
	packageFilei := config.ParsePackagecItem(bytes)
	packageFile := packageFilei.(*config.PackageFile)
	fmt.Println(packageFile.CompressName)
	fmt.Println(packageFile.CompressType)
	fmt.Println(packageFile.SnapIn)
	fmt.Println(packageFile.Add)

}
