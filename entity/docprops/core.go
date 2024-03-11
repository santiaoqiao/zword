package docprops

// `Application-Defined File Properties part` => docProps/app.xml
type CoreProperties struct {
	Created        string `xml:"http://purl.org/dc/terms/ created"`
	Creator        string `xml:"http://purl.org/dc/elements/1.1/ creator"`
	LastModifiedBy string `xml:"http://schemas.openxmlformats.org/package/2006/metadata/core-properties lastModifiedBy"`
	Modified       string `xml:"http://purl.org/dc/terms/ modified"`
	Revision       int    `xml:"http://schemas.openxmlformats.org/package/2006/metadata/core-properties revision"`
}
