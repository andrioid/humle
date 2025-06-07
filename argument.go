package humle

type ArgumentType int

const (
	ArgumentAttribute = iota
	ArgumentProp
	ArgumentText
)

// Arguments are either props or attributes and are used by components
type Argument interface {
	ArgumentType() ArgumentType
}

type Arguments []Argument

func (args Arguments) GetAttributes() Attributes {
	attributes := Attributes{}
	for _, arg := range args {
		switch t := arg.(type) {
		case attribute:
			attributes[t.name] = t
		}
	}
	return attributes
}

// Special cases, like Text() which can be more ergonomic as an argument than child
func (args Arguments) GetNodes() []Node {
	nodes := []Node{}
	for _, arg := range args {
		switch t := arg.(type) {
		case Node:
			nodes = append(nodes, t)
		}
	}
	return nodes
}
