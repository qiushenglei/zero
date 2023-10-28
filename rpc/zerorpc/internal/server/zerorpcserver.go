// Code generated by goctl. DO NOT EDIT.
// Source: zerorpc.proto

package server

import (
	"context"

	"github.com/qiushenglei/gozero/rpc/zerorpc/internal/logic"
	"github.com/qiushenglei/gozero/rpc/zerorpc/internal/svc"
	"github.com/qiushenglei/gozero/rpc/zerorpc/zerorpc"
)

type ZerorpcServer struct {
	svcCtx *svc.ServiceContext
	zerorpc.UnimplementedZerorpcServer
}

func NewZerorpcServer(svcCtx *svc.ServiceContext) *ZerorpcServer {
	return &ZerorpcServer{
		svcCtx: svcCtx,
	}
}

func (s *ZerorpcServer) Ping(ctx context.Context, in *zerorpc.Request) (*zerorpc.Response, error) {
	l := logic.NewPingLogic(ctx, s.svcCtx)
	return l.Ping(in)
}