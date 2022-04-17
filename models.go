package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var (
	fileCheck      = prefixedCheck("File")
	folderCheck    = prefixedCheck("Folder")
	operationCheck = prefixedCheck("Operation")
)

type Operation struct {
	file   *File
	target string
}

func (operation Operation) preview() {
	fmt.Printf(
		"Original Path:\n%s\nTarget Path\n%s\n",
		operation.file.path,
		operation.target,
	)
}

func (operation Operation) commit() {
	origin := operation.file.path
	target := operation.target
	folderCheck(os.MkdirAll(filepath.Dir(target), 0755))
	operationCheck(os.Rename(origin, target))
}

type File struct {
	name      string
	extension string
	path      string
}

func (file File) fullName() string {
	return fmt.Sprintf("%s.%s", file.name, file.extension)
}

func (file *File) getValidName(targetPath string) (validPath string) {
	counter := 1

	if _, err := os.Stat(targetPath); errors.Is(err, os.ErrNotExist) {
		return filepath.Join(targetPath, file.fullName())
	}
	for {
		newName := fmt.Sprintf("%s-%d", file.name, counter)
		validPath = filepath.Join(
			targetPath,
			fmt.Sprintf("%s.%s", newName, file.extension),
		)
		if _, err := os.Stat(validPath); err == nil {
			counter++
		} else {
			file.name = newName
			return
		}
	}
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

func (folder *Folder) findMoveOperations() (operations []Operation) {
	for _, file := range folder.files {
		var targetPath string
		for extensionGroup, extensionList := range folder.config {
			if Contains(extensionList, file.extension) {
				targetPath = filepath.Join(folder.path, extensionGroup, file.extension, file.fullName())
				break
			}
		}
		if targetPath == "" {
			targetPath = filepath.Join(folder.path, "uncategorized", file.extension, file.fullName())
		}

		targetPath = file.getValidName(filepath.Dir(targetPath))
		operations = append(
			operations,
			Operation{&file, targetPath},
		)
	}
	return
}

func (folder Folder) organizeFiles() {
	operations := folder.findMoveOperations()
	for _, operation := range operations {
		operation.preview()
		operation.commit()
	}
}
