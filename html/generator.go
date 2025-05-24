package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"text/template"
)

// List of HTML elements
// - Note: There are no plans to generate all of HTML. Only what we need.
// - Use `NewElement()`or `NewAttribute()“ if you're missing something or send a PR
// Reference: https://developer.mozilla.org/en-US/docs/Glossary/Empty_element
var input = TemplateInput{
	elements: []string{
		"div", "section", "main",
		"button", "a",
		"h1", "h2", "h3", "p", "span",
		"input", "label", "field",
		"pre", "script", "link", "head", "body", "meta", "svg",
	},
	voidElements: []string{"br", "hr"},
	attributes: []string{
		"id", "href", "alt", "placeholder", "src", "rel", "name",
		"content", "charset", "lang", "type", "value", "class",
	},
}

var funcs = template.FuncMap{
	"CasingFunc":      FuncName,
	"CasingElement":   strings.ToLower,
	"CasingAttribute": strings.ToLower,
}

type TemplateInput struct {
	elements     []string
	voidElements []string
	attributes   []string
}

func FuncName(str string) string {
	str = strings.ToLower(str)
	switch str {
	case "id":
		return "ID"
	default:
		letters := strings.Split(str, "")
		return strings.ToUpper(letters[0]) + strings.Join(letters[1:], "")
	}
}

func main() {
	templates := template.New("html")
	templates.Funcs(funcs)
	template.Must(templates.ParseGlob("html/*.go.tmpl"))

	// Create output file
	elF, err := os.Create("html_elements_gen.go")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating file: %v\n", err)
		os.Exit(1)
	}
	defer elF.Close()

	atF, err := os.Create("html_attributes_gen.go")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating file: %v\n", err)
		os.Exit(1)
	}
	defer atF.Close()

	// Write headers
	if err := templates.ExecuteTemplate(elF, "header.go.tmpl", nil); err != nil {
		log.Fatal(err)
	}
	if err := templates.ExecuteTemplate(atF, "header.go.tmpl", nil); err != nil {
		log.Fatal(err)
	}

	// Process Elements
	fmt.Fprintf(elF, "\n// Elements\n")
	for _, elem := range input.elements {
		if err := templates.ExecuteTemplate(elF, "element.go.tmpl", elem); err != nil {
			log.Fatal(err)
		}
	}

	fmt.Fprintf(elF, "\n// Void Elements\n")
	for _, elem := range input.voidElements {
		if err := templates.ExecuteTemplate(elF, "element_void.go.tmpl", elem); err != nil {
			log.Fatal(err)
		}
	}

	// Process Attributes
	fmt.Fprintf(atF, "\n// Attributes\n")
	for _, a := range input.attributes {
		if err := templates.ExecuteTemplate(atF, "attribute.go.tmpl", a); err != nil {
			log.Fatal(err)
		}
	}

}
