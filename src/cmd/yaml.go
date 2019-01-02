package main

type VcsType int

const (
	git VcsType = 0
	hg
	svn
)

var vcsType = [...]string{
	"git",
	"hg",
	"svn",
}

type Config struct {
	Vcs struct {
		Name    VcsType  `yaml:"name"`
		Branch  string   `yaml:"branch"`
		Command []string `yaml:"command"`
	} `yaml:"vcs"`
	Build struct {
		RootPath    string   `yaml:"root_path"`
		Dockerbuild string   `yaml:"dockerbuild"`
		Dockerimage string   `yaml:"dockerimage"`
		CopyTo      []string `yaml:"copy_to"`
		VolumeTo    []string `yaml:"volume_to"`
		ExtractFrom []string `yaml:"extract_from"`
		Test        []string `yaml:"test"`
		Script      []string `yaml:"script"`
	} `yaml:"build"`
	Deploy struct {
		RootPath    string   `yaml:"root_path"`
		Dockerbuild string   `yaml:"dockerbuild"`
		Dockerimage string   `yaml:"dockerimage"`
		CopyTo      []string `yaml:"copy_to"`
		Test        []string `yaml:"test"`
		Script      []string `yaml:"script"`
	} `yaml:"deploy"`
}
