package humle

import (
	"fmt"
)

type Attribute struct {
	name     string
	value    string
	mergeFn  func(v1, v2 string) string
	noEscape bool
}

func NewAttribute(name, value string) *Attribute {
	return &Attribute{
		name:  name,
		value: value,
	}
}

func (a *Attribute) Type() NodeType {
	return NodeAttribute
}

// Renders the attribute
// E.g. with value: `key="value"`, boolean: `enabled`
func (a *Attribute) String() string {
	if a.value == "" {
		return a.name
	}

	if a.noEscape {
		return fmt.Sprintf(`%s="%s"`, a.name, a.value)
	}
	// TODO: Add escaping for HTML entities if needed
	return fmt.Sprintf(`%s="%s"`, a.name, a.value)
	//return fmt.Sprintf(`%s="%s"`, a.name, template.JSEscapeString(a.value))
}

func WithoutEscaping(a *Attribute) *Attribute {
	a.noEscape = true

	return a
}

// What to do when the attribute already exists
func WithMergeStrategy(a *Attribute, mfn func(v1, v2 string) string) *Attribute {
	a.mergeFn = mfn
	return a
}

func (a *Attribute) Merge(v1, v2 string) string {
	if a.mergeFn == nil {
		return v2
	}
	return a.mergeFn(v1, v2)
}
