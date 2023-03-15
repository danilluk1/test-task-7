package commands

import (
	"context"
	"github.com/danilluk1/test-task-7/internal/telegram/types"
	"github.com/mr-linch/go-tg"
	"github.com/mr-linch/go-tg/tgb"
)

type GetInfoCommand struct {
	*tg_types.CommandOpts
}

func (c *GetInfoCommand) HandleCommand(ctx context.Context, msg *tgb.MessageUpdate) error {

}
