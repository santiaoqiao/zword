package docx

type Break struct {
	Val *bool
}

func (b *Break) String() string {
	if *b.Val {
		return "/n"
	}
	return ""
}

func (b *Break) TypeName() RunChildType {
	return RunTypeBreak
}
