package docx

import (
	"encoding/xml"
	"io"
	"santiaoqiao.com/zword/pkg/docx/helper"
	"strconv"
)

type Styles struct {
	DocDefaults  DocDefaults
	LatentStyles LatentStyles
	StyleSheets  map[string]StyleItem
}

type DocDefaults struct {
	RPrDefault *RunProperty
	PPrDefault *ParagraphProperty
}

type LatentStyles struct {
	Count             int
	DefQFormat        bool
	DefUnhideWhenUsed bool
	DefSemiHidden     bool
	DefUIPriority     int
	DefLockedState    bool
	LsdExceptions     []LsdException
}

type LsdException struct {
	Locked         bool
	QFormat        bool
	UnhideWhenUsed bool
	UiPriority     int
	SemiHidden     bool
	Name           string
}

type StyleItem struct {
	TypeFor string
	Default bool
	//rStyleId string  // it has putted into the key of the map
	Name         string
	QFormat      *bool
	AutoRedefine *bool
	UiPriority   int
	BasedOn      string
	Next         string
	RPr          *RunProperty
	PPr          *ParagraphProperty
}

func (s *Styles) UnmarshalXML(d *xml.Decoder, _ xml.StartElement) error {
	items := make(map[string]StyleItem)
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
			case helper.CSpaceW:
				switch t.Name.Local {
				case "docDefaults":
					// <w:docDefaults>...</w:docDefaults>
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
						case helper.CSpaceW:
							switch t.Name.Local {
							case "rPrDefault":
								// <w:rPrDefault>...</w:rPrDefault>
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
									case helper.CSpaceW:
										switch t.Name.Local {
										case "rPr":
											// <w:rPr>...</w:rPr>
											rPr := &RunProperty{}
											err := d.DecodeElement(rPr, &t)
											if err != nil {
												return err
											}
											s.DocDefaults.RPrDefault = rPr
										}
									}
								}
							case "pPrDefault":
								// <w:pPrDefault>...</w:pPrDefault>
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
									case helper.CSpaceW:
										switch t.Name.Local {
										case "pPr":
											// <w:pPr>...</w:pPr>
											pPr := &ParagraphProperty{}
											err := d.DecodeElement(pPr, &t)
											if err != nil {
												return err
											}
											s.DocDefaults.PPrDefault = pPr
										}
									}
								}
							}
						}

					}
				case "latentStyles":
					for _, attr := range t.Attr {
						switch attr.Name.Space {
						case helper.CSpaceW:
							switch attr.Name.Local {
							case "count":
								val, err := helper.Str2Int(attr.Value)
								if err != nil {
									return err
								}
								s.LatentStyles.Count = val
							case "defQFormat":
								val, err := strconv.ParseBool(attr.Value)
								if err != nil {
									return err
								}
								s.LatentStyles.DefQFormat = val
							case "defUnhideWhenUsed":
								val, err := strconv.ParseBool(attr.Value)
								if err != nil {
									return err
								}
								s.LatentStyles.DefUnhideWhenUsed = val
							case "defSemiHidden":
								val, err := strconv.ParseBool(attr.Value)
								if err != nil {
									return err
								}
								s.LatentStyles.DefSemiHidden = val
							case "defUIPriority":
								val, err := helper.Str2Int(attr.Value)
								if err != nil {
									return err
								}
								s.LatentStyles.DefUIPriority = val
							case "defLockedState":
								val, err := strconv.ParseBool(attr.Value)
								if err != nil {
									return err
								}
								s.LatentStyles.DefLockedState = val
							}
						}
					}
					lsdExceptions := make([]LsdException, 0, s.LatentStyles.Count)
				latentStylesInnerLoop:
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
							case helper.CSpaceW:
								switch t.Name.Local {
								case "lsdException":
									// <w:lsdException>...
									lsdException := LsdException{}
									for _, attr := range t.Attr {
										switch attr.Name.Space {
										case helper.CSpaceW:
											// <w:lsdException w:qFormat="1" w:unhideWhenUsed="0" w:uiPriority="0" w:semiHidden="0" w:name="Normal"/>
											switch attr.Name.Local {
											case "unhideWhenUsed":
												val, err := strconv.ParseBool(attr.Value)
												if err != nil {
													return err
												}
												lsdException.UnhideWhenUsed = val
											case "uiPriority":
												val, err := helper.Str2Int(attr.Value)
												if err != nil {
													return err
												}
												lsdException.UiPriority = val
											case "semiHidden":
												val, err := strconv.ParseBool(attr.Value)
												if err != nil {
													return err
												}
												lsdException.SemiHidden = val
											case "name":
												lsdException.Name = attr.Value
											case "qFormat":
												val, err := strconv.ParseBool(attr.Value)
												if err != nil {
													return err
												}
												lsdException.QFormat = val
											}
										}
									}
									lsdExceptions = append(lsdExceptions, lsdException)
								}
							}
						case xml.EndElement:
							// ...</w:latentStyles>
							if t.Name.Space == helper.CSpaceW && t.Name.Local == "latentStyles" {
								break latentStylesInnerLoop
							}
						}
					}
					s.LatentStyles.LsdExceptions = lsdExceptions
				case "style":
					item := StyleItem{}
					key := ""
					for _, attr := range t.Attr {
						switch attr.Name.Space {
						case helper.CSpaceW:
							switch attr.Name.Local {
							case "type":
								item.TypeFor = attr.Value
							case "default":
								val, err := strconv.ParseBool(attr.Value)
								if err != nil {
									return err
								}
								item.Default = val
							case "styleId":
								key = attr.Value
							}
						}
					}
				styleInnerLoop:
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
							case helper.CSpaceW:
								switch t.Name.Local {
								case "name":
									item.Name = helper.UnmarshalSingleVal(t, helper.CSpaceW)
								case "autoRedefine":
									item.AutoRedefine = helper.UnmarshalToggleValToBool(t, helper.CSpaceW)
								case "qFormat":
									item.QFormat = helper.UnmarshalToggleValToBool(t, helper.CSpaceW)
								case "uiPriority":
									val, err := helper.UnmarshalSingleValToInt(t, helper.CSpaceW)
									if err != nil {
										return err
									}
									item.UiPriority = val
								case "next":
									item.Next = helper.UnmarshalSingleVal(t, helper.CSpaceW)
								case "baseOn":
									item.BasedOn = helper.UnmarshalSingleVal(t, helper.CSpaceW)
								case "rPr":
									rPr := &RunProperty{}
									err := d.DecodeElement(rPr, &t)
									if err != nil {
										return err
									}
									item.RPr = rPr
								case "pPr":
									pPr := &ParagraphProperty{}
									err := d.DecodeElement(pPr, &t)
									if err != nil {
										return err
									}
									item.PPr = pPr
								}
							}
						case xml.EndElement:
							if t.Name.Space == helper.CSpaceW && t.Name.Local == "style" {
								if key != "" {
									// add to style sheets
									items[key] = item
								}
								break styleInnerLoop
							}
						}
					}
				}
			}
		case xml.EndElement:
			if t.Name.Local == "styles" {
				// at the end
				s.StyleSheets = items
				return nil
			}
		}
	}
	// if missed tag </styles>, add items to stylesheet
	s.StyleSheets = items
	return nil
}
