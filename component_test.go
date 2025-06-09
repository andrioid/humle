package humle_test

import (
	"testing"

	h "github.com/andrioid/humle"
)

func basicLayout() *h.Component {
	return h.NewComponent("basic-layout", func() h.Node {
		return h.Div(h.Text("Hello"))
	})
}

var testMatrixComponents = h.TestCases{
	"component": h.TestCase{
		Input:    basicLayout(),
		Expected: `<div>Hello</div>`,
	},
}

func TestComponents(t *testing.T) {
	testMatrixComponents.Test(t)
}
