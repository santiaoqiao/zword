package zpackage

import (
	"archive/zip"
)

// 使用zip解压缩文档
func Unpack(src string) (map[string]*zip.File, error) {
	r, err := zip.OpenReader(src)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	ret := make(map[string]*zip.File)
	for _, f := range r.File {
		if !f.FileInfo().IsDir() {
			ret[f.Name] = f
		}
	}
	return ret, nil
}

func Consumer(src string) {}
