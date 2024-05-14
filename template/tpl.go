package template

import (
	"bytes"
	"text/template"

	"github.com/Masterminds/sprig/v3"
)

// include function to include other templates
func include(name string, data interface{}, tpl *template.Template) (string, error) {
	var buf bytes.Buffer
	err := tpl.ExecuteTemplate(&buf, name, data)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

// RenderTpl renders the given Helm .tpl template with the provided bindings.
func RenderTpl(input []byte, bindings map[string]interface{}) ([]byte, error) {
	// Create a new template and add the sprig functions and the include function.
	tpl := template.New("gotpl")
	tpl.Funcs(sprig.TxtFuncMap()).Funcs(template.FuncMap{
		"include": func(name string, data interface{}) (string, error) {
			return include(name, data, tpl)
		},
	})

	// Parse the input template.
	tpl, err := tpl.Parse(string(input))
	if err != nil {
		return nil, err
	}

	// Create a buffer to hold the rendered template.
	var buffer bytes.Buffer

	// Execute the template with the bindings and write to the buffer.
	err = tpl.Execute(&buffer, bindings)
	if err != nil {
		return nil, err
	}

	return buffer.Bytes(), nil
}
