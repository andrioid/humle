package humle

import (
	"io"
	"strings"
)

type Group []Node

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
		wb, err := node.WriteTo(w)
		total += wb
		if err != nil {
			return total, err
		}
	}

	return total, nil

}

func (n Group) String() string {
	var b strings.Builder
	n.WriteTo(&b)
	return b.String()
}
