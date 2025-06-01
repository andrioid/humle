package humle

import "io"

// Components are for creating re-usable composable elements
// When rendered they flatten components and render tags w. attributes
type Component struct {
	name string
	// Props are merged here and Attributes are passed down to all direct tags
	args     []Argument
	children ChildNode
}

type Props map[string]string

func NewComponent(name string, children ChildNode) *Component {
	return &Component{
		name:     name,
		args:     []Argument{},
		children: children,
	}
}

func (c *Component) Type() NodeType {
	return NodeComponent
}

func (c *Component) WriteTo(w io.Writer) (int64, error) {
	return 0, nil
}
