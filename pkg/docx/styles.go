package docx

import "encoding/xml"

type Styles struct {
	DocDefaults  DocDefaults `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main rPrDefault"`
	LatentStyles LatentStyles
	StyleItems   []StyleItem
}

type DocDefaults struct {
	RPrDefault RunProperty
	PPrDefault ParagraphProperty
}

type LatentStyles struct {
	XMLName xml.Name `xml:"w:latentStyles" xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main"`

	Count             int  `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main count,attr"`
	DefQFormat        bool `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main defQFormat,attr"`
	DefUnhideWhenUsed bool `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main defUnhideWhenUsed,attr"`
	DefSemiHidden     bool `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main defSemiHidden,attr"`
	DefUIPriority     int  `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main defUIPriority,attr"`
	DefLockedState    bool `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main defLockedState,attr"`
	LsdExceptions     []LsdException
}

type LsdException struct {
	XMLName        xml.Name `xml:"w:lsdException" xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main"`
	Locked         bool     `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main locked,attr,omitempty"`
	QFormat        bool     `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main qFormat,attr,omitempty"`
	UnhideWhenUsed bool     `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main unhideWhenUsed,attr,omitempty"`
	UiPriority     int      `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main uiPriority,attr,omitempty"`
	SemiHidden     bool     `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main semiHidden,attr,omitempty"`
	Name           string   `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main name,attr"`
}

type StyleItem struct {
	XMLName        xml.Name  `xml:"w:style" xmlns:w="http://schemas.openxmlformats.org/wordprocessingml/2006/main"`
	TypeFor        string    `xml:"w:type,attr,omitempty"`
	Default        bool      `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main default,attr,omitempty"`
	StyleId        string    `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main styleId,attr,omitempty"`
	Name           StyleName `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main name"`
	QFormat        bool      `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main type,attr,omitempty"`
	UiPriority     int       `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main type,attr,omitempty"`
	UnhideWhenUsed bool      `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main type,attr,omitempty"`
	BaseOn         string    `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main type,attr,omitempty"`
	Next           string    `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main type,attr,omitempty"`
}

type StyleName struct {
	Val string `xml:"http://schemas.openxmlformats.org/wordprocessingml/2006/main val"`
}

/*
func (s *Styles) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	for {
		token, err := d.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		switch t := token.(type) {
		case xml.StartElement:
			switch t.Name.Space {
			case cSpaceW:
				switch t.Name.Local {
				}
			}
		case xml.EndElement:
			if t.Name.Local == cTagStyles {
				return nil
			}
		}
	}
	return nil
}
*/
