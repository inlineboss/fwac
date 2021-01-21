package proxy

import (
	"strings"
)

//Detail -
type Detail struct {
	Name string
	Link string
}

// ExtractDetail -
func MakeDetail(path string) (Detail, bool) {

	lastPos := strings.LastIndexByte(path, '/')

	if lastPos == 0 && len(path) > 1 {
		return (Detail{path[lastPos+1:], path}), false
	}
	if lastPos <= 0 {
		return (Detail{"/", "/"}), true
	}

	n := path[lastPos+1:]

	if n == "" {
		d, _ := MakeDetail(path[:lastPos])
		n = d.Name
	}

	return (Detail{n, path}), false
}

//Details -
type Details []Detail

// ExtractDetails -
func MakeDetails(host, path string) Details {

	if len(path) == 0 {
		return nil
	}

	var (
		details Details
		detail  Detail = Detail{Name: "", Link: path}
		end     bool
	)

	for {
		detail, end = MakeDetail(detail.Link)

		details = append(details, detail)

		if detail.Link[len(detail.Link)-1] == '/' {
			detail.Link = strings.TrimRight(detail.Link, detail.Name+"/")
		} else {
			detail.Link = strings.TrimRight(detail.Link, detail.Name)
		}

		if end {
			break
		}
	}

	details = EachDetail(details, func(d Detail) Detail {
		d.Link = host + d.Link
		return d
	})

	return details
}

// Each -
func EachDetail(d Details, f func(Detail) Detail) Details {

	var details Details
	for _, elem := range d {
		details = append(details, f(elem))
	}
	return details
}
