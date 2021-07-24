package logic

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zhuxinlei/micro_book/book/cmd/model/common"
	"strconv"
	"strings"

	"github.com/zhuxinlei/micro_book/book/cmd/rpc/book"
	"github.com/zhuxinlei/micro_book/book/cmd/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetBooksLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetBooksLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBooksLogic {
	return &GetBooksLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetBooksLogic) GetBooks(in *book.IdsReq) (*book.BookInfosReply, error) {
	// todo: add your logic here and delete this line
	fmt.Println("我是book rpc server ---get books的业务逻辑")
	tempIds := strings.Split(in.Ids, ",")
	ids := make([]int, len(tempIds))
	for i, k := range tempIds {
		ids[i], _ = strconv.Atoi(k)
	}
	data, err := l.svcCtx.Book.GetBooks(ids)
	if err != nil {
		return nil, errors.Wrapf(common.GetBooksError, common.ErrCodeMap[common.GetBooksErrorCode])
	}

	return &book.BookInfosReply{
		Code:    200,
		Message: "",
		Data:    data,
	}, nil
}
