package zpackage

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
