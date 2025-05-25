package humle_test

import (
	"testing"

	. "github.com/andrioid/humle"
)

func DocExample(nodes ...Node) Group {
	return Document(
		HTML(
			Head(
			//Slot("head"),
			),
		),
	)
}

var testNameSignals = TestCases{
	"head": TestCase{
		Input: DocExample(
			Signal("head", Title(Text("signal"))),
		),
		Expected: `<button class="bg-blue-500"></button>`,
	},
}

func TestNamedSignals(t *testing.T) {
	testNameSignals.Test(t)
}
