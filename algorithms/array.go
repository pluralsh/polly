package algorithms

func Reverse[T any](arr []T) []T {
	length := len(arr)
	res := make([]T, length)

	for ind, val := range arr {
		res[length-ind-1] = val
	}

	return res
}

func Map[T any, V any](arr []T, f func(T) V) []V {
	res := make([]V, len(arr))
	for ind, val := range arr {
		res[ind] = f(val)
	}

	return res
}

func Filter[T any](arr []T, f func(T) bool) []T {
	res := make([]T, 0)
	for _, v := range arr {
		if f(v) {
			res = append(res, v)
		}
	}

	return res
}
