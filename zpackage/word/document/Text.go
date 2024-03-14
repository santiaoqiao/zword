package document

type Text struct {
	Text          string
	PreserveSpace bool
}

func (t *Text) String() string {
	return t.Text
}
