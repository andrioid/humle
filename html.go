package humle

import (
	"fmt"
	"strings"

	"github.com/Oudwins/tailwind-merge-go/pkg/twmerge"
)

func Data(k, v string) attribute {
	return Attr("data-"+k, v)
}

// Use this if you don't like default Class behavior of merging classes
func AttrClass(v ...string) attribute {
	values := strings.Join(v, " ")
	return WithMergeStrategy(
		Attr("class", values),
		func(v1, v2 string) string {
			if v1 == "" {
				return v2
			}
			if v2 == "" {
				return v1
			}
			return fmt.Sprintf("%s %s", v1, v2)
		},
	)

}

// Class attribute with tailwind-merge functionality
func Class(v ...string) attribute {
	values := strings.Join(v, " ")
	a := Attr("class", twmerge.Merge(values))
	WithMergeStrategy(a, func(s1, s2 string) string {
		return twmerge.Merge(s1, s2)
	})
	return a
}

func DocType() Node {
	return RawHTML("<!DOCTYPE html>")
}

// Prints DocType and treats any nodes as siblings
func Document(nodes ...Node) Group {
	g := Group{
		DocType(),
	}

	for _, node := range nodes {
		g = append(g, node)
	}

	return g
}
