package main

import (
	"flag"
	"os"
	"path"
)

type ExtensionConfiguration struct {
	config map[string][]string
}

func (ec *ExtensionConfiguration) addGroup(key string, values []string) {
	ec.config[key] = values
}

func NewConfig() *ExtensionConfiguration {
	values := make(map[string][]string)
	var ec ExtensionConfiguration
	ec.config = values
	return &ec
}

func getConfig() map[string][]string {
	conf := NewConfig()
	conf.addGroup("code", []string{"go", "mod"})
	conf.addGroup("audio", []string{"cda", "m3u", "mid", "midi", "mp3", "mpa", "ogg", "wav", "wma", "wpl"})
	conf.addGroup("compressed", []string{"7z", "arj", "deb", "gz", "pkg", "rar", "rpm", "tar", "tgz", "xz", "z", "zip"})
	conf.addGroup("data", []string{"csv", "dat", "db", "dbf", "json", "log", "mdb", "ods", "sav", "sql", "xlr", "xls", "xlsx", "xml"})
	conf.addGroup("disc", []string{"bin", "dmg", "iso", "toast", "vcd"})
	conf.addGroup("executable", []string{"apk", "bat", "com", "exe", "gadget", "jar", "wsf"})
	conf.addGroup("font", []string{"fnt", "fon", "otf", "ttf"})
	conf.addGroup("git", []string{"diff"})
	conf.addGroup("image", []string{"CR2", "ai", "bmp", "eps", "exr", "gif", "ico", "jfif", "jpeg", "jpg", "png", "ps", "psd", "raw", "svg", "tif", "tiff"})
	conf.addGroup("internet", []string{"asp", "aspx", "cer", "cfm", "cgi", "css", "htm", "js", "jsp", "part", "php", "pl", "rss", "sass", "scss", "xhtml"})
	conf.addGroup("presentation", []string{"key", "odp", "pps", "ppt", "pptx"})
	conf.addGroup("system", []string{"bak", "cab", "cfg", "conf", "cpl", "cur", "dll", "dmp", "drv", "icns", "ini", "lnk", "msi", "sh", "sys", "tmp"})
	conf.addGroup("text_files", []string{"doc", "docx", "md", "odt ", "pdf", "rtf", "tex", "txt", "wks ", "wpd", "wps"})
	conf.addGroup("video", []string{"3g2", "3gp", "avi", "flv", "h264", "m4v", "mkv", "mov", "mp4", "mpeg", "mpg", "rm", "swf", "vob", "wmv"})
	conf.addGroup("vim", []string{"vba", "vim"})
	return conf.config
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
	config := getConfig()
	folderPath := cliOptions()
	folder := Folder{path: folderPath, files: make([]File, 0), config: config}
	folder.findFiles()
	folder.organizeFiles()
}
