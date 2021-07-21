package logic

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"micro_book/book/cmd/model/common"
	"micro_book/book/cmd/rpc/bookclient"

	"micro_book/book/cmd/testrpc/internal/svc"
	"micro_book/book/cmd/testrpc/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
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
	bookInfo, err := l.svcCtx.BookRpc.GetBook(l.ctx,
		&bookclient.IdReq{
			Id: req.BookID,
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
