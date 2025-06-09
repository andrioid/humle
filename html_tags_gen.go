// Generated from humle/html package
// **DO NOT EDIT**
package humle

// Elements

func HTML(args ...Argument) *TagNode {
	tag := Tag("html", args...)
	return tag
}

func Head(args ...Argument) *TagNode {
	tag := Tag("head", args...)
	return tag
}

func Title(args ...Argument) *TagNode {
	tag := Tag("title", args...)
	return tag
}

func Link(args ...Argument) *TagNode {
	tag := Tag("link", args...)
	return tag
}

func Style(args ...Argument) *TagNode {
	tag := Tag("style", args...)
	return tag
}

func Script(args ...Argument) *TagNode {
	tag := Tag("script", args...)
	return tag
}

func Noscript(args ...Argument) *TagNode {
	tag := Tag("noscript", args...)
	return tag
}

func Template(args ...Argument) *TagNode {
	tag := Tag("template", args...)
	return tag
}

func Article(args ...Argument) *TagNode {
	tag := Tag("article", args...)
	return tag
}

func Aside(args ...Argument) *TagNode {
	tag := Tag("aside", args...)
	return tag
}

func Footer(args ...Argument) *TagNode {
	tag := Tag("footer", args...)
	return tag
}

func Header(args ...Argument) *TagNode {
	tag := Tag("header", args...)
	return tag
}

func Nav(args ...Argument) *TagNode {
	tag := Tag("nav", args...)
	return tag
}

func Figure(args ...Argument) *TagNode {
	tag := Tag("figure", args...)
	return tag
}

func Figcaption(args ...Argument) *TagNode {
	tag := Tag("figcaption", args...)
	return tag
}

func Address(args ...Argument) *TagNode {
	tag := Tag("address", args...)
	return tag
}

func Dl(args ...Argument) *TagNode {
	tag := Tag("dl", args...)
	return tag
}

func Dt(args ...Argument) *TagNode {
	tag := Tag("dt", args...)
	return tag
}

func Dd(args ...Argument) *TagNode {
	tag := Tag("dd", args...)
	return tag
}

func Ol(args ...Argument) *TagNode {
	tag := Tag("ol", args...)
	return tag
}

func Ul(args ...Argument) *TagNode {
	tag := Tag("ul", args...)
	return tag
}

func Li(args ...Argument) *TagNode {
	tag := Tag("li", args...)
	return tag
}

func Table(args ...Argument) *TagNode {
	tag := Tag("table", args...)
	return tag
}

func Caption(args ...Argument) *TagNode {
	tag := Tag("caption", args...)
	return tag
}

func Thead(args ...Argument) *TagNode {
	tag := Tag("thead", args...)
	return tag
}

func Tbody(args ...Argument) *TagNode {
	tag := Tag("tbody", args...)
	return tag
}

func Tfoot(args ...Argument) *TagNode {
	tag := Tag("tfoot", args...)
	return tag
}

func Tr(args ...Argument) *TagNode {
	tag := Tag("tr", args...)
	return tag
}

func Th(args ...Argument) *TagNode {
	tag := Tag("th", args...)
	return tag
}

func Td(args ...Argument) *TagNode {
	tag := Tag("td", args...)
	return tag
}

func Col(args ...Argument) *TagNode {
	tag := Tag("col", args...)
	return tag
}

func Colgroup(args ...Argument) *TagNode {
	tag := Tag("colgroup", args...)
	return tag
}

func Form(args ...Argument) *TagNode {
	tag := Tag("form", args...)
	return tag
}

func Fieldset(args ...Argument) *TagNode {
	tag := Tag("fieldset", args...)
	return tag
}

func Legend(args ...Argument) *TagNode {
	tag := Tag("legend", args...)
	return tag
}

func Label(args ...Argument) *TagNode {
	tag := Tag("label", args...)
	return tag
}

func Input(args ...Argument) *TagNode {
	tag := Tag("input", args...)
	return tag
}

func Select(args ...Argument) *TagNode {
	tag := Tag("select", args...)
	return tag
}

func Option(args ...Argument) *TagNode {
	tag := Tag("option", args...)
	return tag
}

func Optgroup(args ...Argument) *TagNode {
	tag := Tag("optgroup", args...)
	return tag
}

func Div(args ...Argument) *TagNode {
	tag := Tag("div", args...)
	return tag
}

func Section(args ...Argument) *TagNode {
	tag := Tag("section", args...)
	return tag
}

func Main(args ...Argument) *TagNode {
	tag := Tag("main", args...)
	return tag
}

func Button(args ...Argument) *TagNode {
	tag := Tag("button", args...)
	return tag
}

func A(args ...Argument) *TagNode {
	tag := Tag("a", args...)
	return tag
}

func H1(args ...Argument) *TagNode {
	tag := Tag("h1", args...)
	return tag
}

func H2(args ...Argument) *TagNode {
	tag := Tag("h2", args...)
	return tag
}

func H3(args ...Argument) *TagNode {
	tag := Tag("h3", args...)
	return tag
}

func P(args ...Argument) *TagNode {
	tag := Tag("p", args...)
	return tag
}

func Span(args ...Argument) *TagNode {
	tag := Tag("span", args...)
	return tag
}

func Field(args ...Argument) *TagNode {
	tag := Tag("field", args...)
	return tag
}

func Pre(args ...Argument) *TagNode {
	tag := Tag("pre", args...)
	return tag
}

func Body(args ...Argument) *TagNode {
	tag := Tag("body", args...)
	return tag
}

func Svg(args ...Argument) *TagNode {
	tag := Tag("svg", args...)
	return tag
}

// Void Elements
func Br(args ...Argument) *TagNode {
	tag := Tag("br", args...)
    tag = WithVoidElement(tag)
	return tag
}

func Hr(args ...Argument) *TagNode {
	tag := Tag("hr", args...)
    tag = WithVoidElement(tag)
	return tag
}

func Meta(args ...Argument) *TagNode {
	tag := Tag("meta", args...)
    tag = WithVoidElement(tag)
	return tag
}

