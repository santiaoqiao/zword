package docx

// ExtendedProperties is the `Application-Defined File Properties part` => docProps/app.xmldocx
type ExtendedProperties struct {
	Template             string `xmldocx:"Template"`
	Pages                int    `xmldocx:"Pages"`
	Words                int    `xmldocx:"Words"`
	Characters           int    `xmldocx:"Characters"`
	Lines                int    `xmldocx:"Lines"`
	Paragraphs           int    `xmldocx:"Paragraphs"`
	TotalTime            int    `xmldocx:"TotalTime"`
	ScaleCrop            bool   `xmldocx:"ScaleCrop"`
	LinksUpToDate        bool   `xmldocx:"LinksUpToDate"`
	CharactersWithSpaces int    `xmldocx:"CharactersWithSpaces"`
	Application          string `xmldocx:"Application"`
	DocSecurity          int    `xmldocx:"DocSecurity"`
}
