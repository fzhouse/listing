package listing

// Replacer interface
type Replacer interface {
	Len() int
	Replace([]int) Replacer
}

// IntReplacer type
type IntReplacer []int

// Len of IntReplacer
func (il IntReplacer) Len() int {
	return len(il)
}

// Replace of IntReplacer
func (il IntReplacer) Replace(indices []int) Replacer {
	result := make(IntReplacer, len(indices), len(indices))
	for i, idx := range indices {
		result[i] = il[idx]
	}
	return result
}

// RuneReplacer type
type RuneReplacer []rune

// Len of RuneReplacer
func (rl RuneReplacer) Len() int {
	return len(rl)
}

// Replace of RuneReplacer
func (rl RuneReplacer) Replace(indices []int) Replacer {
	result := make(RuneReplacer, len(indices), len(indices))
	for i, idx := range indices {
		result[i] = rl[idx]
	}
	return result
}

// StringReplacer type
type StringReplacer []string

// Len of StringReplacer
func (sl StringReplacer) Len() int {
	return len(sl)
}

// Replace of StringReplacer
func (sl StringReplacer) Replace(indices []int) Replacer {
	result := make(StringReplacer, len(indices), len(indices))
	for i, idx := range indices {
		result[i] = sl[idx]
	}
	return result
}

// Float64Replacer type
type Float64Replacer []string

// Len of Float64Replacer
func (fr Float64Replacer) Len() int {
	return len(fr)
}

// Replace of Float64Replacer
func (fr Float64Replacer) Replace(indices []int) Replacer {
	result := make(Float64Replacer, len(indices), len(indices))
	for i, idx := range indices {
		result[i] = fr[idx]
	}
	return result
}
