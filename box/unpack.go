package box

import (
	"archive/zip"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	log "github.com/sirupsen/logrus"
	packageRels "santiaoqiao.com/zdocx/types/rels"
)

func Unpack(src string, dest string) (interface{}, error) {
	var filenames []string

	r, err := zip.OpenReader(src)
	if err != nil {
		return filenames, err
	}
	defer r.Close()

	// 第一步：读取_rels/.rels
	f := getFileByName(r.File, "_rels/.rels")
	if f == nil {
		return nil, fmt.Errorf("can't find file _rels/.rels")
	}
	// 解析 _rels/.rels
	freader, err := f.Open()
	if err != nil {
		return nil, fmt.Errorf("error in opening file _rels/.rels")
	}
	data, err := io.ReadAll(freader)
	if err != nil {
		return nil, fmt.Errorf("error reading file _rels/.rels")
	}
	packageRelationships_ptr := &packageRels.Relationships{}
	err = xml.Unmarshal(data, packageRelationships_ptr)
	if err != nil {
		return nil, fmt.Errorf("error parse file _rels/.rels")
	}
	log.WithField("_rels/.rels", *packageRelationships_ptr).Debug("完成读取_rels/.rels")
	return nil, nil
}

func Unpack2(src string, dest string) ([]string, error) {
	var filenames []string

	r, err := zip.OpenReader(src)
	if err != nil {
		return filenames, err
	}
	defer r.Close()

	for _, f := range r.File {

		// Store filename/path for returning and using later on
		fpath := filepath.Join(dest, f.Name)

		// Check for ZipSlip. More Info: http://bit.ly/2MsjAWE
		if !strings.HasPrefix(fpath, filepath.Clean(dest)+string(os.PathSeparator)) {
			return filenames, fmt.Errorf("%s: illegal file path", fpath)
		}

		filenames = append(filenames, fpath)

		if f.FileInfo().IsDir() {
			// Make Folder
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		// Make File
		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return filenames, err
		}

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return filenames, err
		}

		rc, err := f.Open()
		if err != nil {
			return filenames, err
		}

		_, err = io.Copy(outFile, rc)

		// Close the file without defer to close before next iteration of loop
		outFile.Close()
		rc.Close()

		if err != nil {
			return filenames, err
		}
	}
	return filenames, nil
}

// 根据文件名获取一个zip.File
func getFileByName(files []*zip.File, filename string) *zip.File {
	for _, f := range files {
		if f.Name == filename {
			return f
		}
	}
	return nil
}
