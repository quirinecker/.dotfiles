package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-git/go-git/v5"
	"io/fs"
	"io/ioutil"
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
	binder := Binder{source, target}
	binder.bind()
	binder.postRun()
}

type Binder struct {
	source string
	target string
}

func (binder Binder) bind() {
	binder.switchToSourceDirectory()
	binder.decryptSecretFiles()
	binder.linkFiles()
}

func (binder Binder) switchToSourceDirectory() {
	err := os.Chdir(binder.source)
	if err != nil {
		println(err.Error())
	}
}

func (binder Binder) decryptSecretFiles() {
	cmd := exec.Command("git", "secret", "reveal", "-f")
	out, err := cmd.Output()

	if err != nil {
		println(err.Error())
	}

	println(string(out))
}

func (binder Binder) linkFiles() {
	err := filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
		binder.tryLink(path, info)
		return nil
	})

	if err != nil {
		println(err.Error())
	}
}

func (binder Binder) tryLink(path string, info os.FileInfo) {
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

func (binder Binder) link(path string) {
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

func (binder Binder) postRun() {
	binder.cloneResources()
}

type GitResources struct {
	Resources []GitResource `json:"resources"`
	Name      string        `json:"name"`
}

type GitResource struct {
	Path string `json:"path"`
	Url  string `json:"url"`
}

func (binder Binder) cloneResources() {
	var resources GitResources
	binder.loadResources(&resources)

	for _, resource := range resources.Resources {
		fmt.Printf("donwloading resource %s to %s", resource.Url, filepath.Join(binder.target, resource.Path))
		_, cloneError := git.PlainClone(filepath.Join(binder.target, resource.Path), false, &git.CloneOptions{
			URL:      resource.Url,
			Progress: os.Stdout,
		})

		if cloneError != nil {
			println(cloneError.Error())
		}
	}
}

func (binder Binder) loadResources(output *GitResources) {
	resourceFile, openError := os.Open("../GitResources.json")

	if openError != nil {
		println(openError.Error())
		return
	}

	resourceBytes, readError := ioutil.ReadAll(resourceFile)

	if readError != nil {
		println(readError.Error())
		return
	}

	marshalError := json.Unmarshal(resourceBytes, &output)

	if marshalError != nil {
		println(marshalError.Error())
		return
	}

}
