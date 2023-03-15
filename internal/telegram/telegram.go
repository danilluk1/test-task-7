package telegram

import (
	"context"
	"github.com/danilluk1/test-task-7/internal/types"
	"github.com/mr-linch/go-tg"
	"github.com/mr-linch/go-tg/tgb"
	"github.com/mr-linch/go-tg/tgb/session"
)

type telegramService struct {
	services *types.Services
	poller   *tgb.Poller
}

func NewTelegram(token string)
