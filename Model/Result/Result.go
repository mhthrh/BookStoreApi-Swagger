package Result

import (
	"encoding/json"
	"fmt"
	"github.com/mhthrh/ApiStore/Model/Book"
	"net/http"
)

type ResultSet struct {
	Header `json:"Header"`
	Body   `json:"Body"`
}

type Header struct {
	Start             uint32 `json:"Start"`
	Finish            uint32 `json:"Finish"`
	ResultCode        uint32 `json:"ResultCode"`
	ResultDescription string `json:"ResultDescription"`
}

type Body struct {
	Books *[]Book.Book `json:"Books"`
}

func New(start, finish, resultCode uint32, resultDesc string, b *[]Book.Book) *ResultSet {
	r := new(ResultSet)
	r.Start = start
	r.Finish = finish
	r.ResultCode = resultCode
	r.ResultDescription = resultDesc
	r.Books = b
	return r
}

func (r *ResultSet) Create(w *http.ResponseWriter, i int) {
	byt, _ := json.Marshal(r)
	(*w).WriteHeader(i)
	fmt.Fprintln(*w, string(byt))

}
