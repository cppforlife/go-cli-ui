package ui

// TextOpts Asking for text options
type TextOpts struct {
	Label        string
	DefaultValue string
}

// ChoiceOpts asking for choice options
type ChoiceOpts struct {
	Label        string
	DefaultValue string
	Choices      []string
}
