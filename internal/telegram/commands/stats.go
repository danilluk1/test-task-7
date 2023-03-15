package commands

import (
	"context"
	"errors"
	"fmt"

	tg_types "github.com/danilluk1/test-task-7/internal/telegram/types"
	"github.com/jackc/pgx/v5"
	"github.com/mr-linch/go-tg/tgb"
	"github.com/rs/zerolog/log"
)

type GetStatsCommand struct {
	*tg_types.CommandOpts
}

func (c *GetStatsCommand) HandleCommand(ctx context.Context, msg *tgb.MessageUpdate) error {
	chatID := msg.Chat.ID.PeerID()

	stats, err := c.Services.Store.GetStats(ctx, chatID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return msg.Answer("We don't have any info about you").DoVoid(ctx)
		} else {
			log.Info().Err(err).Msg("")
			return msg.Answer("Internal server error").DoVoid(ctx)
		}
	}
	return msg.Answer(fmt.Sprintf("First usage: %s, Total: %d", stats.CreatedAt.Time.String(), stats.Count)).DoVoid(ctx)
}

func NewGetStatsCommand(opts *tg_types.CommandOpts) {
	cmd := &GetStatsCommand{
		CommandOpts: opts,
	}

	opts.Router.Message(cmd.HandleCommand, tgb.Command("stats"))
}
