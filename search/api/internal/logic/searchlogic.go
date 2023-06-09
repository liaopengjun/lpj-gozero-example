package logic

import (
	"context"
	"go-zero-demo/search/api/internal/svc"
	"go-zero-demo/search/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchLogic {
	return &SearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchLogic) Search(req *types.SearchReq) (resp *types.SearchReply, err error) {
	// todo: add your logic here and delete this line
	logx.Infof("用户ID: %v", l.ctx.Value("userId")) // 这里的key和生成jwt token时传入的key一致
	return &types.SearchReply{}, nil
	return
}
