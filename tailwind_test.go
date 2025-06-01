package humle_test

import (
	"testing"

	. "github.com/andrioid/humle"
)

var testMatrixTailwind = TestCases{
	"basic template merge test": TestCase{
		Input:    Div(Class("bg-blue-500 bg-pink-500 bg-purple-500", "bg-orange-500")),
		Expected: `<div class="bg-orange-500"></div>`,
	},
}

func TestTailwind(t *testing.T) {
	testMatrixTailwind.Test(t)
}
