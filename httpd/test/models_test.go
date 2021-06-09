package test

import (
	"testing"

	"github.com/joshjwelsh/Social-Library/httpd/model"
)

func TestAddAll(t *testing.T) {

	feed := model.New()
	feed.Add(model.Book{})
	if len(feed.Books) == 0 {
		t.Error("Item was not added")
	}

}

func TestGetAll(t *testing.T) {
	feed := model.New()
	feed.Add(model.Book{})
	results := feed.GetAll()
	if len(results) == 0 {
		t.Error("Item was not added")
	}
}
