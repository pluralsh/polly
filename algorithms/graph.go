package algorithms

import "github.com/pluralsh/polly/containers"

type Prober[T any] interface {
	Push(val T)
	Pop() (T, error)
	Empty() bool
}

func ProbeGraph[T comparable](prober Prober[T], initial T, neighbors func(T) ([]T, error), visit func(T) error) error {
	seen := map[T]bool{}
	prober.Push(initial)

	for !prober.Empty() {
		r, err := prober.Pop()
		if err != nil {
			return err
		}

		if _, ok := seen[r]; ok {
			continue
		}

		seen[r] = true
		if err := visit(r); err != nil {
			return err
		}

		nebs, err := neighbors(r)
		if err != nil {
			return err
		}

		for _, neb := range nebs {
			if _, ok := seen[neb]; !ok {
				prober.Push(neb)
			}
		}
	}

	return nil
}

func DFS[T comparable](initial T, neighbors func(T) ([]T, error), visit func(T) error) error {
	s := containers.NewStack[T]()
	wrapped := func(n T) (ns []T, err error) {
		ns, err = neighbors(n)
		if err != nil {
			return
		}
		ns = Reverse(ns)
		return
	}

	return ProbeGraph[T](s, initial, wrapped, visit)
}

func BFS[T comparable](initial T, neighbors func(T) ([]T, error), visit func(T) error) error {
	q := containers.NewQueue[T]()
	return ProbeGraph[T](q, initial, neighbors, visit)
}

func TopSort[T comparable](vals []T, neighbors func(T) ([]T, error)) ([]T, error) {
	marked := map[T]bool{}
	res := []T{}
	wrapped := func(n T) ([]T, error) {
		ns, err := neighbors(n)
		if err != nil {
			return ns, err
		}

		ns = Filter[T](ns, func(n T) bool {
			_, ok := marked[n]
			return !ok
		})
		return ns, nil
	}

	for _, val := range vals {
		// ignore if already probed (we would have dfs'ed the subgraph already)
		if _, ok := marked[val]; ok {
			continue
		}
		sublist := []T{}
		visit := func(n T) error {
			marked[n] = true
			sublist = append(sublist, n)
			return nil
		}

		// launch a dfs of the unprobed section of the graph
		if err := DFS(val, wrapped, visit); err != nil {
			return res, err
		}

		// since no found items were reached after prior dfs, we know it's safe to prepend (but not necessary append)
		res = append(sublist, res...)
	}

	return res, nil
}

func TopsortGraph[T comparable](g containers.Graph[T]) ([]T, error) {
	return TopSort[T](g.Nodes(), func(n T) ([]T, error) {
		return g.Neighbors(n), nil
	})
}
