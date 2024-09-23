package main

import (
	"os"
	"sort"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig/v3"
	_ "github.com/pluralsh/polly/template"
	tmpl "github.com/pluralsh/polly/template"
	"github.com/samber/lo"
)

const (
	templ = `...

{{ range . }}
##  {{ .Name }}
{{ .Description }}

{{ if .Aliases }}Aliases: {{ .Aliases | join ", " }}{{ end }}
Implementation: {{ .Implementation }}
{{ end }}`
)

func main() {
	filters := lo.Values(tmpl.RegisteredFilters())
	sort.Slice(filters, func(i, j int) bool {
		return strings.Compare(filters[i].Name, filters[j].Name) < 0
	})

	t := template.Must(template.New("").Funcs(sprig.TxtFuncMap()).Parse(templ))
	err := t.Execute(os.Stdout, filters)
	if err != nil {
		panic(err)
	}
}
