package html

import (
	"github.com/inlineboss/fwac/fs"
	"github.com/inlineboss/fwac/proxy"
)

type WEBPresenter struct {
	Proxy   proxy.Proxy
	FSItems []FSItem
}

func MakeWEBPresenter(prx proxy.Proxy) WEBPresenter {
	var self WEBPresenter
	self.Proxy = prx

	elements := fs.ShowDir(prx.FS.LongPath)

	for _, e := range elements {
		self.FSItems = append(self.FSItems, MakeFSItem(prx.URL.Link, e))
	}

	return self
}
