package logic

import (
	"context"
	"errors"

	"github.com/igolang/igozero/app/apprpc"
	"github.com/igolang/igozero/app/internal/svc"

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

func (l *PingLogic) Ping(in *apprpc.PingRequest) (*apprpc.PingResponse, error) {
	if in.Ping == 1 {
		return &apprpc.PingResponse{Pong: "pong!"}, nil
	} else {
		return nil, errors.New("ping error")
	}
}
