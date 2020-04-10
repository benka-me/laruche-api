package rpc

import (
	"context"
	"fmt"
	"github.com/benka-me/laruche-api/go-pkg/larapi"
)

func (c *App) HelloWorld(ctx context.Context, req *larapi.Request) (*larapi.Greeting, error) {

	return &larapi.Greeting{Msg: fmt.Sprintf("Hello %s", req.Msg)}, nil
}
