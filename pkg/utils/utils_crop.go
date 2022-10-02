package utils

import (
	"image"
	"io"
	"os"
	"strconv"
	"sync"
)

type PhotoSize struct {
	W int
	H int
}

type PhotoInfo struct {
	Path        string
	Name        string
	ContentType string
	Error       error
}

type NewPhoto struct {
	List        []*PhotoInfo
	ContentType string
	Error       error
}

var (
	AvatarList  = []*PhotoSize{&PhotoSize{W: 72, H: 72}, &PhotoSize{W: 240, H: 240}, &PhotoSize{W: 640, H: 640}}
	PhotoSuffix = map[int]string{72: "_small", 240: "_medium", 640: "_large"}
)

func CropAvatar(in io.Reader, path string, name string) (new *NewPhoto) {
	new = &NewPhoto{List: make([]*PhotoInfo, 0)}
	var (
		count     = len(AvatarList)
		photoChan = make(chan *PhotoInfo, count)
		wg        = &sync.WaitGroup{}
		photo     *PhotoInfo
		origin    image.Image
		format    string
		i         int
	)

	origin, format, new.Error = image.Decode(in)
	if new.Error != nil {
		return
	}
	new.ContentType = GetContentType(format)

	for i = 0; i < count; i++ {
		wg.Add(1)
		go cropPhoto(wg, origin, format, photoChan, AvatarList[i], path, name)
	}
	wg.Wait()
	for i = 0; i < count; i++ {
		photo = <-photoChan
		if photo.Error != nil {
			new.Error = photo.Error
		}
		photo.ContentType = new.ContentType
		new.List = append(new.List, photo)
	}
	return
}

func cropPhoto(wg *sync.WaitGroup, origin image.Image, fm string, photo chan *PhotoInfo, size *PhotoSize, path string, name string) {
	var (
		pi   = &PhotoInfo{}
		file *os.File
		ok   bool
	)

	defer func() {
		wg.Done()
		photo <- pi
	}()

	if _, ok = PhotoSuffix[size.W]; ok {
		pi.Name = name + PhotoSuffix[size.W]
	} else {
		pi.Name = name + strconv.Itoa(size.W)
	}
	pi.Path = path + pi.Name + "." + fm
	file, pi.Error = os.Create(pi.Path)
	if pi.Error != nil {
		return
	}
	defer func() {
		file.Close()
	}()

	pi.Error = CropPhoto(origin, fm, file, size.W, size.H, 75)
	return
}
