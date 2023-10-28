package logic

import (
	"context"

	"github.com/qiushenglei/gozero/http/zerohttp/internal/svc"
	"github.com/qiushenglei/gozero/http/zerohttp/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ZerohttpLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewZerohttpLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ZerohttpLogic {
	return &ZerohttpLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ZerohttpLogic) Zerohttp(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
