package proxy

import "testing"

func TestMakeDetail(t *testing.T) {

	{
		tPath := "/reg/rules/"
		detail, _ := MakeDetail(tPath)

		if detail.Name != "rules" || detail.Link != tPath {
			t.Errorf("MakeDetail(\"%s\")", tPath)
		}
	}

	{
		tPath := "/rules/"
		detail, _ := MakeDetail(tPath)

		if detail.Name != "rules" || detail.Link != tPath {
			t.Errorf("MakeDetail(\"%s\")", tPath)
		}
	}

	{
		tPath := "/"
		detail, _ := MakeDetail(tPath)

		if detail.Name != "/" || detail.Link != tPath {
			t.Errorf("MakeDetail(\"%s\")", tPath)
		}
	}
}

func TestMakeDetails(t *testing.T) {

	{
		tPath := "/reg/rules/"
		details := MakeDetails("", tPath)

		if details[0].Name != "rules" || details[0].Link != tPath {
			t.Errorf("MakeDetails(\"%s\"): [0]", tPath)
		}

		if details[1].Name != "reg" || details[1].Link != "/" {
			t.Errorf("MakeDetails(\"%s\"): [0]", tPath)
		}
	}

	{
		tPath := "/"
		details := MakeDetails("", tPath)

		if details != nil {
			t.Errorf("ExtractPath(\"%s\")", tPath)
		}
	}
}
