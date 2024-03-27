package docx

import (
	"encoding/xml"
	"fmt"
	"github.com/santiaoqiao/zword/pkg/docx/helper"
	"io"
	"strings"
)

// ParagraphProperty 段落属性
type ParagraphProperty struct {
	// bidi 控制文字方向从右边向左
	bidi *bool
	// jc -justification 段落对齐方式
	justify string
	// 段落缩进
	indent Indent
	// 段落内的run属性
	rPr *RunProperty
	// 标记section
	secPr *SectionProperty
	// 大纲级别
	outLineLvl int // 0-9,9是无级别,没有此属性，默认为9普通文本
	// 样式
	pStyleId string
}

func (pPr *ParagraphProperty) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// 解析xml并给 pPr 赋值
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
				case "bidi":
					// <w:bidi w:val="0"/>
					pPr.bidi = helper.UnmarshalToggleValToBool(t, helper.CSpaceW)
				case "jc":
					// <w:jc w:val="center"/>
					pPr.justify = helper.UnmarshalSingleVal(t, helper.CSpaceW)
				case "ind":
					// <w:ind w:start="1440" w:end="1440" w:hanging="1080" />
					// <w:ind w:left="425" w:leftChars="0" w:hanging="425" w:firstLineChars="0"/>
					for _, attr := range t.Attr {
						switch attr.Name.Space {
						case helper.CSpaceW:
							switch attr.Name.Local {
							case "end":
								val, err := helper.Str2Int(attr.Value)
								if err != nil {
									return err
								}
								pPr.indent.end = val
							case "endChars":
								val, err := helper.Str2Int(attr.Value)
								if err != nil {
									return err
								}
								pPr.indent.endChars = val
							case "start":
								val, err := helper.Str2Int(attr.Value)
								if err != nil {
									return err
								}
								pPr.indent.start = val
							case "startChars":
								val, err := helper.Str2Int(attr.Value)
								if err != nil {
									return err
								}
								pPr.indent.startChars = val
							case "firstLine":
								val, err := helper.Str2Int(attr.Value)
								if err != nil {
									return err
								}
								pPr.indent.firstLine = val
							case "firstLineChars":
								val, err := helper.Str2Int(attr.Value)
								if err != nil {
									return err
								}
								pPr.indent.firstLineChars = val
							case "hanging":
								val, err := helper.Str2Int(attr.Value)
								if err != nil {
									return err
								}
								pPr.indent.hanging = val
							case "hangingChars":
								val, err := helper.Str2Int(attr.Value)
								if err != nil {
									return err
								}
								pPr.indent.hangingChars = val
							case "outlineLvl":
								val, err := helper.Str2Int(attr.Value)
								if err != nil {
									return err
								}
								pPr.outLineLvl = val
							}
						}
					}
				case "rPr":
					rPr := &RunProperty{}
					err := d.DecodeElement(rPr, &t)
					if err != nil {
						return err
					}
					pPr.rPr = rPr
				case "pStyle":
					//<w:pStyle w:val="TestParagraphStyle" />
					pPr.pStyleId = helper.UnmarshalSingleVal(t, helper.CSpaceW)
				}
			}
		case xml.EndElement:
			if t.Name.Local == "pPr" {
				return nil
			}
		}
	}
	return nil
}

func (pPr *ParagraphProperty) String() string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("Bidi: %v, justify:%v, indent:%v\n", pPr.bidi, pPr.justify, pPr.indent))
	sb.WriteString(fmt.Sprintf("%v", pPr.rPr))
	return sb.String()
}

// Indent 设置段落缩进
/* This element specifies the set of indentation properties applied to the current paragraph */
type Indent struct {
	// 指定应放置在本段末尾的缩进
	end int
	// 指定应放置在本段末尾的缩进,此值以字符单位的百分之一指定。
	endChars int
	// 指定应用于父段落第一行的附加缩进
	firstLine int
	// 以字符单位的百分之一指定
	firstLineChars int
	// 从第一行删除缩进
	hanging int
	// 从第一行删除缩进,以字符单位的百分之一指定
	hangingChars int
	//指定应放置在本段开头的缩进
	start int
	//指定应放置在本段开头的缩进,以字符单位的百分之一指定
	startChars int
}

func (pPr *ParagraphProperty) Bidi() *bool {
	return pPr.bidi
}
