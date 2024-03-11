package rels

// `Relationships` => ./_rels/.rels
type Relationships struct {
	Xmlns    string         `xml:"xmlns,attr"`
	Children []Relationship `xml:"Relationship"`
}

type Relationship struct {
	Id     string `xml:"Id,attr"`
	Type   string `xml:"Type,attr"`
	Target string `xml:"Target,attr"`
}
