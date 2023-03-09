package template

import (
	"os"
	"testing"
	"text/template"
)

func TestApacheLicense2(t *testing.T) {
	tpl := template.New("license")
	tpl.Parse(ApacheLicense2())
	tpl.Execute(os.Stdout, []Copyright{{Year: "2023", Owner: "Kami"}, {Year: "2020", Owner: "Leo"}})
}
