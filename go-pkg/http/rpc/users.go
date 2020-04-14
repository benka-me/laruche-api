package rpc

import (
	"context"
	"github.com/benka-me/users/go-pkg/users"
)

func (c *App) Register(ctx context.Context, req *users.RegisterReq) (*users.RegisterRes, error) {
	return c.Users.Register(ctx, req)
}

func (c *App) Login(ctx context.Context, req *users.LoginReq) (*users.LoginRes, error) {
	return c.Users.Login(ctx, req)
}

func (c *App) Get(ctx context.Context, req *users.GetReq) (*users.Data, error) {
	return c.Users.Get(ctx, req)
}

func (c *App) GetAll(ctx context.Context, req *users.Empty) (*users.All, error) {
	return c.Users.GetAll(ctx, req)
}

func (c *App) Auth(ctx context.Context, req *users.Token) (*users.IsAuth, error) {
	return c.Users.Auth(ctx, req)
}
