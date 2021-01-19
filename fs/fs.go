package fs

import "io/ioutil"

type TypeElement int

const (
	Folder TypeElement = iota
	File
)

type Element struct {
	Name string
	Type string
}

func ShowDir(path string) []Element {
	var items []Element

	lst, err := ioutil.ReadDir(path)
	if err != nil {
		return nil
	}
	var elem Element
	for _, val := range lst {
		elem.Name = val.Name()

		if val.IsDir() {
			elem.Type = "Folder"
		} else {
			elem.Type = "File"
		}

		items = append(items, elem)
	}
	return items
}
