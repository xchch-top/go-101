package static_embed

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func Test_Assets_Embed(ts *testing.T) {
	router := gin.New()

	t, err := loadTemplate()
	if err != nil {
		panic(err)
	}
	router.SetHTMLTemplate(t)

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "/html/index.tmpl", gin.H{"Foo": "Foo"})
	})

	router.GET("/bar", func(c *gin.Context) {
		c.HTML(http.StatusOK, "/html/bar.tmpl", gin.H{"Foo": "Foo"})
	})

	router.Run(":8080")

}

// loadTemplate 加载由 go-assets-builder 嵌入的模板
func loadTemplate() (*template.Template, error) {
	t := template.New("")
	for name, file := range Assets.Files {
		if file.IsDir() || !strings.HasSuffix(name, ".tmpl") {
			continue
		}
		h, err := ioutil.ReadAll(file)
		if err != nil {
			return nil, err
		}
		t, err = t.New(name).Parse(string(h))
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}
