package proxy

import (
	"net/http"
)

type Proxy struct {
	URL struct {
		Host     string
		Link     string
		BackLink string
		Road     Details
	}

	FS struct {
		Root      string
		ShortPath string
		LongPath  string
	}
}

func MakeProxy(r *http.Request, dir string) Proxy {
	var prx Proxy
	prx.URL.Host = "http://" + r.Host
	prx.URL.Link = prx.URL.Host + r.URL.Path
	prx.URL.Road = append(Details{}, MakeDetails(prx.URL.Host, r.URL.Path)...)

	prx.FS.Root = dir
	prx.FS.ShortPath = r.URL.Path
	prx.FS.LongPath = dir + r.URL.Path

	return prx
}
