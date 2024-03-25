package docx

// CoreProperties is the `Application-Defined File Properties part` => docProps/core.xmldocx
type CoreProperties struct {
	Created        string `xmldocx:"http://purl.org/dc/terms/ created"`
	Creator        string `xmldocx:"http://purl.org/dc/elements/1.1/ creator"`
	LastModifiedBy string `xmldocx:"http://schemas.openxmlformats.org/package/2006/metadata/core-properties lastModifiedBy"`
	Modified       string `xmldocx:"http://purl.org/dc/terms/ modified"`
	Revision       int    `xmldocx:"http://schemas.openxmlformats.org/package/2006/metadata/core-properties revision"`
}
