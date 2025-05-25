package humle_test

import (
	"testing"

	. "github.com/andrioid/humle"
)

func BaseButton(nodes ...Node) *Tag {
	return Button(
		Class("bg-blue-500"),
		Group(nodes),
	)
}

func PinkButton(nodes ...Node) *Tag {
	return BaseButton(
		Class("bg-pink-800"),
		Group(nodes),
	)
}

var testMatrixTailwind = TestCases{
	"base button": TestCase{
		Input:    BaseButton(),
		Expected: `<button class="bg-blue-500"></button>`,
	},
	"base button with tailwind merge": TestCase{
		Input:    BaseButton(Class("bg-purple-500")),
		Expected: `<button class="bg-purple-500"></button>`,
	},
	"pink button with tailwind merge": TestCase{
		Input: PinkButton(
			Class("text-pink-100"),
		),
		Expected: `<button class="bg-pink-800 text-pink-100"></button>`,
	},
	"basic template merge test": TestCase{
		Input:    Div(Class("bg-blue-500 bg-pink-500 bg-purple-500", "bg-orange-500")),
		Expected: `<div class="bg-orange-500"></div>`,
	},
}

func TestTailwind(t *testing.T) {
	testMatrixTailwind.Test(t)
}
