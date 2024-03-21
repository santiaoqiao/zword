package xmldocx

type ContentTypes struct {
	DefaultItems []ContentTypesDefaultItem  `xmldocx:"Default"`
	OverrideItem []ContentTypesOverrideItem `xmldocx:"Override"`
}

type ContentTypesDefaultItem struct {
	Extension   string `xmldocx:"Extension,attr"`
	ContentType string `xmldocx:"ContentType,attr"`
}

type ContentTypesOverrideItem struct {
	PartName    string `xmldocx:"PartName,attr"`
	ContentType string `xmldocx:"ContentType,attr"`
}
