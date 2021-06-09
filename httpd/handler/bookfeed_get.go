package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joshjwelsh/Social-Library/httpd/model"
)

func BookfeedGet(feed *model.Feed) gin.HandlerFunc {
	return func(c *gin.Context) {
		results := feed.GetAll()
		c.JSON(http.StatusOK, results)

	}

}
