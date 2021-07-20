package model

import (
	"micro_book/book/cmd/model/input"
	"micro_book/book/cmd/model/output"
)

type Book interface {
	Insert(book input.Book) error
	GetBook(book input.Book) (output.Book, error)
}
