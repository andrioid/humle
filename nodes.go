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

type Node interface {
	Type() NodeType
	io.WriterTo
}
