package word

import (
	docx2 "santiaoqiao.com/zword/internal/xmldocx"
	docx "santiaoqiao.com/zword/internal/xmldocx/docProps"
	rels2 "santiaoqiao.com/zword/internal/xmldocx/rels"
	"santiaoqiao.com/zword/internal/xmldocx/word/rels"
	"santiaoqiao.com/zword/internal/xmldocx/word/stroies"
)

// Property 文档属性
type Property struct {
}

type Document struct {
	PackageRelationship *rels2.PackageRelationshipItem
	CoreProperties      *docx.CoreProperties
	CustomProperties    *docx.CustomProperties
	ExtendedProperties  *docx.ExtendedProperties
	ContentTypes        *docx2.ContentTypes
	PartRelationship    *rels.PartRelationship
	FontTable           *FontTable
	Header              *Header
	Numbering           *Numbering
	Settings            *Settings
	Styles              *Styles
	Footer              *Footer
	Property            *Property
	Body                *stroies.Body `xmldocx:"http://schemas.openxmlformats.org/wordprocessingml/2006/main body"`
}
