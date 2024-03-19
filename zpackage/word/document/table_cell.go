package document

import (
	"encoding/xml"
	"io"
	"strings"
)

type TableCell struct {
	Property TableCellProperty
	Children []BodyChild
}

func (c *TableCell) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			// <w:p>.....</w:p>，交给 Paragraph 处理
			if t.Name.Local == "p" {
				p := &Paragraph{}
				err := p.UnmarshalXML(d, token.(xml.StartElement))
				if err != nil {
					return err
				}
				c.Children = append(c.Children, p)
			}
			// <w:tcPr>....<w:tcPr>，交给 TableCellProperty 处理
			if t.Name.Local == "tcPr" {
				tcPr := TableCellProperty{}
				err := tcPr.UnmarshalXML(d, token.(xml.StartElement))
				if err != nil {
					return err
				}
				c.Property = tcPr
			}
			// 表格套表格
			if t.Name.Local == "tbl" {
				table := &Table{}
				err := table.UnmarshalXML(d, token.(xml.StartElement))
				if err != nil {
					return err
				}
				c.Children = append(c.Children, table)
			}
		case xml.EndElement:
			if t.Name.Local == "tc" {
				return nil
			}
		}
	}
	return nil
}

func (c *TableCell) String() string {
	sb := strings.Builder{}
	for _, child := range c.Children {
		sb.WriteString(c.Property.String())
		sb.WriteString(child.String())
	}
	return sb.String()
}
