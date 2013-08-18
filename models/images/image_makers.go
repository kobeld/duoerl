package images

import (
	"github.com/kobeld/duoerl/global"
	"github.com/sunfmin/tenpu"
	"github.com/sunfmin/tenpu/gridfs"
	"github.com/sunfmin/tenpu/mgometa"
	"net/http"
)

var TheImageMaker = NewImageMaker()

type ImageMaker struct {
}

func NewImageMaker() *ImageMaker {
	return &ImageMaker{}
}

func (this *ImageMaker) MakeForRead(r *http.Request) (storage tenpu.BlobStorage, meta tenpu.MetaStorage, input tenpu.Input, err error) {
	var imageInput *ImageInput
	storage, meta, imageInput, err = this.make(r)
	imageInput.Id = r.FormValue(":id")
	input = imageInput
	return
}

func (this *ImageMaker) MakeForUpload(r *http.Request) (storage tenpu.BlobStorage, meta tenpu.MetaStorage, input tenpu.UploadInput, err error) {
	var imageInput *ImageInput
	storage, meta, imageInput, err = this.make(r)
	imageInput.OwnerId = r.URL.Query().Get(":uid")
	input = imageInput
	return
}

func (this *ImageMaker) make(r *http.Request) (storage tenpu.BlobStorage, meta tenpu.MetaStorage, input *ImageInput, err error) {
	storage = gridfs.NewStorage(global.ImageDatabase)
	meta = mgometa.NewStorage(global.ImageDatabase, IMAGES)
	input = &ImageInput{}
	return
}
