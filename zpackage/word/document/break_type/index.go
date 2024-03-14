package break_type

type BreakTypeEnum int

const (
	Page BreakTypeEnum = iota
	Column
	TextWrapping
	Unsupported
)
