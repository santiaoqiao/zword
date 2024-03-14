package document

import (
	"encoding/xml"
	"fmt"
	"io"
	"santiaoqiao.com/zoffice/zpackage/helper"
	"strings"
)

type RunChild interface{}

type Run struct {
	//RunProperty interface{}
	Children []RunChild
}

// UnmarshalXML 解析<w:r>...</w:r>标签下的内容
func (r *Run) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			// <w:t>hello</w:t>
			if t.Name.Local == "t" {
				text := &Text{}
				if err := d.DecodeElement(&text.Text, &t); err != nil {
					return err
				}
				r.Children = append(r.Children, text)
			}
			// <w:br />
			if t.Name.Local == "br" {
				b := &Break{}
				if value, ok := helper.Unwrap(t, "type"); ok {
					b.BreakType = value
				}
				//for _, attr := range t.Attr {
				//	// <w:br w:type="textWrapping" />
				//	if attr.Name.Local == "type" {
				//		switch attr.Value {
				//		case "page":
				//			b.BreakType = break_type.Page
				//		case "column":
				//			b.BreakType = break_type.Column
				//		case "textWrapping":
				//			b.BreakType = break_type.TextWrapping
				//		default:
				//			b.BreakType = break_type.Unsupported
				//		}
				//	}
				//}
				r.Children = append(r.Children, b)
			}
		case xml.EndElement:
			if t.Name.Local == "r" {
				return nil
			}
		}
	}
	return nil
}

func (r Run) String() string {
	sb := strings.Builder{}
	for index, child := range r.Children {
		sb.WriteString(fmt.Sprintf("R%d - %v\n", index, child))
	}
	return sb.String()
}
