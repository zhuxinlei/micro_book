package svc

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/tal-tech/go-zero/core/logx"
	"micro_book/book/cmd/api/internal/config"
	"micro_book/book/cmd/model/table"
)

type ServiceContext struct {
	Config config.Config
	Book   table.BookModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	DB, err := gorm.Open("mysql", "root:12345678@tcp(127.0.0.1:3306)/micro_product?charset=utf8&parseTime=true&loc=UTC")

	if err != nil {
		logx.Errorf("连接数据库出错%s", err.Error())
	}
	book := table.NewBookModel(DB)
	return &ServiceContext{
		Config: c,
		Book:   book,
	}
}
