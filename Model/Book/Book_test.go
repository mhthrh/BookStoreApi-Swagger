package Book

import (
	"reflect"
	"testing"
	"time"
)

func TestAddBook(t *testing.T) {
	type args struct {
		b Book
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "",
			args: args{b: Book{
				Id:          0,
				ISBN:        "EU98",
				Title:       "Book23",
				Authors:     []string{"aut1", "aut2"},
				Publisher:   "",
				PublishDate: time.Now(),
				Pages:       0,
			}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AddBook(tt.args.b)
		})
	}
}

func TestDeleteBook(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "test5",
			args:    args{id: 3},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DeleteBook(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteBook() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetBookByID(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		args    args
		want    *Book
		wantErr bool
	}{
		{
			name:    "test3",
			args:    args{id: 1},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetBookByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBookByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBookByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetBookByRange(t *testing.T) {
	type args struct {
		i int
		j int
	}
	tests := []struct {
		name    string
		args    args
		want    Books
		wantErr bool
	}{
		{
			name: "test2",
			args: args{
				i: 0,
				j: 3,
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetBookByRange(tt.args.i, tt.args.j)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetBookByRange() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBookByRange() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetBooks(t *testing.T) {
	tests := []struct {
		name string
		want Books
	}{
		{
			name: "test1",
			want: AllBooks,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetBooks(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBooks() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpdateBook(t *testing.T) {
	type args struct {
		b Book
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UpdateBook(tt.args.b); (err != nil) != tt.wantErr {
				t.Errorf("UpdateBook() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
