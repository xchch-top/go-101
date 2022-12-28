package upload

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"testing"
)

func Test_Upload_Single(t *testing.T) {
	router := gin.Default()
	// 为multipart forms 设置较低的内存限制(默认是32MiB)
	router.MaxMultipartMemory = 8 << 20

	router.POST("/upload", func(ctx *gin.Context) {
		file, _ := ctx.FormFile("file")
		log.Printf("filename: %v\n", file.Filename)

		dst := "./" + file.Filename
		// 上传文件至指定的的完整文件路径
		ctx.SaveUploadedFile(file, dst)
		ctx.String(http.StatusOK, fmt.Sprintf("%s uploaded", file.Filename))
	})

	router.Run(":8080")
}

func Test_Upload_Multiple(t *testing.T) {
	router := gin.Default()
	// 为multipart forms 设置较低的内存限制(默认是32MiB)
	router.MaxMultipartMemory = 8 << 20

	router.POST("/upload", func(ctx *gin.Context) {
		form, _ := ctx.MultipartForm()
		files := form.File["file"]
		for _, file := range files {
			log.Printf("filename: %v\n", file.Filename)
			dst := "./" + file.Filename
			// 上传文件至指定的的完整文件路径
			ctx.SaveUploadedFile(file, dst)
		}
		ctx.String(http.StatusOK, fmt.Sprintf("%d files uploaded", len(files)))
	})

	router.Run(":8080")
}
