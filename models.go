package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
)

var (
	fileCheck   = prefixedCheck("File")
	folderCheck = prefixedCheck("Folder")
)

type File struct {
	name      string
	extension string
	path      string
}

func (file File) fullName() string {
	return fmt.Sprintf("%s.%s", file.name, file.extension)
}

func (file *File) rename() {

}

type Folder struct {
	path  string
	files []File
}

func (folder *Folder) addFile(file File) {
	folder.files = append(folder.files, file)
}

func (folder *Folder) findFiles() {
	files, err := ioutil.ReadDir(folder.path)
	folder.path, _ = filepath.Abs(folder.path)
	if err != nil {
		log.Fatal(err)
	}
	for _, info := range files {
		res, _ := folderCheck(err)
		if res {
			if info.IsDir() == false {
				filePtr := new(File)
				fullPath := fmt.Sprintf("%s/%s", folder.path, info.Name())
				filePtr.path = fullPath
				nameInfo := strings.Split(info.Name(), ".")
				filePtr.name = nameInfo[0]
				if len(nameInfo) > 1 {
					filePtr.extension = nameInfo[1]
				}
				folder.addFile(*filePtr)
			}
		}
	}
	folderCheck(err)
}

func (folder *Folder) organize(config map[string][]string, persist bool) (operations []string) {
	for _, file := range folder.files {
		for extensionGroup, extensionList := range config {
			if Contains(extensionList, file.extension) {
				targetPath := filepath.Join(folder.path, extensionGroup, file.extension, file.fullName())
				operations = append(
					operations,
					fmt.Sprintf("%s -> %s\n", file.path, targetPath),
				)
			}
		}
	}
	return
}
