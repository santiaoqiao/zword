package stroies

import "strings"

type Text struct {
	Text  string `xml_parser:",chardata"`
	Space string `xml_parser:"space,attr,omitempty"`
}

func (t *Text) String() string {
	if t.Space == "preserve" {
		return t.Text
	}
	return strings.TrimSpace(t.Text)
}
