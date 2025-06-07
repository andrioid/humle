// Component allows us to more deeply control how arguments and children are handled
package component

import (
	"io"

	h "github.com/andrioid/humle"
)

type renderFunc func(args []h.Argument, children h.Node) *h.Tag

type componentModel struct {
	// Name is mostly for debugging and maybe future web component
	Name       string
	args       []h.Argument
	childNode  h.Node
	renderFunc renderFunc
}

func NewComponent(name string, render renderFunc) *componentModel {
	return &componentModel{
		Name:       name,
		renderFunc: render,
		args:       []h.Argument{},
	}
}

func (c *componentModel) WriteTo(w io.Writer) (int64, error) {
	t := c.renderFunc(c.args, c.childNode)
	return t.WriteTo(w)
}
