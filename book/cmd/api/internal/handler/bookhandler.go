package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tal-tech/go-zero/rest/httpx"
	"github.com/zhuxinlei/micro_book/book/cmd/api/internal/logic"
	"github.com/zhuxinlei/micro_book/book/cmd/api/internal/svc"
	"github.com/zhuxinlei/micro_book/book/cmd/api/internal/types"
	"github.com/zhuxinlei/micro_book/book/cmd/model/common"
	"net/http"
)

func bookHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.InsertBookReq
		/*if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}*/
		c := &gin.Context{
			Request: r,
		}
		err := c.ShouldBindJSON(&req)

		if err != nil {
			fmt.Println("参数错误", err.Error())
			httpx.Error(w, err)
			return
		}
		l := logic.NewBookLogic(r.Context(), ctx)
		resp, err := l.Book(req)
		if err != nil {
			common.ServerError(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
