package main

import (
	"fmt"
	"santiaoqiao.com/zoffice/zdocx"

	"santiaoqiao.com/zoffice/zpackage"
)

func main() {
	ret, err := zpackage.DocxUnpack("./tmp/官僚资本主义.docx")
	if err != nil {
		println(err.Error())
	}
	file, ok := ret["[Content_Types].xml"]
	if ok {
		v := &zdocx.ContentTypes{}
		err := zpackage.UnmarshalFile(file, v)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		fmt.Println(*v)
	}
	for k := range ret {
		fmt.Println(k)
	}
}
