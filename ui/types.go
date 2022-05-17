package ui

// TextOpts Asking for text options
type TextOpts struct {
	Label   string
	Default string
}

// ChoiceOpts asking for choice options
type ChoiceOpts struct {
	Label   string
	Default string
	Choices []string
}
