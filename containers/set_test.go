package containers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLooksLikeASet(t *testing.T) {
	s := NewSet[int]()

	assert.False(t, s.Has(1), false)
	s.Add(1)
	assert.True(t, s.Has(1))
	s.Add(1)
	s.Add(2)
	assert.True(t, s.Has(2))
	s.Remove(1)
	assert.False(t, s.Has(1))
	assert.Equal(t, s.List(), []int{2})
}

func TestToSet(t *testing.T) {
	vals := []int{1, 2, 5}
	s := ToSet(vals)

	for _, v := range vals {
		assert.True(t, s.Has(v))
	}

	assert.False(t, s.Has(3))
}

func TestEqual(t *testing.T) {
	vals := []int{1, 2, 3}
	s := ToSet(vals)

	assert.True(t, s.Equal(ToSet(vals)))
	assert.False(t, s.Equal(ToSet([]int{1, 2})))
	assert.False(t, s.Equal(ToSet([]int{1, 2, 3, 4})))
	assert.False(t, s.Equal(ToSet([]int{1, 2, 5})))
}
