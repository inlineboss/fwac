package url

import "testing"

func TestExtractLast(t *testing.T) {

	{
		detail, _ := ExtractLast("/reg/rules/", '/')

		if detail.Name != "rules" { //&& path != "/reg" && !end {
			t.Errorf("ExtractLastDir(\"/reg/rules/\") (n = %s, p = %s): name != \"rules\" && path != \"/reg\"", detail.Name, detail.Name)
		}
	}
}

func TestExtractPath(t *testing.T) {

	{
		tPath := "/reg/rules/"
		detail, _ := ExtractPath(tPath, '/')

		if detail.Name != "rules" || detail.Path != tPath {
			t.Errorf("ExtractPath(\"%s\")", tPath)
		}
	}

	{
		tPath := "/rules/"
		detail, _ := ExtractPath(tPath, '/')

		if detail.Name != "rules" || detail.Path != tPath {
			t.Errorf("ExtractPath(\"%s\")", tPath)
		}
	}

	{
		tPath := "/"
		detail, _ := ExtractPath(tPath, '/')

		if detail.Name != "/" || detail.Path != tPath {
			t.Errorf("ExtractPath(\"%s\")", tPath)
		}
	}
}

func TestExtractPaths(t *testing.T) {

	{
		tPath := "/reg/rules/"
		details := ExtractPaths(tPath, '/')

		if details[0].Name != "rules" || details[0].Path != tPath { //&& path != "/reg" && !end {
			t.Errorf("ExtractPaths(\"%s\"): [0]", tPath)
		}

		if details[1].Name != "reg" || details[1].Path != "/" { //&& path != "/reg" && !end {
			t.Errorf("ExtractPaths(\"%s\"): [0]", tPath)
		}
	}

	{
		tPath := "/"
		details := ExtractPaths(tPath, '/')

		if details != nil {
			t.Errorf("ExtractPath(\"%s\")", tPath)
		}
	}
}
