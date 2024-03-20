package docx

import (
	"encoding/xml"
	"fmt"
	"io"
	"santiaoqiao.com/zword/internal/docx/helper"
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
	bold bool
	// 粗体（复杂脚本）
	boldCs bool
	// 字体颜色
	color Color
	// 是否为标记 Complex Script
	complexScript bool
	// 双横线穿过
	doubleStrikethrough bool
	// 强调 <w:em w:val="dot"/>
	emphasisMark string
	// 斜体（简单文字）
	italics bool
	// 斜体（复杂脚本）
	italicsCs bool
	// 浮雕
	imprint bool
	// 字符字距
	fontKerning int
	// 拼写检查的语言
	lang Language
	// 外轮廓
	outline bool
	// 文字在垂直方向上上下偏移的距离
	position int
	// 字体
	fonts RunFonts
	// 样式ID
	styleId string
	// 字体大小（简单文字）
	size int
	// 字体大小（复杂脚本）
	sizeCs int
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
					rPr.bold = helper.UnmarshalToggleValToBool(t, cSpaceW)
				case cTagBoldCs:
					// <w:bCs w:val="false"/> | <w:bCs />
					rPr.boldCs = helper.UnmarshalToggleValToBool(t, cSpaceW)
				case cTagColor:
					//<w:color w:themeColor="accent3"  w:val="FF0000"/>
					for _, attr := range t.Attr {
						switch {
						case attr.Name.Space == cSpaceW && attr.Name.Local == cAttrVal:
							rPr.color.value = attr.Value
						case attr.Name.Space == cSpaceW && attr.Name.Local == cAttrThemeColor:
							rPr.color.theme = attr.Value
						}
					}
				case cTagCs:
					//<w:cs/>
					rPr.complexScript = helper.UnmarshalToggleValToBool(t, cSpaceW)
				case cTagDStrike:
					//<w:dstrike w:val="true"/>
					rPr.doubleStrikethrough = helper.UnmarshalToggleValToBool(t, cSpaceW)
				case cTagEmphasisMark:
					//<w:em w:val="dot"/>
					rPr.emphasisMark = helper.UnmarshalSingleAttr(t, cSpaceW, cAttrVal)
				case cTagItalics:
					//	<w:i />
					rPr.italics = helper.UnmarshalToggleValToBool(t, cSpaceW)
				case cTagItalicsCs:
					// <w:iCs w:val="true"/>
					rPr.italicsCs = helper.UnmarshalToggleValToBool(t, cSpaceW)
				case cTagImprint:
					// <w:imprint w:val="true"/>
					rPr.imprint = helper.UnmarshalToggleValToBool(t, cSpaceW)
				case cTagKern:
					// <w:kern w:val="28" />
					if val, err := helper.UnmarshalSingleAttrToInt(t, cSpaceW, cAttrVal); err != nil {
						return err
					} else {
						rPr.fontKerning = val
					}
				case cTagLang:
					// <w:lang w:val="fr-CA" w:bidi="he-IL" />
					for _, attr := range t.Attr {
						switch attr.Name.Space {
						case cSpaceW:
							switch attr.Name.Local {
							case cAttrBidi:
								rPr.lang.bidi = attr.Value
							case cAttrVal:
								rPr.lang.value = attr.Value
							case cAttrEastAsia:
								rPr.lang.eastAsian = attr.Value
							}
						}
					}
				case cTagOutline:
					//<w:outline w:val="false"/>
					rPr.outline = helper.UnmarshalToggleValToBool(t, cSpaceW)
				case cTagPosition:
					// <w:position w:val="24" />
					if val, err := helper.UnmarshalSingleValToInt(t, cSpaceW); err != nil {
						return err
					} else {
						rPr.position = val
					}
				case cTagRFonts:
					// <w:rFonts w:ascii="Courier New" w:cs="Times New Roman" />
					// <w:rFonts w:hint="eastAsia" w:ascii="黑体" w:hAnsi="黑体" w:eastAsia="黑体" w:cs="黑体"/>
					// <w:rFonts w:hint="default" w:asciiTheme="minorAscii" w:hAnsiTheme="minorAscii" w:eastAsiaTheme="minorEastAsia"/>
					for _, attr := range t.Attr {
						switch attr.Name.Space {
						case cSpaceW:
							switch attr.Name.Local {
							case cAttrHint:
								rPr.fonts.hint = attr.Value
							case cAttrAscii:
								rPr.fonts.ascii = attr.Value
							case cAttrCs:
								rPr.fonts.cs = attr.Value
							case cAttrEastAsia:
								rPr.fonts.eastAsia = attr.Value
							case cAttrHAnsi:
								rPr.fonts.hAnsi = attr.Value
							case cAttrAsciiTheme:
								rPr.fonts.asciiTheme = attr.Value
							case cAttrEastAsiaTheme:
								rPr.fonts.eastAsiaTheme = attr.Value
							case cAttrHAnsiTheme:
								rPr.fonts.hAnsiTheme = attr.Value
							}
						}

					}
				case cTagRStyle:
					// <w:rStyle w:val="14"/>
					rPr.styleId = helper.UnmarshalSingleVal(t, cSpaceW)
				case cTagSize:
					// <w:sz w:val="27"/>
					val, err := helper.UnmarshalSingleValToInt(t, cSpaceW)
					if err != nil {
						return err
					}
					rPr.size = val
				case cTagSizeCs:
					//<w:szCs w:val="20"/>
					val, err := helper.UnmarshalSingleValToInt(t, cSpaceW)
					if err != nil {
						return err
					}
					rPr.sizeCs = val
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

func (rPr *RunProperty) GetBold() bool {
	return rPr.bold
}

func (rPr *RunProperty) GetFont() RunFonts {
	return rPr.fonts
}

func (rPr *RunProperty) GetFontSize() int {
	return rPr.size
}

// Color 字体颜色
type Color struct {
	// 字体颜色值，如 D4F4F2，前面不带#号
	value string
	// 字体的主题颜色，如运用了主题，以主题为主
	theme string
}

// Language 字体语言
type Language struct {
	// 指定在处理使用拉丁字符的运行内容时(由运行内容的Unicode字符值决定)应用于检查拼写和语法(如果请求)的语言
	value string
	// 指定在处理使用复杂脚本字符的运行内容时应使用的语言，由运行内容的Unicode字符值决定。
	bidi string
	// 指定在处理使用东亚字符的运行内容时应使用的语言
	eastAsian string
}

// RunFonts 最多有4种字体槽
type RunFonts struct {
	// 默认提示所用的子图
	hint string
	// 处理Ascii字符时所使用的字体
	ascii string
	// 处理 High ANSI 字符时所使用的字体
	hAnsi string
	// 处理东南亚 East Asian 文字所使用的字体，包括中文
	eastAsia string
	// 处理 Complex Script 字符时所使用的字体
	cs string
	// Ascii字符所使用的主题
	asciiTheme string
	// High ANSI字符所使用的主题
	hAnsiTheme string
	// 东南亚文字所使用的主题
	eastAsiaTheme string
}
