package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func main() {
	linker := FileLinker{}
	linker.linkAll()
}

type FileLinker struct {
	root string
}

func (linker FileLinker) linkAll() {
	linker.changeToHome()
	err := filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
		if !info.IsDir() {
			linker.link(path)
		}
		return nil
	})

	if err != nil {
		println(err)
	}
}

func (linker FileLinker) changeToHome() {
	err := os.Chdir("home")
	if err != nil {
		println("could not change directory to home")
	}
}

func (linker FileLinker) link(path string) {
	source := filepath.Join("./", path)
	target := filepath.Join("/home/quirinecker", path)
	fmt.Printf("%s -> %s  \n", target, source)

	deleteError := os.Remove(target)
	linkError := os.Link(source, target)
	if linkError != nil {
		println(linkError.Error())
	}

	if deleteError != nil {
		println(deleteError.Error())
	}

}
