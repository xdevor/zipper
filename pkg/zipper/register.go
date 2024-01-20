package zipper

import (
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"
)

type ZipOps struct {
	Name    string
	Operate func()
}

var AllOperations []ZipOps

func AddOperations(operation ZipOps) {
	AllOperations = append(AllOperations, operation)
}

type osFS struct{}

func (osFS) Open(name string) (fs.File, error) { return os.Open(filepath.FromSlash(name)) }

func (osFS) ReadDir(name string) ([]fs.DirEntry, error) { return os.ReadDir(filepath.FromSlash(name)) }

func (osFS) Stat(name string) (fs.FileInfo, error) { return os.Stat(filepath.FromSlash(name)) }

func (osFS) ReadFile(name string) ([]byte, error) { return os.ReadFile(filepath.FromSlash(name)) }

func (osFS) Glob(pattern string) ([]string, error) { return filepath.Glob(filepath.FromSlash(pattern)) }

func Execute() {
	goFiles, err := fs.Glob(osFS{}, path.Join("onetimeops", "*.go"))
	if err != nil {
		return
	}

	for _, file := range goFiles {
		fmt.Println(file)
	}

	for _, operation := range AllOperations {
		operation.Operate()
	}
}
