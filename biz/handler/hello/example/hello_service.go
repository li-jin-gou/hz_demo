// this is my custom handler.
package example

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/li-jin-gou/hz_demo/biz/service"
	"github.com/li-jin-gou/hz_demo/biz/utils"
	example "github.com/li-jin-gou/hz_demo/hertz_gen/hello/example"
)

// HelloMethod1 .
// @router /hello1 [GET]
func HelloMethod1(ctx context.Context, c *app.RequestContext) {
	var err error
	var req example.HelloReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewHelloMethod1Service(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// HelloMethod2 .
// @router /hello2 [GET]
func HelloMethod2(ctx context.Context, c *app.RequestContext) {
	var err error
	var req example.HelloReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewHelloMethod2Service(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}

// HelloMethod3 .
// @router /hello3 [GET]
func HelloMethod3(ctx context.Context, c *app.RequestContext) {
	var err error
	var req example.HelloReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	resp, err := service.NewHelloMethod3Service(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, resp)
}
