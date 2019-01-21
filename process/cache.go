package process

import (
	"fmt"
	"github.com/cemacs/devops-engine/util"
	"github.com/docker/docker/api/types/volume"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
)

type CacheType int

const (
	grp CacheType = 1 + iota
	usr
	prj
	tmp
)

var cacheTypes = [...]string{
	"grp", "usr", "prj", "tmp",
}

func (cacheType CacheType) String() string {
	return cacheTypes[(cacheType-1)%4]
}

func cacheName(cacheType CacheType, middleName string, volumeName string) string {
	return fmt.Sprintf("%s-%s-%s", cacheType.String(), middleName, volumeName)
}

func CacheCreate(cli client.APIClient, cacheType CacheType, middleName string, volumeName string) string {
	ctx := context.Background()
	v, err := cli.VolumeCreate(ctx, volume.VolumeCreateBody{
		Name: cacheName(cacheType, middleName, volumeName),
		Labels: map[string]string{
			"build_no":     "prj",
			"project_name": "prj",
		},
	})
	util.OMG(err)
	return v.Name
}

func CacheRemove(cli client.APIClient, vid string) {
	ctx := context.Background()
	err := cli.VolumeRemove(ctx, vid, true)
	util.OMG(err)
}
