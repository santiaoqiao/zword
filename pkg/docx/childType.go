package docx

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
	TypeName() RunChildType
}

type ParagraphChild interface {
	String() string
	TypeName() ParagraphChildType
}

// BodyChild <body>中的元素：p、tbl、sectPr
type BodyChild interface {
	String() string
	TypeName() BodyChildType
}
