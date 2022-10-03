package ximaging

import (
	"bytes"
	"github.com/chai2010/webp"
	"image"
	"io"
	"io/ioutil"
)

func OpenWebp(filename string) (image.Image, error) {
	srcData, _ := ioutil.ReadFile(filename)
	img, err := webp.Decode(bytes.NewReader(srcData))
	if err != nil {
		return nil, err
	}
	return img, nil
}

func DecodeWebp(r io.Reader) (image.Image, error) {
	return webp.Decode(r)
}

func EncodeWebp(w io.Writer, img image.Image) error {
	return webp.Encode(w, img, &webp.Options{Lossless: true})
}
