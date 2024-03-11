package box

import "archive/zip"

// 根据文件名获取一个zip.File
func getFileByName(files []*zip.File, filename string) *zip.File {
	for _, f := range files {
		if f.Name == filename {
			return f
		}
	}
	return nil
}
