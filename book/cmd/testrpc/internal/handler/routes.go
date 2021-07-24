// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"github.com/zhuxinlei/micro_book/book/cmd/testrpc/internal/svc"

	"github.com/tal-tech/go-zero/rest"
)

func RegisterHandlers(engine *rest.Server, serverCtx *svc.ServiceContext) {
	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/search/book",
				Handler: search_bookHandler(serverCtx),
			},
		},
	)
	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/search/books",
				Handler: search_booksHandler(serverCtx),
			},
		},
	)
	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/search/user",
				Handler: search_userHandler(serverCtx),
			},
		},
	)
}
