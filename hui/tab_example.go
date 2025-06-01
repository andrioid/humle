package hui

import h "github.com/andrioid/humle"

func TabExample() *h.Component {
	return Tabs(h.Arguments{
		// Passes items to Tabs
		TabProps{
			Items: []string{"bleh"},
		},
		// Overrides the Tabs root element class
		h.Class("border border-pink-200"),
	})
}
