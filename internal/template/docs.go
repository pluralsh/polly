package main

import (
	"io"
	"os"
	"sort"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig/v3"
	tmpl "github.com/pluralsh/polly/template"
	"github.com/samber/lo"
)

const (
	docsPath     = "docs/liquid-filters.md"
	docsTemplate = `# Supported Liquid Filters
{{ range . }}
##  {{ .Name }}
{{ .Description }}

{{ if .Aliases }}_Aliases_: {{ .Aliases | join ", " }}{{ end }}

_Implementation_: {{ .Implementation }}
{{ end }}`
)

func main() {
	f, err := os.Create(docsPath)
	if err != nil {
		panic(err)
	}

	if err = generateFilterDocs(f, registeredFilters()); err != nil {
		panic(err)
	}
}

func registeredFilters() []tmpl.FilterFunction {
	filters := lo.Values(tmpl.RegisteredFilters())
	sort.Slice(filters, func(i, j int) bool {
		return strings.Compare(filters[i].Name, filters[j].Name) < 0
	})

	return filters
}

func generateFilterDocs(writer io.Writer, filters []tmpl.FilterFunction) error {
	t := template.Must(template.New("").Funcs(sprig.TxtFuncMap()).Parse(docsTemplate))
	return t.Execute(writer, filters)
}
