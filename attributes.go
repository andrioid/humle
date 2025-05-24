package humle

import (
	"slices"
	"strings"
)

type Attributes map[string]Attribute

// Sets or merges (according to merge strategy) an attribute
func (a *Attributes) Set(k string, v Attribute) error {

	if val, ok := (*a)[k]; ok {
		val.value = val.Merge(v.value)
		return nil
	}
	(*a)[k] = v

	return nil
}

func (a *Attributes) String() string {
	var attributes = make([]Attribute, len(*a))
	for _, attr := range *a {
		attributes = append(attributes, attr)
	}
	slices.SortFunc(attributes, func(a, b Attribute) int {
		if a.name == "id" {
			return -1
		}
		if b.name == "id" {
			return 1
		}
		if n := strings.Compare(a.name, b.name); n != 0 {
			return n
		}
		return 0
	})
	attrStrings := []string{}
	for _, as := range attributes {
		str := as.String()
		if str != "" {
			attrStrings = append(attrStrings, str)
		}

	}

	return strings.Join(attrStrings, " ")
}
