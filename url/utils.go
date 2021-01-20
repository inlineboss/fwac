package url

import (
	"strings"
)

func lst(str string) byte {
	if len(str) == 0 {
		return 0
	}

	return str[len(str)-1]
}

func ExtractLasts(path string, sep byte) Details {

	if len(path) == 0 {
		return nil
	}

	var (
		details Details
		detail  Detail = Detail{Name: "", Path: path}
		end     bool
	)

	for {
		detail, end = ExtractLast(detail.Path, sep)

		if end {
			break
		}

		details = append(details, detail)
	}

	return details
}

func ExtractLast(path string, sep byte) (Detail, bool) {

	if len(path) == 0 {
		return (Detail{}), true
	}

	if path == string(sep) {
		return (Detail{string(sep), string(sep)}), true
	}

	lastPos := strings.LastIndexByte(path, sep)

	if lastPos <= 0 {
		return (Detail{string(sep), string(sep)}), true
	}

	n := path[lastPos+1:]
	p := path[0:lastPos]

	return (Detail{n, p}), false
}

func ExtractPaths(path string, sep byte) Details {

	if len(path) == 0 {
		return nil
	}

	var (
		details Details
		detail  Detail = Detail{Name: "", Path: path}
		end     bool
	)

	for {
		detail, end = ExtractPath(detail.Path, sep)

		if end {
			break
		}

		details = append(details, detail)

		if lst(detail.Path) == '/' {
			detail.Path = strings.TrimRight(detail.Path, detail.Name+"/")
		} else {
			detail.Path = strings.TrimRight(detail.Path, detail.Name)
		}

	}

	return details
}

func ExtractPath(path string, sep byte) (Detail, bool) {

	lastPos := strings.LastIndexByte(path, sep)

	if lastPos == 0 && len(path) > 1 {
		return (Detail{path[lastPos+1:], path}), true
	}
	if lastPos <= 0 {
		return (Detail{string(sep), string(sep)}), true
	}

	n := path[lastPos+1:]

	if n == "" {
		d, _ := ExtractPath(path[:lastPos], sep)
		n = d.Name
	}

	return (Detail{n, path}), false
}
