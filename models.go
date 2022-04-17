package main

import (
	"fmt"
	"io/ioutil"
	"os"
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
	path   string
	files  []File
	config map[string][]string
}

func (folder *Folder) addFile(file File) {
	folder.files = append(folder.files, file)
}

func (folder *Folder) findFiles() {
	files, err := ioutil.ReadDir(folder.path)
	folder.path, _ = filepath.Abs(folder.path)
	folderCheck(err)
	for _, info := range files {
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
	folderCheck(err)
}

func (folder *Folder) findMoveOperations() (operations []string) {
	for _, file := range folder.files {
		for extensionGroup, extensionList := range folder.config {
			if Contains(extensionList, file.extension) {
				targetPath := filepath.Join(folder.path, extensionGroup, file.extension, file.fullName())
				operations = append(
					operations,
					fmt.Sprintf("%s -> %s", file.path, targetPath),
				)
			}
		}
	}
	return
}

func (folder Folder) organizeFiles() {
	operations := folder.findMoveOperations()
	for _, operation := range operations {
		paths := strings.Split(operation, " -> ")
		origin := paths[0]
		target := paths[1]
		fmt.Printf("%s\nwill move to\n%s\n", origin, target)
		folderCheck(os.MkdirAll(filepath.Dir(target), 0755))
		fileCheck(os.Rename(origin, target))
	}
}
