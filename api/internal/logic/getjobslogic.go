package logic

import (
	"context"

	"cron-service/api/internal/svc"
	"cron-service/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetJobsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetJobsLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetJobsLogic {
	return GetJobsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetJobsLogic) GetJobs(req types.QueryJobsRequest) (*types.QueryJobsResponse, error) {
	// todo: add your logic here and delete this line

	return &types.QueryJobsResponse{}, nil
}
