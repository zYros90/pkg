package srvctx

import "context"

// server context
type SrvCtx struct {
	ctx      context.Context
	Username string
}

func NewSrvCtx(ctx context.Context, username string) *SrvCtx {
	return &SrvCtx{
		ctx:      ctx,
		Username: username,
	}
}
