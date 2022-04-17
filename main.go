package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"os"
	"path"
)

func readExtensionsConfig() (result map[string][]string) {
	jsonFile, err := os.Open("extension_configuration.json")
	check := prefixedCheck("Read Extensions")
	res, _ := check(err)
	if res {
		defer jsonFile.Close()
		byteValue, _ := ioutil.ReadAll(jsonFile)
		json.Unmarshal([]byte(byteValue), &result)
	} else {
		panic(err)
	}
	return
}

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
	config := readExtensionsConfig()
	folderPath := cliOptions()
	folder := Folder{path: folderPath, files: make([]File, 0), config: config}
	folder.findFiles()
	folder.organizeFiles()
}
