package logic

import (
	"context"

	"github.com/qiushenglei/gozero/rpc/zerorpc/internal/svc"
	"github.com/qiushenglei/gozero/rpc/zerorpc/zerorpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *zerorpc.Request) (*zerorpc.Response, error) {
	// todo: add your logic here and delete this line

	return &zerorpc.Response{}, nil
}
