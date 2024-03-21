package stroies

import "strings"

type Text struct {
	Text  string `xmldocx:",chardata"`
	Space string `xmldocx:"space,attr,omitempty"`
}

func (t *Text) String() string {
	if t.Space == "preserve" {
		return t.Text
	}
	return strings.TrimSpace(t.Text)
}
