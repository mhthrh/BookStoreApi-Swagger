package Test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"testing"
)

func TestManual(t *testing.T) {
	var data bytes.Buffer
	json.NewEncoder(&data).Encode(b)

	//res, err := http.Post("http://localhost:8080/api/v1/book/add", "application/json", &data)
	res, err := http.Get("http://localhost:8080/api/v1/book/books/0-10")
	//res, err := http.Get("http://localhost:8080/api/v1/book/findbook/123456")
	//res, err := http.Get("http://localhost:8080/api/v1/book/delete/5454123456")
	//res, err := http.Post("http://localhost:8080/api/v1/book/update", "application/json", &data)

	if err != nil || res.StatusCode != 200 {
		log.Fatalf("Error : %v    | Status : %s", err, res.Status)
	}
	fmt.Println("Ok")
	fmt.Println(res.Body)

}
