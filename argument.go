package humle

type ArgumentType int

const (
	ArgumentAttribute = iota
	ArgumentProp
)

// Arguments are either props or attributes and are used by components
type Argument interface {
	ArgumentType() ArgumentType
}

type Arguments []Argument

func (args Arguments) AddAttributes(attrs ...Attribute) Arguments {
	return args
}

func (args Arguments) GetAttributes() []Attribute {
	attributes := []Attribute{}
	for _, arg := range args {
		switch t := arg.(type) {
		case Attribute:
			attributes = append(attributes, t)
		}
	}
	return attributes
}
