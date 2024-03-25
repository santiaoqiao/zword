package stroies

import (
	"encoding/xml"
	"io"
	"santiaoqiao.com/zword/pkg/xml_parser/properties"
	"strings"
)

// Body 对应主文档中的 <w:body>...</w:body>
type Body struct {
	Children []BodyChild
}

// BodyChild <body>中的元素：p、tbl、sectPr
type BodyChild interface {
	String() string
	TypeName() string
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
			case cSpaceW:
				switch t.Name.Local {
				case cTagP:
					//<w:p>.....</w:p>
					p := &Paragraph{}
					err := d.DecodeElement(p, &t)
					if err != nil {
						return err
					}
					b.Children = append(b.Children, p)
				case cTagTbl:
					// <w:tbl>....</w:tbl>
					table := &Table{}
					err := d.DecodeElement(table, &t)
					if err != nil {
						return err
					}
					b.Children = append(b.Children, table)
				case cTagSectPr:
					// <w:sectPr>....</w:sectPr>
					secPr := &properties.SectionProperty{}
					err := d.DecodeElement(secPr, &t)
					if err != nil {
						return err
					}
					b.Children = append(b.Children, secPr)
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
