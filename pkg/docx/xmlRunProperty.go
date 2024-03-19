package docx

import (
	"encoding/xml"
	"fmt"
	"io"
	"santiaoqiao.com/zoffice/pkg/helper"
	"strings"
)

const (
	nsW string = "http://schemas.openxmlformats.org/wordprocessingml/2006/main"

	tagB   string = "b"
	tagBcs        = "bCs"
)

/*
	xmlRunProperty Run的属性，与 XML 文档对应

This element specifies a set of run properties which shall be applied to the contents of the parent run after all
style formatting has been applied to the text. These properties are defined as direct formatting, since they are
directly applied to the run and supersede any formatting from styles
*/
type xmlRunProperty struct {
	// 粗体（简单文字）
	Bold bool
	// 粗体（复杂脚本）
	BoldCs bool
	// 字体颜色
	Color Color
	// 是否为标记 Complex Script
	ComplexScript bool
	// 双横线穿过
	DoubleStrikethrough bool
	// 强调 <w:em w:val="dot"/>
	EmphasisMark string
	// 斜体（简单文字）
	Italics bool
	// 斜体（复杂脚本）
	ItalicsCs bool
	// 浮雕
	Imprint bool
	// 字符字距
	Kern int
	// 拼写检查的语言
	Lang Language
	// 外轮廓
	Outline bool
	// 文字在垂直方向上上下偏移的距离
	Position int
	// 字体
	RFonts RunFonts
	// 样式ID
	RStyleId string
	// 字体大小（简单文字）
	Sz int
	// 字体大小（复杂脚本）
	SzCs int
}

// UnmarshalXML 解析XML文档
func (rPr *xmlRunProperty) UnmarshalXML(d *xml.Decoder, _ xml.StartElement) error {
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
			switch {
			case t.Name.Space == nsW && t.Name.Local == tagB:
				// <w:b w:val="false"/> | <w:b "/>
				rPr.Bold = helper.UnwrapValToToggle(t)
			case t.Name.Space == nsW && t.Name.Local == tagBcs:
				// <w:bCs w:val="false"/> | <w:bCs />
				rPr.BoldCs = helper.UnwrapValToToggle(t)
			}
			//switch t.Name.Local {
			//case tagB:
			//	// <w:b w:val="false"/> | <w:b "/>
			//	rPr.Bold = helper.UnwrapValToToggle(t)
			//case "bCs":
			//	// <w:bCs w:val="false"/> | <w:bCs />
			//	rPr.BoldCs = helper.UnwrapValToToggle(t)
			//case "color":
			//	//<w:color w:themeColor="accent3" />
			//	if val, ok := helper.Unwrap(t, "themeColor"); ok {
			//		rPr.Color.Theme = val
			//	}
			//	// <w:color w:val="FF0000"/>
			//	if val, ok := helper.Unwrap(t, "val"); ok {
			//		rPr.Color.Value = val
			//	}
			//case "cs":
			//	rPr.ComplexScript = true
			//case "dstrike":
			//	rPr.DoubleStrikethrough = helper.UnwrapValToToggle(t)
			//case "em":
			//	if value, ok := helper.Unwrap(t, "em"); ok {
			//		rPr.EmphasisMark = value
			//	}
			//case "i":
			//	rPr.Italics = helper.UnwrapValToToggle(t)
			//case "iCs":
			//	rPr.ItalicsCs = helper.UnwrapValToToggle(t)
			//case "imprint":
			//	rPr.Imprint = helper.UnwrapValToToggle(t)
			//case "kern":
			//	if value, _, err := helper.UnwrapValToInt(t); err != nil {
			//		return err
			//	} else {
			//		rPr.Kern = value
			//	}
			//case "lang":
			//	for _, attr := range t.Attr {
			//		switch attr.Name.Local {
			//		case "bidi":
			//			rPr.Lang.Bidi = attr.Value
			//		case "val":
			//			rPr.Lang.Value = attr.Value
			//		case "eastAsia":
			//			rPr.Lang.EastAsian = attr.Value
			//		}
			//	}
			//case "outline":
			//	rPr.Outline = helper.UnwrapValToToggle(t)
			//case "position":
			//	val, _, err := helper.UnwrapValToInt(t)
			//	if err != nil {
			//		return err
			//	}
			//	rPr.Position = val
			//case "rFonts":
			//	if val, ok := helper.Unwrap(t, "hint"); ok {
			//		rPr.RFonts.Hint = val
			//	} else if val, ok = helper.Unwrap(t, "ascii"); ok {
			//		rPr.RFonts.Ascii = val
			//	} else if val, ok = helper.Unwrap(t, "hAnsi"); ok {
			//		rPr.RFonts.HAnsi = val
			//	} else if val, ok = helper.Unwrap(t, "eastAsia"); ok {
			//		rPr.RFonts.EastAsia = val
			//	} else if val, ok = helper.Unwrap(t, "cs"); ok {
			//		rPr.RFonts.Cs = val
			//	} else if val, ok = helper.Unwrap(t, "asciiTheme"); ok {
			//		rPr.RFonts.AsciiTheme = val
			//	} else if val, ok = helper.Unwrap(t, "hAnsiTheme"); ok {
			//		rPr.RFonts.HAnsiTheme = val
			//	} else if val, ok = helper.Unwrap(t, "eastAsiaTheme"); ok {
			//		rPr.RFonts.EastAsiaTheme = val
			//	}
			//case "rStyle":
			//	if val, ok := helper.UnwrapVal(t); ok {
			//		rPr.RStyleId = val
			//	}
			//case "sz":
			//	if val, _, err := helper.UnwrapValToInt(t); err != nil {
			//		return err
			//	} else {
			//		rPr.Sz = val
			//	}
			//case "szCs":
			//	if val, _, err := helper.UnwrapValToInt(t); err != nil {
			//		return err
			//	} else {
			//		rPr.SzCs = val
			//	}
			//}
		case xml.EndElement:
			if t.Name.Local == "rPr" {
				return nil
			}
		}
	}
	return nil
}

// String 输出为字符串
func (rPr *xmlRunProperty) String() string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("%#v", rPr))
	return sb.String()
}
