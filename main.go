package main

import (
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	bind("./home", home())
}

func home() string {
	home, err := os.UserHomeDir()

	if err != nil {
		println(err.Error())
		return ""
	} else {
		return home
	}
}

func bind(source string, target string) {
	binder := binder{source, target}
	binder.bind()
}

type binder struct {
	source string
	target string
}

func (binder binder) bind() {
	binder.switchToSourceDirectory()
	binder.decryptSecretFiles()
	binder.linkFiles()
}

func (binder binder) switchToSourceDirectory() {
	err := os.Chdir(binder.source)
	if err != nil {
		println(err.Error())
	}
}

func (binder binder) decryptSecretFiles() {
	cmd := exec.Command("git", "secret", "reveal", "-f")
	out, err := cmd.Output()

	if err != nil {
		println(err)
	}

	println(string(out))
}

func (binder binder) linkFiles() {
	err := filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
		binder.tryLink(path, info)
		return nil
	})

	if err != nil {
		println(err.Error())
	}
}

func (binder binder) tryLink(path string, info os.FileInfo) {
	if !info.IsDir() {
		binder.link(path)
	} else {
		err := os.Mkdir(filepath.Join(binder.target, path), os.ModePerm)
		if err != nil && !os.IsExist(err) {
			println(err.Error())
		} else if err == nil {
			binder.tryLink(path, info)
		}
	}
}

func (binder binder) link(path string) {
	source := filepath.Join("./", path)
	target := filepath.Join(binder.target, path)
	fmt.Printf("linking %s -> %s  \n", source, target)

	deleteError := os.Remove(target)
	linkError := os.Link(source, target)
	if linkError != nil {
		println(linkError.Error())
	}

	if deleteError != nil && !os.IsNotExist(deleteError) {
		println(deleteError.Error())
	}
}
