package algorithms_test

import (
	"testing"

	"github.com/pluralsh/polly/algorithms"
	"github.com/stretchr/testify/assert"
)

func TestSortMap(t *testing.T) {
	input := map[string]interface{}{
		"z": "last",
		"a": "first",
		"m": map[string]interface{}{
			"c": 3,
			"a": 1,
			"b": 2,
		},
		"list": []interface{}{
			map[string]interface{}{"z": 1, "a": 2},
			"plain string",
			123,
		},
	}

	expected := map[string]interface{}{
		"a": "first",
		"list": []interface{}{
			map[string]interface{}{"a": 2, "z": 1},
			"plain string",
			123,
		},
		"m": map[string]interface{}{
			"a": 1,
			"b": 2,
			"c": 3,
		},
		"z": "last",
	}

	sorted := algorithms.SortMap(input)
	assert.Equal(t, expected, sorted)
}
