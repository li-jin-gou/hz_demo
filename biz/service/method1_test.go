package service

import (
	"context"
	"testing"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	hello "github.com/hertz/hello/biz/hertz_gen/hertz/hello"
)

func TestMethod1Service_Run(t *testing.T) {
	ctx := context.Background()
	c := app.NewContext(1)
	s := NewMethod1Service(ctx, c)
	// init req and assert value
	req := &hello.HelloReq{}
	resp := s.Run(req)
	assert.DeepEqual(t, nil, resp)
	// todo edit your unit test.
}
