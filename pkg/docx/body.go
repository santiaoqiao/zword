package docx

import (
	"encoding/xml"
	"github.com/santiaoqiao/zword/pkg/docx/helper"
	"io"
	"strings"
)

// Body 对应主文档中的 <w:body>...</w:body>
type Body struct {
	Children []BodyChild
}

// UnmarshalXML 解析<body>元素
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
			switch t.Name.Space {
			case helper.CSpaceW:
				switch t.Name.Local {
				case "p":
					//<w:p>.....</w:p>
					p := &Paragraph{}
					//err := d.DecodeElement(p, &t)
					err := p.UnmarshalXML(d, t)
					if err != nil {
						return err
					}
					b.Children = append(b.Children, p)
				case "tbl":
					// <w:tbl>....</w:tbl>
					table := &Table{}
					//err := d.DecodeElement(table, &t)
					err := table.UnmarshalXML(d, t)
					if err != nil {
						return err
					}
					b.Children = append(b.Children, table)
				case "sectPr":
					// <w:sectPr>....</w:sectPr>
					secPr := &SectionProperty{}
					err := d.DecodeElement(secPr, &t)
					//err := secPr.UnmarshalXML(d, t)
					if err != nil {
						return err
					}
					b.Children = append(b.Children, secPr)
				default:
					break
				}
			}
		case xml.EndElement:
			if t.Name.Local == "body" {
				return nil
			}
		}
	}
	return nil
}

func (b *Body) String() string {
	sb := strings.Builder{}
	for _, child := range b.Children {
		sb.WriteString(child.String())
		sb.WriteString("\n")
	}
	return sb.String()
}
