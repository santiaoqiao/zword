package docx

import (
	"encoding/xml"
	"io"
	"strings"
)

type Paragraph struct {
	Children     []ParagraphChild
	Property     *ParagraphProperty
	HasNumbering bool
}

type ParagraphChild interface {
	String() string
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
			//fmt.Printf("paragraph:\t<%s>\n", t.Name.Local)
			// <w:r>.....</w:r>，交给 Run 处理
			if t.Name.Local == "r" {
				r := &Run{}
				err := r.UnmarshalXML(d, token.(xml.StartElement))
				if err != nil {
					return err
				}
				//fmt.Println(*r)
				p.Children = append(p.Children, r)
			}
			// <w:pPr>....<w:pPr>，交给 ParagraphProperty 处理
			if t.Name.Local == "pPr" {
				pRp := &ParagraphProperty{}
				err := pRp.UnmarshalXML(d, token.(xml.StartElement))
				if err != nil {
					return err
				}
				p.Property = pRp
			}
		case xml.EndElement:
			//fmt.Printf("paragraph:\t</%s>\n", t.Name.Local)
			if t.Name.Local == "p" {
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
