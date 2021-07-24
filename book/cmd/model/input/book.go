package input

import (
	"github.com/jinzhu/gorm"
	"github.com/shopspring/decimal"
	"github.com/zhuxinlei/micro_book/book/cmd/rpc/book"
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

func (b *Book) ToOutput() *book.BookInfoReply {

	price, _ := b.Price.Float64()
	return &book.BookInfoReply{
		Id:          int64(b.ID),
		BookName:    b.BookName,
		Price:       price,
		Author:      b.Author,
		PublishTime: b.PublishTime.Unix(),
		Desc:        b.Desc,
	}

}
