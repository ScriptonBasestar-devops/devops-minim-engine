package process

import (
	"fmt"
	"github.com/docker/docker/api/types/volume"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
	"testing"
	"util"
)

func TestMainCreateCache(t *testing.T) {
	name := "test1"

	cli, err := client.NewClientWithOpts(client.WithVersion("1.39"), client.WithScheme("http"))
	util.OMG(err)
	ctx := context.Background()

	v, err := cli.VolumeCreate(ctx, volume.VolumeCreateBody{
		Name: name,
	})
	util.OMG(err)
	fmt.Println("=======================")
	fmt.Println(v.Name)
	fmt.Println(v.Driver)
	fmt.Println(v.Scope)
	fmt.Println(v.Status)
	fmt.Println(v.Mountpoint)
}

func TestMainRemoveCache(t *testing.T) {
	vid := "test1"

	cli, err := client.NewClientWithOpts(client.WithVersion("1.39"), client.WithScheme("http"))
	util.OMG(err)
	ctx := context.Background()

	err = cli.VolumeRemove(ctx, vid, true)
	util.OMG(err)
	fmt.Println("=======================")
}
