package template

import (
	"bytes"
	"text/template"

	"github.com/Masterminds/sprig/v3"
)

func RenderTpl(input []byte, bindings map[string]interface{}) ([]byte, error) {
	tpl, err := template.New("gotpl").Funcs(sprig.TxtFuncMap()).Parse(string(input))
	if err != nil {
		return nil, err
	}

	var buffer bytes.Buffer
	err = tpl.Execute(&buffer, bindings)
	return buffer.Bytes(), err
}
