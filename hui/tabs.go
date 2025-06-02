package hui

import h "github.com/andrioid/humle"

type TabProps struct {
	Items []string
}

func (tp TabProps) ArgumentType() h.ArgumentType {
	return h.ArgumentProp
}

func Tabs(args h.Arguments) *h.Component {
	// TODO: How do I get a list of TabItems?
	// Could generics help here?
	return h.NewComponent("hui.Tabs",
		h.Div(
			// Passes attributes along to root tag
			args.GetAttributes()...,
		).Children(
			h.Div(),
			h.Div(),
		),
	)
}
