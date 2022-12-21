// this is my custom handler.

package hello

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	hello "github.com/hertz/hello/biz/hertz_gen/hertz/hello"
	"github.com/hertz/hello/biz/service"
)

// Method1 .
// @router /hello [GET]
func Method1(ctx context.Context, c *app.RequestContext) {
	//  you can code something
	var err error
	var req hello.HelloReq
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(400, err.Error())
		return
	}

	resp := service.NewMethod1Service(ctx, c).Run(&req)
	c.JSON(200, resp)
}
