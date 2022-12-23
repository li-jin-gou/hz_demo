package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	example "github.com/li-jin-gou/hz_demo/hertz_gen/hello/example"
)

type HelloMethod2Service struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewHelloMethod2Service(Context context.Context, RequestContext *app.RequestContext) *HelloMethod2Service {
	return &HelloMethod2Service{RequestContext: RequestContext, Context: Context}
}

func (h *HelloMethod2Service) Run(req *example.HelloReq) (resp *example.HelloResp, err error) {
	//defer func() {
	//	hlog.CtxInfof(h.Context, "req = %+v", req)
	//	hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	return
}
