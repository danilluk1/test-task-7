package tg_types

import (
	"github.com/danilluk1/test-task-7/internal/types"
	"github.com/mr-linch/go-tg/tgb"
	"github.com/mr-linch/go-tg/tgb/session"
)

type Session struct {
	ChatID string
}

type CommandOpts struct {
	Services       *types.Services
	Router         *tgb.Router
	SessionManager *session.Manager[Session]
}

type MiddlewareOpts struct {
	Services       *types.Services
	SessionManager *session.Manager[Session]
}
