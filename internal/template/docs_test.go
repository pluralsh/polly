package main

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLiquidFunctionDocs(t *testing.T) {
	b := new(bytes.Buffer)
	err := generateFilterDocs(b, registeredFilters())
	assert.NoError(t, err)

	f, err := os.ReadFile("../../docs/liquid-filters.md")
	assert.NoError(t, err)
	assert.Equal(t, b.String(), string(f), "filter docs are not up-to-date")
}
