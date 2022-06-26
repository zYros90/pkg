package srvctx

import "context"

// server context
type Ctx struct {
	context.Context
	Username string
}

func New(ctx context.Context, username string) *Ctx {
	return &Ctx{
		ctx,
		username,
	}
}
