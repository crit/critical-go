package folder

import (
	"io/ioutil"
	"os"
	"strings"
)

type Folder struct {
	Path string
}

func New(path string) Folder {
	return Folder{Path: path}
}

func (f Folder) Put(fileName string, data []byte) error {
	return ioutil.WriteFile(location(f.Path, fileName), data, 0644)
}

func (f Folder) Get(fileName string) []byte {
	data, _ := ioutil.ReadFile(location(f.Path, fileName))
	return data
}

func (f Folder) Delete(fileName string) {
	os.Remove(location(f.Path, fileName))
}

func (f Folder) Flush() {
	os.RemoveAll(f.Path)
}

func location(path, fileName string) string {
	return strings.TrimRight(path, "/") + "/" + strings.TrimLeft(fileName, "/")
}
