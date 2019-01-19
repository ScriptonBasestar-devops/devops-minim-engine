package process

import (
	"fmt"
	"github.com/docker/docker/api/types/volume"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
	"util"
)

/*
cache:
#  grp:
#  usr:
  prj:
    m2: ${USER_HOME}/.m2
    gradle: ${USER_HOME}/.gradle
    npm: ${WORK_ROOT}/app/node_modules
  tmp:
    pass_package: # no default
*/
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

func CacheCreate(cacheType CacheType, middleName string, volumeName string) string {
	cli, err := client.NewClientWithOpts(client.WithVersion("1.39"), client.WithScheme("http"))
	util.OMG(err)
	ctx := context.Background()

	v, err := cli.VolumeCreate(ctx, volume.VolumeCreateBody{
		Name: cacheName(cacheType, middleName, volumeName),
	})
	util.OMG(err)
	//fmt.Println(v.Name)
	return v.Name
}

func CacheRemove(vid string) {
	cli, err := client.NewClientWithOpts(client.WithVersion("1.39"), client.WithScheme("http"))
	util.OMG(err)
	ctx := context.Background()

	err = cli.VolumeRemove(ctx, vid, true)
	util.OMG(err)
}
