package zword

import (
	"encoding/xml"
	"io"
	"strings"
)

type Table struct {
	Rows []TableRow
}

func (tbl *Table) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			// <w:tr>.....</w:tr>，交给 Paragraph 处理
			if t.Name.Local == "tr" {
				r := TableRow{}
				err := r.UnmarshalXML(d, token.(xml.StartElement))
				if err != nil {
					return err
				}
				tbl.Rows = append(tbl.Rows, r)
			}
			//todo:两个属性没做
			//if t.Name.Local == "tblPr" {
			//}
			//if t.Name.Local == "tblGrid"{
			//
			//}
		case xml.EndElement:
			if t.Name.Local == "tbl" {
				return nil
			}
		}
	}
	return nil
}

func (t *Table) String() string {
	sb := strings.Builder{}
	for _, row := range t.Rows {
		sb.WriteString(row.String())
		sb.WriteString("\n")
	}
	return sb.String()
}

func (t *Table) TypeName() string {
	return "tbl"
}
