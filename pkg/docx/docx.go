package docx

// Document 代表了word整个文档
type Docx struct {
	PackageRelationship *PackageRelationshipItem
	CoreProperties      *CoreProperties
	CustomProperties    *CustomProperties
	ExtendedProperties  *ExtendedProperties
	ContentTypes        *ContentTypes
	PartRelationship    *PartRelationship
	//FontTable           *FontTable
	//Header              *Header
	//Numbering           *Numbering
	//Settings            *Settings
	Styles *Styles
	//Footer              *Footer
	Document *Document
}

// Document 代表了word整个文档
type Document struct {
	Body *Body `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main body"`
}

type BodyChildType uint8

const (
	BodyTypeParagraph BodyChildType = iota
	BodyTypeTable
	BodyTypeSecPr
)

type ParagraphChildType uint8

const (
	ParagraphTypeRun ParagraphChildType = iota
	ParagraphTypePPr
)

type RunChildType uint8

const (
	RunTypeText RunChildType = iota
	RunTypeRpr
	RunTypeBreak
)

type RunChild interface {
	String() string
}

type ParagraphChild interface {
	String() string
}

// BodyChild <body>中的元素：p、tbl、sectPr
type BodyChild interface {
	String() string
	TypeName() BodyChildType
}

// xml中的 namespace
const (
	cSpaceW string = "http://schemas.openxmlformats.org/wordprocessingml/2006/main"
)

// xml 中的 tag标签
const (
	// RunProperty tags
	cTagBold         string = "b"
	cTagBoldCs              = "bCs"
	cTagColor               = "Color"
	cTagCs                  = "Cs"
	cTagDStrike             = "dstrike"
	cTagEmphasisMark        = "em"
	cTagItalics             = "i"
	cTagItalicsCs           = "iCs"
	cTagImprint             = "Imprint"
	cTagKern                = "kern"
	cTagLang                = "Lang"
	cTagOutline             = "Outline"
	cTagPosition            = "Position"
	cTagRFonts              = "rFonts"
	cTagRStyle              = "rStyle"
	cTagSize                = "sz"
	cTagSizeCs              = "szCs"
	// Paragraph tags
	cTagBidi           = "Bidi"
	cTagJustify        = "jc"
	cTagInd            = "ind"
	cTagEnd            = "end"
	cTagEndChars       = "endChars"
	cTagFirstLine      = "firstLine"
	cTagFirstLineChars = "firstLineChars"
	cTagHanging        = "hanging"
	cTagHangingChars   = "hangingChars"
	cTagStart          = "start"
	cTagStartChars     = "startChars"
	cTagRPr            = "rPr"
	cTagOutlineLvl     = "outlineLvl"
	cTagPP             = "pPr"
	cTagR              = "r"
	// Body tags
	cTagP      = "p"
	cTagTbl    = "tbl"
	cTagSectPr = "sectPr"

	// styles
	cTagStyles      = "styles"
	cTagdocDefaults = "docDefaults"
	cTagrPrDefault  = "rPrDefault"
	cTagpPrDefault  = "pPrDefault"
)

// xml 中的 attr 属性
const (
	cAttrVal           = "val"
	cAttrThemeColor    = "themeColor"
	cAttrBidi          = "bidi"
	cAttrEastAsia      = "eastAsia"
	cAttrAscii         = "ascii"
	cAttrCs            = "cs"
	cAttrHAnsi         = "hAnsi"
	cAttrHint          = "hint"
	cAttrAsciiTheme    = "asciiTheme"
	cAttrHAnsiTheme    = "hAnsiTheme"
	cAttrEastAsiaTheme = "eastAsiaTheme"
)
