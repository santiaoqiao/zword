package docx

import (
	"encoding/xml"
	"github.com/santiaoqiao/zword/pkg/docx/helper"
	"io"
	"strconv"
)

type TableCellProperty struct {
	// Table Cell Width,如果为小于0的值，则代表未取到该值
	Tcw float64
	// dxa |  pct(百分比)
	TcwType string
}

func (c *TableCellProperty) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			// <w:tcW w:w="2840" w:customtype="dxa"/>
			if t.Name.Local == "tcW" {
				if w, ok := helper.UnmarshalSingleAttrWithOk(t, helper.CSpaceW, "w"); ok {
					wf, err := strconv.ParseFloat(w, 32)
					if err != nil {
						return err
					}
					c.Tcw = wf
				} else {
					c.Tcw = -1.0
				}

				if wt, ok := helper.UnmarshalSingleAttrWithOk(t, helper.CSpaceW, "customtype"); ok {
					c.TcwType = wt
				}
			}

			// <w:tcPr>....<w:tcPr>，交给 TableCellProperty 处理
			if t.Name.Local == "tbl" {
				// todo: table处理
			}
		case xml.EndElement:
			if t.Name.Local == "tcPr" {
				return nil
			}
		}
	}
	return nil
}

func (c *TableCellProperty) String() string {
	return "tableProperty"
}
