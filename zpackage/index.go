package zpackage

import (
	"archive/zip"
	"encoding/xml"
	"fmt"
	"io"
)

// DocxUnpack 使用zip解压缩.docx文档
func DocxUnpack(src string) (map[string]*zip.File, error) {
	r, err := zip.OpenReader(src)
	if err != nil {
		return nil, err
	}
	defer func(r *zip.ReadCloser) {
		_ = r.Close()
	}(r)
	ret := make(map[string]*zip.File)
	for _, f := range r.File {
		if !f.FileInfo().IsDir() {
			ret[f.Name] = f
		}
	}
	return ret, nil
	// 1. 读取 [Content_Types].xml，从中可以得到各个部分在什么地方
	// 2. 读取 主要的 document.main+xml 内容类型，获取所在路径，并解析它

}

// UnmarshalFile 从指定的xml文件中，读取数据到一个对象中
func UnmarshalFile(filePtr *zip.File, ptr interface{}) error {
	reader, err := filePtr.Open()
	if err != nil {
		return fmt.Errorf("error in opening file %s: %s", filePtr.Name, err.Error())
	}
	data, err := io.ReadAll(reader)
	if err != nil {
		return fmt.Errorf("error reading file %s", filePtr.Name)
	}
	err = xml.Unmarshal(data, ptr)
	if err != nil {
		return fmt.Errorf("error parse file %s", filePtr.Name)
	}
	return nil
}
