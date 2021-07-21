package logic

import (
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/zhuxinlei/micro_book/book/cmd/model/common"
	"github.com/zhuxinlei/micro_book/book/cmd/model/input"
	"github.com/zhuxinlei/micro_book/book/cmd/rpc/book"
	"github.com/zhuxinlei/micro_book/book/cmd/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetBookLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetBookLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBookLogic {
	return &GetBookLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetBookLogic) GetBook(in *book.IdReq) (*book.BookInfoReply, error) {
	// todo: add your logic here and delete this line
	fmt.Println("我是book rpc server的业务逻辑")

	data := input.Book{
		Model: gorm.Model{
			ID: uint(in.Id),
		},
	}
	res, err := l.svcCtx.Book.GetBook(data)
	if err != nil {
		return nil, errors.Wrapf(common.GetBookError, common.ErrCodeMap[common.GetBookErrorCode])
	}
	price, _ := res.Price.Float64()
	return &book.BookInfoReply{
		Id:                   int64(res.Id),
		BookName:             res.BookName,
		Price:                price,
		Author:               res.Author,
		PublishTime:          res.PublishTime.Unix(),
		Desc:                 res.Desc,
		XXX_NoUnkeyedLiteral: struct{}{},
		XXX_unrecognized:     nil,
		XXX_sizecache:        0,
	}, nil
}
