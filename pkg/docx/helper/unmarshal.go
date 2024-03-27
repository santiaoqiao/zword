package helper

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

// UnmarshalSingleAttrWithOk 获取tag中单个属性的值，返回值中ok表示是否存在该属性
func UnmarshalSingleAttrWithOk(t xml.StartElement, space string, local string) (val string, ok bool) {
	if space == "" {
		for _, attr := range t.Attr {
			if attr.Name.Local == local {
				return attr.Value, true
			}
		}
	} else {
		for _, attr := range t.Attr {
			if attr.Name.Space == space && attr.Name.Local == local {
				return attr.Value, true
			}
		}
	}
	return "", false
}

// UnmarshalSingleAttr 获取tag中单个属性的值，如果tag中不存在该属性，则返回空字符串
func UnmarshalSingleAttr(t xml.StartElement, space string, local string) (val string) {
	if space == "" {
		for _, attr := range t.Attr {
			if attr.Name.Local == local {
				return attr.Value
			}
		}
	} else {
		for _, attr := range t.Attr {
			if attr.Name.Space == space && attr.Name.Local == local {
				return attr.Value
			}
		}
	}
	return ""
}

// UnmarshalSingleValWithOk 获取标签中单个val属性的值，返回值中ok表示是否存在该属性
func UnmarshalSingleValWithOk(t xml.StartElement, space string) (val string, ok bool) {
	return UnmarshalSingleAttrWithOk(t, space, "val")
}

// UnmarshalSingleVal 获取标签中单个val属性的值，如果tag中不存在该属性，则返回空字符串
func UnmarshalSingleVal(t xml.StartElement, space string) (val string) {
	return UnmarshalSingleAttr(t, space, "val")
}

// UnmarshalSingleAttrToInt 获取tag中的唯一attr属性值，并将其转换为int类型
func UnmarshalSingleAttrToInt(t xml.StartElement, space string, local string) (val int, err error) {
	if space == "" {
		for _, attr := range t.Attr {
			if attr.Name.Local == local {
				i, err := strconv.ParseInt(attr.Value, 10, 0)
				if err != nil {
					return None, err
				}
				return int(i), nil
			}
		}
	} else {
		for _, attr := range t.Attr {
			if attr.Name.Space == space && attr.Name.Local == local {
				i, err := strconv.ParseInt(attr.Value, 10, 0)
				if err != nil {
					return 0, err
				}
				return int(i), nil
			}
		}
	}
	return None, fmt.Errorf("can't find attr [%s] in tag [%s]", local, t.Name.Local)
}

// UnmarshalSingleValToInt 获取tag中的唯一attr为val的属性值，并将其转换为int类型
func UnmarshalSingleValToInt(t xml.StartElement, space string) (val int, err error) {
	return UnmarshalSingleAttrToInt(t, space, "val")
}

var toggleTrue = true
var toggleFalse = false

// UnmarshalToggleValToBool 获取切换属性的值
func UnmarshalToggleValToBool(t xml.StartElement, space string) (val *bool) {
	s, ok := UnmarshalSingleValWithOk(t, space)
	if ok {
		v, err := strconv.ParseBool(s)
		if err != nil {
			return &toggleFalse
		}
		if v {
			return &toggleTrue
		} else {
			return &toggleFalse
		}
	} else {
		// 没有val属性，在toggle中，直接为true
		return &toggleTrue
	}
}
