// Package painkiller ...
package painkiller

// Pill piankiller type
//go:generate stringer -type=Pill
type Pill int

const (
	// Placebo ...
	Placebo Pill = iota
	// Aspirin ...
	Aspirin
	// Ibuprofen ..
	Ibuprofen
	// Paracetamol ...
	Paracetamol
	// Acetaminophen ...
	Acetaminophen = Paracetamol
)
