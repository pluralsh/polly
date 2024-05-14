package template

import (
	"bytes"
	"os"
	"path/filepath"
	"text/template"

	"github.com/Masterminds/sprig/v3"
)

// include function to include other templates
func include(name string, data interface{}, tpl *template.Template, baseDir string) (string, error) {
	var buf bytes.Buffer

	// Construct the file path for the included template
	includePath := filepath.Join(baseDir, name)

	// Read and parse the included template
	includeContent, err := os.ReadFile(includePath)
	if err != nil {
		return "", err
	}

	// Parse the included template content
	includeTpl, err := tpl.New(name).Parse(string(includeContent))
	if err != nil {
		return "", err
	}

	// Execute the included template
	err = includeTpl.ExecuteTemplate(&buf, name, data)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

// RenderTpl renders the given Helm .tpl template with the provided bindings and automatically includes additional templates from a directory.
func RenderTpl(filePath string, bindings map[string]interface{}) ([]byte, error) {
	// Read the input template file.
	input, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	// Create a new template and add the sprig functions and the include function.
	tpl := template.New("gotpl")
	baseDir := filepath.Dir(filePath)
	tpl.Funcs(sprig.TxtFuncMap()).Funcs(template.FuncMap{
		"include": func(name string, data interface{}) (string, error) {
			return include(name, data, tpl, baseDir)
		},
	})

	// Parse the input template.
	tpl, err = tpl.Parse(string(input))
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
