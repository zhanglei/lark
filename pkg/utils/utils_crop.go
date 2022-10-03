package utils

import (
	"image"
	"io"
	"os"
	"sync"
)

type PhotoSize struct {
	Tag string
	W   int
	H   int
}

type PhotoInfo struct {
	Name  string
	Key   string
	Path  string
	Tag   string
	Error error
}

type Photos struct {
	Maps        map[string]*PhotoInfo
	ContentType string
	Format      string
	Error       error
}

var (
	sizeList = []*PhotoSize{
		&PhotoSize{Tag: "small", W: 72, H: 72},
		&PhotoSize{Tag: "medium", W: 240, H: 240},
		&PhotoSize{Tag: "large", W: 640, H: 640}}
)

func CropAvatar(in io.Reader, path string) (photos *Photos) {
	photos = &Photos{Maps: map[string]*PhotoInfo{}}
	var (
		count     = len(sizeList)
		photoChan = make(chan *PhotoInfo, count)
		wg        = &sync.WaitGroup{}
		photo     *PhotoInfo
		origin    image.Image
		format    string
		i         int
	)

	origin, format, photos.Error = image.Decode(in)
	if photos.Error != nil {
		return
	}
	photos.ContentType = GetContentType(format)
	photos.Format = format

	for i = 0; i < count; i++ {
		wg.Add(1)
		go cropPhoto(wg, origin, format, photoChan, sizeList[i], path)
	}
	wg.Wait()
	for i = 0; i < count; i++ {
		photo = <-photoChan
		if photo.Error != nil {
			photos.Error = photo.Error
		}
		photos.Maps[photo.Name] = photo
	}
	return
}

func cropPhoto(wg *sync.WaitGroup, origin image.Image, fm string, photo chan *PhotoInfo, size *PhotoSize, path string) {
	var (
		pi = &PhotoInfo{
			Key: NewUUID(),
			Tag: size.Tag,
		}
		file *os.File
	)
	pi.Name = pi.Key + "." + fm
	pi.Path = path + pi.Name

	defer func() {
		wg.Done()
		photo <- pi
	}()

	file, pi.Error = os.Create(pi.Path)
	if pi.Error != nil {
		return
	}
	defer func() {
		file.Close()
	}()

	pi.Error = CropPhoto(origin, fm, file, size.W, size.H, 100)
	return
}
