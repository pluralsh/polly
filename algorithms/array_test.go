package algorithms

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	reversed := Reverse([]int{1, 2, 3})

	assert.Equal(t, reversed, []int{3, 2, 1})
}

func TestMap(t *testing.T) {
	assert.Equal(t, Map([]int{1, 2, 3}, func(v int) int { return v * v }), []int{1, 4, 9})
}

func TestFilter(t *testing.T) {
	assert.Equal(t, Filter([]int{1, 2, 3}, func(v int) bool { return v == 2 }), []int{2})
}
