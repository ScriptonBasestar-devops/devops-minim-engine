package build_test

import (
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
	"testing"
	"util"
)

func TestBuild(t *testing.T) {
	//func ExampleDocker() {
	fmt.Println("==================")
	fmt.Println("start Example Docker")
	cli, err := client.NewClientWithOpts(client.WithVersion("1.39"), client.WithScheme("http"))
	util.OMG(err)

	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})

	for _, container := range containers {
		fmt.Printf("%s %s\n", container.ID[:10], container.Image)
	}

	images, err := cli.ImageList(context.Background(), types.ImageListOptions{})
	util.OMG(err)
	for _, image := range images {
		//fmt.Println(image)
		//fmt.Printf("ID: %s, Name: %s\n", image.ID[:10], strings.Split(image.RepoDigests[0], ":")[0])
		fmt.Printf("ID: %s, Name: %s\n", image.ID[:10], image.RepoTags)
	}
	// Output: MOOOO!
}
