// Generated from humle/html package
// **DO NOT EDIT**
package humle

// Elements
func Div(nodes ...Node) *Element {
	return NewElement("div", nodes...)
}
func Section(nodes ...Node) *Element {
	return NewElement("section", nodes...)
}
func Main(nodes ...Node) *Element {
	return NewElement("main", nodes...)
}
func Button(nodes ...Node) *Element {
	return NewElement("button", nodes...)
}
func A(nodes ...Node) *Element {
	return NewElement("a", nodes...)
}
func H1(nodes ...Node) *Element {
	return NewElement("h1", nodes...)
}
func H2(nodes ...Node) *Element {
	return NewElement("h2", nodes...)
}
func H3(nodes ...Node) *Element {
	return NewElement("h3", nodes...)
}
func P(nodes ...Node) *Element {
	return NewElement("p", nodes...)
}
func Span(nodes ...Node) *Element {
	return NewElement("span", nodes...)
}
func Input(nodes ...Node) *Element {
	return NewElement("input", nodes...)
}
func Label(nodes ...Node) *Element {
	return NewElement("label", nodes...)
}
func Field(nodes ...Node) *Element {
	return NewElement("field", nodes...)
}
func Pre(nodes ...Node) *Element {
	return NewElement("pre", nodes...)
}
func Script(nodes ...Node) *Element {
	return NewElement("script", nodes...)
}
func Link(nodes ...Node) *Element {
	return NewElement("link", nodes...)
}
func Head(nodes ...Node) *Element {
	return NewElement("head", nodes...)
}
func Body(nodes ...Node) *Element {
	return NewElement("body", nodes...)
}
func Meta(nodes ...Node) *Element {
	return NewElement("meta", nodes...)
}
func Svg(nodes ...Node) *Element {
	return NewElement("svg", nodes...)
}

// Void Elements
func Br(nodes ...Node) *Element {
    el := NewElement("br", nodes...)
    el = WithVoidElement(el)
	return el
}
func Hr(nodes ...Node) *Element {
    el := NewElement("hr", nodes...)
    el = WithVoidElement(el)
	return el
}
