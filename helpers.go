package humle

func Iterate[T any](ts []T, cb func(int, T) Node) Group {
	nodes := make([]Node, 0, len(ts))
	for i, t := range ts {
		nodes = append(nodes, cb(i, t))
	}
	return nodes
}

func If(condition bool, n Node) Node {
	if condition {
		return n
	}
	return nil
}

func Iff(condition bool, f func() Node) Node {
	if condition {
		return f()
	}
	return nil
}

func Map[T any](ts []T, cb func(T) Node) Group {
	nodes := make([]Node, 0, len(ts))
	for _, t := range ts {
		nodes = append(nodes, cb(t))
	}
	return nodes
}
