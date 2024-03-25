package zword

import (
	"archive/zip"
	"encoding/xml"
	"fmt"
	"io"
	"santiaoqiao.com/zword/pkg/xml_parser"
	"santiaoqiao.com/zword/pkg/xml_parser/properties"
)

func OpenDocxFile(filename string) (*xml_parser.Document, error) {
	doc := &xml_parser.Document{}
	// 解压文件
	r, err := zip.OpenReader(filename)
	if err != nil {
		return nil, err
	}
	defer func(r *zip.ReadCloser) {
		_ = r.Close()
	}(r)

	// 获取目录中的所有文件，与路径相对应，保存在 fileMap 中
	fileMap := make(map[string]*zip.File)
	for _, f := range r.File {
		if !f.FileInfo().IsDir() {
			fileMap[f.Name] = f
		}
	}
	// 🚩 读取 [Content_Types].xml_parser，从中可以得到各个部分在什么地方
	contentTypesXMLFile, ok := fileMap["[Content_Types].xml"]
	if ok {
		ptr := &properties.ContentTypes{}
		err := unmarshalFile(contentTypesXMLFile, ptr)
		if err != nil {
			return nil, err
		}
		doc.ContentTypes = ptr
	}

	// 🚩 读取 主要的 word.main+xml_parser 内容类型，获取所在路径，并解析它
	documentXMLLFile, ok := fileMap["word/document.xml"]
	if ok {
		ptr := &xml_parser.Document{}
		err := unmarshalFile(documentXMLLFile, ptr)
		if err != nil {
			return nil, err
		}
		doc = ptr
	}
	return doc, nil
}

// 解析XML文件到指定的对象
func unmarshalFile(filePtr *zip.File, ptr interface{}) error {
	reader, err := filePtr.Open()
	defer func() {
		_ = reader.Close()
	}()
	if err != nil {
		return fmt.Errorf("error in opening file %s, errors: %s", filePtr.Name, err.Error())
	}
	data, err := io.ReadAll(reader)
	if err != nil {
		return fmt.Errorf("error reading file %s, errors: %s", filePtr.Name, err.Error())
	}
	err = xml.Unmarshal(data, ptr)
	if err != nil {
		return fmt.Errorf("error parse file %s, errors: %s", filePtr.Name, err.Error())
	}
	return nil
}
