package telegrambot

// https://core.telegram.org/bots/api#updating-messages

import "fmt"

type EditMessageTextParams struct {
	// Optional. Required if inline_message_id is not specified. Unique
	// identifier for the target chat or username of the target channel (in the
	// format @channelusername)
	ChatID ChatIDOrUsername `json:"chat_id,omitempty"`
	// Optional. Required if inline_message_id is not specified. Identifier of
	// the message to edit
	MessageID MessageID `json:"message_id,omitempty"`
	// Optional. Required if chat_id and message_id are not specified.
	// Identifier of the inline message
	InlineMessageID InlineMessageID `json:"inline_message_id,omitempty"`
	// New text of the message, 1-4096 characters after entities parsing
	Text string `json:"text"`
	// Optional. Mode for parsing entities in the message text. See formatting
	// options for more details.
	ParseMode ParseMode `json:"parse_mode,omitempty"`
	// Optional. A JSON-serialized list of special entities that appear in
	// message text, which can be specified instead of parse_mode
	Entities []*MessageEntity `json:"entities,omitempty"`
	// Optional. Disables link previews for links in this message
	DisableWebPagePreview bool `json:"disable_web_page_preview,omitempty"`
	// Optional. A JSON-serialized object for an inline keyboard.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// Use this method to edit text and game messages. On success, if the edited
// message is not an inline message, the edited Message is returned, otherwise
// True is returned. https://core.telegram.org/bots/api#games
// https://core.telegram.org/bots/api#message
//
// https://core.telegram.org/bots/api#editmessagetext
func (api *API) EditMessageText(params *EditMessageTextParams) (*Message, error) {
	var msg *Message

	if params.InlineMessageID != "" {
		msg = &Message{}
	}

	migrateToChatID, err := api.makeAPICall("editMessageText", params, nil, msg)
	if err != nil {
		if migrateToChatID != 0 {
			params.ChatID = migrateToChatID
			_, err = api.makeAPICall("editMessageText", params, nil, msg)
			if err != nil {
				return nil, fmt.Errorf("EditMessageText: %w", err)
			}
		} else {
			return nil, fmt.Errorf("EditMessageText: %w", err)
		}
	}

	return msg, nil
}

type EditMessageCaptionParams struct {
	// Optional. Required if inline_message_id is not specified. Unique
	// identifier for the target chat or username of the target channel (in the
	// format @channelusername)
	ChatID ChatIDOrUsername `json:"chat_id,omitempty"`
	// Optional. Required if inline_message_id is not specified. Identifier of
	// the message to edit
	MessageID MessageID `json:"message_id,omitempty"`
	// Optional. Required if chat_id and message_id are not specified.
	// Identifier of the inline message
	InlineMessageID InlineMessageID `json:"inline_message_id,omitempty"`
	// Optional. New caption of the message, 0-1024 characters after entities
	// parsing
	Caption string `json:"caption,omitempty"`
	// Optional. Mode for parsing entities in the message caption. See
	// formatting options for more details.
	// https://core.telegram.org/bots/api#formatting-options
	ParseMode ParseMode `json:"parse_mode,omitempty"`
	// Optional. A JSON-serialized list of special entities that appear in the
	// caption, which can be specified instead of parse_mode
	CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"`
	// Optional. A JSON-serialized object for an inline keyboard.
	// https://core.telegram.org/bots#inline-keyboards-and-on-the-fly-updating
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// Use this method to edit captions of messages. On success, if the edited
// message is not an inline message, the edited Message is returned, otherwise
// True is returned. https://core.telegram.org/bots/api#message
//
// https://core.telegram.org/bots/api#editmessagecaption
func (api *API) EditMessageCaption(params *EditMessageCaptionParams) (*Message, error) {
	var msg *Message

	if params.InlineMessageID != "" {
		msg = &Message{}
	}

	migrateToChatID, err := api.makeAPICall("editMessageCaption", params, nil, msg)
	if err != nil {
		if migrateToChatID != 0 {
			params.ChatID = migrateToChatID
			_, err = api.makeAPICall("editMessageCaption", params, nil, msg)
			if err != nil {
				return nil, fmt.Errorf("EditMessageCaption: %w", err)
			}
		} else {
			return nil, fmt.Errorf("EditMessageCaption: %w", err)
		}
	}

	return msg, nil
}

type EditMessageMediaParams struct {
	// Optional. Required if inline_message_id is not specified. Unique
	// identifier for the target chat or username of the target channel (in the
	// format @channelusername)
	ChatID ChatIDOrUsername `json:"chat_id,omitempty"`
	// Optional. Required if inline_message_id is not specified. Identifier of
	// the message to edit
	MessageID MessageID `json:"message_id,omitempty"`
	// Optional. Required if chat_id and message_id are not specified.
	// Identifier of the inline message
	InlineMessageID InlineMessageID `json:"inline_message_id,omitempty"`
	// A JSON-serialized object for a new media content of the message
	Media *InputMedia `json:"media"`
	// Optional. A JSON-serialized object for a new inline keyboard.
	// https://core.telegram.org/bots#inline-keyboards-and-on-the-fly-updating
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// Use this method to edit animation, audio, document, photo, or video messages.
// If a message is part of a message album, then it can be edited only to an
// audio for audio albums, only to a document for document albums and to a photo
// or a video otherwise. When an inline message is edited, a new file can't be
// uploaded; use a previously uploaded file via its file_id or specify a URL. On
// success, if the edited message is not an inline message, the edited Message
// is returned, otherwise True is returned.
// https://core.telegram.org/bots/api#message
//
// https://core.telegram.org/bots/api#editmessagemedia
func (api *API) EditMessageMedia(params *EditMessageMediaParams) (*Message, error) {
	var msg *Message

	if params.InlineMessageID != "" {
		msg = &Message{}
	}

	migrateToChatID, err := api.makeAPICall("editMessageMedia", params, []InputFile{params.Media.Media, params.Media.Thumb}, msg)
	if err != nil {
		if migrateToChatID != 0 {
			params.ChatID = migrateToChatID
			_, err = api.makeAPICall("editMessageMedia", params, []InputFile{params.Media.Media, params.Media.Thumb}, msg)
			if err != nil {
				return nil, fmt.Errorf("EditMessageMedia: %w", err)
			}
		} else {
			return nil, fmt.Errorf("EditMessageMedia: %w", err)
		}
	}

	return msg, nil
}

type EditMessageReplyMarkupParams struct {
	// Optional. Required if inline_message_id is not specified. Unique
	// identifier for the target chat or username of the target channel (in the
	// format @channelusername)
	ChatID ChatIDOrUsername `json:"chat_id,omitempty"`
	// Optional. Required if inline_message_id is not specified. Identifier of
	// the message to edit
	MessageID MessageID `json:"message_id,omitempty"`
	// Optional. Required if chat_id and message_id are not specified.
	// Identifier of the inline message
	InlineMessageID InlineMessageID `json:"inline_message_id,omitempty"`
	// Optional. A JSON-serialized object for an inline keyboard.
	// https://core.telegram.org/bots#inline-keyboards-and-on-the-fly-updating
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// Use this method to edit only the reply markup of messages. On success, if the
// edited message is not an inline message, the edited Message is returned,
// otherwise True is returned. https://core.telegram.org/bots/api#message
//
// https://core.telegram.org/bots/api#editmessagereplymarkup
func (api *API) EditMessageReplyMarkup(params *EditMessageReplyMarkupParams) (*Message, error) {
	var msg *Message

	if params.InlineMessageID != "" {
		msg = &Message{}
	}

	migrateToChatID, err := api.makeAPICall("editMessageReplyMarkup", params, nil, msg)
	if err != nil {
		if migrateToChatID != 0 {
			params.ChatID = migrateToChatID
			_, err = api.makeAPICall("editMessageReplyMarkup", params, nil, msg)
			if err != nil {
				return nil, fmt.Errorf("EditMessageReplyMarkup: %w", err)
			}
		} else {
			return nil, fmt.Errorf("EditMessageReplyMarkup: %w", err)
		}
	}

	return msg, nil
}

type StopPollParams struct {
	// Unique identifier for the target chat or username of the target channel
	// (in the format @channelusername)
	ChatID ChatIDOrUsername `json:"chat_id"`
	// Identifier of the original message with the poll
	MessageID MessageID `json:"message_id"`
	// Optional. A JSON-serialized object for a new message inline keyboard.
	// https://core.telegram.org/bots#inline-keyboards-and-on-the-fly-updating
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// Use this method to stop a poll which was sent by the bot. On success, the
// stopped Poll is returned. https://core.telegram.org/bots/api#poll
//
// https://core.telegram.org/bots/api#stoppoll
func (api *API) StopPoll(params *StopPollParams) (*Poll, error) {
	poll := &Poll{}

	migrateToChatID, err := api.makeAPICall("stopPoll", params, nil, poll)
	if err != nil {
		if migrateToChatID != 0 {
			params.ChatID = migrateToChatID
			_, err = api.makeAPICall("stopPoll", params, nil, poll)
			if err != nil {
				return nil, fmt.Errorf("StopPoll: %w", err)
			}
		} else {
			return nil, fmt.Errorf("StopPoll: %w", err)
		}
	}

	return poll, nil
}

type DeleteMessageParams struct {
	// Unique identifier for the target chat or username of the target channel
	// (in the format @channelusername)
	ChatID ChatIDOrUsername `json:"chat_id"`
	// Identifier of the message to delete
	MessageID MessageID `json:"message_id"`
}

// Use this method to delete a message, including service messages, with the
// following limitations:
//   - A message can only be deleted if it was sent less than 48 hours ago.
//   - A dice message in a private chat can only be deleted if it was sent more than 24 hours ago.
//   - Bots can delete outgoing messages in private chats, groups, and supergroups.
//   - Bots can delete incoming messages in private chats.
//   - Bots granted can_post_messages permissions can delete outgoing messages in channels.
//   - If the bot is an administrator of a group, it can delete any message there.
//   - If the bot has can_delete_messages permission in a supergroup or a channel, it can delete any message there.
//
// Returns True on success.
//
// https://core.telegram.org/bots/api#deletemessage
func (api *API) DeleteMessage(params *DeleteMessageParams) error {
	migrateToChatID, err := api.makeAPICall("deleteMessage", params, nil, nil)
	if err != nil {
		if migrateToChatID != 0 {
			params.ChatID = migrateToChatID
			_, err = api.makeAPICall("deleteMessage", params, nil, nil)
			if err != nil {
				return fmt.Errorf("DeleteMessage: %w", err)
			}
		} else {
			return fmt.Errorf("DeleteMessage: %w", err)
		}
	}

	return nil
}
