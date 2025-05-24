package humle

func Data(k, v string) *Attribute {
	return NewAttribute("data-"+k, v)
}
