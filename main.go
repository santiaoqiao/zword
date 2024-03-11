package main

import (
	"fmt"

	"santiaoqiao.com/zoffice/zpackage"
)

func main() {
	ret, err := zpackage.Unpack("./tmp/官僚资本主义.docx")
	if err != nil {
		println(err.Error())
	}
	// fmt.Println(docx.PackageRels.Xmlns)
	// for _, r := range docx.PackageRels.Children {
	// 	fmt.Println(r)
	// }
	// box.Unpack2("./tmp/官僚资本主义.docx", "./tmp/unpack")
	for k, v := range ret {
		fmt.Println(k, *v)
	}

}
