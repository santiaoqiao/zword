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
	"santiaoqiao.com/zdocx/entity"
	"santiaoqiao.com/zdocx/entity/docprops"
	packageRels "santiaoqiao.com/zdocx/entity/rels"
)

func Unpack(src string, dest string) (*entity.Docx, error) {
	// 定义一个 Docx, 最终返回其地址
	docx_ptr := &entity.Docx{}

	// 使用zip解压缩docx文档
	r, err := zip.OpenReader(src)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	// 第一步：读取_rels/.rels
	packageRels_ptr := &packageRels.Relationships{}
	readPartFromZipFile(r, "_rels/.rels", packageRels_ptr)
	log.Debug("完成读取_rels/.rels")
	docx_ptr.PackageRels = packageRels_ptr
	// 第二步；根据_rels/.rels中的描述，分别读取extend(app.xml),core,custom
	for index := range packageRels_ptr.Children {
		typeNmae := packageRels_ptr.Children[index].Type
		target := packageRels_ptr.Children[index].Target
		// 1- 如果存在extended-properties，则读取
		if strings.HasSuffix(typeNmae, "extended-properties") {
			extendProperties_ptr := &docprops.ExtendedProperties{}
			readPartFromZipFile(r, target, extendProperties_ptr)
			log.Debug("完成读取extended-properties")
		}
		// 2- 如果存在core-properties，则读取
		if strings.HasSuffix(typeNmae, "core-properties") {
			coreProperties_ptr := &docprops.CoreProperties{}
			err := readPartFromZipFile(r, target, coreProperties_ptr)
			if err != nil {
				log.Error(err.Error())
			}
			log.Debug("完成读取core-properties")
		}
		// 3- 如果存在custom-properties，则读取
		if strings.HasSuffix(typeNmae, "custom-properties") {
			customProperties_ptr := &docprops.CustomProperties{}
			err := readPartFromZipFile(r, target, customProperties_ptr)
			if err != nil {
				log.Error(err.Error())
			}
			log.Debug("完成读取_custom-properties")
		}
	}
	// 第三步：读取

	return docx_ptr, nil
}

func readPartFromZipFile(r *zip.ReadCloser, path string, ptr interface{}) error {
	f := getFileByName(r.File, path)
	if f == nil {
		return fmt.Errorf("can't find file _rels/.rels")
	}
	// 解析 _rels/.rels
	freader, err := f.Open()
	if err != nil {
		return fmt.Errorf("error in opening file _rels/.rels")
	}
	data, err := io.ReadAll(freader)
	if err != nil {
		return fmt.Errorf("error reading file _rels/.rels")
	}
	err = xml.Unmarshal(data, ptr)
	if err != nil {
		return fmt.Errorf("error parse file _rels/.rels")
	}
	return nil
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
