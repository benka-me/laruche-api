package rpc

import (
	"context"
	"github.com/benka-me/laruche/go-pkg/laruche"
)

func (c *App) PublishBee(ctx context.Context, req *laruche.PushBee) (*laruche.PushBeeRes, error) {
	if c.IsAuth(ctx, req.Token) {
		return c.Larsrv.PublishBee(ctx, req)
	}
	return &laruche.PushBeeRes{}, Unauthenticated()
}

func (c *App) PublishHive(ctx context.Context, req *laruche.PushHive) (*laruche.PushHiveRes, error) {
	if c.IsAuth(ctx, req.Token) {
		return c.Larsrv.PublishHive(ctx, req)
	}
	return &laruche.PushHiveRes{}, Unauthenticated()
}

func (c *App) PrivatizeBee(ctx context.Context, req *laruche.PushBee) (*laruche.PushBeeRes, error) {
	if c.IsAuth(ctx, req.Token) {
		return c.Larsrv.PrivatizeBee(ctx, req)
	}
	return &laruche.PushBeeRes{}, Unauthenticated()
}

func (c *App) PrivatizeHive(ctx context.Context, req *laruche.PushHive) (*laruche.PushHiveRes, error) {
	if c.IsAuth(ctx, req.Token) {
		return c.Larsrv.PrivatizeHive(ctx, req)
	}
	return &laruche.PushHiveRes{}, Unauthenticated()
}

func (c *App) SetBee(ctx context.Context, req *laruche.PushBee) (*laruche.PushBeeRes, error) {
	if c.IsAuth(ctx, req.Token) {
		return c.Larsrv.SetBee(ctx, req)
	}
	return &laruche.PushBeeRes{}, Unauthenticated()
}

func (c *App) SetHive(ctx context.Context, req *laruche.PushHive) (*laruche.PushHiveRes, error) {
	if c.IsAuth(ctx, req.Token) {
		return c.Larsrv.SetHive(ctx, req)
	}
	return &laruche.PushHiveRes{}, Unauthenticated()
}

func (c *App) GetBee(ctx context.Context, req *laruche.BeeReq) (*laruche.Bee, error) {
	if c.IsAuth(ctx, req.Token) {
		return c.Larsrv.GetBee(ctx, req)
	}
	return &laruche.Bee{}, Unauthenticated()
}

func (c *App) GetBees(ctx context.Context, req *laruche.BeesReq) (*laruche.Bees, error) {
	if c.IsAuth(ctx, req.Token) {
		return c.Larsrv.GetBees(ctx, req)
	}
	return &laruche.Bees{}, Unauthenticated()
}
