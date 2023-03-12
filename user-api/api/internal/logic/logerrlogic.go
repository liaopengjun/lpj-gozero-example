package logic

import (
	"context"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
	"go-zero-demo/user-api/api/internal/svc"
)

type LogerrLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLogerrLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LogerrLogic {
	return &LogerrLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LogerrLogic) Logerr() error {
	// todo: add your logic here and delete this line
	if err := l.printLog(); err != nil {
		logx.Errorf("err:%+v \n", err)
	}
	return nil
}

func (l *LogerrLogic) printLog() error {
	return l.printLog2()
}

func (l *LogerrLogic) printLog2() error {
	return l.printLog3()
}

func (l *LogerrLogic) printLog3() error {
	return errors.Wrap(errors.New("这是一个日志错误"), " log failed")
}
