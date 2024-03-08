package main

import (
	"santiaoqiao.com/zdocx/box"
)

func main() {
	box.Unpack("./tmp/官僚资本主义.docx", "./tmp/unpack")
	// a, _ := os.Open("./tmp/官僚资本主义/_rels/.rels")
	// b, _ := io.ReadAll(a)

	// var v = rels.Relationships{}
	// err := xml.Unmarshal(b, &v)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// for _, c := range v.Children {
	// 	fmt.Printf("%v\n", c)
	// }
}
