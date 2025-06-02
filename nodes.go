package humle

import "io"

type NodeType int

const (
	NodeTag = iota
	NodeTextValue
	NodeGroup
	NodeRawHTML
	NodeComponent
)

type ChildNode interface {
	Type() NodeType
}

// A node capable of writing HTML. Can be used as children
// E.g. Element, Text
type NodeWriter interface {
	ChildNode
	io.WriterTo
}
