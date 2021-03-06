// Code generated by goctl. DO NOT EDIT.
package types

type QueryJobsRequest struct {
	Query string `query:"query,optional=1"`
	Page  int    `query:"page,optional=1,default=1"`
	Count int    `query:"count,optional=1,default=10"`
}

type AddJobRequest struct {
	Name    string            `form:"name,optional=1" desc:"任务名,可选"`
	Express string            `form:"express" desc:"定时表达式"`
	Handler string            `form:"handler" desc:"任务处理器"`
	Params  map[string]string `form:"params,optional=1" desc:"任务参数"`
}

type JobEntry struct {
	Id           int64             `json:"id"`
	Name         string            `json:"name"`
	Count        int64             `json:"count"`
	Status       int               `json:"status"`
	Express      string            `json:"express"`
	Handler      string            `json:"handler"`
	NextExecTime int64             `json:"nextExecTime"`
	Params       map[string]string `json:"params"`
}

type Aggregation struct {
	Total int `json:"total"`
	More  int `json:"more"`
	Page  int `json:"page"`
	Size  int `json:"size"`
	Count int `json:"count"`
}

type QueryResponseData struct {
	Items []*JobEntry  `json:"items"`
	Agg   *Aggregation `json:"agg"`
}

type QueryJobsResponse struct {
	Code int                `json:"code"`
	Msg  string             `json:"msg"`
	Data *QueryResponseData `json:"data"`
}

type AddJobResponseData struct {
	Id int64 `json:"id"`
}

type AddJobResponse struct {
	Code int                 `json:"code"`
	Msg  string              `json:"msg"`
	Data *AddJobResponseData `json:"data"`
}
