package multiple_service

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"testing"
	"time"
)

// 运行多个服务
func Test_Multiple_Service(t *testing.T) {
	server01 := &http.Server{
		Addr:         ":8080",
		Handler:      router01(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server02 := &http.Server{
		Addr:         ":8081",
		Handler:      router02(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	var g errgroup.Group
	g.Go(func() error {
		return server01.ListenAndServe()
	})
	g.Go(func() error {
		return server02.ListenAndServe()
	})

	err := g.Wait()
	if err != nil {
		log.Fatal(err)
	}

}

func router01() http.Handler {
	router := gin.New()
	router.Use(gin.Recovery())
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome server 01"})
	})
	return router
}

func router02() http.Handler {
	router := gin.New()
	router.Use(gin.Recovery())
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome server 02"})
	})
	return router
}
