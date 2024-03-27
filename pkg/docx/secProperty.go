package docx

type SectionProperty struct {
}

func (s *SectionProperty) String() string {
	return ""
}

func (s *SectionProperty) TypeName() BodyChildType {
	return BodyTypeSecPr
}
