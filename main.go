package main

import (
	"flag"
	"os"
	"path"
)

func cliOptions() (folderPath string) {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	flag.StringVar(&folderPath, "path", path.Join(home, "Downloads"), "Path to organize")
	flag.Parse()
	return
}

func main() {
	config := getConfig()
	folderPath := cliOptions()
	folder := Folder{path: folderPath, files: make([]File, 0), config: config}
	folder.findFiles()
	folder.organizeFiles()
}
