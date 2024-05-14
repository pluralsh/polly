package template

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	testTplTemplate = "a basic {{ .Template }}"
)

func TestReverse(t *testing.T) {
	tplFile := filepath.Join("..", "test", "_simple.tpl")
	res, err := RenderTpl(tplFile, map[string]interface{}{
		"Template": "template",
	})

	assert.NoError(t, err)
	assert.Equal(t, string(res), "a basic template")
}
