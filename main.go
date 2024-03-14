package main

import (
	"fmt"
)

func main() {
	docx := &Docx{}
	err := docx.Read("./tmp/aaa.docx")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, p := range docx.Document.Body.Children {
		fmt.Printf("%v", p)
	}

	//file, err := os.Open("./tmp/unpack/_rels/.rels")
	//if err != nil {
	//	fmt.Println("Error opening file:", err)
	//	return
	//}
	//defer file.Close()
	//// 创建Decoder
	//decoder := xml.NewDecoder(file)
	//
	//for {
	//	// 读取下一个令牌
	//	token, err := decoder.Token()
	//	if err == io.EOF {
	//		break
	//	}
	//	if err != nil {
	//		fmt.Println("Error parsing XML:", err)
	//		return
	//	}
	//
	//	// 检查令牌类型
	//	switch t := token.(type) {
	//	case xml.StartElement:
	//		//fmt.Println(t.Name.Local)
	//		if t.Name.Local == "Relationship" {
	//			for _, attr := range t.Attr {
	//				fmt.Printf("%s -> %s\n", attr.Name.Local, attr.Value)
	//			}
	//
	//		}
	//	}
	//}
}
