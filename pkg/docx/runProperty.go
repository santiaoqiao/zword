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

This element specifies a set of run properties which shall be applied to the contents of the parent run after all
style formatting has been applied to the text. These properties are defined as direct formatting, since they are
directly applied to the run and supersede any formatting from styles
*/
type RunProperty struct {
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
			case cSpaceW:
				switch t.Name.Local {
				case cTagBold:
					// <w:b w:val="false"/> | <w:b "/>
					rPr.Bold = helper.UnmarshalToggleValToBool(t, cSpaceW)
				case cTagBoldCs:
					// <w:bCs w:val="false"/> | <w:bCs />
					rPr.BoldCs = helper.UnmarshalToggleValToBool(t, cSpaceW)
				case cTagColor:
					//<w:Color w:themeColor="accent3"  w:val="FF0000"/>
					for _, attr := range t.Attr {
						switch {
						case attr.Name.Space == cSpaceW && attr.Name.Local == cAttrVal:
							rPr.Color.Value = attr.Value
						case attr.Name.Space == cSpaceW && attr.Name.Local == cAttrThemeColor:
							rPr.Color.Theme = attr.Value
						}
					}
				case cTagCs:
					//<w:Cs/>
					rPr.ComplexScript = helper.UnmarshalToggleValToBool(t, cSpaceW)
				case cTagDStrike:
					//<w:dstrike w:val="true"/>
					rPr.DoubleStrikethrough = helper.UnmarshalToggleValToBool(t, cSpaceW)
				case cTagEmphasisMark:
					//<w:em w:val="dot"/>
					rPr.EmphasisMark = helper.UnmarshalSingleAttr(t, cSpaceW, cAttrVal)
				case cTagItalics:
					//	<w:i />
					rPr.Italics = helper.UnmarshalToggleValToBool(t, cSpaceW)
				case cTagItalicsCs:
					// <w:iCs w:val="true"/>
					rPr.ItalicsCs = helper.UnmarshalToggleValToBool(t, cSpaceW)
				case cTagImprint:
					// <w:Imprint w:val="true"/>
					rPr.Imprint = helper.UnmarshalToggleValToBool(t, cSpaceW)
				case cTagKern:
					// <w:kern w:val="28" />
					if val, err := helper.UnmarshalSingleAttrToInt(t, cSpaceW, cAttrVal); err != nil {
						return err
					} else {
						rPr.FontKerning = val
					}
				case cTagLang:
					// <w:Lang w:val="fr-CA" w:Bidi="he-IL" />
					for _, attr := range t.Attr {
						switch attr.Name.Space {
						case cSpaceW:
							switch attr.Name.Local {
							case cAttrBidi:
								rPr.Lang.Bidi = attr.Value
							case cAttrVal:
								rPr.Lang.Value = attr.Value
							case cAttrEastAsia:
								rPr.Lang.EastAsian = attr.Value
							}
						}
					}
				case cTagOutline:
					//<w:Outline w:val="false"/>
					rPr.Outline = helper.UnmarshalToggleValToBool(t, cSpaceW)
				case cTagPosition:
					// <w:Position w:val="24" />
					if val, err := helper.UnmarshalSingleValToInt(t, cSpaceW); err != nil {
						return err
					} else {
						rPr.Position = val
					}
				case cTagRFonts:
					// <w:rFonts w:Ascii="Courier New" w:Cs="Times New Roman" />
					// <w:rFonts w:Hint="EastAsia" w:Ascii="黑体" w:HAnsi="黑体" w:EastAsia="黑体" w:Cs="黑体"/>
					// <w:rFonts w:Hint="default" w:AsciiTheme="minorAscii" w:HAnsiTheme="minorAscii" w:EastAsiaTheme="minorEastAsia"/>
					for _, attr := range t.Attr {
						switch attr.Name.Space {
						case cSpaceW:
							switch attr.Name.Local {
							case cAttrHint:
								rPr.Fonts.Hint = attr.Value
							case cAttrAscii:
								rPr.Fonts.Ascii = attr.Value
							case cAttrCs:
								rPr.Fonts.Cs = attr.Value
							case cAttrEastAsia:
								rPr.Fonts.EastAsia = attr.Value
							case cAttrHAnsi:
								rPr.Fonts.HAnsi = attr.Value
							case cAttrAsciiTheme:
								rPr.Fonts.AsciiTheme = attr.Value
							case cAttrEastAsiaTheme:
								rPr.Fonts.EastAsiaTheme = attr.Value
							case cAttrHAnsiTheme:
								rPr.Fonts.HAnsiTheme = attr.Value
							}
						}

					}
				case cTagRStyle:
					// <w:rStyle w:val="14"/>
					rPr.StyleId = helper.UnmarshalSingleVal(t, cSpaceW)
				case cTagSize:
					// <w:sz w:val="27"/>
					val, err := helper.UnmarshalSingleValToInt(t, cSpaceW)
					if err != nil {
						return err
					}
					rPr.Size = val
				case cTagSizeCs:
					//<w:szCs w:val="20"/>
					val, err := helper.UnmarshalSingleValToInt(t, cSpaceW)
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
