package docx

import (
	"archive/zip"
	"encoding/xml"
	"fmt"
	"io"
)

type Docx struct {
	ContentTypes *ContentTypes
	Document     *Document
	//Styles       *Styles
}

func (docx *Docx) Read(filename string) error {
	r, err := zip.OpenReader(filename)
	if err != nil {
		return err
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
	// 🚩 读取 [Content_Types].xml，从中可以得到各个部分在什么地方
	contentTypesXMLFile, ok := fileMap["[Content_Types].xml"]
	if ok {
		ptr := &ContentTypes{}
		err := unmarshalFile(contentTypesXMLFile, ptr)
		if err != nil {
			return err
		}
		docx.ContentTypes = ptr
	}

	// 🚩 读取 主要的 document.main+xml 内容类型，获取所在路径，并解析它
	documentXMLLFile, ok := fileMap["word/document.xml"]
	if ok {
		ptr := &Document{}
		err := unmarshalFile(documentXMLLFile, ptr)
		if err != nil {
			return err
		}
		docx.Document = ptr
	}

	return nil
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
