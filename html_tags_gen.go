// Generated from humle/html package
// **DO NOT EDIT**
package humle

// Elements

func HTML(args ...Argument) *Tag {
	tag := NewTag("html", args...)
	return tag
}

func Head(args ...Argument) *Tag {
	tag := NewTag("head", args...)
	return tag
}

func Title(args ...Argument) *Tag {
	tag := NewTag("title", args...)
	return tag
}

func Link(args ...Argument) *Tag {
	tag := NewTag("link", args...)
	return tag
}

func Style(args ...Argument) *Tag {
	tag := NewTag("style", args...)
	return tag
}

func Script(args ...Argument) *Tag {
	tag := NewTag("script", args...)
	return tag
}

func Noscript(args ...Argument) *Tag {
	tag := NewTag("noscript", args...)
	return tag
}

func Template(args ...Argument) *Tag {
	tag := NewTag("template", args...)
	return tag
}

func Article(args ...Argument) *Tag {
	tag := NewTag("article", args...)
	return tag
}

func Aside(args ...Argument) *Tag {
	tag := NewTag("aside", args...)
	return tag
}

func Footer(args ...Argument) *Tag {
	tag := NewTag("footer", args...)
	return tag
}

func Header(args ...Argument) *Tag {
	tag := NewTag("header", args...)
	return tag
}

func Nav(args ...Argument) *Tag {
	tag := NewTag("nav", args...)
	return tag
}

func Figure(args ...Argument) *Tag {
	tag := NewTag("figure", args...)
	return tag
}

func Figcaption(args ...Argument) *Tag {
	tag := NewTag("figcaption", args...)
	return tag
}

func Address(args ...Argument) *Tag {
	tag := NewTag("address", args...)
	return tag
}

func Dl(args ...Argument) *Tag {
	tag := NewTag("dl", args...)
	return tag
}

func Dt(args ...Argument) *Tag {
	tag := NewTag("dt", args...)
	return tag
}

func Dd(args ...Argument) *Tag {
	tag := NewTag("dd", args...)
	return tag
}

func Ol(args ...Argument) *Tag {
	tag := NewTag("ol", args...)
	return tag
}

func Ul(args ...Argument) *Tag {
	tag := NewTag("ul", args...)
	return tag
}

func Li(args ...Argument) *Tag {
	tag := NewTag("li", args...)
	return tag
}

func Table(args ...Argument) *Tag {
	tag := NewTag("table", args...)
	return tag
}

func Caption(args ...Argument) *Tag {
	tag := NewTag("caption", args...)
	return tag
}

func Thead(args ...Argument) *Tag {
	tag := NewTag("thead", args...)
	return tag
}

func Tbody(args ...Argument) *Tag {
	tag := NewTag("tbody", args...)
	return tag
}

func Tfoot(args ...Argument) *Tag {
	tag := NewTag("tfoot", args...)
	return tag
}

func Tr(args ...Argument) *Tag {
	tag := NewTag("tr", args...)
	return tag
}

func Th(args ...Argument) *Tag {
	tag := NewTag("th", args...)
	return tag
}

func Td(args ...Argument) *Tag {
	tag := NewTag("td", args...)
	return tag
}

func Col(args ...Argument) *Tag {
	tag := NewTag("col", args...)
	return tag
}

func Colgroup(args ...Argument) *Tag {
	tag := NewTag("colgroup", args...)
	return tag
}

func Form(args ...Argument) *Tag {
	tag := NewTag("form", args...)
	return tag
}

func Fieldset(args ...Argument) *Tag {
	tag := NewTag("fieldset", args...)
	return tag
}

func Legend(args ...Argument) *Tag {
	tag := NewTag("legend", args...)
	return tag
}

func Label(args ...Argument) *Tag {
	tag := NewTag("label", args...)
	return tag
}

func Input(args ...Argument) *Tag {
	tag := NewTag("input", args...)
	return tag
}

func Select(args ...Argument) *Tag {
	tag := NewTag("select", args...)
	return tag
}

func Option(args ...Argument) *Tag {
	tag := NewTag("option", args...)
	return tag
}

func Optgroup(args ...Argument) *Tag {
	tag := NewTag("optgroup", args...)
	return tag
}

func Div(args ...Argument) *Tag {
	tag := NewTag("div", args...)
	return tag
}

func Section(args ...Argument) *Tag {
	tag := NewTag("section", args...)
	return tag
}

func Main(args ...Argument) *Tag {
	tag := NewTag("main", args...)
	return tag
}

func Button(args ...Argument) *Tag {
	tag := NewTag("button", args...)
	return tag
}

func A(args ...Argument) *Tag {
	tag := NewTag("a", args...)
	return tag
}

func H1(args ...Argument) *Tag {
	tag := NewTag("h1", args...)
	return tag
}

func H2(args ...Argument) *Tag {
	tag := NewTag("h2", args...)
	return tag
}

func H3(args ...Argument) *Tag {
	tag := NewTag("h3", args...)
	return tag
}

func P(args ...Argument) *Tag {
	tag := NewTag("p", args...)
	return tag
}

func Span(args ...Argument) *Tag {
	tag := NewTag("span", args...)
	return tag
}

func Field(args ...Argument) *Tag {
	tag := NewTag("field", args...)
	return tag
}

func Pre(args ...Argument) *Tag {
	tag := NewTag("pre", args...)
	return tag
}

func Body(args ...Argument) *Tag {
	tag := NewTag("body", args...)
	return tag
}

func Svg(args ...Argument) *Tag {
	tag := NewTag("svg", args...)
	return tag
}

// Void Elements
func Br(args ...Argument) *Tag {
	tag := NewTag("br", args...)
    tag = WithVoidElement(tag)
	return tag
}

func Hr(args ...Argument) *Tag {
	tag := NewTag("hr", args...)
    tag = WithVoidElement(tag)
	return tag
}

func Meta(args ...Argument) *Tag {
	tag := NewTag("meta", args...)
    tag = WithVoidElement(tag)
	return tag
}

