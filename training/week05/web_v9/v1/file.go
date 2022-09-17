package v1

import (
	lru "github.com/hashicorp/golang-lru"
	"io"
	"io/ioutil"
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
	// 可以在这里做点事情
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

type StaticResourceHandler struct {
	dir                     string
	extensionContentTypeMap map[string]string

	// 我要缓存, 但不是缓存所有的, 并且要控制住内存消耗
	cache       *lru.Cache
	maxFileSize int
}

func NewStaticResourceHandler(opts ...StaticResourceOption) *StaticResourceHandler {
	c, _ := lru.New(100)
	res := &StaticResourceHandler{
		cache: c,
		// 5M
		maxFileSize: 5 * 1024 * 1024,
		extensionContentTypeMap: map[string]string{
			// 这里根据自己的需要不断添加
			"jpeg": "image/jpeg",
			"jpe":  "image/jpeg",
			"jpg":  "image/jpeg",
			"png":  "image/png",
			"pdf":  "image/pdf",
		},
	}
	for _, opt := range opts {
		opt(res)
	}
	return res
}

type StaticResourceOption func(handler *StaticResourceHandler)

func WithStaticResourceDir(dir string) StaticResourceOption {
	return func(handler *StaticResourceHandler) {
		handler.dir = dir
	}
}

func WithMaxFileSize(max int) StaticResourceOption {
	return func(handler *StaticResourceHandler) {
		handler.maxFileSize = max
	}
}

func WithMoreExtension(extMap map[string]string) StaticResourceOption {
	return func(h *StaticResourceHandler) {
		for ext, contentType := range extMap {
			h.extensionContentTypeMap[ext] = contentType
		}
	}
}

func (s *StaticResourceHandler) Handle(ctx *Context) {
	file, ok := ctx.PathParams["file"]
	if !ok {
		ctx.RespData = []byte("文件不存在")
		ctx.StatusCode = http.StatusNotFound
		return
	}

	header := ctx.Resp.Header()
	// 从缓存中读
	itm, ok := s.cache.Get(file)
	if ok {
		header.Set("Content-Type", itm.(*cacheItem).contentType)
		ctx.Resp.WriteHeader(http.StatusOK)
		ctx.RespData = itm.(*cacheItem).data
		return
	}

	path := filepath.Join(s.dir, filepath.Clean(file))
	f, err := os.Open(path)
	if err != nil {
		ctx.RespData = []byte("文件不存在")
		ctx.StatusCode = http.StatusNotFound
		return
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		ctx.RespData = []byte("文件不存在")
		ctx.StatusCode = http.StatusNotFound
		return
	}

	// 写入缓存
	contentType := s.extensionContentTypeMap[filepath.Ext(file)]
	if len(data) < s.maxFileSize {
		newItem := &cacheItem{
			data:        data,
			contentType: contentType,
		}
		s.cache.Add(file, newItem)
	}
	header.Set("Content-Type", contentType)
	ctx.Resp.WriteHeader(http.StatusOK)
	ctx.RespData = data
}

type cacheItem struct {
	data        []byte
	contentType string
}
