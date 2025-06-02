package humle

import (
	"io"
)

type Group []ChildNode

func (g Group) Type() NodeType {
	return NodeGroup
}

func (g Group) WriteTo(w io.Writer) (int64, error) {
	var total int64
	for _, node := range g {
		if group, ok := node.(Group); ok {
			wb, err := group.WriteTo(w)
			total += wb
			if err != nil {
				return total, err
			}
		}
		if writable, ok := node.(NodeWriter); ok {
			wb, err := writable.WriteTo(w)
			total += wb
			if err != nil {
				return total, err
			}
		}
	}

	return total, nil

}
