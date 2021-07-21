package logic

import (
	"context"
	"github.com/pkg/errors"
	"github.com/zhuxinlei/micro_book/book/cmd/model/common"
	"github.com/zhuxinlei/micro_book/book/cmd/model/input"

	"github.com/tal-tech/go-zero/core/logx"
	"github.com/zhuxinlei/micro_book/book/cmd/api/internal/svc"
	"github.com/zhuxinlei/micro_book/book/cmd/api/internal/types"
)

type Search_bookLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearch_bookLogic(ctx context.Context, svcCtx *svc.ServiceContext) Search_bookLogic {
	return Search_bookLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Search_bookLogic) Search_book(req types.SearchBookReq) (*types.Reply, error) {
	// todo: add your logic here and delete this line
	book := input.Book{BookName: req.BookName}
	data,err := l.svcCtx.Book.GetBook(book)
	if err != nil{
		return nil, errors.Wrapf(common.GetBookError, common.ErrCodeMap[common.GetBookErrorCode])
	}


	return &types.Reply{
		Code:    200,
		Message: "",
		Data: data,
	}, nil
}
