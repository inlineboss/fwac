package html

import "github.com/inlineboss/fwac/fs"

type FSItem struct {
	Item fs.Element
	Link string
}

func MakeFSItem(linkDir string, elem fs.Element) FSItem {
	return FSItem{elem, linkDir + "/" + elem.Name}
}
