package output

import (
	"github.com/shopspring/decimal"
	"time"
)

type Book struct {
	Id          uint             `json:"id"`
	CreateAt    int64             `json:"create_at"`
	BookName    string          `json:"book_name"`
	Price       decimal.Decimal `json:"price"`
	Author      string          `json:"author"`
	PublishTime time.Time       `json:"publish_time"`
	Desc        string          `json:"desc"`
}

