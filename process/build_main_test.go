package process_test

import (
	"fmt"
	"github.com/cemacs/devops-engine/process"
	"github.com/cemacs/devops-engine/util"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
	"io"
	"io/ioutil"
	"os"
	"path"
	"testing"
)

func newClient() *client.Client {
	cli, err := client.NewClientWithOpts(client.WithVersion("1.39"), client.WithScheme("http"))
	util.OMG(err)
	return cli
}

// ==================== image use ====================
func TestImagePull(t *testing.T) {
	fmt.Println("======s TestExampleImagePull s======")
	ctx := context.Background()
	cli := newClient()

	resp, err := cli.ImagePull(ctx, "alpine", types.ImagePullOptions{})
	util.OMG(err)
	body, err := ioutil.ReadAll(resp)
	util.OMG(err)
	fmt.Println("======s body s======")
	fmt.Println(string(body))
	fmt.Println("======e body e======")
	fmt.Println("======e TestExampleImagePull e======")
}

func TestContainerStart(t *testing.T) {
	fmt.Println("======s TestContainerStart s======")
	ctx := context.Background()
	cli := newClient()

	fmt.Println("container create")
	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: "alpine",
		Cmd:   []string{"echo", "hello world in container"},
		//Tty:   true,
	}, &container.HostConfig{
		//AutoRemove: true,
	}, nil, "")
	util.OMG(err)

	fmt.Println("container start")
	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	fmt.Println("container wait")
	statusCh, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			panic(err)
		}
	case <-statusCh:
	}

	fmt.Println("container logs")
	out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
	util.OMG(err)
	fmt.Println("======s body s======")
	_, err = io.Copy(os.Stdout, out)
	fmt.Println("======e body e======")
	util.OMG(err)

	err = cli.ContainerRemove(ctx, resp.ID, types.ContainerRemoveOptions{})
	util.OMG(err)

	fmt.Println("======e TestContainerStart e======")
}

// 컨테이너 이름 수동
func TestContainerFind(t *testing.T) {
	containerName := "elegant_galories"
	fmt.Println("======s TestContainerFind s======")
	ctx := context.Background()
	cli := newClient()

	filterArgs := filters.NewArgs()
	filterArgs.Add("name", containerName)
	// find container
	fmt.Println("container list")
	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{
		Size:    true,
		All:     true,
		Since:   "container",
		Filters: filterArgs,
	})
	util.OMG(err)

	fmt.Println(containers)
	fmt.Println("======e TestContainerFind e======")
}

// container id 수동으로 가져오기 7e91e7fe39bc39a65d418b0bc644a2c93fd7dc7f969cc9abf12d03ce037b2bd1
func TestContainerExec(t *testing.T) {
	containerId := "7e91e7fe39bc39a65d418b0bc644a2c93fd7dc7f969cc9abf12d03ce037b2bd1"
	fmt.Println("======s TestContainerExec s======")
	ctx := context.Background()
	cli := newClient()

	// exec 1
	fmt.Println("======s exec1 s======")
	fmt.Println("container exec create")
	idResp, err := cli.ContainerExecCreate(ctx, containerId, types.ExecConfig{
		//Cmd: []string{"touch", "/root/aaaaa"},
		Cmd: []string{"echo", "hello world in container - with exec"},
		Tty: true,
		//AttachStdin:  true,
		AttachStdout: true,
		//Detach:       false,
	})
	util.OMG(err)

	fmt.Println(idResp.ID)
	//fmt.Println("container exec start")
	//err = cli.ContainerExecStart(context.Background(), idResp.ID, types.ExecStartCheck{
	//	//Detach: true,
	//	//Tty:    true,
	//})
	//util.OMG(err)
	fmt.Println("container exec attach start")
	hjResp, err := cli.ContainerExecAttach(context.Background(), idResp.ID, types.ExecStartCheck{
		Tty: true,
	})
	fmt.Println("======s body s======")
	_, err = io.Copy(os.Stdout, hjResp.Reader)
	util.OMG(err)
	fmt.Println("======e body e======")
	fmt.Println("======e exec1 e======")
	fmt.Println("======e TestContainerExec e======")
}

func ExampleContainerStartAndExec() {
	fmt.Println("ExampleContainerStartAndExec start")
	ctx := context.Background()
	cli := newClient()

	fmt.Println("container create")
	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: "alpine",
		Cmd:   []string{"echo", "hello world in container - with create"},
		Tty:   true,
	}, nil, nil, "")
	util.OMG(err)

	fmt.Println("container start")
	err = cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{})
	util.OMG(err)

	// exec 1
	fmt.Println("exec1 ===============")
	fmt.Println("container exec create")
	idResp, err := cli.ContainerExecCreate(ctx, resp.ID, types.ExecConfig{
		Cmd: []string{"echo", "hello world in container - with exec"},
		Tty: true,
	})
	util.OMG(err)

	fmt.Println("container exec start")
	err = cli.ContainerExecStart(context.Background(), idResp.ID, types.ExecStartCheck{
		Tty: true,
	})
	util.OMG(err)

	fmt.Println("container wait")
	statusCh, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			panic(err)
		}
	case <-statusCh:
	}

	fmt.Println("container logs")
	out, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
	util.OMG(err)
	_, err = io.Copy(os.Stdout, out)
	util.OMG(err)

	// exec 2
	//fmt.Println("exec2 ===============")
	//fmt.Println("container exec create")
	//idResp, err = cli.ContainerExecCreate(ctx, resp.ID, types.ExecConfig{
	//	Cmd: []string{"echo", "hello world in container"},
	//})
	//util.OMG(err)
	//
	//fmt.Println("container exec start")
	//err = cli.ContainerExecStart(context.Background(), idResp.ID, types.ExecStartCheck{})
	//util.OMG(err)
	//
	//fmt.Println("container logs")
	//out, err = cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{ShowStdout: true})
	//util.OMG(err)
	//_, err = io.Copy(os.Stdout, out)
	//util.OMG(err)

	//fmt.Println("container stop")
	//timeout, err := time.ParseDuration("5s")
	//util.OMG(err)
	//err = cli.ContainerStop(ctx, resp.ID, &timeout)
	//util.OMG(err)

	fmt.Println("ExampleContainerStartAndExec end")
	// Output: hello
}

func ExampleDockerImage() {
	ctx := context.Background()
	cli := newClient()

	ctx.Value("ubuntu")
	image0, err := cli.ImageList(ctx, types.ImageListOptions{})
	util.OMG(err)

	fmt.Println("image0 ==================")
	fmt.Println(image0)
	image, err := cli.ImagePull(context.Background(), "ubuntu:18.04", types.ImagePullOptions{})
	util.OMG(err)
	fmt.Println("image ===================")
	fmt.Println(image)
	fmt.Println("oo")
	// Output: oo
}

func ExampleImageBuild() {
	ctx := context.Background()
	cli := newClient()

	//cli.ImageCreate(ctx, "", types.ImageCreateOptions{})
	//	fmt.Println("============")
	_, err := cli.ImageBuild(ctx, nil, types.ImageBuildOptions{
		Dockerfile: path.Join("./", "MinG.Build.Dockerfile"),
		PullParent: true,
		NoCache:    true,
		Tags:       []string{"testproj1" + "-build"},
	})
	fmt.Println("============")
	util.OMG(err)
	fmt.Println("kkk")
	// Output: kkk
}

func TestMainList(t *testing.T) {
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
}

func TestImgaePrepare(t *testing.T) {
	process.Build()
}
