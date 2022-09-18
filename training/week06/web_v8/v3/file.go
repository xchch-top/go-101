package v3

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

func (f *FileUploader) Handle() HandleFunc {
	return func(ctx *Context) {
		file, header, err := ctx.Req.FormFile(f.FileField)
		if err != nil {
			ctx.StatusCode = http.StatusInternalServerError
			ctx.RespData = []byte("上传失败")
			return
		}

		dst, err := os.OpenFile(f.DstPathFunc(header), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0o666)
		if err != nil {
			ctx.StatusCode = http.StatusInternalServerError
			ctx.RespData = []byte("上传失败")
			return
		}

		io.CopyBuffer(dst, file, nil)
	}
}

type FileDownloader struct {
	Param string
	// 文件下载的前缀路径
	Dir string
}

func (f *FileDownloader) Handle() HandleFunc {
	return func(ctx *Context) {
		fq := ctx.QueryValue(f.Param)
		if fq.err != nil {
			ctx.StatusCode = http.StatusInternalServerError
			ctx.RespData = []byte("文件下载的参数错误")
			return
		}
		file, _ := fq.AsString()
		// file可能是一些乱七八糟的东西 //////ab.txt
		path := filepath.Join(f.Dir, filepath.Clean(file))
		fn := filepath.Base(path)
		header := ctx.Resp.Header()

		header.Set("Content-Disposition", "attachment;filename="+fn)
		header.Set("Content-Description", "File Transfer")
		header.Set("Content-Type", "application/octet-stream")
		header.Set("Content-Transfer-Encoding", "binary")
		header.Set("Expires", "0")
		header.Set("Cache-Control", "must-revalidate")
		header.Set("Pragma", "public")

		// 文件下载本质上, 就是把文件写到这里
		http.ServeFile(ctx.Resp, ctx.Req, path)
	}
}
