package lets_encrypt

import (
	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/acme/autocert"
	"log"
	"net/http"
	"testing"
)

// 支持 Let's Encrypt -- 没搞懂
func Test_Lets_Encrypt(t *testing.T) {
	router := gin.Default()

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong")
	})

	m := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("example.top"),
		Cache:      autocert.DirCache("/var/www/.cache"),
	}

	err := autotls.RunWithManager(router, &m)
	log.Fatal(err)
}
