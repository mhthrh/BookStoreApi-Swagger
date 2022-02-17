package Book

import (
	"encoding/json"
	"errors"
	"github.com/mhthrh/ApiStore/Helper"
	"path/filepath"
	"time"
)

var (
	path  string
	Books *[]Book
)

func init() {
	//for test you must change path
	d, _ := Helper.GetPath()
	path = filepath.Join(d, "Files/books.json")
	//path = "D:\\Projects\\Go-lang\\BookStore\\Files\\Books.json"
	Books = load()
}

type Book struct {
	Id          uint32    `json:"Id"`
	ISBN        string    `json:"ISBN"`
	Title       string    `json:"Title"`
	Authors     []string  `json:"Authors"`
	Publisher   string    `json:"Publisher"`
	PublishDate time.Time `json:"PublishDate"`
	Pages       int       `json:"Pages"`
}

func load() *[]Book {
	var books []Book
	s, err := Helper.Read(path)
	if err != nil {

	}
	err = json.Unmarshal([]byte(s), &books)
	if err != nil {

	}
	return &books
}

func Load(i, j int) *[]Book {
	var b []Book
	for _, book := range *Books {
		if book.Id >= uint32(i) && uint32(j) >= book.Id {
			b = append(b, book)
		}
	}
	return &b
}

func Find(isbn string) (*Book, int) {
	for i, book := range *Books {
		if book.ISBN == isbn {
			return &book, i
		}
	}
	return nil, 0
}

func Add(b Book) error {
	if bo, _ := Find(string(b.ISBN)); bo != nil {
		return errors.New("already exist")
	}
	*Books = append(*Books, b)
	a, err := json.Marshal(*Books)
	if err != nil {
		return errors.New("check data")
	}
	Helper.Write(string(a), path)
	Books = load()
	return nil
}

func Update(b Book) error {
	if err := Delete(b.ISBN); err != nil {
		return err
	}
	if err := Add(b); err != nil {
		return err
	}
	return nil
}

func Delete(isbn string) error {
	b, i := Find(isbn)
	if b == nil {
		return errors.New("Not found, ")
	}
	*Books = removeIndex(*Books, i)

	a, err := json.Marshal(Books)
	if err != nil {
		return err
	}
	Helper.Write(string(a), path)
	Books = load()

	return nil
}

func removeIndex(s []Book, index int) []Book {
	return append(s[:index], s[index+1:]...)
}
