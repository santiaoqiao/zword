package docx

// Docx 代表了word整个文档
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
