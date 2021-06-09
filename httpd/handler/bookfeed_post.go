package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joshjwelsh/Social-Library/httpd/model"
)

type bookfeedPostRequest struct {
	UserId     string `json:"Id"`
	Title      string `json:"title"`
	Desc       string `json:"description"`
	Author     string `json:"author"`
	DisplayPic []byte `json:"displayPic"`
}

func NewsfeedPost(feed *model.Feed) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := bookfeedPostRequest{}
		c.Bind(&requestBody)
		item := model.Book{
			UserId:     requestBody.UserId,
			Title:      requestBody.Title,
			Desc:       requestBody.Desc,
			Author:     requestBody.Author,
			DisplayPic: requestBody.DisplayPic,
		}
		feed.Add(item)
		c.Status(http.StatusNoContent) // 204 Response - no body but all is well
	}

}
