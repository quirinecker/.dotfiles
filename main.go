package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func main() {
	linker := &FileLinker{}
	linker.linkAll()
}

type FileLinker struct {
	root string
	home string
}

func (linker FileLinker) linkAll() {
	linker.init()
	linker.changeToHome()
	err := filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
		linker.tryLink(path, info)
		return nil
	})

	if err != nil {
		println(err)
	}
}

func (linker FileLinker) changeToHome() {
	err := os.Chdir("home")
	if err != nil {
		println(err.Error())
	}
}

func (linker FileLinker) link(path string) {
	source := filepath.Join("./", path)
	target := filepath.Join(linker.home, path)
	fmt.Printf("linking %s -> %s  \n", target, source)

	deleteError := os.Remove(target)
	linkError := os.Link(source, target)
	if linkError != nil {
		println(linkError.Error())
	}

	if deleteError != nil {
		println(deleteError.Error())
	}

}

func (linker FileLinker) init() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		println(err.Error)
	} else {
		linker.home = homeDir

		fmt.Printf("local: %s | global: %s", homeDir, linker.home)
	}
}

func (linker FileLinker) tryLink(path string, info os.FileInfo) {
	if !info.IsDir() {
		println("linking")
		linker.link(path)
	} else {
		err := os.Mkdir(filepath.Join(linker.home, path), os.ModePerm)
		if err != nil && !os.IsExist(err) {
			println(err.Error())
		} else if err == nil {
			linker.tryLink(path, info)
		}
	}
}
