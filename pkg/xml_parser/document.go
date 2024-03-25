package xml_parser

import (
	"santiaoqiao.com/zword/pkg/xml_parser/properties"
	"santiaoqiao.com/zword/pkg/xml_parser/stroies"
)

// Property 文档属性
type Property struct {
}

type Document struct {
	PackageRelationship *properties.PackageRelationshipItem
	CoreProperties      *properties.CoreProperties
	CustomProperties    *properties.CustomProperties
	ExtendedProperties  *properties.ExtendedProperties
	ContentTypes        *properties.ContentTypes
	PartRelationship    *properties.PartRelationship
	//FontTable           *FontTable
	//Header              *Header
	//Numbering           *Numbering
	//Settings            *Settings
	//Styles              *Styles
	//Footer              *Footer
	Property *Property
	Body     *stroies.Body `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main body"`
}
