# Telegrambot

Telegram Bot API library in Go.

The easest to use, the most clean, and strongly typed Telegram Bot API library in Go.

This repository also contains directory `tools`, which contains useful functions for Telegram bot development with this library. See https://pkg.go.dev/github.com/nickname76/telegrambot/tools

Also check out Telegram deep links: https://t.me/DeepLink

Telegram Bot API documentaion: https://core.telegram.org/bots/api

Documentation: https://pkg.go.dev/github.com/nickname76/telegrambot

*Please, **star** this repository, if you found this library useful.*

## Example usage

```Go
package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/nickname76/telegrambot"
)

func main() {
	api, me, err := telegrambot.NewAPI("YOUR_TELEGRAM_BOT_API_TOKEN")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	stop := telegrambot.StartReceivingUpdates(api, func(update *telegrambot.Update, err error) {
		if err != nil {
			log.Printf("Error: %v", err)
			return
		}

		msg := update.Message
		if msg == nil {
			return
		}

		_, err = api.SendMessage(&telegrambot.SendMessageParams{
			ChatID: msg.Chat.ID,
			Text:   fmt.Sprintf("Hello %v, I am %v", msg.From.FirstName, me.FirstName),
			ReplyMarkup: &telegrambot.ReplyKeyboardMarkup{
				Keyboard: [][]*telegrambot.KeyboardButton{{
					{
						Text: "Hello",
					},
				}},
				ResizeKeyboard:  true,
				OneTimeKeyboard: true,
			},
		})

		if err != nil {
			log.Printf("Error: %v", err)
			return
		}
	})

	log.Printf("Started on %v", me.Username)

	exitCh := make(chan os.Signal, 1)
	signal.Notify(exitCh, os.Interrupt)

	<-exitCh

	// Waits for all updates handling to complete
	stop()
}

```

## See also

If you want to develop Telegram bot, you should also see these libraries, which might be useful for you

- [Repeater](https://github.com/nickname76/repeater) - Go library for creating repeating function calls
- [Instorage](https://github.com/nickname76/instorage) - Simple, easy to use database for faster development of small projects and MVPs in Go. Uses Badger as a storage.
- [Locstrs](https://github.com/nickname76/locstrs) - Strings localisation library in Go
