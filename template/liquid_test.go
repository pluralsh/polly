package template_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pluralsh/polly/template"
)

func TestAppendFunctionNotOverridden(t *testing.T) {
	// Test data
	input := `{% assign fruits = "apple" | append: ", banana" %}{{ fruits }}`
	expected := "apple, banana"

	// Render the template
	result, err := template.RenderLiquid([]byte(input), map[string]interface{}{})

	// Assert that rendering was successful and produced the expected output
	assert.NoError(t, err)
	assert.Equal(t, expected, string(result))

	// Additionally, verify that the function wasn't registered from Sprig
	// by checking if it's in the excluded functions list
	filters := template.RegisteredFilters()
	_, exists := filters["append"]
	assert.True(t, !exists, "append function should not be registered as it is excluded from Sprig functions")
}
