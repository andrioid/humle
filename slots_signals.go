package humle

import "io"

// Allows delayed rendering at a named slot
type SignalNode struct {
	name string
	node NodeWriter
}

func Signal(name string, node NodeWriter) *SignalNode {
	return &SignalNode{
		name: name,
		node: node,
	}
}

func (s *SignalNode) Type() NodeType {
	return NodeSignal
}

type SlotNode struct {
	name string
}

func Slot(name string) *SlotNode {
	return &SlotNode{
		name: name,
	}
}

func (s *SlotNode) Type() NodeType {
	return NodeSlot
}

func (s *SlotNode) WriteTo(w io.Writer) (int64, error) {
	return 0, nil
}
