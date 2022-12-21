package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	hello "github.com/hertz/hello/biz/hertz_gen/hertz/hello"
)

type Method1Service struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewMethod1Service(Context context.Context, RequestContext *app.RequestContext) *Method1Service {
	return &Method1Service{RequestContext: RequestContext, Context: Context}
}

func (h *Method1Service) Run(req *hello.HelloReq) (resp *hello.HelloResp) {
	//defer func() {
	//	hlog.CtxInfof(h.Context, "req = %+v", req)
	//	hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()

	// todo edit your code
	return
}
