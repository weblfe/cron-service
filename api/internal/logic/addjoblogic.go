package logic

import (
	"context"

	"cron-service/api/internal/svc"
	"cron-service/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type AddJobLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddJobLogic(ctx context.Context, svcCtx *svc.ServiceContext) AddJobLogic {
	return AddJobLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddJobLogic) AddJob(req types.AddJobRequest) (*types.AddJobResponse, error) {
	// todo: add your logic here and delete this line

	return &types.AddJobResponse{}, nil
}
