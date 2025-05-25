package humle

import (
	"slices"
	"strings"
)

type Attributes map[string]*Attribute

// Sets or merges (according to merge strategy) an attribute
func (a Attributes) Set(k string, newAttribute *Attribute) {

	// If attribute already exists, we need to merge it
	if existing, ok := (a)[k]; ok {
		// Using the overriding attribute's merge function we merge with the existing value
		// - This has the advantage of merging with classes that don't use merge
		newValue := newAttribute.Merge(existing.value, newAttribute.value)
		existing.value = newValue
		(a)[k] = existing
		return
	}
	(a)[k] = newAttribute
}

func (a Attributes) String() string {
	var attributes = make([]Attribute, len(a))
	for _, attr := range a {
		attributes = append(attributes, *attr)
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
