package model

import (
	"github.com/zhuxinlei/micro_book/book/cmd/model/input"
	"github.com/zhuxinlei/micro_book/book/cmd/model/output"
)

type Book interface {
	Insert(book input.Book) error
	GetBook(book input.Book) (output.Book, error)
}
