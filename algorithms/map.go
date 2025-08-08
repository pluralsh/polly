package algorithms

import (
	"sort"

	"github.com/samber/lo"
)

func MapValues[K comparable, V any](m map[K]V) []V {
	s := make([]V, 0, len(m))
	for _, v := range m {
		s = append(s, v)
	}
	return s
}

func MapKeys[K comparable, V any](m map[K]V) []K {
	s := make([]K, 0, len(m))
	for k := range m {
		s = append(s, k)
	}
	return s
}

func Merge(maps ...map[string]interface{}) map[string]interface{} {
	res := maps[0]
	for _, m := range maps[1:] {
		res = deepMerge(res, m)
	}

	return res
}

func deepMerge(m1, m2 map[string]interface{}) map[string]interface{} {
	// lifted from helm's merge code
	out := make(map[string]interface{}, len(m1))
	for k, v := range m1 {
		out[k] = v
	}

	for k, v := range m2 {
		if v, ok := v.(map[string]interface{}); ok {
			if bv, ok := out[k]; ok {
				if bv, ok := bv.(map[string]interface{}); ok {
					out[k] = deepMerge(bv, v)
					continue
				}
			}
		}
		out[k] = v
	}
	return out
}

func SortMap(input map[string]interface{}) map[string]interface{} {
	sorted := make(map[string]interface{})
	keys := lo.Keys(input)
	sort.Strings(keys)

	for _, k := range keys {
		v := input[k]
		switch t := v.(type) {
		case map[string]interface{}:
			sorted[k] = SortMapRecursive(t)
		case []interface{}:
			sortedSlice := make([]interface{}, len(t))
			for i, item := range t {
				if m, ok := item.(map[string]interface{}); ok {
					sortedSlice[i] = SortMapRecursive(m)
				} else {
					sortedSlice[i] = item
				}
			}
			sorted[k] = sortedSlice
		default:
			sorted[k] = v
		}
	}
	return sorted
}
