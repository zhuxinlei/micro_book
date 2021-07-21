package logic

import (
	"context"

	"micro_book/book/cmd/testrpc/internal/svc"
	"micro_book/book/cmd/testrpc/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type Search_userLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearch_userLogic(ctx context.Context, svcCtx *svc.ServiceContext) Search_userLogic {
	return Search_userLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Search_userLogic) Search_user(req types.SearchUserReq) (*types.Reply, error) {
	// todo: add your logic here and delete this line

	return &types.Reply{}, nil
}
