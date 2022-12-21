package service

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/test/assert"
	example "github.com/hertz/hello/biz/hertz_gen/hello/example"
	"testing"
)

func TestHelloMethod3Service_Run(t *testing.T) {
	ctx := context.Background()
	c := app.NewContext(1)
	s := NewHelloMethod3Service(ctx, c)
	// init req and assert value
	req := &example.HelloReq{}
	resp := s.Run(req)
	assert.DeepEqual(t, nil, resp)
	// todo edit your unit test.
}
