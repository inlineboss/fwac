package present

import "github.com/inlineboss/fwac/filesys"

type FileInfo struct {
	File filesys.File
	Link string
}

func ExtractInfo(linkDir string, f filesys.File) FileInfo {
	return FileInfo{f, linkDir + "/" + f.Name}
}
