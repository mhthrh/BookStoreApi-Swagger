package Test

import (
	"github.com/mhthrh/ApiStore/Model/Book"
	"reflect"
	"testing"
	"time"
)

func TestAdd(t *testing.T) {
	type args struct {
		b Book.Book
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Add test",
			args: args{b: Book.Book{
				Id:          3,
				ISBN:        "EU-3456",
				Title:       "book3",
				Authors:     []string{"Aut1", "Aut2", "Aut2"},
				Publisher:   "pub3",
				PublishDate: time.Now(),
				Pages:       220,
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		if err := Book.Add(tt.args.b); (err != nil) != tt.wantErr {
			t.Errorf("Add() error = %v, wantErr %v", err, tt.wantErr)
		}

	}
}

func TestDelete(t *testing.T) {
	type args struct {
		isbn string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "test Delete",
			args:    args{isbn: "a9a28dd8-8f1a-11ec-901c-f8a9634eef20"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		if err := Book.Delete(tt.args.isbn); (err != nil) != tt.wantErr {
			t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
		}
	}
}

func TestFind(t *testing.T) {
	type args struct {
		isbn string
	}
	tests := []struct {
		name  string
		args  args
		want  *Book.Book
		want1 int
	}{
		{
			name:  "Find",
			args:  args{isbn: "b17e88f4-8f1b-11ec-9834-f8a9634eef20"},
			want:  nil,
			want1: 0,
		},
	}
	for _, tt := range tests {
		got, got1 := Book.Find(tt.args.isbn)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("Find() got = %v, want %v", got, tt.want)
		}
		if got1 != tt.want1 {
			t.Errorf("Find() got1 = %v, want %v", got1, tt.want1)
		}
	}
}

func TestUpdate(t *testing.T) {
	type args struct {
		b Book.Book
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Update test",
			args: args{b: Book.Book{
				Id:          0,
				ISBN:        "",
				Title:       "",
				Authors:     nil,
				Publisher:   "",
				PublishDate: time.Time{},
				Pages:       0,
			}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		if err := Book.Update(tt.args.b); (err != nil) != tt.wantErr {
			t.Errorf("Update() error = %v, wantErr %v", err, tt.wantErr)
		}
	}
}
