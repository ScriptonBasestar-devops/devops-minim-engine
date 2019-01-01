package src

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
		Name    VcsType
		Branch  string
		Command []string
	}
	Build struct {
		RootPath    string
		Dockerbuild string
		Dockerimage string
		CopyTo      []string
		VolumeTo    []string
		ExtractFrom []string
		Test        []string
		Script      []string
	}
	Deploy struct {
		RootPath    string
		Dockerbuild string
		Dockerimage string
		CopyTo      []string
		Test        []string
		Script      []string
	}
}
