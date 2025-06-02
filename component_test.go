package humle_test

import (
	"testing"

	. "github.com/andrioid/humle"
)

func BaseButton(args ...Argument) *Component {
	newArgs := Arguments(args).AddAttributes(Class("bg-blue-500"))
	return NewComponent(
		"basebutton",
		Button(newArgs.GetAttributes()...),
	)

}

func PinkButton(args ...Argument) *Component {
	newArgs := Arguments(args).AddAttributes(Class("text-pink-800"))
	comp := NewComponent("PinkButton", BaseButton(newArgs...))
	return comp
}

var testMatrixComponents = TestCases{
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
}

func TestComponents(t *testing.T) {
	testMatrixComponents.Test(t)
}
