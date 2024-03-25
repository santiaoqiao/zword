package zword

// Property 文档属性
type Property struct {
}

type Document struct {
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
	//Styles              *Styles
	//Footer              *Footer
	Property *Property
	Body     *Body `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main body"`
}
