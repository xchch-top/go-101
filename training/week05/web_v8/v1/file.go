package v1

import (
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

type FileUploader struct {
	// FileField 对应于文件在表单中的字段名字
	FileField string

	// DstPathFunc 用于计算目标路径
	DstPathFunc func(fh *multipart.FileHeader) string
}

func (f *FileUploader) Handle(ctx *Context) HandleFunc {
	return func(ctx *Context) {
		file, header, err := ctx.Req.FormFile(f.FileField)
		if err != nil {
			ctx.StatusCode = http.StatusInternalServerError
			ctx.RespData = []byte("上传失败")
			return
		}

		dst, err := os.OpenFile(filepath.Join(f.DstPathFunc(header), header.Filename), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0o666)
		if err != nil {
			ctx.StatusCode = http.StatusInternalServerError
			ctx.RespData = []byte("上传失败")
			return
		}

		io.CopyBuffer(dst, file, nil)
	}
}
