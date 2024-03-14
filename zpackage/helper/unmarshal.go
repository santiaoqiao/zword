package helper

import (
	"encoding/xml"
	"strconv"
)

// <w:bidi w:val="0"/>
func Unwrap(t xml.StartElement, attrName string) (value string, haveValue bool) {
	for _, attr := range t.Attr {
		if attr.Name.Local == attrName {
			return attr.Value, true
		}
	}
	return "", false
}

func UnwrapVal(t xml.StartElement) (value string, haveValue bool) {
	return Unwrap(t, "val")
}

func UnwrapValToInt(t xml.StartElement) (value int, haveValue bool, err error) {
	s, ok := Unwrap(t, "val")
	if ok {
		i, err := strconv.ParseInt(s, 10, 0)
		if err != nil {
			return 0, true, err
		}
		return int(i), true, nil
	} else {
		return 0, false, nil
	}
}

func UnwrapValToBool(t xml.StartElement) (value bool, haveValue bool) {
	s, ok := Unwrap(t, "val")
	if ok {
		v, err := strconv.ParseBool(s)
		if err != nil {
			return false, ok
		}
		return v, ok
	} else {
		return true, ok
	}
}
