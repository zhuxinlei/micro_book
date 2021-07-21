package handler

import (
	"net/http"

	"micro_book/book/cmd/testrpc/internal/logic"
	"micro_book/book/cmd/testrpc/internal/svc"
	"micro_book/book/cmd/testrpc/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func search_userHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SearchUserReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewSearch_userLogic(r.Context(), ctx)
		resp, err := l.Search_user(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
