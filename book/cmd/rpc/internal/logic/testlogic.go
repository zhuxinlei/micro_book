package logic

import (
	"context"

	"micro_book/book/cmd/rpc/book"
	"micro_book/book/cmd/rpc/internal/svc"

	"github.com/tal-tech/go-zero/core/logx"
)

type TestLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TestLogic {
	return &TestLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TestLogic) Test(in *book.IdReq) (*book.BookInfoReply, error) {
	// todo: add your logic here and delete this line

	return &book.BookInfoReply{}, nil
}
