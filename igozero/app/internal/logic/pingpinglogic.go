package logic

import (
	"context"
	"errors"
	"github.com/igolang/igozero/app/apprpc"
	"github.com/igolang/igozero/app/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingPingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingPingLogic {
	return &PingPingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingPingLogic) PingPing(in *apprpc.PingRequest) (*apprpc.Response, error) {
	if in.Ping == 1 {
		return ok(&apprpc.PingResponse{Pong: "pong!"})
	} else {
		return abort(errors.New("ping error"))
	}
}
