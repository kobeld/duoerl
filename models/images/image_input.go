package images

import (
	"bytes"
	"github.com/sunfmin/tenpu"
	"io"
	"mime/multipart"
)

var IMAGES = "images"

type ImageInput struct {
	Id          string
	FileName    string
	ContentType string
	Thumb       string
	Download    bool
	OwnerId     string
	Category    string
}

func (this *ImageInput) GetFileMeta() (filename string, contentType string, contentId string) {
	filename = this.FileName
	contentType = this.ContentType
	return
}

func (this *ImageInput) GetViewMeta() (id string, thumb string, download bool) {
	id = this.Id
	thumb = this.Thumb
	download = this.Download
	return
}

func (this *ImageInput) LoadAttachments() (atts []*tenpu.Attachment, err error) {
	return
}

func (this *ImageInput) SetMultipart(part *multipart.Part) (isFile bool) {
	if part.FileName() != "" {
		this.FileName = part.FileName()
		this.ContentType = part.Header["Content-Type"][0]
		isFile = true
		return
	}

	switch part.FormName() {
	case "OwnerId":
		this.OwnerId = formValue(part)
	}
	return
}

func (this *ImageInput) SetAttrsForCreate(att *tenpu.Attachment) (err error) {
	if this.OwnerId != "" {
		att.OwnerId = []string{this.OwnerId}
	}
	att.Category = this.Category
	return
}

func (this *ImageInput) SetAttrsForDelete(att *tenpu.Attachment) (shouldUpdate bool, shouldDelete bool, err error) {
	shouldDelete = true
	return
}

// ---- Private ----
func formValue(p *multipart.Part) string {
	var b bytes.Buffer
	io.CopyN(&b, p, int64(1<<20)) // Copy max: 1 MiB
	return b.String()
}
