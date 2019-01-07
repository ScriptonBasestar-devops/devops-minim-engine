package vcs

import (
	"config"
	"fmt"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"os"
	"path"
	"util"
)

func clone(directory, url, branch string) string {
	//examples.Info("git clone %s %s --recursive", url, directory)
	r, err := git.PlainClone(directory, false, &git.CloneOptions{
		URL:               url,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
		SingleBranch:      true,
		ReferenceName:     plumbing.ReferenceName(fmt.Sprintf("refs/heads/%s", branch)),
	})
	util.OMG(err)

	ref, err := r.Head()
	util.OMG(err)

	return ref.Hash().String()
}

func latest(directory, branch string) string {
	r, err := git.PlainOpen(directory)
	util.OMG(err)

	_, err = r.Branch(branch)
	util.OMG(err)
	err = r.Fetch(&git.FetchOptions{})
	util.OMG(err)

	ref, err := r.Head()
	util.OMG(err)

	w, err := r.Worktree()
	util.OMG(err)

	err = w.Reset(&git.ResetOptions{
		Commit: ref.Hash(),
		Mode:   git.HardReset,
	})
	util.OMG(err)

	err = w.Clean(&git.CleanOptions{
		Dir: true,
	})
	util.OMG(err)

	//TODO 서브모듈도 처리
	//s, err := w.Submodule("")
	//repo, err := s.Repository()
	//repo.

	return ref.Hash().String()
}

func PrepareVCS(sc *config.SystemConfig, yc *config.YamlConfig) {
	repo := yc.Vcs.Repo
	branch := yc.Vcs.Branch
	rootPath := sc.Project.RootPath
	directory := sc.Source.Path
	command := yc.Vcs.Command

	if _, err := os.Stat(path.Join(rootPath, directory)); os.IsNotExist(err) {
		clone(directory, repo, branch)
		return
	}

	if command != nil {
		util.RunCommands(command)
		return
	}

	latest(directory, branch)
	//util.RunCommands([]string{
	//	"git fetch origin",
	//	"git checkout master",
	//	//"git reset --hard HEAD",
	//	fmt.Sprintf("git reset --hard origin/%s", branch),
	//	"git clean -dfx",
	//})
}
