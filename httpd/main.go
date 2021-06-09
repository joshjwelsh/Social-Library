package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joshjwelsh/Social-Library/httpd/handler"
	"github.com/joshjwelsh/Social-Library/httpd/model"
)

func main() {
	feed := model.New()
	r := gin.Default()

	r.GET("/bookfeed", handler.BookfeedGet(feed))
	r.POST("/bookfeed", handler.NewsfeedPost(feed))

	r.Run()
}
