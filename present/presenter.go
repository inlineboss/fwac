package present

import (
	"github.com/inlineboss/fwac/filesys"
	"github.com/inlineboss/fwac/proxy"
)

type Presenter struct {
	Proxy    proxy.Proxy
	FileInfo []FileInfo
}

func MakePresenter(prx proxy.Proxy) Presenter {
	var self Presenter
	self.Proxy = prx

	elements := filesys.ShowDir(prx.FS.LongPath)

	for _, e := range elements {
		self.FileInfo = append(self.FileInfo, MakeFSItem(prx.URL.Link, e))
	}

	return self
}
