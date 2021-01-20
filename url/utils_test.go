package url

import "testing"

func TestExtractDetail(t *testing.T) {

	{
		tPath := "/reg/rules/"
		detail, _ := ExtractDetail(tPath)

		if detail.Name != "rules" || detail.Path != tPath {
			t.Errorf("ExtractDetail(\"%s\")", tPath)
		}
	}

	{
		tPath := "/rules/"
		detail, _ := ExtractDetail(tPath)

		if detail.Name != "rules" || detail.Path != tPath {
			t.Errorf("ExtractDetail(\"%s\")", tPath)
		}
	}

	{
		tPath := "/"
		detail, _ := ExtractDetail(tPath)

		if detail.Name != "/" || detail.Path != tPath {
			t.Errorf("ExtractDetail(\"%s\")", tPath)
		}
	}
}

func TestExtractDetails(t *testing.T) {

	{
		tPath := "/reg/rules/"
		details := ExtractDetails(tPath)

		if details[0].Name != "rules" || details[0].Path != tPath {
			t.Errorf("ExtractDetails(\"%s\"): [0]", tPath)
		}

		if details[1].Name != "reg" || details[1].Path != "/" {
			t.Errorf("ExtractDetails(\"%s\"): [0]", tPath)
		}
	}

	{
		tPath := "/"
		details := ExtractDetails(tPath)

		if details != nil {
			t.Errorf("ExtractPath(\"%s\")", tPath)
		}
	}
}
