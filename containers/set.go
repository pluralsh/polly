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

func (s Set[V]) Difference(other Set[V]) Set[V] {
	new := NewSet[V]()
	for v, _ := range s {
		if !other.Has(v) {
			new.Add(v)
		}
	}
	return new
}

func (s Set[V]) Union(other Set[V]) Set[V] {
	new := ToSet(s.List())
	for v, _ := range other {
		new.Add(v)
	}
	return new
}

func Union[V comparable](sets ...Set[V]) Set[V] {
	res := NewSet[V]()

	// use nested loops for a bit of extra efficiency
	for _, s := range sets {
		for v, _ := range s {
			res.Add(v)
		}
	}

	return res
}

func (s Set[V]) Intersect(other Set[V]) Set[V] {
	res := NewSet[V]()
	for v, _ := range s {
		if other.Has(v) {
			res.Add(v)
		}
	}

	return res
}

func Intersect[V comparable](sets ...Set[V]) Set[V] {
	res := NewSet[V]()
	if len(sets) == 0 {
		return res
	}
	first, rest := sets[0], sets[1:]
	for v, _ := range first {
		if lo.EveryBy(rest, func(s Set[V]) bool { return s.Has(v) }) {
			res.Add(v)
		}
	}

	return res
}

func (s Set[V]) SymmetricDifference(other Set[V]) Set[V] {
	un := s.Union(other)
	in := s.Intersect(other)
	return un.Difference(in)
}
