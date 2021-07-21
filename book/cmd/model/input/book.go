package input

import (
	"github.com/jinzhu/gorm"
	"github.com/shopspring/decimal"
	"github.com/zhuxinlei/micro_book/book/cmd/model/output"
	"time"
)

type Book struct {
	gorm.Model
	BookName    string          `json:"book_name"`
	Price       decimal.Decimal `json:"price"`
	Author      string          `json:"author"`
	PublishTime time.Time       `json:"publish_time"`
	Desc        string          `json:"desc"`
}

func (b *Book) ToOutput() output.Book {

	return output.Book{
		Id:          b.ID,
		CreateAt:    b.CreatedAt.Unix(),
		BookName:    b.BookName,
		Price:       b.Price,
		Author:      b.Author,
		PublishTime: b.PublishTime,
		Desc:        b.Desc,
	}
}
