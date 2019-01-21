package process

import (
	"archive/tar"
	"bufio"
	"compress/gzip"
	"fmt"
	"github.com/cemacs/devops-engine/config"
	"github.com/cemacs/devops-engine/util"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
	"io"
	"log"
	"os"
	"strings"
)

func addFile(tw *tar.Writer, path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer util.OMG(file.Close())
	if stat, err := file.Stat(); err == nil {
		// now lets create the header as needed for this file within the tarball
		header := new(tar.Header)
		header.Name = path
		header.Size = stat.Size()
		header.Mode = int64(stat.Mode())
		header.ModTime = stat.ModTime()
		// write the header to the tarball archive
		if err := tw.WriteHeader(header); err != nil {
			return err
		}
		// copy the file data to the tarball
		if _, err := io.Copy(tw, file); err != nil {
			return err
		}
	}
	return nil
}

// 이미지 생성
// return imgPrepare 이미지 이름 반환
func imgPrepare(cli *client.Client, projectName string, dockerfile string, arg map[string]*string, paths []string) string {
	// set up the output file
	tarPath := "/work/tmp/$projectName/output.tar.gz"
	file, err := os.Create(tarPath)
	if err != nil {
		log.Fatalln(err)
	}
	defer util.OMG(file.Close())
	// set up the gzip writer
	gw := gzip.NewWriter(file)
	defer util.OMG(gw.Close())
	tw := tar.NewWriter(gw)
	defer util.OMG(tw.Close())

	// add each file as needed into the current tar archive
	for i := range paths {
		if err := addFile(tw, paths[i]); err != nil {
			log.Fatalln(err)
		}
	}

	//context tarfile 맞는지 확인 필요
	dockerBuildContext, err := os.Open(tarPath)
	util.OMG(err)
	defer util.OMG(dockerBuildContext.Close())

	//이미지 가져오기 또는 이미지 생성
	imgName := projectName + "-build"
	ctx := context.Background()

	_, err = cli.ImageBuild(ctx, dockerBuildContext, types.ImageBuildOptions{
		Dockerfile: dockerfile,
		PullParent: true,
		NoCache:    true,
		Tags:       []string{imgName},
		BuildArgs:  arg,
		Remove:     true,
	})

	util.OMG(err)
	return imgName
}

//컨테이너 실행
func ctnStart(cli *client.Client, imageName string, volume map[string]struct{}) string {
	ctx := context.Background()

	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image:      imageName,
		Env:        []string{},
		Cmd:        []string{},
		Entrypoint: []string{},
		WorkingDir: "",
		Volumes:    volume,
		//TODO add container name with project pattern
	}, &container.HostConfig{
		VolumesFrom: []string{"", ""},
	}, nil, "")
	util.OMG(err)

	err = cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{})
	util.OMG(err)

	_, err = cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{
		ShowStdout: true,
	})
	util.OMG(err)
	return resp.ID
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

func Build(c *config.YamlConfig, projectRoot string) {
	cli, err := client.NewClientWithOpts(client.WithVersion("1.39"), client.WithScheme("http"))
	util.OMG(err)

	imgName := imgPrepare(cli, c.Meta.ProjectName, c.Build.Dockerfile, util.ArgDataArrToMap(c.Build.Arg))

	ctnId := ctnStart(cli, imgName, util.VolumeDataArrToMap(c.Build.Volume))
	util.OMG(err)
	for _, execKV := range c.Build.Exec {
		switch execKV.Act {
		case "command":
			outExec := ctnExec(cli, ctnId, execKV.Value)
			_, err = io.Copy(os.Stdout, outExec)
			util.OMG(err)
		case "copy_in":
			break
		case "copy_out":
			break
		default:
			panic("설정오류")
		}
	}
	err = cli.ContainerRemove(context.Background(), ctnId, types.ContainerRemoveOptions{})
	util.OMG(err)
}
