package properties

import (
	"encoding/xml"
	"fmt"
	"io"
	"santiaoqiao.com/zword/pkg/xml_parser/helper"
	"santiaoqiao.com/zword/pkg/xml_parser/stroies"
	"strings"
)

// ParagraphProperty 段落属性
type ParagraphProperty struct {
	// bidi 控制文字方向从右边向左
	bidi bool
	// jc -justification 段落对齐方式
	/* both	两端对齐
	center	居中
	distribute	平均分配所有字符（分散对齐）
	end	右对齐
	highKashida	最宽 Kashida长度，用于类似阿拉伯语中，详见《ECMA-376-1:2016》 p1399
	lowKashida	最窄 Kashida长度
	mediumKashida	中等 Kashida长度
	numTab	与列表选项卡对齐
	start	左对齐，Align To Leading Edge
	thaiDistribute	泰语对齐方式
	*/
	justify string
	// 段落缩进
	indent Indent
	// 段落内的run属性
	rPr *RunProperty
	// 标记section
	secPr *SectionProperty
}

func (pPr *ParagraphProperty) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// 初始化 pPr
	pPr.bidi = false
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
			case stroies.cSpaceW:
				switch t.Name.Local {
				case stroies.cTagBidi:
					// <w:Bidi w:val="0"/>
					pPr.bidi = helper.UnmarshalToggleValToBool(t, stroies.cSpaceW)
				case stroies.cTagJustify:
					// <w:jc w:val="center"/>
					pPr.justify = helper.UnmarshalSingleVal(t, stroies.cSpaceW)
				case stroies.cTagInd:
					// <w:ind w:start="1440" w:end="1440" w:hanging="1080" />
					// <w:ind w:left="425" w:leftChars="0" w:hanging="425" w:firstLineChars="0"/>
					for _, attr := range t.Attr {
						switch attr.Name.Space {
						case stroies.cSpaceW:
							switch attr.Name.Local {
							case stroies.cTagEnd:
								val, err := helper.Str2Int(attr.Value)
								if err != nil {
									return err
								}
								pPr.indent.end = val
							case stroies.cTagEndChars:
								val, err := helper.Str2Int(attr.Value)
								if err != nil {
									return err
								}
								pPr.indent.endChars = val
							case stroies.cTagStart:
								val, err := helper.Str2Int(attr.Value)
								if err != nil {
									return err
								}
								pPr.indent.start = val
							case stroies.cTagStartChars:
								val, err := helper.Str2Int(attr.Value)
								if err != nil {
									return err
								}
								pPr.indent.startChars = val
							case stroies.cTagFirstLine:
								val, err := helper.Str2Int(attr.Value)
								if err != nil {
									return err
								}
								pPr.indent.firstLine = val
							case stroies.cTagFirstLineChars:
								val, err := helper.Str2Int(attr.Value)
								if err != nil {
									return err
								}
								pPr.indent.firstLineChars = val
							case stroies.cTagHanging:
								val, err := helper.Str2Int(attr.Value)
								if err != nil {
									return err
								}
								pPr.indent.hanging = val
							case stroies.cTagHangingChars:
								val, err := helper.Str2Int(attr.Value)
								if err != nil {
									return err
								}
								pPr.indent.hangingChars = val
							}
						}
					}
				case stroies.cTagRPr:
					rPr := &RunProperty{}
					err := d.DecodeElement(rPr, &t)
					if err != nil {
						return err
					}
					pPr.rPr = rPr
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
