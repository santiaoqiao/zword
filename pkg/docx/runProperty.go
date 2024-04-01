package docx

import (
	"encoding/xml"
	"fmt"
	"github.com/santiaoqiao/zword/pkg/docx/helper"
	"io"
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
	bold *bool
	// 粗体（复杂脚本）
	boldCs *bool
	// 字体颜色
	color *Color
	// 是否为标记 Complex Script
	complexScript *bool
	// 双横线穿过
	doubleStrikethrough *bool
	// 强调 <w:em w:val="dot"/>
	emphasisMark *string
	// 斜体（简单文字）
	italics *bool
	// 斜体（复杂脚本）
	italicsCs *bool
	// 浮雕
	imprint *bool
	// 字符字距
	fontKerning *int
	// 拼写检查的语言
	lang *Language
	// 外轮廓
	outline *bool
	// 文字在垂直方向上上下偏移的距离
	position *int
	// 字体
	fonts *RunFonts
	// 样式ID
	rStyleId string
	// 字体大小（简单文字）
	size *int
	// 字体大小（复杂脚本）
	sizeCs *int
	// 指向样式表中的指针
	runStyleRPr       *RunProperty
	paragraphStyleRPr *RunProperty
	//numberingStyle *StyleItem
	//tableStyle     *StyleItem
	// 指向父亲的指针
	parentRun *Run
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
					rPr.boldCs = helper.UnmarshalToggleValToBool(t, helper.CSpaceW)
				case "color":
					//<w:color w:themeColor="accent3"  w:val="FF0000"/>
					color := &Color{}
					for _, attr := range t.Attr {
						switch {
						case attr.Name.Space == helper.CSpaceW && attr.Name.Local == "val":
							color.Value = attr.Value
						case attr.Name.Space == helper.CSpaceW && attr.Name.Local == "themeColor":
							color.Theme = attr.Value
						}
					}
					rPr.color = color
				case "cs":
					//<w:cs/>
					rPr.complexScript = helper.UnmarshalToggleValToBool(t, helper.CSpaceW)
				case "dstrike":
					//<w:dstrike w:val="true"/>
					rPr.doubleStrikethrough = helper.UnmarshalToggleValToBool(t, helper.CSpaceW)
				case "em":
					//<w:em w:val="dot"/>
					val := helper.UnmarshalSingleVal(t, helper.CSpaceW)
					rPr.emphasisMark = &val
				case "i":
					//	<w:i />
					rPr.italics = helper.UnmarshalToggleValToBool(t, helper.CSpaceW)
				case "iCs":
					// <w:iCs w:val="true"/>
					rPr.italicsCs = helper.UnmarshalToggleValToBool(t, helper.CSpaceW)
				case "imprint":
					// <w:imprint w:val="true"/>
					rPr.imprint = helper.UnmarshalToggleValToBool(t, helper.CSpaceW)
				case "kern":
					// <w:kern w:val="28" />
					if val, err := helper.UnmarshalSingleValToInt(t, helper.CSpaceW); err != nil {
						return err
					} else {
						rPr.fontKerning = &val
					}
				case "lang":
					// <w:lang w:val="fr-CA" w:bidi="he-IL" />
					lang := &Language{}
					for _, attr := range t.Attr {
						switch attr.Name.Space {
						case helper.CSpaceW:
							switch attr.Name.Local {
							case "bidi":
								lang.Bidi = attr.Value
							case "val":
								lang.Value = attr.Value
							case "eastAsia":
								lang.EastAsian = attr.Value
							}
						}
					}
					rPr.lang = lang
				case "outline":
					//<w:outline w:val="false"/>
					rPr.outline = helper.UnmarshalToggleValToBool(t, helper.CSpaceW)
				case "position":
					// <w:position w:val="24" />
					if val, err := helper.UnmarshalSingleValToInt(t, helper.CSpaceW); err != nil {
						return err
					} else {
						rPr.position = &val
					}
				case "rFonts":
					// <w:rFonts w:Ascii="Courier New" w:Cs="Times New Roman" />
					// <w:rFonts w:Hint="EastAsia" w:Ascii="黑体" w:HAnsi="黑体" w:EastAsia="黑体" w:Cs="黑体"/>
					// <w:rFonts w:Hint="default" w:AsciiTheme="minorAscii" w:HAnsiTheme="minorAscii" w:EastAsiaTheme="minorEastAsia"/>
					fonts := &RunFonts{}
					for _, attr := range t.Attr {
						switch attr.Name.Space {
						case helper.CSpaceW:
							switch attr.Name.Local {
							case "hint":
								fonts.Hint = attr.Value
							case "ascii":
								fonts.Ascii = attr.Value
							case "cs":
								fonts.Cs = attr.Value
							case "eastAsia":
								fonts.EastAsia = attr.Value
							case "hAnsi":
								fonts.HAnsi = attr.Value
							case "asciiTheme":
								fonts.AsciiTheme = attr.Value
							case "eastAsiaTheme":
								fonts.EastAsiaTheme = attr.Value
							case "hAnsiTheme":
								fonts.HAnsiTheme = attr.Value
							}
						}

					}
					rPr.fonts = fonts
				case "rStyle":
					// <w:rStyle w:val="14"/>
					rPr.rStyleId = helper.UnmarshalSingleVal(t, helper.CSpaceW)
					// 获取 样式
					style := getStyle(rPr.rStyleId)
					if style != nil {
						rPr.runStyleRPr = style.RPr
					}
				case "sz":
					// <w:sz w:val="27"/>
					val, err := helper.UnmarshalSingleValToInt(t, helper.CSpaceW)
					if err != nil {
						return err
					}
					rPr.size = &val
				case "szCs":
					//<w:szCs w:val="20"/>
					val, err := helper.UnmarshalSingleValToInt(t, helper.CSpaceW)
					if err != nil {
						return err
					}
					rPr.sizeCs = &val
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

func (rPr *RunProperty) Bold() *bool {
	// 1 先看 Direct formatting
	if rPr.bold != nil {
		return rPr.bold
	}
	// 2 Character styles
	if rPr.runStyleRPr != nil && rPr.runStyleRPr.bold != nil {
		return rPr.runStyleRPr.bold
	}
	//3 Paragraph style
	if rPr.paragraphStyleRPr != nil && rPr.paragraphStyleRPr.bold != nil {
		return rPr.paragraphStyleRPr.bold
	}
	//4 Numbering styles
	//5 Table styles
	//6 Document defaults
	if docFile.Styles.DocDefaults.RPrDefault != nil && docFile.Styles.DocDefaults.RPrDefault.bold != nil {
		return docFile.Styles.DocDefaults.RPrDefault.bold
	}
	return nil
}

func (rPr *RunProperty) BoldCs() *bool {
	// 1 先看 Direct formatting
	if rPr.boldCs != nil {
		return rPr.boldCs
	}
	// 2 Character styles
	if rPr.runStyleRPr != nil && rPr.runStyleRPr.boldCs != nil {
		return rPr.runStyleRPr.boldCs
	}
	//3 Paragraph style
	if rPr.paragraphStyleRPr != nil && rPr.paragraphStyleRPr.boldCs != nil {
		return rPr.paragraphStyleRPr.boldCs
	}
	//4 Numbering styles
	//5 Table styles
	//6 Document defaults
	if docFile.Styles.DocDefaults.RPrDefault != nil && docFile.Styles.DocDefaults.RPrDefault.boldCs != nil {
		return docFile.Styles.DocDefaults.RPrDefault.boldCs
	}
	return nil
}

func (rPr *RunProperty) Color() *Color {
	// 1 先看 Direct formatting
	if rPr.color != nil {
		return rPr.color
	}
	// 2 Character styles
	if rPr.runStyleRPr != nil && rPr.runStyleRPr.color != nil {
		return rPr.runStyleRPr.color
	}
	//3 Paragraph style
	if rPr.paragraphStyleRPr != nil && rPr.paragraphStyleRPr.color != nil {
		return rPr.paragraphStyleRPr.color
	}
	//4 Numbering styles
	//5 Table styles
	//6 Document defaults
	if docFile.Styles.DocDefaults.RPrDefault != nil && docFile.Styles.DocDefaults.RPrDefault.color != nil {
		return docFile.Styles.DocDefaults.RPrDefault.color
	}
	return nil
}

func (rPr *RunProperty) ComplexScript() *bool {
	// 1 先看 Direct formatting
	if rPr.complexScript != nil {
		return rPr.complexScript
	}
	// 2 Character styles
	if rPr.runStyleRPr != nil && rPr.runStyleRPr.complexScript != nil {
		return rPr.runStyleRPr.complexScript
	}
	//3 Paragraph style
	if rPr.paragraphStyleRPr != nil && rPr.paragraphStyleRPr.complexScript != nil {
		return rPr.paragraphStyleRPr.complexScript
	}
	//4 Numbering styles
	//5 Table styles
	//6 Document defaults
	if docFile.Styles.DocDefaults.RPrDefault != nil && docFile.Styles.DocDefaults.RPrDefault.complexScript != nil {
		return docFile.Styles.DocDefaults.RPrDefault.complexScript
	}
	return nil
}

func (rPr *RunProperty) DoubleStrikethrough() *bool {
	// 1 先看 Direct formatting
	if rPr.doubleStrikethrough != nil {
		return rPr.doubleStrikethrough
	}
	// 2 Character styles
	if rPr.runStyleRPr != nil && rPr.runStyleRPr.doubleStrikethrough != nil {
		return rPr.runStyleRPr.doubleStrikethrough
	}
	//3 Paragraph style
	if rPr.paragraphStyleRPr != nil && rPr.paragraphStyleRPr.doubleStrikethrough != nil {
		return rPr.paragraphStyleRPr.doubleStrikethrough
	}
	//4 Numbering styles
	//5 Table styles
	//6 Document defaults
	if docFile.Styles.DocDefaults.RPrDefault != nil && docFile.Styles.DocDefaults.RPrDefault.doubleStrikethrough != nil {
		return docFile.Styles.DocDefaults.RPrDefault.doubleStrikethrough
	}
	return nil
}

func (rPr *RunProperty) EmphasisMark() *string {
	// 1 先看 Direct formatting
	if rPr.emphasisMark != nil {
		return rPr.emphasisMark
	}
	// 2 Character styles
	if rPr.runStyleRPr != nil && rPr.runStyleRPr.emphasisMark != nil {
		return rPr.runStyleRPr.emphasisMark
	}
	//3 Paragraph style
	if rPr.paragraphStyleRPr != nil && rPr.paragraphStyleRPr.emphasisMark != nil {
		return rPr.paragraphStyleRPr.emphasisMark
	}
	//4 Numbering styles
	//5 Table styles
	//6 Document defaults
	if docFile.Styles.DocDefaults.RPrDefault != nil && docFile.Styles.DocDefaults.RPrDefault.emphasisMark != nil {
		return docFile.Styles.DocDefaults.RPrDefault.emphasisMark
	}
	return nil
}

func (rPr *RunProperty) Italics() *bool {
	// 1 先看 Direct formatting
	if rPr.italics != nil {
		return rPr.italics
	}
	// 2 Character styles
	if rPr.runStyleRPr != nil && rPr.runStyleRPr.italics != nil {
		return rPr.runStyleRPr.italics
	}
	//3 Paragraph style
	if rPr.paragraphStyleRPr != nil && rPr.paragraphStyleRPr.italics != nil {
		return rPr.paragraphStyleRPr.italics
	}
	//4 Numbering styles
	//5 Table styles
	//6 Document defaults
	if docFile.Styles.DocDefaults.RPrDefault != nil && docFile.Styles.DocDefaults.RPrDefault.italics != nil {
		return docFile.Styles.DocDefaults.RPrDefault.italics
	}
	return nil
}

func (rPr *RunProperty) ItalicsCs() *bool {
	// 1 先看 Direct formatting
	if rPr.italicsCs != nil {
		return rPr.italicsCs
	}
	// 2 Character styles
	if rPr.runStyleRPr != nil && rPr.runStyleRPr.italicsCs != nil {
		return rPr.runStyleRPr.italicsCs
	}
	//3 Paragraph style
	if rPr.paragraphStyleRPr != nil && rPr.paragraphStyleRPr.italicsCs != nil {
		return rPr.paragraphStyleRPr.italicsCs
	}
	//4 Numbering styles
	//5 Table styles
	//6 Document defaults
	if docFile.Styles.DocDefaults.RPrDefault != nil && docFile.Styles.DocDefaults.RPrDefault.italicsCs != nil {
		return docFile.Styles.DocDefaults.RPrDefault.italicsCs
	}
	return nil
}

func (rPr *RunProperty) Imprint() *bool {
	// 1 先看 Direct formatting
	if rPr.imprint != nil {
		return rPr.imprint
	}
	// 2 Character styles
	if rPr.runStyleRPr != nil && rPr.runStyleRPr.imprint != nil {
		return rPr.runStyleRPr.imprint
	}
	//3 Paragraph style
	if rPr.paragraphStyleRPr != nil && rPr.paragraphStyleRPr.imprint != nil {
		return rPr.paragraphStyleRPr.imprint
	}
	//4 Numbering styles
	//5 Table styles
	//6 Document defaults
	if docFile.Styles.DocDefaults.RPrDefault != nil && docFile.Styles.DocDefaults.RPrDefault.imprint != nil {
		return docFile.Styles.DocDefaults.RPrDefault.imprint
	}
	return nil
}

func (rPr *RunProperty) FontKerning() *int {
	// 1 先看 Direct formatting
	if rPr.fontKerning != nil {
		return rPr.fontKerning
	}
	// 2 Character styles
	if rPr.runStyleRPr != nil && rPr.runStyleRPr.fontKerning != nil {
		return rPr.runStyleRPr.fontKerning
	}
	//3 Paragraph style
	if rPr.paragraphStyleRPr != nil && rPr.paragraphStyleRPr.fontKerning != nil {
		return rPr.paragraphStyleRPr.fontKerning
	}
	//4 Numbering styles
	//5 Table styles
	//6 Document defaults
	if docFile.Styles.DocDefaults.RPrDefault != nil && docFile.Styles.DocDefaults.RPrDefault.fontKerning != nil {
		return docFile.Styles.DocDefaults.RPrDefault.fontKerning
	}
	return nil
}

func (rPr *RunProperty) Lang() *Language {
	// 1 先看 Direct formatting
	if rPr.lang != nil {
		return rPr.lang
	}
	// 2 Character styles
	if rPr.runStyleRPr != nil && rPr.runStyleRPr.lang != nil {
		return rPr.runStyleRPr.lang
	}
	//3 Paragraph style
	if rPr.paragraphStyleRPr != nil && rPr.paragraphStyleRPr.lang != nil {
		return rPr.paragraphStyleRPr.lang
	}
	//4 Numbering styles
	//5 Table styles
	//6 Document defaults
	if docFile.Styles.DocDefaults.RPrDefault != nil && docFile.Styles.DocDefaults.RPrDefault.lang != nil {
		return docFile.Styles.DocDefaults.RPrDefault.lang
	}
	return nil
}

func (rPr *RunProperty) Outline() *bool {
	// 1 先看 Direct formatting
	if rPr.outline != nil {
		return rPr.outline
	}
	// 2 Character styles
	if rPr.runStyleRPr != nil && rPr.runStyleRPr.outline != nil {
		return rPr.runStyleRPr.outline
	}
	//3 Paragraph style
	if rPr.paragraphStyleRPr != nil && rPr.paragraphStyleRPr.outline != nil {
		return rPr.paragraphStyleRPr.outline
	}
	//4 Numbering styles
	//5 Table styles
	//6 Document defaults
	if docFile.Styles.DocDefaults.RPrDefault != nil && docFile.Styles.DocDefaults.RPrDefault.outline != nil {
		return docFile.Styles.DocDefaults.RPrDefault.outline
	}
	return nil
}

func (rPr *RunProperty) Size() *int {
	// 1 先看 Direct formatting
	if rPr.size != nil {
		return rPr.size
	}
	// 2 Character styles
	if rPr.runStyleRPr != nil && rPr.runStyleRPr.size != nil {
		return rPr.runStyleRPr.size
	}
	//3 Paragraph style
	if rPr.paragraphStyleRPr != nil && rPr.paragraphStyleRPr.size != nil {
		return rPr.paragraphStyleRPr.size
	}
	//4 Numbering styles
	//5 Table styles
	//6 Document defaults
	if docFile.Styles.DocDefaults.RPrDefault != nil && docFile.Styles.DocDefaults.RPrDefault.size != nil {
		return docFile.Styles.DocDefaults.RPrDefault.size
	}
	return nil
}

func (rPr *RunProperty) SizeCs() *int {
	// 1 先看 Direct formatting
	if rPr.sizeCs != nil {
		return rPr.sizeCs
	}
	// 2 Character styles
	if rPr.runStyleRPr != nil && rPr.runStyleRPr.sizeCs != nil {
		return rPr.runStyleRPr.sizeCs
	}
	//3 Paragraph style
	if rPr.paragraphStyleRPr != nil && rPr.paragraphStyleRPr.sizeCs != nil {
		return rPr.paragraphStyleRPr.sizeCs
	}
	//4 Numbering styles
	//5 Table styles
	//6 Document defaults
	if docFile.Styles.DocDefaults.RPrDefault != nil && docFile.Styles.DocDefaults.RPrDefault.sizeCs != nil {
		return docFile.Styles.DocDefaults.RPrDefault.sizeCs
	}
	return nil
}

func (rPr *RunProperty) Fonts() *RunFonts {
	// 1 先看 Direct formatting
	if rPr.fonts != nil {
		// 只有hint的情况
		if rPr.fonts.Hint != "" &&
			rPr.fonts.EastAsia == "" &&
			rPr.fonts.Cs == "" &&
			rPr.fonts.Ascii == "" &&
			rPr.fonts.HAnsi == "" {
		} else {
			return rPr.fonts
		}
	}
	// 2 Character styles
	if rPr.runStyleRPr != nil && rPr.runStyleRPr.fonts != nil {
		return rPr.runStyleRPr.fonts
	}
	//3 Paragraph style
	if rPr.paragraphStyleRPr != nil && rPr.paragraphStyleRPr.fonts != nil {
		return rPr.paragraphStyleRPr.fonts
	}
	//4 Numbering styles
	//5 Table styles
	//6 Document defaults
	if docFile.Styles.DocDefaults.RPrDefault != nil && docFile.Styles.DocDefaults.RPrDefault.fonts != nil {
		return docFile.Styles.DocDefaults.RPrDefault.fonts
	}
	return nil
}

func (rPr *RunProperty) Position() *int {
	// 1 先看 Direct formatting
	if rPr.position != nil {
		return rPr.position
	}
	// 2 Character styles
	if rPr.runStyleRPr != nil && rPr.runStyleRPr.position != nil {
		return rPr.runStyleRPr.position
	}
	//3 Paragraph style
	if rPr.paragraphStyleRPr != nil && rPr.paragraphStyleRPr.position != nil {
		return rPr.paragraphStyleRPr.position
	}
	//4 Numbering styles
	//5 Table styles
	//6 Document defaults
	if docFile.Styles.DocDefaults.RPrDefault != nil && docFile.Styles.DocDefaults.RPrDefault.position != nil {
		return docFile.Styles.DocDefaults.RPrDefault.position
	}
	return nil
}
