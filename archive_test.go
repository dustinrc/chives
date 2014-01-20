package chives

import (
	. "launchpad.net/gocheck"
	"os"
	"path/filepath"
)

var dirsAndFiles = []struct {
	name    string
	content string
}{
	{"file1.txt", "blah 1"},
	{"a/filea1.txt", "blah a1"},
	{"b/fileb1.txt", "blah b1"},
	{"a/c/fileac1.txt", "blah ac1"},
	{"d", ""},
}

func createDirsAndFiles(tempDir string) (root string) {
	// tempDir must already exist
	root = filepath.Join(tempDir, "r")
	os.Mkdir(root, 0777)

	for _, o := range dirsAndFiles {
		oPath := filepath.Join(root, o.name)

		if o.content == "" {
			os.MkdirAll(oPath, 0777)
		} else {
			dir, _ := filepath.Split(oPath)
			os.MkdirAll(dir, 0777)
			f, _ := os.Create(oPath)
			defer f.Close()
			f.WriteString(o.content)
		}

	}

	return
}

func (s *S) TestDirTarUntar(c *C) {
	root := createDirsAndFiles(c.MkDir())

	dt := NewDirTar(root, "myfile.tar")
	err := dt.Create()
	c.Assert(err, IsNil)
}
