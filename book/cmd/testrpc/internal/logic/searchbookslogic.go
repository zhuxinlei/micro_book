package logic

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"github.com/zhuxinlei/micro_book/book/cmd/model/common"
	"github.com/zhuxinlei/micro_book/book/cmd/rpc/bookclient"

	"github.com/zhuxinlei/micro_book/book/cmd/testrpc/internal/svc"
	"github.com/zhuxinlei/micro_book/book/cmd/testrpc/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type Search_booksLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearch_booksLogic(ctx context.Context, svcCtx *svc.ServiceContext) Search_booksLogic {
	return Search_booksLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Search_bookLogic) Search_books(req types.SearchBooksReq) (*types.Reply, error) {
	// todo: add your logic here and delete this line
	bookInfo, err := l.svcCtx.BookRpc.GetBooks(l.ctx,
		&bookclient.IdsReq{
			Ids: req.BookIDs,
		},
	)
	if err != nil {
		fmt.Println("book rpc server获取book info error")
		return nil, errors.Wrapf(common.GetBookError, common.ErrCodeMap[common.GetBookErrorCode])
	}
	return &types.Reply{
		Code:    200,
		Message: "",
		Data:    bookInfo,
	}, nil
}
