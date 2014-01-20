package chives

import (
	"archive/tar"
	"io"
	"os"
	"path/filepath"
)

// DirTar associates a directory and tarfile.
type DirTar struct {
	DirPath   string
	TarPath   string
	tarWriter *tar.Writer
}

// NewDirTar creates a new DirTar instance.
func NewDirTar(dirPath, tarPath string) *DirTar {
	return &DirTar{dirPath, tarPath, nil}
}

// addToTar satisfies the WalkFunc type.
func (dt *DirTar) addToTar(path string, info os.FileInfo, e error) (err error) {
	fr, err := os.Open(path)
	if err != nil {
		return
	}
	defer fr.Close()

	hdr, err := tar.FileInfoHeader(info, path)
	if err != nil {
		return
	}
	// info.Name() only gives the basename of path, make sure to get the correct path
	hdr.Name = path

	if err = dt.tarWriter.WriteHeader(hdr); err != nil {
		return
	}

	if info.Mode().IsDir() {
		return
	} else if _, err = io.Copy(dt.tarWriter, fr); err != nil {
		return
	}

	return
}

// Create creates the tarfile by walking through the directory structure.
func (dt *DirTar) Create() (err error) {
	tarFile, err := os.Create(dt.TarPath)
	if err != nil {
		return
	}
	defer tarFile.Close()

	dt.tarWriter = tar.NewWriter(tarFile)

	if err = filepath.Walk(dt.DirPath, dt.addToTar); err != nil {
		return
	}
	if err = dt.tarWriter.Close(); err != nil {
		return
	}

	return
}
