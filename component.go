// Component allows us to more deeply control how arguments and children are handled
package humle

import (
	"io"
)

type RenderFunc func(args []Argument, children Node) *Tag

type componentModel struct {
	// Name is mostly for debugging and maybe future web component
	Name       string
	args       []Argument
	childNode  Node
	renderFunc RenderFunc
}

func NewComponent(name string, render RenderFunc) *componentModel {
	return &componentModel{
		Name:       name,
		renderFunc: render,
		args:       []Argument{},
	}
}

func (c *componentModel) WriteTo(w io.Writer) (int64, error) {
	t := c.renderFunc(c.args, c.childNode)
	return t.WriteTo(w)
}
