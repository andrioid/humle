// Generated from humle/html package
// **DO NOT EDIT**
package humle

// Elements
func HTML(nodes ...Node) *Tag {
	return NewTag("html", nodes...)
}
func Head(nodes ...Node) *Tag {
	return NewTag("head", nodes...)
}
func Title(nodes ...Node) *Tag {
	return NewTag("title", nodes...)
}
func Meta(nodes ...Node) *Tag {
	return NewTag("meta", nodes...)
}
func Link(nodes ...Node) *Tag {
	return NewTag("link", nodes...)
}
func Style(nodes ...Node) *Tag {
	return NewTag("style", nodes...)
}
func Script(nodes ...Node) *Tag {
	return NewTag("script", nodes...)
}
func Noscript(nodes ...Node) *Tag {
	return NewTag("noscript", nodes...)
}
func Template(nodes ...Node) *Tag {
	return NewTag("template", nodes...)
}
func Article(nodes ...Node) *Tag {
	return NewTag("article", nodes...)
}
func Aside(nodes ...Node) *Tag {
	return NewTag("aside", nodes...)
}
func Footer(nodes ...Node) *Tag {
	return NewTag("footer", nodes...)
}
func Header(nodes ...Node) *Tag {
	return NewTag("header", nodes...)
}
func Nav(nodes ...Node) *Tag {
	return NewTag("nav", nodes...)
}
func Figure(nodes ...Node) *Tag {
	return NewTag("figure", nodes...)
}
func Figcaption(nodes ...Node) *Tag {
	return NewTag("figcaption", nodes...)
}
func Address(nodes ...Node) *Tag {
	return NewTag("address", nodes...)
}
func Dl(nodes ...Node) *Tag {
	return NewTag("dl", nodes...)
}
func Dt(nodes ...Node) *Tag {
	return NewTag("dt", nodes...)
}
func Dd(nodes ...Node) *Tag {
	return NewTag("dd", nodes...)
}
func Ol(nodes ...Node) *Tag {
	return NewTag("ol", nodes...)
}
func Ul(nodes ...Node) *Tag {
	return NewTag("ul", nodes...)
}
func Li(nodes ...Node) *Tag {
	return NewTag("li", nodes...)
}
func Table(nodes ...Node) *Tag {
	return NewTag("table", nodes...)
}
func Caption(nodes ...Node) *Tag {
	return NewTag("caption", nodes...)
}
func Thead(nodes ...Node) *Tag {
	return NewTag("thead", nodes...)
}
func Tbody(nodes ...Node) *Tag {
	return NewTag("tbody", nodes...)
}
func Tfoot(nodes ...Node) *Tag {
	return NewTag("tfoot", nodes...)
}
func Tr(nodes ...Node) *Tag {
	return NewTag("tr", nodes...)
}
func Th(nodes ...Node) *Tag {
	return NewTag("th", nodes...)
}
func Td(nodes ...Node) *Tag {
	return NewTag("td", nodes...)
}
func Col(nodes ...Node) *Tag {
	return NewTag("col", nodes...)
}
func Colgroup(nodes ...Node) *Tag {
	return NewTag("colgroup", nodes...)
}
func Form(nodes ...Node) *Tag {
	return NewTag("form", nodes...)
}
func Fieldset(nodes ...Node) *Tag {
	return NewTag("fieldset", nodes...)
}
func Legend(nodes ...Node) *Tag {
	return NewTag("legend", nodes...)
}
func Label(nodes ...Node) *Tag {
	return NewTag("label", nodes...)
}
func Input(nodes ...Node) *Tag {
	return NewTag("input", nodes...)
}
func Select(nodes ...Node) *Tag {
	return NewTag("select", nodes...)
}
func Option(nodes ...Node) *Tag {
	return NewTag("option", nodes...)
}
func Optgroup(nodes ...Node) *Tag {
	return NewTag("optgroup", nodes...)
}
func Div(nodes ...Node) *Tag {
	return NewTag("div", nodes...)
}
func Section(nodes ...Node) *Tag {
	return NewTag("section", nodes...)
}
func Main(nodes ...Node) *Tag {
	return NewTag("main", nodes...)
}
func Button(nodes ...Node) *Tag {
	return NewTag("button", nodes...)
}
func A(nodes ...Node) *Tag {
	return NewTag("a", nodes...)
}
func H1(nodes ...Node) *Tag {
	return NewTag("h1", nodes...)
}
func H2(nodes ...Node) *Tag {
	return NewTag("h2", nodes...)
}
func H3(nodes ...Node) *Tag {
	return NewTag("h3", nodes...)
}
func P(nodes ...Node) *Tag {
	return NewTag("p", nodes...)
}
func Span(nodes ...Node) *Tag {
	return NewTag("span", nodes...)
}
func Field(nodes ...Node) *Tag {
	return NewTag("field", nodes...)
}
func Pre(nodes ...Node) *Tag {
	return NewTag("pre", nodes...)
}
func Body(nodes ...Node) *Tag {
	return NewTag("body", nodes...)
}
func Svg(nodes ...Node) *Tag {
	return NewTag("svg", nodes...)
}

// Void Elements
func Br(nodes ...Node) *Tag {
    el := NewTag("br", nodes...)
    el = WithVoidElement(el)
	return el
}
func Hr(nodes ...Node) *Tag {
    el := NewTag("hr", nodes...)
    el = WithVoidElement(el)
	return el
}
