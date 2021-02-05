package handler

import (
	"net/http"

	"cron-service/api/internal/logic"
	"cron-service/api/internal/svc"
	"cron-service/api/internal/types"

	"github.com/tal-tech/go-zero/rest/httpx"
)

func GetJobsHandler(ctx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.QueryJobsRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGetJobsLogic(r.Context(), ctx)
		resp, err := l.GetJobs(req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
