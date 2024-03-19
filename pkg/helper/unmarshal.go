package helper

import (
	"encoding/xml"
	"santiaoqiao.com/zoffice/types/triple"
	"strconv"
)

// <w:bidi w:val="0"/>
func Unwrap(t xml.StartElement, attrName string) (val string, ok bool) {
	for _, attr := range t.Attr {
		if attr.Name.Local == attrName {
			return attr.Value, true
		}
	}
	return "", false
}

func UnwrapVal(t xml.StartElement) (val string, ok bool) {
	return Unwrap(t, "val")
}

func UnwrapValToInt(t xml.StartElement) (val int, ok bool, err error) {
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

func UnwrapValToBool(t xml.StartElement) (val bool, ok bool) {
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

func UnwrapValToToggle(t xml.StartElement) (val bool) {
	s, ok := Unwrap(t, "val")
	if ok {
		v, err := strconv.ParseBool(s)
		if err != nil {
			return false
		}
		return v
	} else {
		// 没有val属性，在toggle中，直接为true
		return true
	}
}

func UnwrapValToTriple(t xml.StartElement) (val triple.Triple) {
	s, ok := Unwrap(t, "val")
	if ok {
		v, err := strconv.ParseBool(s)
		if err != nil {
			return triple.False
		}
		if v {
			return triple.True
		} else {
			return triple.False
		}
	} else {
		return triple.None
	}
}
