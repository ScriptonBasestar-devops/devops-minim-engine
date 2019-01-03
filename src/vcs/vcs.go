package vcs

import (
	"config"
	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/_examples"
	"os"
	"util"
)

func clone(sc *config.SystemConfig, yc *config.YamlConfig) string {
	//examples.CheckArgs("<url>", "<directory>")
	url := yc.Vcs.Repo
	directory := sc.Source.Path

	// Clone the given repository to the given directory
	examples.Info("git clone %s %s --recursive", url, directory)

	r, err := git.PlainClone(directory, false, &git.CloneOptions{
		URL:               url,
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
	})
	util.OMG(err)

	// ... retrieving the branch being pointed by HEAD
	ref, err := r.Head()
	util.OMG(err)

	return ref.Hash().String()
}

func PrepareVCS(sc *config.SystemConfig, yc *config.YamlConfig) {
	directory := sc.Source.Path
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		clone(sc, yc)
	} else {
		util.RunCommands(yc.Vcs.Command)
	}
}
