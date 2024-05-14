package template

import (
	"bytes"
	"os"
	"path/filepath"
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

// RenderTpl renders the given Helm .tpl template with the provided bindings and automatically includes additional templates from a directory.
func RenderTpl(filePath string, bindings map[string]interface{}) ([]byte, error) {

	// Read the input template file.
	input, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	// Create a new template and add the sprig functions and the include function.
	tpl := template.New("gotpl")
	tpl.Funcs(sprig.TxtFuncMap()).Funcs(template.FuncMap{
		"include": func(name string, data interface{}) (string, error) {
			return include(name, data, tpl)
		},
	})

	// Parse the input template.
	tpl, err = tpl.Parse(string(input))
	if err != nil {
		return nil, err
	}

	// Load and parse all templates from the directory.
	templateDir := filepath.Join(filepath.Dir(filePath))
	files, err := os.ReadDir(templateDir)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		filePath := filepath.Join(templateDir, file.Name())
		content, err := os.ReadFile(filePath)
		if err != nil {
			return nil, err
		}
		_, err = tpl.New(file.Name()).Parse(string(content))
		if err != nil {
			return nil, err
		}
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
