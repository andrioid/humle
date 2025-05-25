package humle_test

import (
	"testing"

	. "github.com/andrioid/humle"
)

var testMatrix2 = TestCases{
	"with value": TestCase{
		Input:    Div(Text("test")),
		Expected: `<div>test</div>`,
	},
	"empty non-void": TestCase{
		Input:    Div(),
		Expected: `<div></div>`,
	},
	"with class": TestCase{
		Input:    Div(Class("bg-pink-500")),
		Expected: `<div class="bg-pink-500"></div>`,
	},
	"br": TestCase{
		Input:    Br(),
		Expected: `<br>`,
	},
	"with text": TestCase{
		Input:    Div(Text("hello")),
		Expected: `<div>hello</div>`,
	},
	"with id and data": TestCase{
		Input:    Div(ID("divid"), Data("show", `$tab === 'logs'`)),
		Expected: `<div id="divid" data-show="$tab === 'logs'"></div>`,
	},
	"raw": TestCase{
		Input:    RawHTML("<div class=\"bg-blue-500\"></div>"),
		Expected: `<div class="bg-blue-500"></div>`,
	},
	"raw inside div": TestCase{
		Input:    Div(RawHTML("<div class=\"bg-blue-500\"></div>")),
		Expected: `<div><div class="bg-blue-500"></div></div>`,
	},
	"document empty": TestCase{
		Input:    Document(),
		Expected: `<!DOCTYPE html>`,
	},
	"document with nodes": TestCase{
		Input: Document(
			HTML(
				Head(),
				Body(),
			)),
		Expected: `<!DOCTYPE html><html><head></head><body></body></html><html><head></head><body></body></html>`,
	},
}

func TestRender(t *testing.T) {
	testMatrix2.Test(t)
}
