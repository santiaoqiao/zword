package rels

// `PackageRelationships` => ./rels/.rels
type PackageRelationships struct {
	Xmlns    string                    `xmldocx:"xmlns,attr"`
	Children []PackageRelationshipItem `xmldocx:"PackageRelationshipItem"`
}

type PackageRelationshipItem struct {
	Id     string `xmldocx:"Id,attr"`
	Type   string `xmldocx:"TypeName,attr"`
	Target string `xmldocx:"Target,attr"`
}
