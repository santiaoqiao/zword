package main

import (
	"santiaoqiao.com/zdocx/box"
)

func main() {
	_, err := box.Unpack("./tmp/官僚资本主义.docx", "./tmp/unpack")
	if err != nil {
		println(err.Error())
	}
	// fmt.Println(docx.PackageRels.Xmlns)
	// for _, r := range docx.PackageRels.Children {
	// 	fmt.Println(r)
	// }
	// box.Unpack2("./tmp/官僚资本主义.docx", "./tmp/unpack")
}
