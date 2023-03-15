package telegram

import (
	"context"

	"github.com/danilluk1/test-task-7/internal/telegram/commands"
	"github.com/danilluk1/test-task-7/internal/telegram/middlewares"
	"github.com/danilluk1/test-task-7/internal/telegram/types"
	"github.com/danilluk1/test-task-7/internal/types"
	"github.com/mr-linch/go-tg"
	"github.com/mr-linch/go-tg/tgb"
	"github.com/mr-linch/go-tg/tgb/session"
)

type telegramService struct {
	services *types.Services
	poller   *tgb.Poller
}

func NewTelegram(token string, services *types.Services) *telegramService {
	client := tg.New(token)

	sessionManager := session.NewManager(tg_types.Session{})

	router := tgb.NewRouter().
		Use(sessionManager).
		Use(&middlewares.LoggMiddleware{
			Services: services,
		})

	commandOpts := &tg_types.CommandOpts{
		Services:       services,
		Router:         router,
		SessionManager: sessionManager,
	}
	commands.NewGetInfoCommand(commandOpts)
	commands.NewGetStatsCommand(commandOpts)

	poller := tgb.NewPoller(router, client)

	return &telegramService{
		poller:   poller,
		services: services,
	}
}

func (t *telegramService) StartPolling(ctx context.Context) {
	go t.poller.Run(ctx)
}
