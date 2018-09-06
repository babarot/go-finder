package finder

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/b4b4r07/go-finder/installer"
)

// Fzf represents the filter instance
type Fzf struct {
	*Command
}

// Install installs the command
func (c Fzf) Install(path string) error {
	bin := filepath.Join(path, "fzf")
	if _, err := os.Stat(bin); err == nil {
		// Already installed, no need to install
		return nil
	}

	release := installer.GitHubRelease{
		Owner:   "junegunn",
		Repo:    "fzf-bin",
		Version: "0.17.4",
		Package: fmt.Sprintf("fzf-%s-%s_%s.tgz", "0.17.4", runtime.GOOS, runtime.GOARCH),
	}

	tarball, err := release.Grab()
	if err != nil {
		return err
	}

	pkg := installer.New(tarball)
	if err := pkg.Unpack(); err != nil {
		return err
	}

	err = pkg.Install(path)
	if err == nil {
		c.Path = bin
	}

	return err
}
