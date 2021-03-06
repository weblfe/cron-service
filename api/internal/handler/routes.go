// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"cron-service/api/internal/svc"

	"github.com/tal-tech/go-zero/rest"
)

func RegisterHandlers(engine *rest.Server, serverCtx *svc.ServiceContext) {
	engine.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/service/cron/lists",
				Handler: GetJobsHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/service/cron/add",
				Handler: AddJobHandler(serverCtx),
			},
		},
	)
}
