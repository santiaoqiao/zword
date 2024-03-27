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

// docFile 是一个全局变量
var docFile = &Docx{}

// getStyle 根据样式表获取一套样式属性
func getStyle(id string) *StyleItem {
	if s, ok := docFile.Styles.StyleSheets[id]; ok {
		return &s
	}
	return nil
}
