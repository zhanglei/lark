package xminio

import (
	"lark/pkg/utils"
	"sync"
)

func FPutPhotoListToMinio(photos []*utils.PhotoInfo) (resultList *PutResultList) {
	resultList = &PutResultList{List: make([]*PutResult, 0)}
	var (
		wg         = &sync.WaitGroup{}
		length     = len(photos)
		resultChan = make(chan *PutResult, length)
		result     *PutResult
		pt         *utils.PhotoInfo
		i          int
	)
	for _, pt = range photos {
		wg.Add(1)
		go FPutPhotoToMinio(pt, resultChan, wg)
	}
	wg.Wait()

	for i = 0; i < length; i++ {
		result = <-resultChan
		if result.Err != nil {
			resultList.Err = result.Err
		}
		resultList.List = append(resultList.List, result)
	}
	return
}

func FPutPhotoToMinio(photo *utils.PhotoInfo, resultChan chan *PutResult, wg *sync.WaitGroup) {
	result := new(PutResult)
	defer func() {
		wg.Done()
		resultChan <- result
	}()
	result.Info, result.Err = FPut(FILE_TYPE_PHOTO, photo.Name, photo.Path, photo.ContentType)
	return
}
