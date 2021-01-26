package proxy

import (
	"testing"
)

func TestMakeDetails(t *testing.T) {
	host := "http://localhost:8080"
	folders := []string{"/", "work", "topcon", "copies", "source", "AntList"}
	path := "/" + folders[1] + "/" + folders[2] + "/" + folders[3] + "/" + folders[4] + "/" + folders[5]
	dtls := MakeDetails(host, path)
	for i, d := range dtls {
		if folders[i] != d.Name {
			t.Errorf("MakeDetails %s: %s != %s", host+path, folders[i], d.Name)
		}
	}
}
