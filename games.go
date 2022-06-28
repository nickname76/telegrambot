// https://core.telegram.org/bots/api#games
package telegrambot

import "fmt"

type SendGameParams struct {
	// Unique identifier for the target chat
	ChatID ChatID `json:"chat_id"`
	// Short name of the game, serves as the unique identifier for the game. Set
	// up your games via Botfather. https://t.me/botfather
	GameShortName GameShortName `json:"game_short_name"`
	// Optional. Sends the message silently. Users will receive a notification
	// with no sound. https://telegram.org/blog/channels-2-0#silent-messages
	DisableNotification bool `json:"disable_notification,omitempty"`
	// Optional. Protects the contents of the sent message from forwarding and
	// saving
	ProtectContent bool `json:"protect_content,omitempty"`
	// Optional. If the message is a reply, ID of the original message
	ReplyToMessageID MessageID `json:"reply_to_message_id,omitempty"`
	// Optional. Pass True, if the message should be sent even if the specified
	// replied-to message is not found
	AllowSendingWithoutReply bool `json:"allow_sending_without_reply,omitempty"`
	// Optional. A JSON-serialized object for an inline keyboard. If empty, one
	// 'Play game_title' button will be shown. If not empty, the first button
	// must launch the game.
	// https://core.telegram.org/bots#inline-keyboards-and-on-the-fly-updating
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// Use this method to send a game. On success, the sent Message is returned.
// https://core.telegram.org/bots/api#message
//
// https://core.telegram.org/bots/api#sendgame
func (api *API) SendGame(params *SendGameParams) (*Message, error) {
	msg := &Message{}

	migrateToChatID, err := api.makeAPICall("sendGame", params, nil, msg)
	if err != nil {
		if migrateToChatID != 0 {
			params.ChatID = migrateToChatID
			_, err = api.makeAPICall("sendGame", params, nil, msg)
			if err != nil {
				return nil, fmt.Errorf("SendGame: %w", err)
			}
		} else {
			return nil, fmt.Errorf("SendGame: %w", err)
		}
	}

	return msg, nil
}

// This object represents a game. Use BotFather to create and edit games, their
// short names will act as unique identifiers.
//
// https://core.telegram.org/bots/api#game
type Game struct {
	// Title of the game
	Title string `json:"title"`
	// Description of the game
	Description string `json:"description"`
	// Photo that will be displayed in the game message in chats.
	Photo []*PhotoSize `json:"photo"`
	// Optional. Brief description of the game or high scores included in the
	// game message. Can be automatically edited to include current high scores
	// for the game when the bot calls setGameScore, or manually edited using
	// editMessageText. 0-4096 characters.
	// https://core.telegram.org/bots/api#setgamescore
	// https://core.telegram.org/bots/api#editmessagetext
	Text string `json:"text,omitempty"`
	// Optional. Special entities that appear in text, such as usernames, URLs,
	// bot commands, etc.
	RTxtEntities []*MessageEntity `json:"text_entities,omitempty"`
	// Optional. Animation that will be displayed in the game message in chats.
	// Upload via BotFather https://t.me/botfather
	Animation *Animation `json:"animation,omitempty"`
}

// A placeholder, currently holds no information. Use BotFather to set up your
// game. https://t.me/botfather
//
// https://core.telegram.org/bots/api#callbackgame
type CallbackGame struct{}

type SetGameScoreParams struct {
	// User identifier
	UserID UserID `json:"user_id"`
	// New score, must be non-negative
	Score int `json:"score"`
	// Optional. Pass True, if the high score is allowed to decrease. This can
	// be useful when fixing mistakes or banning cheaters
	Force bool `json:"force,omitempty"`
	// Optional. Pass True, if the game message should not be automatically
	// edited to include the current scoreboard
	DisableEditMessage bool `json:"disable_edit_message,omitempty"`
	// Optional. Required if inline_message_id is not specified. Unique
	// identifier for the target chat
	ChatID ChatID `json:"chat_id,omitempty"`
	// Optional. Required if inline_message_id is not specified. Identifier of
	// the sent message
	MessageID MessageID `json:"message_id,omitempty"`
	// Optional. Required if chat_id and message_id are not specified.
	// Identifier of the inline message
	InlineMessageID InlineMessageID `json:"inline_message_id,omitempty"`
}

// Use this method to set the score of the specified user in a game message. On
// success, if the message is not an inline message, the Message is returned,
// otherwise True is returned. Returns an error, if the new score is not greater
// than the user's current score in the chat and force is False.
// https://core.telegram.org/bots/api#message
//
// https://core.telegram.org/bots/api#setgamescore
func (api *API) SetGameScore(params *SetGameScoreParams) (*Message, error) {
	var msg *Message

	if params.InlineMessageID != "" {
		msg = &Message{}
	}

	migrateToChatID, err := api.makeAPICall("setGameScore", params, nil, msg)
	if err != nil {
		if migrateToChatID != 0 {
			params.ChatID = migrateToChatID
			_, err = api.makeAPICall("setGameScore", params, nil, msg)
			if err != nil {
				return nil, fmt.Errorf("SetGameScore: %w", err)
			}
		} else {
			return nil, fmt.Errorf("SetGameScore: %w", err)
		}
	}

	return msg, nil
}

type GetGameHighScoresParams struct {
	// Target user id
	UserID UserID `json:"user_id"`
	// Optional. Required if inline_message_id is not specified. Unique
	// identifier for the target chat
	ChatID ChatID `json:"chat_id,omitempty"`
	// Optional. Required if inline_message_id is not specified. Identifier of
	// the sent message
	MessageID MessageID `json:"message_id,omitempty"`
	// Optional. Required if chat_id and message_id are not specified.
	// Identifier of the inline message
	InlineMessageID InlineMessageID `json:"inline_message_id,omitempty"`
}

// Use this method to get data for high score tables. Will return the score of
// the specified user and several of their neighbors in a game. On success,
// returns an Array of GameHighScore objects.
// https://core.telegram.org/bots/api#gamehighscore
//
// This method will currently return scores for the target user, plus two of
// their closest neighbors on each side. Will also return the top three users if
// the user and his neighbors are not among them. Please note that this behavior
// is subject to change.
//
// https://core.telegram.org/bots/api#getgamehighscores
func (api *API) GetGameHighScores(params *GetGameHighScoresParams) ([]*GameHighScore, error) {
	gameHighScores := []*GameHighScore{}

	migrateToChatID, err := api.makeAPICall("getGameHighScores", params, nil, &gameHighScores)
	if err != nil {
		if migrateToChatID != 0 {
			params.ChatID = migrateToChatID
			_, err = api.makeAPICall("getGameHighScores", params, nil, &gameHighScores)
			if err != nil {
				return nil, fmt.Errorf("GetGameHighScores: %w", err)
			}
		} else {
			return nil, fmt.Errorf("GetGameHighScores: %w", err)
		}
	}

	return gameHighScores, nil
}

// This object represents one row of the high scores table for a game.
//
// https://core.telegram.org/bots/api#gamehighscore
type GameHighScore struct {
	// Position in high score table for the game
	Position int `json:"position"`
	// User
	User *User `json:"user"`
	// Score
	Score int `json:"score"`
}
