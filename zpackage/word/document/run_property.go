package document

import (
	"encoding/xml"
	"fmt"
	"io"
	"santiaoqiao.com/zoffice/zpackage/helper"
	"strings"
)

/*
This element specifies a set of run properties which shall be applied to the contents of the parent run after all
style formatting has been applied to the text. These properties are defined as direct formatting, since they are
directly applied to the run and supersede any formatting from styles

This formatting is applied at the following location in the style hierarchy:
 Document defaults
 Table styles
 Numbering styles
 Paragraph styles
 Character styles
 Direct formatting (this element)
*/

type RunProperty struct {
	/*Bold  non complex script*/
	Bold bool
	/*Complex Script Bold*/
	BoldCs bool
	/*font color
	If the themeColor attribute is specified,
	then the val attribute is ignored for this run.*/
	Color Color
	/*标记是否为 Complex Script
	Use Complex Script Formatting on Run*/
	ComplexScript       bool
	DoubleStrikethrough bool
	EmphasisMark        string
	Italics             bool
	ItalicsCs           bool
	Imprint             bool
	// 字符字距
	Kern    int
	Lang    Language
	Outline bool
	// Vertically Raised or Lowered Text
	Position int
	// Run Fonts
	RFonts   RunFonts
	RStyleId string
	// sz (Non-Complex Script Font Size)
	Sz   int
	SzCs int
}

func (rPr *RunProperty) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// 解析xml并给 rPr 赋值
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
			switch t.Name.Local {
			case "b":
				// <w:b w:val="false"/> | <w:b "/>
				rPr.Bold = helper.UnwrapValToToggle(t)
			case "bCs":
				// <w:bCs w:val="false"/> | <w:bCs />
				rPr.BoldCs = helper.UnwrapValToToggle(t)
			case "color":
				//<w:color w:themeColor="accent3" />
				if val, ok := helper.Unwrap(t, "themeColor"); ok {
					rPr.Color.Theme = val
				}
				// <w:color w:val="FF0000"/>
				if val, ok := helper.Unwrap(t, "val"); ok {
					rPr.Color.Value = val
				}
			case "cs":
				rPr.ComplexScript = true
			case "dstrike":
				rPr.DoubleStrikethrough = helper.UnwrapValToToggle(t)
			case "em":
				if value, ok := helper.Unwrap(t, "em"); ok {
					rPr.EmphasisMark = value
				}
			case "i":
				rPr.Italics = helper.UnwrapValToToggle(t)
			case "iCs":
				rPr.ItalicsCs = helper.UnwrapValToToggle(t)
			case "imprint":
				rPr.Imprint = helper.UnwrapValToToggle(t)
			case "kern":
				if value, _, err := helper.UnwrapValToInt(t); err != nil {
					return err
				} else {
					rPr.Kern = value
				}
			case "lang":
				for _, attr := range t.Attr {
					switch attr.Name.Local {
					case "bidi":
						rPr.Lang.Bidi = attr.Value
					case "val":
						rPr.Lang.Value = attr.Value
					case "eastAsia":
						rPr.Lang.EastAsian = attr.Value
					}
				}
			case "outline":
				rPr.Outline = helper.UnwrapValToToggle(t)
			case "position":
				val, _, err := helper.UnwrapValToInt(t)
				if err != nil {
					return err
				}
				rPr.Position = val
			case "rFonts":
				if val, ok := helper.Unwrap(t, "hint"); ok {
					rPr.RFonts.Hint = val
				} else if val, ok = helper.Unwrap(t, "ascii"); ok {
					rPr.RFonts.Ascii = val
				} else if val, ok = helper.Unwrap(t, "hAnsi"); ok {
					rPr.RFonts.HAnsi = val
				} else if val, ok = helper.Unwrap(t, "eastAsia"); ok {
					rPr.RFonts.EastAsia = val
				} else if val, ok = helper.Unwrap(t, "cs"); ok {
					rPr.RFonts.Cs = val
				} else if val, ok = helper.Unwrap(t, "asciiTheme"); ok {
					rPr.RFonts.AsciiTheme = val
				} else if val, ok = helper.Unwrap(t, "hAnsiTheme"); ok {
					rPr.RFonts.HAnsiTheme = val
				} else if val, ok = helper.Unwrap(t, "eastAsiaTheme"); ok {
					rPr.RFonts.EastAsiaTheme = val
				}
			case "rStyle":
				if val, ok := helper.UnwrapVal(t); ok {
					rPr.RStyleId = val
				}
			case "sz":
				if val, _, err := helper.UnwrapValToInt(t); err != nil {
					return err
				} else {
					rPr.Sz = val
				}
			case "szCs":
				if val, _, err := helper.UnwrapValToInt(t); err != nil {
					return err
				} else {
					rPr.SzCs = val
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

func (rPr *RunProperty) String() string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf("%#v", rPr))
	return sb.String()
}

type Color struct {
	Value string
	Theme string
}
type Language struct {
	Value     string
	Bidi      string
	EastAsian string
}

// RunFonts 最多有4种字体槽
type RunFonts struct {
	// 提示
	Hint string
	// the first 128 Unicode code points
	Ascii string
	// High ANSI
	HAnsi string
	// East Asian
	EastAsia string
	// Complex Script
	Cs            string
	AsciiTheme    string
	HAnsiTheme    string
	EastAsiaTheme string
}
