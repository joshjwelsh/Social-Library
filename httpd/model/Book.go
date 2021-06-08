package model

type Book struct {
	UserId     string `json:"Id"`
	Title      string `json:"title"`
	Desc       string `json:"description"`
	Author     string `json:"author"`
	DisplayPic []byte `json:"displayPic"`
}
