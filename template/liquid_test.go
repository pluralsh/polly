package template

import (
	"testing"

	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
)

func TestLiquidFunctionDocs(t *testing.T) {
	missingDocs, redundantDocs := lo.Difference(lo.Keys(registeredFunctions), lo.Keys(functionDocs))
	assert.Empty(t, missingDocs, "found liquid functions without documentation: %s", missingDocs)
	assert.Empty(t, redundantDocs, "found documentation for functions that are not registered: %s", redundantDocs)
}
