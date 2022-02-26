package Book

import (
	"fmt"
	"github.com/mhthrh/ApiStore/Utility/DirectUtil"
	"github.com/mhthrh/ApiStore/Utility/FileUtil"
	"github.com/mhthrh/ApiStore/Utility/JsonUtil"
	"time"
)

// BookNotFound is an error raised when a Book can not be found in disk
var BookNotFound = fmt.Errorf("book not found")

var (
	path     string
	Name     string
	AllBooks Books
)

// Books defines a slice of book
type Books []*Book

// KeyBook is a key for the Book obj in the context
type KeyBook struct{}

func init() {
	dir := DirectUtil.Ut{}
	path, _ = dir.GetPath()
	Name = "Files/books.json"
	AllBooks = GetBooks()
}

// Book defines the structure for an API book
// swagger:model
type Book struct {
	// The ID for the book
	//
	// required: false
	// min: 1
	Id int `json:"Id" validate:"required"` // Unique identifier for the book
	// the ISBN for this book
	//
	// required: true
	// max length: 100
	ISBN string `json:"ISBN" validate:"required"`
	// the Title for this book
	//
	// required: true
	// max length: 10000
	Title string `json:"Title" validate:"required"`
	// the Authors for the book
	//
	// required: true
	// max length: 10000
	Authors []string `json:"Authors" validate:"required"`
	// the Publisher for the book
	//
	// required: true
	// max length: 10000
	Publisher string `json:"Publisher" validate:"required"`
	// the PublishDate for the book
	//
	// required: true
	// max length: 100
	PublishDate time.Time `json:"PublishDate" validate:"required"`
	// the Pages for the book
	//
	// required: true
	// max length: 5000
	Pages int `json:"Pages" validate:"required"`
}

// GetBooks returns all books from disk
func GetBooks() Books {

	var books []*Book
	s, err := FileUtil.New(path, Name).Read()
	if err != nil {

	}
	JsonUtil.New(nil, nil).Json2Struct(s, &books)

	return books

}

// GetBookByID returns a single book which matches the id from disk.
// If a book is not found this function returns a BookNotFound error
func GetBookByID(id int) (*Book, error) {
	if findBookIndex(id) == -1 {
		return nil, BookNotFound
	}
	return AllBooks[id], nil
}

// UpdateBook replaces a book in file with the given item.
// If a book with the given id does not exist in disk
// this function returns a BookNotFound error
func UpdateBook(b Book) error {
	if findBookIndex(int(b.Id)) == -1 {
		return BookNotFound
	}

	AllBooks[b.Id] = &b
	FileUtil.New(path, Name).Write(JsonUtil.New(nil, nil).Struct2Json(AllBooks))
	return nil
}

// AddBook adds a new book to disk
func AddBook(b Book) {
	if len(AllBooks) == 0 {
		b.Id = 0
	} else {
		b.Id = AllBooks[len(AllBooks)-1].Id + 1
	}
	AllBooks = append(AllBooks, &b)
	FileUtil.New(path, Name).Write(JsonUtil.New(nil, nil).Struct2Json(AllBooks))
}

// DeleteBook delete a book from disk
func DeleteBook(id int) error {
	i := findBookIndex(id)
	if i == -1 {
		return BookNotFound
	}
	AllBooks = append(AllBooks[:i], AllBooks[i+1])
	FileUtil.New(path, Name).Write(JsonUtil.New(nil, nil).Struct2Json(AllBooks))
	return nil
}

// findBookIndex finds the index of a book in disk
// returns -1 when no book can be found
func findBookIndex(id int) int {
	for i, p := range AllBooks {
		if p.Id == id {
			return i
		}
	}
	return -1
}
