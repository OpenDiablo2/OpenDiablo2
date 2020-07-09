package d2enum

// TermCategory applies styles to the lines in the  Terminal
type TermCategory int

// Terminal Category types
const (
	TermCategoryNone TermCategory = iota
	TermCategoryInfo
	TermCategoryWarning
	TermCategoryError
)
