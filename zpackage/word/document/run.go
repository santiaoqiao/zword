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
	RunProperty *RunProperty
	Children    []RunChild
}

// UnmarshalXML 解析<w:r>...</w:r>标签下的内容
func (r *Run) UnmarshalXML(d *xml.Decoder, _ xml.StartElement) error {
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
				r.Children = append(r.Children, b)
			}
			if t.Name.Local == "rPr" {
				rPr := &RunProperty{}
				err := rPr.UnmarshalXML(d, token.(xml.StartElement))
				if err != nil {
					return err
				}
				r.RunProperty = rPr
			}
		case xml.EndElement:
			if t.Name.Local == "r" {
				return nil
			}
		}
	}
	return nil
}

func (r *Run) String() string {
	sb := strings.Builder{}
	// 将 <w:r>...</w:r>中的文本全部合并起来，一般一个run中只有一个text
	for _, child := range r.Children {
		sb.WriteString(r.RunProperty.String())
		sb.WriteString(fmt.Sprintf("%v", child))
	}
	return sb.String()
}
