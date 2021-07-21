package logic

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/tal-tech/go-zero/core/logx"
	"github.com/zhuxinlei/micro_book/book/cmd/api/internal/svc"
	"github.com/zhuxinlei/micro_book/book/cmd/api/internal/types"
	"github.com/zhuxinlei/micro_book/book/cmd/model/common"
	"github.com/zhuxinlei/micro_book/book/cmd/model/input"
)

type BookLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBookLogic(ctx context.Context, svcCtx *svc.ServiceContext) BookLogic {
	return BookLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BookLogic) Book(req types.InsertBookReq) (*types.Reply, error) {
	// todo: add your logic here and delete this line
	data := input.Book{
		BookName:    req.BookName,
		Price:       req.Price,
		Author:      req.Author,
		PublishTime: req.PublishTime,
		Desc:        req.Desc,
	}
	err := l.svcCtx.Book.Insert(data)
	if err != nil {
		fmt.Println(err)
		return nil, errors.Wrapf(common.InsertBookError, common.ErrCodeMap[common.InsertBookErrorCode])
	}
	return &types.Reply{
		Code:    200,
		Message: "add book success",
	}, nil
}
