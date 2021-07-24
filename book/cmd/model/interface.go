package model

import (
	"github.com/zhuxinlei/micro_book/book/cmd/model/input"
	"github.com/zhuxinlei/micro_book/book/cmd/rpc/book"
)

type Book interface {
	Insert(book input.Book) error
	GetBook(book input.Book) (book.BookInfoReply, error)
	GetBooks(ids []int) ([]*book.BookInfoReply, error)
}
