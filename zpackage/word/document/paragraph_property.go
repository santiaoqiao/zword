package document

import (
	"encoding/xml"
	"io"
	"santiaoqiao.com/zoffice/zpackage/helper"
)

type ParagraphProperty struct {
	// Bidi 控制文字方向从右边向左
	/* Right to Left Paragraph Layout
	This element specifies that this paragraph shall be displayed from right to left.*/
	Bidi bool
	// Jc -justification 段落排列方式
	/* Paragraph Alignment
	This element specifies the paragraph alignment which shall be applied to text in this paragraph.

	可能的值为如下：
	both	两端对齐
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
	Jc string
}

func (pPr *ParagraphProperty) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// 初始化 pPr
	pPr.Bidi = false
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
			if t.Name.Local == "bidi" {
				if val, ok := helper.UnwrapValToBool(t); ok {
					pPr.Bidi = val
				} else {
					pPr.Bidi = true
				}
			}
			if t.Name.Local == "jc" {
				if val, ok := helper.UnwrapVal(t); ok {
					pPr.Jc = val
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
