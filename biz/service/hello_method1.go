package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	example "github.com/hertz/hello/biz/hertz_gen/hello/example"
)

type HelloMethod1Service struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewHelloMethod1Service(Context context.Context, RequestContext *app.RequestContext) *HelloMethod1Service {
	return &HelloMethod1Service{RequestContext: RequestContext, Context: Context}
}

func (h *HelloMethod1Service) Run(req *example.HelloReq) (resp *example.HelloResp) {
	//defer func() {
	//	hlog.CtxInfof(h.Context, "req = %+v", req)
	//	hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()

	// todo edit your code
	return
}
