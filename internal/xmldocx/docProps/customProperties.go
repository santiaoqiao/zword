package docx

// CustomProperties is the `Application-Defined File Properties part` => docProps/custom.xmldocx
type CustomProperties struct {
	Xmlns    string               `xmldocx:"xmlns,attr"`
	Children []CustomPropertyItem `xmldocx:"property"`
}

// CustomPropertyItem is the item of the docProps/custom.xmldocx file
type CustomPropertyItem struct {
	Fmtid  string `xmldocx:"fmtid,attr"`
	Pid    string `xmldocx:"pid,attr"`
	Name   string `xmldocx:"name,attr"`
	Lpwstr string `xmldocx:"http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes lpwstr"`
}
