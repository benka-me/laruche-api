package rpc

import (
	"context"
	"errors"
	"github.com/benka-me/users/go-pkg/users"
)

func (c *App) IsAuth(ctx context.Context, req *users.Token) bool {
	res, _ := c.Users.Auth(ctx, req)
	return res.Val
}

func Unauthenticated() error {
	return errors.New("not authenticated")
}
