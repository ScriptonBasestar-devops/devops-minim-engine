package main

import (
	"archive/tar"
	"bytes"
	"fmt"
	"github.com/cemacs/devops-engine/util"
	"github.com/docker/cli/cli/command/image/build"
	"github.com/docker/docker/pkg/archive"
	"github.com/docker/docker/pkg/stringid"
	"io"
	"io/ioutil"
	"os"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
)

func AddDockerfileToBuildContext(dockerfileCtx io.ReadCloser, buildCtx io.ReadCloser) (io.ReadCloser, string, error) {
	file, err := ioutil.ReadAll(dockerfileCtx)
	dockerfileCtx.Close()
	if err != nil {
		return nil, "", err
	}
	now := time.Now()
	hdrTmpl := &tar.Header{
		Mode:       0600,
		Uid:        0,
		Gid:        0,
		ModTime:    now,
		Typeflag:   tar.TypeReg,
		AccessTime: now,
		ChangeTime: now,
	}
	randomName := ".dockerfile." + stringid.GenerateRandomID()[:20]

	buildCtx = archive.ReplaceFileTarWrapper(buildCtx, map[string]archive.TarModifierFunc{
		// Add the dockerfile with a random filename
		randomName: func(_ string, h *tar.Header, content io.Reader) (*tar.Header, []byte, error) {
			return hdrTmpl, file, nil
		},
		// Update .dockerignore to include the random filename
		".dockerignore": func(_ string, h *tar.Header, content io.Reader) (*tar.Header, []byte, error) {
			if h == nil {
				h = hdrTmpl
			}

			b := &bytes.Buffer{}
			if content != nil {
				if _, err := b.ReadFrom(content); err != nil {
					return nil, nil, err
				}
			} else {
				b.WriteString(".dockerignore")
			}
			b.WriteString("\n" + randomName + "\n")
			return h, b.Bytes(), nil
		},
	})
	return buildCtx, randomName, nil
}

func main() {
	cli, err := client.NewClientWithOpts(client.WithVersion("1.39"), client.WithScheme("http"))
	util.OMG(err)

	//dockerfileCtx0, err := os.Open("/work/devops/devops-MinG-engine/src/cmd/Dockerfile")
	//util.OMG(err)
	//defer dockerfileCtx0.Close()
	buildCtx0, err := ioutil.ReadAll(os.Open("/work/devops/devops-MinG-engine/src/cmd/"))
	util.OMG(err)
	defer buildCtx0.Close()
	//defaultHeaders := map[string]string{"User-Agent": "ego-v-0.0.1"}

	//buildCtx, relDockerfile, err := AddDockerfileToBuildContext(dockerfileCtx0, buildCtx0)
	//util.OMG(err)
	buildCtx, relDockerfile, err := build.GetContextFromReader(buildCtx0, "Dockerfile")
	util.OMG(err)
	fmt.Println("=================================")
	fmt.Println(buildCtx)
	fmt.Println(relDockerfile)
	fmt.Println("=================================")

	options := types.ImageBuildOptions{
		Dockerfile:     relDockerfile,
		SuppressOutput: false,
		Remove:         true,
		ForceRemove:    true,
		PullParent:     true,
	}
	buildResponse, err := cli.ImageBuild(context.Background(), buildCtx, options)
	if err != nil {
		fmt.Printf("%s", err.Error())
	}
	fmt.Printf("********* %s **********", buildResponse.OSType)
	response, err := ioutil.ReadAll(buildResponse.Body)
	if err != nil {
		fmt.Printf("%s", err.Error())
	}
	fmt.Println(string(response))
}
