package Test

import (
	"bytes"
	"encoding/json"
	"fmt"
	Book2 "github.com/mhthrh/ApiStore/Model/Book"
	"log"
	"net/http"
	"testing"
	"time"
)

func TestManual(t *testing.T) {

	var data bytes.Buffer
	json.NewEncoder(&data).Encode(Book2.Book{
		Id:          33,
		ISBN:        "SU8790",
		Title:       "title",
		Authors:     []string{"au1", "au2"},
		Publisher:   "Publisher",
		PublishDate: time.Now(),
		Pages:       200,
	})
	for i := 0; i < 10; i++ {
		res, err := http.Post("http://localhost:8080/books", "application/json", &data)

		//res, err := http.Get("http://localhost:8080/books/0-10")
		//res, err := http.Get("http://localhost:8080/books/123456")
		//res, err := http.Get("http://localhost:8080/books/5454123456")
		//res, err := http.Post("http://localhost:8080/books", "application/json", &data)

		if err != nil || res.StatusCode != 200 {
			log.Fatalf("Error : %v    | Status : %s", err, res.Status)
		}
		fmt.Println("Ok")
		fmt.Println(res.Body)
	}
}
