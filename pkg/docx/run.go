package docx

import (
	"encoding/xml"
	"fmt"
	"github.com/santiaoqiao/zword/pkg/docx/helper"
	"io"
	"strings"
)

type Run struct {
	RunProperty     *RunProperty
	Children        []RunChild
	ParentParagraph *Paragraph
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
			switch t.Name.Space {
			case helper.CSpaceW:
				switch t.Name.Local {
				case "t":
					// <w:t>hello</w:t>
					text := &Text{}
					if err := d.DecodeElement(text, &t); err != nil {
						return err
					}
					r.Children = append(r.Children, text)
				case "br":
					// <w:br />
					b := &Break{}
					b.Val = helper.UnmarshalToggleValToBool(t, helper.CSpaceW)
					r.Children = append(r.Children, b)
				case "rPr":
					// 获取paragraph的样式，用于做最后的样式计算
					var paragraphStyleRPr *RunProperty = nil
					if r.ParentParagraph.Property.pStyleId != "" {
						paragraphStyle := getStyle(r.ParentParagraph.Property.pStyleId)
						if paragraphStyle != nil {
							paragraphStyleRPr = paragraphStyle.RPr
						}
					}
					rPr := &RunProperty{parentRun: r, paragraphStyleRPr: paragraphStyleRPr}
					err := d.DecodeElement(rPr, &t)
					if err != nil {
						return err
					}
					r.RunProperty = rPr
				}
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
	// 每个run的属性
	//sb.WriteString(fmt.Sprintf("%v\n", r.RunProperty))
	for _, child := range r.Children {
		sb.WriteString(fmt.Sprintf("%v", child))
	}
	return sb.String()
}

func (r *Run) TypeName() ParagraphChildType {
	return ParagraphTypeRun
}
