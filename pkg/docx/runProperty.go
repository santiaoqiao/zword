package docx

import (
	"encoding/xml"
	"fmt"
	"io"
	"santiaoqiao.com/zword/pkg/docx/helper"
	"strings"
)

/*
	RunProperty Run的属性，与 XML 文档对应

This element specifies a set of run properties which shall be applied to the contents of the parentRun run after all
style formatting has been applied to the text. These properties are defined as direct formatting, since they are
directly applied to the run and supersede any formatting from styles
*/
type RunProperty struct {
	// 粗体（简单文字）
	bold bool
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
	FontKerning int
	// 拼写检查的语言
	Lang Language
	// 外轮廓
	Outline bool
	// 文字在垂直方向上上下偏移的距离
	Position int
	// 字体
	Fonts RunFonts
	// 样式ID
	StyleId string
	// 字体大小（简单文字）
	Size int
	// 字体大小（复杂脚本）
	SizeCs int
	// 指向样式表中的指针
	parentRun *Run
}

// UnmarshalXML 解析XML文档
func (rPr *RunProperty) UnmarshalXML(d *xml.Decoder, _ xml.StartElement) error {
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
			//space为w的tag <w:....>
			case helper.CSpaceW:
				switch t.Name.Local {
				case "b":
					// <w:b w:val="false"/> | <w:b "/>
					rPr.bold = helper.UnmarshalToggleValToBool(t, helper.CSpaceW)
				case "bCs":
					// <w:bCs w:val="false"/> | <w:bCs />
					rPr.BoldCs = helper.UnmarshalToggleValToBool(t, helper.CSpaceW)
				case "color":
					//<w:color w:themeColor="accent3"  w:val="FF0000"/>
					for _, attr := range t.Attr {
						switch {
						case attr.Name.Space == helper.CSpaceW && attr.Name.Local == "val":
							rPr.Color.Value = attr.Value
						case attr.Name.Space == helper.CSpaceW && attr.Name.Local == "themeColor":
							rPr.Color.Theme = attr.Value
						}
					}
				case "cs":
					//<w:cs/>
					rPr.ComplexScript = helper.UnmarshalToggleValToBool(t, helper.CSpaceW)
				case "dstrike":
					//<w:dstrike w:val="true"/>
					rPr.DoubleStrikethrough = helper.UnmarshalToggleValToBool(t, helper.CSpaceW)
				case "em":
					//<w:em w:val="dot"/>
					rPr.EmphasisMark = helper.UnmarshalSingleVal(t, helper.CSpaceW)
				case "i":
					//	<w:i />
					rPr.Italics = helper.UnmarshalToggleValToBool(t, helper.CSpaceW)
				case "iCs":
					// <w:iCs w:val="true"/>
					rPr.ItalicsCs = helper.UnmarshalToggleValToBool(t, helper.CSpaceW)
				case "imprint":
					// <w:imprint w:val="true"/>
					rPr.Imprint = helper.UnmarshalToggleValToBool(t, helper.CSpaceW)
				case "kern":
					// <w:kern w:val="28" />
					if val, err := helper.UnmarshalSingleValToInt(t, helper.CSpaceW); err != nil {
						return err
					} else {
						rPr.FontKerning = val
					}
				case "lang":
					// <w:lang w:val="fr-CA" w:bidi="he-IL" />
					for _, attr := range t.Attr {
						switch attr.Name.Space {
						case helper.CSpaceW:
							switch attr.Name.Local {
							case "bidi":
								rPr.Lang.Bidi = attr.Value
							case "val":
								rPr.Lang.Value = attr.Value
							case "eastAsia":
								rPr.Lang.EastAsian = attr.Value
							}
						}
					}
				case "outline":
					//<w:outline w:val="false"/>
					rPr.Outline = helper.UnmarshalToggleValToBool(t, helper.CSpaceW)
				case "position":
					// <w:position w:val="24" />
					if val, err := helper.UnmarshalSingleValToInt(t, helper.CSpaceW); err != nil {
						return err
					} else {
						rPr.Position = val
					}
				case "rFonts":
					// <w:rFonts w:Ascii="Courier New" w:Cs="Times New Roman" />
					// <w:rFonts w:Hint="EastAsia" w:Ascii="黑体" w:HAnsi="黑体" w:EastAsia="黑体" w:Cs="黑体"/>
					// <w:rFonts w:Hint="default" w:AsciiTheme="minorAscii" w:HAnsiTheme="minorAscii" w:EastAsiaTheme="minorEastAsia"/>
					for _, attr := range t.Attr {
						switch attr.Name.Space {
						case helper.CSpaceW:
							switch attr.Name.Local {
							case "hint":
								rPr.Fonts.Hint = attr.Value
							case "ascii":
								rPr.Fonts.Ascii = attr.Value
							case "cs":
								rPr.Fonts.Cs = attr.Value
							case "eastAsia":
								rPr.Fonts.EastAsia = attr.Value
							case "hAnsi":
								rPr.Fonts.HAnsi = attr.Value
							case "asciiTheme":
								rPr.Fonts.AsciiTheme = attr.Value
							case "eastAsiaTheme":
								rPr.Fonts.EastAsiaTheme = attr.Value
							case "hAnsiTheme":
								rPr.Fonts.HAnsiTheme = attr.Value
							}
						}

					}
				case "rStyle":
					// <w:rStyle w:val="14"/>
					rPr.StyleId = helper.UnmarshalSingleVal(t, helper.CSpaceW)
				case "sz":
					// <w:sz w:val="27"/>
					val, err := helper.UnmarshalSingleValToInt(t, helper.CSpaceW)
					if err != nil {
						return err
					}
					rPr.Size = val
				case "szCs":
					//<w:szCs w:val="20"/>
					val, err := helper.UnmarshalSingleValToInt(t, helper.CSpaceW)
					if err != nil {
						return err
					}
					rPr.SizeCs = val
				}
			}

		case xml.EndElement:
			if t.Name.Local == "rPr" {
				return nil
			}
		}
	}
	return nil
}

