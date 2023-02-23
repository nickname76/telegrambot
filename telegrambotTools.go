package telegrambot

import (
	"fmt"
	"sort"
	"strings"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/nickname76/repeater"
)

// Starts receiving all possible updates from Telegram
func StartReceivingUpdates(api *API, receiver func(update *Update, err error)) (stop func()) {
	// By default telegram sends most types of updates, but not all, so here are
	// specified all types of updates
	return StartReceivingUpdatesWithParams(api, GetUpdatesParams{
		Timeout: 2,
		AllowedUpdates: []UpdateType{
			UpdateTypeMessage,
			UpdateTypeEditedMessage,
			UpdateTypeChannelPost,
			UpdateTypeEditedChannelPost,
			UpdateTypeInlineQuery,
			UpdateTypeChosenInlineResult,
			UpdateTypeCallbackQuery,
			UpdateTypeShippingQuery,
			UpdateTypePreCheckoutQuery,
			UpdateTypePoll,
			UpdateTypePollAnswer,
			UpdateTypeMyChatMember,
			UpdateTypeChatMember,
			UpdateTypeChatJoinRequest,
		},
	}, receiver)
}

// Starts receiving updates from Telegram with custom parameters. You should not
// pass offset field in params.
func StartReceivingUpdatesWithParams(api *API, params GetUpdatesParams, receiver func(update *Update, err error)) (stop func()) {
	stop = repeater.StartRepeater(0, func() {
		updates, err := api.GetUpdates(&params)
		if err != nil {
			receiver(nil, err)
			return
		}

		if len(updates) == 0 {
			return
		}

		updates = SortUpdates(updates)

		for _, update := range updates {
			receiver(update, nil)
		}

		params.Offset = updates[len(updates)-1].UpdateID + 1
	})

	return stop
}

// Use to parse body from Webhook request, used to receive updates
func ParseWebhookUpdate(body []byte) (*Update, error) {
	jsoniterCfg := jsoniter.Config{
		OnlyTaggedField:               true,
		ObjectFieldMustBeSimpleString: true,
		CaseSensitive:                 true,
	}.Froze()

	update := new(Update)

	err := jsoniterCfg.Unmarshal(body, update)
	if err != nil {
		return nil, fmt.Errorf("ParseWebhookUpdate: %w", err)
	}

	return update, nil
}

type updatesSortInterface []*Update

func (usi updatesSortInterface) Len() int {
	return len(usi)
}
func (usi updatesSortInterface) Less(i, j int) bool {
	return usi[i].UpdateID < usi[j].UpdateID
}
func (usi updatesSortInterface) Swap(i, j int) {
	usi[i], usi[j] = usi[j], usi[i]
}

// Used internally by StartReceivingUpdates. You can use it in custom update
// receivers, to sort updates by their UpdateID
func SortUpdates(updates []*Update) []*Update {
	sortedUpdates := updatesSortInterface(updates)
	sort.Sort(sortedUpdates)
	return []*Update(sortedUpdates)
}

// Compiles callback data in command-args type.
// Concatenates command and args with \x00 symbol
func CompileCbQryData(command, args string) string {
	if args == "" {
		return command
	}

	return command + "\x00" + args
}

// Decompiles callback data in command-args type.
// Use to decompile output from CompileCbQryData.
func DecompileCbQryData(cbQryData string) (command, args string) {
	data := strings.SplitN(cbQryData, "\x00", 2)
	command = data[0]
	if len(data) == 2 {
		args = data[1]
	}
	return
}

// Starts continually send chat action every 4 seconds until stop function is called
func StartChatAction(api API, params *SendChatActionParams) (stop func(), err error) {
	err = api.SendChatAction(params)
	if err != nil {
		return nil, fmt.Errorf("StartChatAction: %w", err)
	}

	return repeater.StartRepeater(time.Second*4, func() {
		api.SendChatAction(params)
	}), nil
}

// Returns command name from msg.
// If command not found, return nothing.
// If command is not placed at the start of a message, returns nothing.
func ParseMessageCommand(msg *Message) (command string, args string) {
	var (
		text         string
		textEntities []*MessageEntity
	)

	switch {
	case msg.Text != "":
		text = msg.Text
		textEntities = msg.Entities
	case msg.Caption != "":
		text = msg.Caption
		textEntities = msg.CaptionEntities
	default:
		return
	}

	for _, entity := range textEntities {
		if entity.Type != MessageEntityTypeBotCommand || entity.Offset != 0 {
			continue
		}

		command = text[1:entity.Length]

		usernameIndex := strings.Index(command, "@")
		if usernameIndex != -1 {
			command = command[:usernameIndex]
		}

		args = strings.TrimSpace(text[entity.Length:])

		break
	}

	return
}
