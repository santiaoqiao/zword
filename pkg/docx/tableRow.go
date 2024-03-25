package zword

import (
	"encoding/xml"
	"io"
	"strings"
)

type TableRow struct {
	Cells []TableCell
}

func (r *TableRow) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			// <w:tc>.....</w:tc>，交给 TableCell 处理
			if t.Name.Local == "tc" {
				c := TableCell{}
				err := c.UnmarshalXML(d, token.(xml.StartElement))
				if err != nil {
					return err
				}
				r.Cells = append(r.Cells, c)
			}
		case xml.EndElement:
			if t.Name.Local == "tr" {
				return nil
			}
		}
	}
	return nil
}

func (r *TableRow) String() string {
	sb := strings.Builder{}
	for _, cell := range r.Cells {
		sb.WriteString(cell.String())
		sb.WriteString("\t")
	}
	//sb.WriteString("\n")
	return sb.String()
}
