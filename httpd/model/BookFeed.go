package model

type Book struct {
	UserId     string `json:"Id"`
	Title      string `json:"title"`
	Desc       string `json:"description"`
	Author     string `json:"author"`
	DisplayPic []byte `json:"displayPic"`
}

type Feed struct {
	Books []Book
}

func New() *Feed {
	return &Feed{
		Books: []Book{},
	}
}

func (f *Feed) Add(book Book) {
	f.Books = append(f.Books, book)
}

func (f *Feed) GetAll() []Book {
	return f.Books
}
