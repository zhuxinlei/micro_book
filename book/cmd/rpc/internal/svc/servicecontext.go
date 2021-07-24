package svc

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/zhuxinlei/micro_book/book/cmd/model"
	"github.com/zhuxinlei/micro_book/book/cmd/model/table"
	"github.com/zhuxinlei/micro_book/book/cmd/rpc/internal/config"
)

type ServiceContext struct {
	Config config.Config
	Book   model.Book
}

func NewServiceContext(c config.Config) *ServiceContext {
	DB, err := gorm.Open("mysql", "root:12345678@tcp(127.0.0.1:3306)/micro_product?charset=utf8&parseTime=true")
	if err != nil {
		logx.Errorf("连接数据库出错%s", err.Error())
	}
	book := table.NewBookModel(DB)
	return &ServiceContext{
		Config: c,
		Book:   book,
	}
}
