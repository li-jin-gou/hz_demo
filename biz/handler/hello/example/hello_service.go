// this is my custom handler.

package example

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	example "github.com/hertz/hello/biz/model/hello/example"
	"github.com/hertz/hello/biz/service"
)

// HelloMethod1 .
// @router /hello1 [GET]
func HelloMethod1(ctx context.Context, c *app.RequestContext) {
	//  you can code something
	var err error
	var req example.HelloReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	resp := service.NewHelloMethod1Service(ctx, c).Run(&req)
	c.JSON(200, resp)
}

// HelloMethod2 .
// @router /hello2 [GET]
func HelloMethod2(ctx context.Context, c *app.RequestContext) {
	//  you can code something
	var err error
	var req example.HelloReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	resp := service.NewHelloMethod2Service(ctx, c).Run(&req)
	c.JSON(200, resp)
}

// HelloMethod3 .
// @router /hello3 [GET]
func HelloMethod3(ctx context.Context, c *app.RequestContext) {
	//  you can code something
	var err error
	var req example.HelloReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	resp := service.NewHelloMethod3Service(ctx, c).Run(&req)
	c.JSON(200, resp)
}
