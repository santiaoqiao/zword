package docx

type ContentTypes struct {
	DefaultItems []ContentTypesDefaultItem  `xml:"Default"`
	OverrideItem []ContentTypesOverrideItem `xml:"Override"`
}

type ContentTypesDefaultItem struct {
	Extension   string `xml:"Extension,attr"`
	ContentType string `xml:"ContentType,attr"`
}

type ContentTypesOverrideItem struct {
	PartName    string `xml:"PartName,attr"`
	ContentType string `xml:"ContentType,attr"`
}

// CoreProperties is the `Application-Defined File Properties part` => docProps/core.xml
type CoreProperties struct {
	Created        string `xml:"http://purl.org/dc/terms/ created"`
	Creator        string `xml:"http://purl.org/dc/elements/1.1/ creator"`
	LastModifiedBy string `xml:"http://schemas.openxmlformats.org/pkg/2006/metadata/core-properties lastModifiedBy"`
	Modified       string `xml:"http://purl.org/dc/terms/ modified"`
	Revision       int    `xml:"http://schemas.openxmlformats.org/pkg/2006/metadata/core-properties revision"`
}

// CustomProperties is the `Application-Defined File Properties part` => docProps/custom.xml
type CustomProperties struct {
	Xmlns    string               `xml:"xmlns,attr"`
	Children []CustomPropertyItem `xml:"property"`
}

// CustomPropertyItem is the item of the docProps/custom.xml file
type CustomPropertyItem struct {
	Fmtid  string `xml:"fmtid,attr"`
	Pid    string `xml:"pid,attr"`
	Name   string `xml:"name,attr"`
	Lpwstr string `xml:"http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes lpwstr"`
}

// ExtendedProperties is the `Application-Defined File Properties part` => docProps/app.xml
type ExtendedProperties struct {
	Template             string `xml:"Template"`
	Pages                int    `xml:"Pages"`
	Words                int    `xml:"Words"`
	Characters           int    `xml:"Characters"`
	Lines                int    `xml:"Lines"`
	Paragraphs           int    `xml:"Paragraphs"`
	TotalTime            int    `xml:"TotalTime"`
	ScaleCrop            bool   `xml:"ScaleCrop"`
	LinksUpToDate        bool   `xml:"LinksUpToDate"`
	CharactersWithSpaces int    `xml:"CharactersWithSpaces"`
	Application          string `xml:"Application"`
	DocSecurity          int    `xml:"DocSecurity"`
}

// PackageRelationships => ./rels/.rels
type PackageRelationships struct {
	Xmlns    string                    `xml:"xmlns,attr"`
	Children []PackageRelationshipItem `xml:"PackageRelationshipItem"`
}

type PackageRelationshipItem struct {
	Id     string `xml:"Id,attr"`
	Type   string `xml:"TypeName,attr"`
	Target string `xml:"Target,attr"`
}

type PartRelationship struct {
}
