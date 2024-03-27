package docx

import (
	"encoding/xml"
	"io"
	"santiaoqiao.com/zword/pkg/docx/helper"
	"strings"
)

type Paragraph struct {
	Children     []ParagraphChild
	Property     *ParagraphProperty
	HasNumbering bool
}

func (p *Paragraph) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			switch t.Name.Space {
			case helper.CSpaceW:
				switch t.Name.Local {
				case "r":
					// <w:r>.....</w:r>
					r := &Run{ParentParagraph: p}
					err := r.UnmarshalXML(d, token.(xml.StartElement))
					if err != nil {
						return err
					}
					p.Children = append(p.Children, r)
				case "pPr":
					// <w:pPr>....<w:pPr>
					pRp := &ParagraphProperty{}
					err := pRp.UnmarshalXML(d, token.(xml.StartElement))
					if err != nil {
						return err
					}
					p.Property = pRp
				}
			}
		case xml.EndElement:
			if t.Name.Space == helper.CSpaceW && t.Name.Local == "p" {
				return nil
			}
		}
	}
	return nil
}

func (p *Paragraph) String() string {
	sb := strings.Builder{}
	//sb.WriteString(fmt.Sprintf("%v\n", p.Property))
	for _, child := range p.Children {
		sb.WriteString(child.String())
	}
	return sb.String()
}

func (p *Paragraph) TypeName() BodyChildType {
	return BodyTypeParagraph
}
