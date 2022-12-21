package service

import (
	"context"
	"testing"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	example "github.com/hertz/hello/biz/hertz_gen/hello/example"
)

func TestHelloMethod1Service_Run(t *testing.T) {
	ctx := context.Background()
	c := app.NewContext(1)
	s := NewHelloMethod1Service(ctx, c)
	// init req and assert value
	req := &example.HelloReq{}
	resp := s.Run(req)
	assert.DeepEqual(t, nil, resp)
	// todo edit your unit test.
}
