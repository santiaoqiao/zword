package docprops

// `Application-Defined File Properties part` => docProps/app.xml
type CustomProperties struct {
	Xmlns    string     `xml:"xmlns,attr"`
	Children []Property `xml:"property"`
}

type Property struct {
	Fmtid  string `xml:"fmtid,attr"`
	Pid    string `xml:"pid,attr"`
	Name   string `xml:"name,attr"`
	Lpwstr string `xml:"http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes lpwstr"`
}
