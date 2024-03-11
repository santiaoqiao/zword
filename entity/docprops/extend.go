package docprops

// `Application-Defined File Properties part` => docProps/app.xml
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
