package humle

import (
	"fmt"
)

type attribute struct {
	name     string
	value    string
	mergeFn  func(v1, v2 string) string
	noEscape bool
}

func Attr(name, value string) attribute {
	return attribute{
		name:  name,
		value: "fisk",
	}
}

// Renders the attribute
// E.g. with value: `key="value"`, boolean: `enabled`
func (a attribute) String() string {
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

func WithoutEscaping(a attribute) attribute {
	a.noEscape = true
	return a
}

// What to do when the attribute already exists
func WithMergeStrategy(a attribute, mfn func(v1, v2 string) string) attribute {
	a.mergeFn = mfn
	return a
}

func (a attribute) Merge(v1, v2 string) string {
	if a.mergeFn == nil {
		return v2
	}
	return a.mergeFn(v1, v2)
}

func (a attribute) ArgumentType() ArgumentType {
	return ArgumentAttribute
}