// String 输出为字符串
func (rPr *RunProperty) String() string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("%#v", rPr))
	return sb.String()
}

func (rPr *RunProperty) Bold() bool {
	//pStyle := rPr.parentRun.ParentParagraph.Property.pStyleId
	if DocFile.Styles != nil && rPr.StyleId != "" {
		//if s, ok := DocFile.Styles.StyleSheets[rPr.StyleId]; ok {
		//
		//}
	}
	return rPr.bold
}

// Color 字体颜色
type Color struct {
	// 字体颜色值，如 D4F4F2，前面不带#号
	Value string
	// 字体的主题颜色，如运用了主题，以主题为主
	Theme string
}

// Language 字体语言
type Language struct {
	// 指定在处理使用拉丁字符的运行内容时(由运行内容的Unicode字符值决定)应用于检查拼写和语法(如果请求)的语言
	Value string
	// 指定在处理使用复杂脚本字符的运行内容时应使用的语言，由运行内容的Unicode字符值决定。
	Bidi string
	// 指定在处理使用东亚字符的运行内容时应使用的语言
	EastAsian string
}

// RunFonts 最多有4种字体槽
type RunFonts struct {
	// 默认提示所用的子图
	Hint string
	// 处理Ascii字符时所使用的字体
	Ascii string
	// 处理 High ANSI 字符时所使用的字体
	HAnsi string
	// 处理东南亚 East Asian 文字所使用的字体，包括中文
	EastAsia string
	// 处理 Complex Script 字符时所使用的字体
	Cs string
	// Ascii字符所使用的主题
	AsciiTheme string
	// High ANSI字符所使用的主题
	HAnsiTheme string
	// 东南亚文字所使用的主题
	EastAsiaTheme string
}
