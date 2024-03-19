package docx

import "strings"

type Text struct {
	Text          string
	PreserveSpace bool
}

func (t *Text) String() string {
	if t.PreserveSpace {
		return t.Text
	}
	return strings.TrimSpace(t.Text)
}
