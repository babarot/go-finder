package installer

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

// Package represents the archive package
type Package struct {
	Archive    string
	Binary     string
	Permission int64
}

// New create the package
func New(archive string) Package {
	return Package{
		Archive: archive,
	}
}

// Unpack unpacks the archive package
func (p *Package) Unpack() error {
	file, err := os.Open(p.Archive)
	if err != nil {
		return err
	}
	defer file.Close()

	var fileReader io.ReadCloser = file
	if fileReader, err = gzip.NewReader(file); err != nil {
		return err
	}
	defer fileReader.Close()

	defer os.Remove(p.Archive)

	tarBallReader := tar.NewReader(fileReader)
	for {
		header, err := tarBallReader.Next()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		// get the individual filename and extract to the current directory
		filename := header.Name

		p.Binary = header.Name
		p.Permission = header.Mode

		switch header.Typeflag {
		case tar.TypeDir:
			err = os.MkdirAll(filename, os.FileMode(header.Mode))
			if err != nil {
				return err
			}
		case tar.TypeReg:
			writer, err := os.Create(filename)
			if err != nil {
				return err
			}
			io.Copy(writer, tarBallReader)
			err = os.Chmod(filename, os.FileMode(header.Mode))
			if err != nil {
				return err
			}
			writer.Close()
		default:
			return fmt.Errorf("Unable to untar type: %c in file %s", header.Typeflag, filename)
		}
	}
	return nil
}

// Install installs the file to PATH
func (p *Package) Install(path string) error {
	if path == "" {
		return fmt.Errorf("%s: invalid path", path)
	}
	file := filepath.Join(path, p.Binary)
	err := copy(p.Binary, file)
	if err != nil {
		return err
	}
	defer os.Remove(p.Binary)
	return os.Chmod(file, os.FileMode(p.Permission))
}

func copy(srcName, dstName string) error {
	src, err := os.Open(srcName)
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	return err
}

// GitHubRelease represents the GitHub Releases
type GitHubRelease struct {
	Owner   string
	Repo    string
	Version string
	Package string
}

// Grab grabs the release binary from GitHub Releases
func (r GitHubRelease) Grab() (string, error) {
	resp, err := http.Get(fmt.Sprintf("https://github.com/%s/%s/releases/download/%s/%s",
		r.Owner, r.Repo, r.Version, r.Package,
	))
	if err != nil {
		return r.Package, err
	}
	defer resp.Body.Close()

	file, err := os.Create(r.Package)
	if err != nil {
		return r.Package, err
	}
	defer file.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return r.Package, err
	}
	file.Write(body)

	return r.Package, nil
}
