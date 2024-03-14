package document

import (
	"encoding/xml"
	"io"
)

type Body struct {
	Children []BodyChild
}

type BodyChild interface{}

func (b *Body) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for {
		token, err := d.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		switch t := token.(type) {
		case xml.StartElement:
			// <w:p>.....</w:p>，交给 Paragraph 处理
			if t.Name.Local == "p" {
				p := &Paragraph{}
				err := p.UnmarshalXML(d, token.(xml.StartElement))
				if err != nil {
					return err
				}
				b.Children = append(b.Children, p)
			}
			// <w:tbl>....<w:tbl>，交给 Table 处理
			if t.Name.Local == "tbl" {
				// todo: table处理
			}
		case xml.EndElement:
			if t.Name.Local == "body" {
				return nil
			}
		}
	}
	return nil
}
