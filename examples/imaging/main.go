package main

import (
	"fmt"
	"io/ioutil"
	"lark/pkg/common/ximaging"
	"os"
	"path/filepath"
)

func main() {
	var (
		fileName = "./examples/imaging/lark.jpeg"
		ext      = filepath.Ext(fileName)
		file     *os.File
		buf      []byte
		err      error
	)
	file, err = os.OpenFile(fileName, os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()

	buf, err = ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	ximaging.ReSizeImage(buf, ext, true, func(szType string, localId int, w, h int32, b []byte) error {
		return ioutil.WriteFile(fmt.Sprintf("%s.%s%s", fileName, szType, ext), b, 0644)
	})
}
