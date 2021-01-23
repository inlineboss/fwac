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
func MakeDetail(path string) Detail {

	lastPos := strings.LastIndexByte(path, '/')

	if lastPos == 0 && len(path) > 1 {
		return (Detail{path[lastPos+1:], path})
	}
	if lastPos <= 0 {
		return (Detail{"/", "/"})
	}

	n := path[lastPos+1:]

	if n == "" {
		n = MakeDetail(path[:lastPos]).Name
	}

	return (Detail{n, path})
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
	)

	for {
		detail = MakeDetail(detail.Link)

		details = append(details, detail)

		if len(detail.Link) == 1 {
			break
		}

		if detail.Link[len(detail.Link)-1] == '/' {
			detail.Link = detail.Link[:len(detail.Link)-len(detail.Name+"/")]
		} else {
			detail.Link = detail.Link[:len(detail.Link)-len(detail.Name)]
		}
	}

	details = EachDetail(details, func(d Detail) Detail {
		d.Link = host + d.Link
		return d
	})

	reversed := Details{}

	// reverse order
	// and append into new slice
	for i := range details {
		n := details[len(details)-1-i]
		//fmt.Println(n) -- sanity check
		reversed = append(reversed, n)
	}

	return reversed
}

// Each -
func EachDetail(d Details, f func(Detail) Detail) Details {

	var details Details
	for _, elem := range d {
		details = append(details, f(elem))
	}
	return details
}
