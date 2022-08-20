# [<img src="https://user-images.githubusercontent.com/41116859/185259360-b9b44eb2-6e47-4451-8d1e-90e3e4f34eef.png" height="32" /> Telegrambot](https://github.com/nickname76/telegrambot#example-usage) [<img class="badge" tag="github.com/nickname76/telegrambot" align="right" src="https://goreportcard.com/badge/github.com/nickname76/telegrambot">](https://goreportcard.com/report/github.com/nickname76/telegrambot)

<img src="https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fstatic.wixstatic.com%2Fmedia%2F950c70_eb49b9b040b14b70972c9777d736f7ea~mv2_d_2112_2112_s_2.gif&f=1&nofb=1" align="right" height="240" />

The most clean and strongly typed Telegram Bot API library in Go

Completely covers the latest Bot API version - **6.2**

Telegram Bot API documentaion: https://core.telegram.org/bots/api

Telegram deep links list: https://corefork.telegram.org/api/links (https://t.me/DeepLink)

**DISCORD**: [https://discord.gg/golang](https://discord.gg/rX4EhxsW6X) (#telegrambot channel)

Documentation on this library:
- **API** https://pkg.go.dev/github.com/nickname76/telegrambot
- **Tools** https://pkg.go.dev/github.com/nickname76/telegrambot/tools

*Please, **star** this repository, if you found this library useful!*

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
