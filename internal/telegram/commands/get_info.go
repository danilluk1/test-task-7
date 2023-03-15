package commands

import (
	"context"
	"strings"

	"github.com/danilluk1/test-task-7/internal/telegram/types"
	"github.com/rs/zerolog/log"
	// "github.com/mr-linch/go-tg"
	"github.com/mr-linch/go-tg/tgb"
)

type GetInfoCommand struct {
	*tg_types.CommandOpts
}

func (c *GetInfoCommand) HandleCommand(ctx context.Context, msg *tgb.MessageUpdate) error {
	//chat := c.SessionManager.Get(ctx).ChatID

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
	return msg.Answer(forecast.Name).DoVoid(ctx)
}

func NewGetInfoCommand(opts *tg_types.CommandOpts) {
	cmd := &GetInfoCommand{
		CommandOpts: opts,
	}

	opts.Router.Message(cmd.HandleCommand, tgb.Command("info"))
}
