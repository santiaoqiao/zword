package zword

type Break struct {
	BreakType string
}

func (b *Break) String() string {
	return b.BreakType
}
