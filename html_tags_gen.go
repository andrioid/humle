// Generated from humle/html package
// **DO NOT EDIT**
package humle

// Elements

func HTML(attrs ...Attribute) *Tag {
	tag := NewTag("html", attrs...)
	return tag
}

func Head(attrs ...Attribute) *Tag {
	tag := NewTag("head", attrs...)
	return tag
}

func Title(attrs ...Attribute) *Tag {
	tag := NewTag("title", attrs...)
	return tag
}

func Meta(attrs ...Attribute) *Tag {
	tag := NewTag("meta", attrs...)
	return tag
}

func Link(attrs ...Attribute) *Tag {
	tag := NewTag("link", attrs...)
	return tag
}

func Style(attrs ...Attribute) *Tag {
	tag := NewTag("style", attrs...)
	return tag
}

func Script(attrs ...Attribute) *Tag {
	tag := NewTag("script", attrs...)
	return tag
}

func Noscript(attrs ...Attribute) *Tag {
	tag := NewTag("noscript", attrs...)
	return tag
}

func Template(attrs ...Attribute) *Tag {
	tag := NewTag("template", attrs...)
	return tag
}

func Article(attrs ...Attribute) *Tag {
	tag := NewTag("article", attrs...)
	return tag
}

func Aside(attrs ...Attribute) *Tag {
	tag := NewTag("aside", attrs...)
	return tag
}

func Footer(attrs ...Attribute) *Tag {
	tag := NewTag("footer", attrs...)
	return tag
}

func Header(attrs ...Attribute) *Tag {
	tag := NewTag("header", attrs...)
	return tag
}

func Nav(attrs ...Attribute) *Tag {
	tag := NewTag("nav", attrs...)
	return tag
}

func Figure(attrs ...Attribute) *Tag {
	tag := NewTag("figure", attrs...)
	return tag
}

func Figcaption(attrs ...Attribute) *Tag {
	tag := NewTag("figcaption", attrs...)
	return tag
}

func Address(attrs ...Attribute) *Tag {
	tag := NewTag("address", attrs...)
	return tag
}

func Dl(attrs ...Attribute) *Tag {
	tag := NewTag("dl", attrs...)
	return tag
}

func Dt(attrs ...Attribute) *Tag {
	tag := NewTag("dt", attrs...)
	return tag
}

func Dd(attrs ...Attribute) *Tag {
	tag := NewTag("dd", attrs...)
	return tag
}

func Ol(attrs ...Attribute) *Tag {
	tag := NewTag("ol", attrs...)
	return tag
}

func Ul(attrs ...Attribute) *Tag {
	tag := NewTag("ul", attrs...)
	return tag
}

func Li(attrs ...Attribute) *Tag {
	tag := NewTag("li", attrs...)
	return tag
}

func Table(attrs ...Attribute) *Tag {
	tag := NewTag("table", attrs...)
	return tag
}

func Caption(attrs ...Attribute) *Tag {
	tag := NewTag("caption", attrs...)
	return tag
}

func Thead(attrs ...Attribute) *Tag {
	tag := NewTag("thead", attrs...)
	return tag
}

func Tbody(attrs ...Attribute) *Tag {
	tag := NewTag("tbody", attrs...)
	return tag
}

func Tfoot(attrs ...Attribute) *Tag {
	tag := NewTag("tfoot", attrs...)
	return tag
}

func Tr(attrs ...Attribute) *Tag {
	tag := NewTag("tr", attrs...)
	return tag
}

func Th(attrs ...Attribute) *Tag {
	tag := NewTag("th", attrs...)
	return tag
}

func Td(attrs ...Attribute) *Tag {
	tag := NewTag("td", attrs...)
	return tag
}

func Col(attrs ...Attribute) *Tag {
	tag := NewTag("col", attrs...)
	return tag
}

func Colgroup(attrs ...Attribute) *Tag {
	tag := NewTag("colgroup", attrs...)
	return tag
}

func Form(attrs ...Attribute) *Tag {
	tag := NewTag("form", attrs...)
	return tag
}

func Fieldset(attrs ...Attribute) *Tag {
	tag := NewTag("fieldset", attrs...)
	return tag
}

func Legend(attrs ...Attribute) *Tag {
	tag := NewTag("legend", attrs...)
	return tag
}

func Label(attrs ...Attribute) *Tag {
	tag := NewTag("label", attrs...)
	return tag
}

func Input(attrs ...Attribute) *Tag {
	tag := NewTag("input", attrs...)
	return tag
}

func Select(attrs ...Attribute) *Tag {
	tag := NewTag("select", attrs...)
	return tag
}

func Option(attrs ...Attribute) *Tag {
	tag := NewTag("option", attrs...)
	return tag
}

func Optgroup(attrs ...Attribute) *Tag {
	tag := NewTag("optgroup", attrs...)
	return tag
}

func Div(attrs ...Attribute) *Tag {
	tag := NewTag("div", attrs...)
	return tag
}

func Section(attrs ...Attribute) *Tag {
	tag := NewTag("section", attrs...)
	return tag
}

func Main(attrs ...Attribute) *Tag {
	tag := NewTag("main", attrs...)
	return tag
}

func Button(attrs ...Attribute) *Tag {
	tag := NewTag("button", attrs...)
	return tag
}

func A(attrs ...Attribute) *Tag {
	tag := NewTag("a", attrs...)
	return tag
}

func H1(attrs ...Attribute) *Tag {
	tag := NewTag("h1", attrs...)
	return tag
}

func H2(attrs ...Attribute) *Tag {
	tag := NewTag("h2", attrs...)
	return tag
}

func H3(attrs ...Attribute) *Tag {
	tag := NewTag("h3", attrs...)
	return tag
}

func P(attrs ...Attribute) *Tag {
	tag := NewTag("p", attrs...)
	return tag
}

func Span(attrs ...Attribute) *Tag {
	tag := NewTag("span", attrs...)
	return tag
}

func Field(attrs ...Attribute) *Tag {
	tag := NewTag("field", attrs...)
	return tag
}

func Pre(attrs ...Attribute) *Tag {
	tag := NewTag("pre", attrs...)
	return tag
}

func Body(attrs ...Attribute) *Tag {
	tag := NewTag("body", attrs...)
	return tag
}

func Svg(attrs ...Attribute) *Tag {
	tag := NewTag("svg", attrs...)
	return tag
}

// Void Elements
func Br(attrs ...Attribute) *Tag {
	tag := NewTag("br", attrs...)
    tag = WithVoidElement(tag)
	return tag
}

func Hr(attrs ...Attribute) *Tag {
	tag := NewTag("hr", attrs...)
    tag = WithVoidElement(tag)
	return tag
}

