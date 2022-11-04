package containers

import (
	"github.com/samber/lo"
)

type Set[V comparable] map[V]bool

func NewSet[V comparable]() Set[V] {
	return map[V]bool{}
}

func ToSet[V comparable](vals []V) Set[V] {
	s := NewSet[V]()
	for _, v := range vals {
		s[v] = true
	}

	return s
}

func (s Set[V]) Add(v V) {
	s[v] = true
}

func (s Set[V]) Has(v V) bool {
	return s[v]
}

func (s Set[V]) Remove(v V) {
	delete(s, v)
}

func (s Set[V]) List() []V {
	return lo.Keys(s)
}

func (s Set[V]) Len() int {
	return len(s)
}

func (s Set[V]) Equal(other Set[V]) bool {
	if s.Len() != other.Len() {
		return false
	}

	for k, _ := range s {
		if !other.Has(k) {
			return false
		}
	}

	return true
}
