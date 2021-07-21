package svc

import (
	"github.com/tal-tech/go-zero/zrpc"
	"micro_book/book/cmd/rpc/bookclient"
	"micro_book/book/cmd/testrpc/internal/config"
)

type ServiceContext struct {
	Config  config.Config
	BookRpc bookclient.Book
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		BookRpc: bookclient.NewBook(zrpc.MustNewClient(c.BookRpc)),
	}
}
