package humle_test

import (
	"fmt"
	"testing"

	. "github.com/andrioid/humle"
)

var testMatrix = map[Node]string{
	Div(Text("test")):         `<div>test</div>`,
	Div():                     `<div></div>`,
	Div(Class("bg-pink-500")): `<div class="bg-pink-500"></div>`,
	Br():                      `<br>`,
	Div(Text("hello")):        `<div>hello</div>`,
	Div(ID("divid"), Data("show", `$tab === 'logs'`)): `<div id="divid" data-show="$tab === 'logs'"></div>`,
}

func TestRender(t *testing.T) {
	for el, expected := range testMatrix {
		if fmt.Sprintf("%s", el) != expected {
			t.Errorf("DOM mismatch\n\033[32m%s\033[0m\n\033[31m%s\033[0m", expected, el)
		}
	}
}
