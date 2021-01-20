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

// ExtractDetails -
func ExtractDetails(path string) Details {

	if len(path) == 0 {
		return nil
	}

	var (
		details Details
		detail  Detail = Detail{Name: "", Path: path}
		end     bool
	)

	for {
		detail, end = ExtractDetail(detail.Path)

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

// ExtractDetail -
func ExtractDetail(path string) (Detail, bool) {

	lastPos := strings.LastIndexByte(path, '/')

	if lastPos == 0 && len(path) > 1 {
		return (Detail{path[lastPos+1:], path}), true
	}
	if lastPos <= 0 {
		return (Detail{string('/'), string('/')}), true
	}

	n := path[lastPos+1:]

	if n == "" {
		d, _ := ExtractDetail(path[:lastPos])
		n = d.Name
	}

	return (Detail{n, path}), false
}

// Each -
func Each(d Details, f func(Detail) Detail) Details {

	var details Details
	for _, elem := range d {
		details = append(details, f(elem))
	}
	return details
}
