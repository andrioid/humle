// Component allows us to more deeply control how arguments and children are handled
package humle

import (
	"errors"
	"io"
	"strings"
)

type RenderFunc func() Node

type Component struct {
	// Name is mostly for debugging and maybe future web component
	Name       string
	args       []Argument
	renderFunc RenderFunc
}

func NewComponent(name string, render RenderFunc) *Component {
	return &Component{
		Name:       name,
		renderFunc: render,
		args:       []Argument{},
	}
}

func (c *Component) WriteTo(w io.Writer) (int64, error) {
	if c.renderFunc == nil {
		return 0, errors.New("Component has no render method")
	}
	t := c.renderFunc()
	return t.WriteTo(w)
}

func (n *Component) String() string {
	var b strings.Builder
	n.WriteTo(&b)
	return b.String()
}

func (c *Component) Type() NodeType {
	return NodeComponent
}
