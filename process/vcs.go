package process

import (
	"fmt"
	"github.com/cemacs/devops-engine/util"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing"
)

// clone git repository
// $ git clone --single-branch --branch $branch $url $directory --recursive
func Clone(directory, url, branch string) string {
	//examples.Info("git clone --single-branch --branch %s %s %s --recursive", branch, url, directory)
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

// reset repo HEAD
// git fetch origin
// git branch $branch
// git reset --hard HEAD
// git clean -dfx
// //git submodule update
func Latest(directory, branch string) string {
	r, err := git.PlainOpen(directory)
	util.OMG(err)

	err = r.Fetch(&git.FetchOptions{})
	util.OMG(err)

	_, err = r.Branch(branch)
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

	//TODO 서브모듈도 똑같이 업데이트
	//s, err := w.Submodule("")
	//repo, err := s.Repository()
	//repo.

	return ref.Hash().String()
}
