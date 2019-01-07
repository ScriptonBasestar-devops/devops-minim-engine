package build

import (
	"bufio"
	"config"
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/volume"
	"github.com/docker/docker/client"
	"io"
	"log"
	"os"
	"path"
	"strings"
	"util"
)

// 이미지 생성
// return imgPrepare 이미지 이름 반환
func imgPrepare(cli *client.Client, projectName string, projectRoot string, dockerBuild string, dockerImage string, volumes []string) string {
	//이미지 가져오기 또는 이미지 생성
	var imgName string
	if dockerBuild != "" {
		imgName = projectName + "-build"
		ctx := context.Background()
		_, err := cli.ImageBuild(ctx, nil, types.ImageBuildOptions{
			Dockerfile: path.Join(projectRoot, dockerBuild),
			PullParent: true,
			NoCache:    true,
			Tags:       []string{imgName},
		})
		util.OMG(err)
	} else {
		imgName = dockerImage
		_, err := cli.ImagePull(context.Background(), dockerImage, types.ImagePullOptions{})
		util.OMG(err)
	}
	return imgName
}

//컨테이너 실행
func ctnStart(cli *client.Client, imageName string) (string, *bufio.Reader) {
	ctx := context.Background()

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: imageName,
		Tty:   true,
		//TODO add container name with project pattern
	}, nil, nil, "")
	util.OMG(err)

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
	util.OMG(err)
	return resp.ID, bufio.NewReader(out)
}

//컨테이너 명령 실행
func ctnExec(cli *client.Client, containerId string, command []string) *bufio.Reader {
	ctx := context.Background()
	idResp, err := cli.ContainerExecCreate(ctx, containerId, types.ExecConfig{
		Cmd: command,
		Tty: true,
		//AttachStdin:  true,
		AttachStdout: true,
		//Detach:       false,
	})
	util.OMG(err)

	hjResp, err := cli.ContainerExecAttach(context.Background(), idResp.ID, types.ExecStartCheck{
		Tty: true,
	})
	return hjResp.Reader
}

func ctnCopy(cli *client.Client, containerId string, copyTargets []string) {
	ctx := context.Background()
	for _, copyTarget := range copyTargets {
		arr := strings.Split(":", copyTarget)
		fileOrDir := arr[0]
		hostPath := arr[1]
		containerPath := arr[2]
		fileinfo, err := os.Stat(hostPath)
		if os.IsNotExist(err) {
			log.Fatal(fmt.Sprintln("%s 위치에 파일/디렉토리가 없습니다.", hostPath))
			panic(err)
		}
		if fileinfo.IsDir() {
			if fileOrDir != "dir" {
				panic("copy_to dir로 되어 있는데 디렉토리가 아닙니다.")
			}
		}

		f, err := os.Open(hostPath)
		util.OMG(err)
		err = cli.CopyToContainer(ctx, containerId, containerPath, f, types.CopyToContainerOptions{AllowOverwriteDirWithFile: true})
		util.OMG(err)
	}
}

func createVolume(cli *client.Client, volumes []string) {
	ctx := context.Background()

	for _, v := range volumes {
		arr := strings.Split(":", v)
		fileOrDir := arr[0]
		vName := arr[1]
		cPath := arr[2]

		v, err := cli.VolumeCreate(ctx, volume.VolumeCreateBody{Name: vName})
		util.OMG(err)

		v.Mountpoint = cPath
	}
}

func Build(sc *config.SystemConfig, yc *config.YamlConfig) {
	projectRoot := sc.Project.RootPath
	cli, err := client.NewClientWithOpts(client.WithVersion("1.39"), client.WithScheme("http"))
	util.OMG(err)

	imgName := imgPrepare(cli, sc.Project.Name, projectRoot, yc.Build.Dockerbuild, yc.Build.Dockerimage, yc.Build.VolumeTo)
	ctnId, out := ctnStart(cli, imgName)
	_, err = io.Copy(os.Stdout, out)
	util.OMG(err)
	for _, command := range yc.Build.Script {
		outExec := ctnExec(cli, ctnId, []string{command})
		_, err = io.Copy(os.Stdout, outExec)
		util.OMG(err)
	}
	err = cli.ContainerRemove(context.Background(), ctnId, types.ContainerRemoveOptions{})
	util.OMG(err)
}
