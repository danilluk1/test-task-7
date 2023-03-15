package commands

import (
	"context"
	"encoding/json"
	"errors"
	"strings"

	tg_types "github.com/danilluk1/test-task-7/internal/telegram/types"
	"github.com/jackc/pgx/v5"
	"github.com/mr-linch/go-tg/tgb"
	"github.com/rs/zerolog/log"
)

type GetInfoCommand struct {
	*tg_types.CommandOpts
}

func (c *GetInfoCommand) HandleCommand(ctx context.Context, msg *tgb.MessageUpdate) error {
	chatID := c.SessionManager.Get(ctx).ChatID

	_, err := c.Services.Store.GetStats(ctx, chatID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			c.Services.Store.CreateStats(ctx, chatID)
		} else {
			log.Info().Err(err)
			return msg.Answer("Internal server error").DoVoid(ctx)
		}
	}

	parts := strings.Split(msg.Text, " ")
	if len(parts) != 2 {
		return msg.Answer("Invalid request").DoVoid(ctx)
	}
	city := parts[1]

	forecast, err := c.Services.WeatherService.GetCurrentWeather(city)
	if err != nil {
		log.Info().Err(err)
		return err
	}

	c.Services.Store.UpdateCounter(ctx, chatID)
	forecastStr, err := json.Marshal(forecast)
	if err != nil {
		log.Info().Err(err)
		return err
	}
	return msg.Answer(string(forecastStr)).DoVoid(ctx)
}

func NewGetInfoCommand(opts *tg_types.CommandOpts) {
	cmd := &GetInfoCommand{
		CommandOpts: opts,
	}

	opts.Router.Message(cmd.HandleCommand, tgb.Command("info"))
}
