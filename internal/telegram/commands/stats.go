package commands

import (
	"context"
	"github.com/danilluk1/test-task-7/internal/telegram/types"
	// "github.com/mr-linch/go-tg"
	"github.com/mr-linch/go-tg/tgb"
)

type GetStatsCommand struct {
	*tg_types.CommandOpts
}

func (c *GetStatsCommand) HandleCommand(ctx context.Context, msg *tgb.MessageUpdate) error {
	return nil
}

func NewGetStatsCommand(opts *tg_types.CommandOpts) {
	cmd := &GetStatsCommand{
		CommandOpts: opts,
	}

	opts.Router.Message(cmd.HandleCommand, tgb.Command("stats"))
}
