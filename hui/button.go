package hui

import h "github.com/andrioid/humle"

func Button(args h.Arguments) *h.Component {
	// TODO: How do I get a list of TabItems?
	return h.NewComponent("hui.button",
		h.Button(
			// Passes attributes along to root tag
			args.GetAttributes()...,
		).Children(
			h.Div(),
			h.Div(),
		),
	)
}

func PinkButton(args h.Arguments) *h.Component {
	newArgs := h.Arguments(args).AddAttributes(
		// This appends or merges the class attribute to the base button
		h.Class("text-pink-800"),
	)

	return Button(newArgs)
}
