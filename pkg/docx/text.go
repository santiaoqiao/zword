package docx

import "strings"

type Text struct {
	Text  string `xml:",chardata"`
	Space string `xml:"space,attr,omitempty"`
}

func (t *Text) String() string {
	if t.Space == "preserve" {
		return t.Text
	}
	return strings.TrimSpace(t.Text)
}

func (t *Text) TypeName() RunChildType {
	return RunTypeText
}
