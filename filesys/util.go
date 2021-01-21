package filesys

import "io/ioutil"

func ShowDir(path string) Files {
	var files Files

	fsi, err := ioutil.ReadDir(path)

	if err != nil {
		return nil
	}

	var file File
	for _, fi := range fsi {
		file.Name = fi.Name()

		if fi.IsDir() {
			file.Type = "Folder"
		} else {
			file.Type = "File"
		}

		files = append(files, file)
	}

	return files
}
