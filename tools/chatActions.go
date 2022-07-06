package tbtools

import (
	"fmt"
	"time"

	"github.com/nickname76/telegrambot"

	"github.com/nickname76/repeater"
)

// Starts continually send chat action every 4 seconds until stop function is called
func StartChatAction(api telegrambot.API, params *telegrambot.SendChatActionParams) (stop func(), err error) {
	err = api.SendChatAction(params)
	if err != nil {
		return nil, fmt.Errorf("StartChatAction: %w", err)
	}

	return repeater.StartRepeater(time.Second*4, func() {
		api.SendChatAction(params)
	}), nil
}
