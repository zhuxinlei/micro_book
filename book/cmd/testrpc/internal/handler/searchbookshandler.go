package handler

import (
	"net/http"

	"github.com/zhuxinlei/micro_book/book/cmd/testrpc/internal/logic"
	"github.com/zhuxinlei/micro_book/book/cmd/testrpc/internal/svc"
	"github.com/zhuxinlei/micro_book/book/cmd/testrpc/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func search_booksHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SearchBooksReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewSearch_bookLogic(r.Context(), ctx)
		resp, err := l.Search_books(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
