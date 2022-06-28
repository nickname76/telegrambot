// https://core.telegram.org/bots/api#getting-updates
package telegrambot

import "fmt"

// This object represents an incoming update. At most one of the optional
// parameters can be present in any given update.
// https://core.telegram.org/bots/api#available-types
//
// https://core.telegram.org/bots/api#update
type Update struct {
	// The update's unique identifier. Update identifiers start from a certain
	// positive number and increase sequentially. This ID becomes especially
	// handy if you're using Webhooks, since it allows you to ignore repeated
	// updates or to restore the correct update sequence, should they get out of
	// order. If there are no new updates for at least a week, then identifier
	// of the next update will be chosen randomly instead of sequentially.
	// https://core.telegram.org/bots/api#setwebhook
	UpdateID UpdateID `json:"update_id"`
	// Optional. New incoming message of any kind — text, photo, sticker, etc.
	Message *Message `json:"message,omitempty"`
	// Optional. New version of a message that is known to the bot and was
	// edited
	EditedMessage *Message `json:"edited_message,omitempty"`
	// Optional. New incoming channel post of any kind — text, photo, sticker,
	// etc.
	ChannelPost *Message `json:"channel_post,omitempty"`
	// Optional. New version of a channel post that is known to the bot and was
	// edited
	EditedChannelPost *Message `json:"edited_channel_post,omitempty"`
	// Optional. New incoming inline query
	// https://core.telegram.org/bots/api#inline-mode
	InlineQuery *InlineQuery `json:"inline_query,omitempty"`
	// Optional. The result of an inline query that was chosen by a user and
	// sent to their chat partner. Please see our documentation on the feedback
	// collecting for details on how to enable these updates for your bot.
	// https://core.telegram.org/bots/api#inline-mode
	// https://core.telegram.org/bots/inline#collecting-feedback
	ChosenInlineResult *ChosenInlineResult `json:"chosen_inline_result,omitempty"`
	// Optional. New incoming callback query
	CallbackQuery *CallbackQuery `json:"callback_query,omitempty"`
	// Optional. New incoming shipping query. Only for invoices with flexible
	// price
	ShippingQuery *ShippingQuery `json:"shipping_query,omitempty"`
	// Optional. New incoming pre-checkout query. Contains full information
	// about checkout
	PreCheckoutQuery *PreCheckoutQuery `json:"pre_checkout_query,omitempty"`
	// Optional. New poll state. Bots receive only updates about stopped polls
	// and polls, which are sent by the bot
	Poll *Poll `json:"poll,omitempty"`
	// Optional. A user changed their answer in a non-anonymous poll. Bots
	// receive new votes only in polls that were sent by the bot itself.
	PollAnswer *PollAnswer `json:"poll_answer,omitempty"`
	// Optional. The bot's chat member status was updated in a chat. For private
	// chats, this update is received only when the bot is blocked or unblocked
	// by the user.
	MyChatMember *ChatMemberUpdated `json:"my_chat_member,omitempty"`
	// Optional. A chat member's status was updated in a chat. The bot must be
	// an administrator in the chat and must explicitly specify "chat_member" in
	// the list of allowed_updates to receive these updates.
	ChatMember *ChatMemberUpdated `json:"chat_member,omitempty"`
	// Optional. A request to join the chat has been sent. The bot must have the
	// can_invite_users administrator right in the chat to receive these
	// updates.
	ChatJoinRequest *ChatJoinRequest `json:"chat_join_request,omitempty"`
}

type GetUpdatesParams struct {
	// Optional. Identifier of the first update to be returned. Must be greater
	// by one than the highest among the identifiers of previously received
	// updates. By default, updates starting with the earliest unconfirmed
	// update are returned. An update is considered confirmed as soon as
	// getUpdates is called with an offset higher than its update_id. The
	// negative offset can be specified to retrieve updates starting from
	// -offset update from the end of the updates queue. All previous updates
	// will forgotten. https://core.telegram.org/bots/api#getupdates
	Offset UpdateID `json:"offset,omitempty"`
	// Optional. Limits the number of updates to be retrieved. Values between
	// 1-100 are accepted. Defaults to 100.
	Limit int `json:"limit,omitempty"`
	// Optional. Timeout in seconds for long polling. Defaults to 0, i.e. usual
	// short polling. Should be positive, short polling should be used for
	// testing purposes only.
	Timeout int `json:"timeout,omitempty"`
	// Optional. A JSON-serialized list of the update types you want your bot to
	// receive. For example, specify [“message”, “edited_channel_post”,
	// “callback_query”] to only receive updates of these types. See Update for
	// a complete list of available update types. Specify an empty list to
	// receive all update types except chat_member (default). If not specified,
	// the previous setting will be used.
	// https://core.telegram.org/bots/api#update
	//
	// Please note that this parameter doesn't affect updates created before the
	// call to the getUpdates, so unwanted updates may be received for a short
	// period of time.
	AllowedUpdates []UpdateType `json:"allowed_updates,omitempty"`
}

// Use this method to receive incoming updates using long polling (wiki). An
// Array of Update objects is returned.
// https://en.wikipedia.org/wiki/Push_technology#Long_polling
// https://core.telegram.org/bots/api#update
//
// Notes
//   1. This method will not work if an outgoing webhook is set up.
//   2. In order to avoid getting duplicate updates, recalculate offset after each server response.
//
// https://core.telegram.org/bots/api#getupdates
func (api *API) GetUpdates(params *GetUpdatesParams) ([]*Update, error) {
	updates := []*Update{}

	_, err := api.makeAPICall("getUpdates", params, nil, &updates)
	if err != nil {
		return nil, fmt.Errorf("GetUpdates: %w", err)
	}

	return updates, nil
}

type SetWebhookParams struct {
	// HTTPS url to send updates to. Use an empty string to remove webhook
	// integration
	URL string `json:"url"`
	// Optional. Upload your public key certificate so that the root certificate
	// in use can be checked. See our self-signed guide for details.
	// https://core.telegram.org/bots/self-signed
	Certificate InputFile `json:"certificate,omitempty"`
	// Optional. The fixed IP address which will be used to send webhook
	// requests instead of the IP address resolved through DNS
	IPAddress string `json:"ip_address,omitempty"`
	// Optional. Maximum allowed number of simultaneous HTTPS connections to the
	// webhook for update delivery, 1-100. Defaults to 40. Use lower values to
	// limit the load on your bot's server, and higher values to increase your
	// bot's throughput.
	MaxConnections int `json:"max_connections,omitempty"`
	// Optional. A JSON-serialized list of the update types you want your bot to
	// receive. For example, specify [“message”, “edited_channel_post”,
	// “callback_query”] to only receive updates of these types. See Update for
	// a complete list of available update types. Specify an empty list to
	// receive all update types except chat_member (default). If not specified,
	// the previous setting will be used.
	// https://core.telegram.org/bots/api#update Please note that this parameter
	// doesn't affect updates created before the call to the setWebhook, so
	// unwanted updates may be received for a short period of time.
	AllowedUpdates []UpdateType `json:"allowed_updates,omitempty"`
	// Optional. Pass True to drop all pending updates
	DropPendingUpdates bool `json:"drop_pending_updates,omitempty"`
	// A secret token to be sent in a header “X-Telegram-Bot-Api-Secret-Token”
	// in every webhook request, 1-256 characters. Only characters `A-Z`, `a-z`,
	// `0-9`, `_` and `-` are allowed. The header is useful to ensure that the
	// request comes from a webhook set by you.
	SecretToken string `json:"secret_token,omitempty"`
}

// Use this method to specify a url and receive incoming updates via an outgoing
// webhook. Whenever there is an update for the bot, we will send an HTTPS POST
// request to the specified url, containing a JSON-serialized Update. In case of
// an unsuccessful request, we will give up after a reasonable amount of
// attempts. Returns True on success. https://core.telegram.org/bots/api#update
//
// If you'd like to make sure that the webhook was set by you, you can specify
// secret data in the parameter secret_token. If specified, the request will
// contain a header “X-Telegram-Bot-Api-Secret-Token” with the secret token as
// content.
//
// Notes
//   1. You will not be able to receive updates using getUpdates for as long as an outgoing webhook is set up.
//   2. To use a self-signed certificate, you need to upload your public key certificate using certificate parameter. Please upload as InputFile, sending a String will not work.
//   3. Ports currently supported for webhooks: 443, 80, 88, 8443.
//
// https://core.telegram.org/bots/api#getupdates
// https://core.telegram.org/bots/self-signed
//
// If you're having any trouble setting up webhooks, please check out this
// amazing guide to webhooks. https://core.telegram.org/bots/webhooks
//
// https://core.telegram.org/bots/api#setwebhook
func (api *API) SetWebhook(params *SetWebhookParams) error {
	_, err := api.makeAPICall("setWebhook", params, []InputFile{params.Certificate}, nil)
	if err != nil {
		return fmt.Errorf("SetWebhook: %w", err)
	}

	return nil
}

type DeleteWebhookParams struct {
	// Optional. Pass True to drop all pending updates
	DropPendingUpdates bool `json:"drop_pending_updates,omitempty"`
}

// Use this method to remove webhook integration if you decide to switch back to
// getUpdates. Returns True on success.
// https://core.telegram.org/bots/api#getupdates
//
// https://core.telegram.org/bots/api#deletewebhook
func (api *API) DeleteWebhook(params *DeleteWebhookParams) error {
	_, err := api.makeAPICall("deleteWebhook", params, nil, nil)
	if err != nil {
		return fmt.Errorf("DeleteWebhook: %w", err)
	}

	return nil
}

// Use this method to get current webhook status. Requires no parameters. On
// success, returns a WebhookInfo object. If the bot is using getUpdates, will
// return an object with the url field empty.
// https://core.telegram.org/bots/api#webhookinfo
// https://core.telegram.org/bots/api#getupdates
//
// https://core.telegram.org/bots/api#getwebhookinfo
func (api *API) GetWebhookInfo() (*WebhookInfo, error) {
	webhookInfo := &WebhookInfo{}

	_, err := api.makeAPICall("getWebhookInfo", nil, nil, webhookInfo)
	if err != nil {
		return nil, fmt.Errorf("GetWebhookInfo: %w", err)
	}

	return webhookInfo, nil
}

// Contains information about the current status of a webhook.
//
// https://core.telegram.org/bots/api#webhookinfo
type WebhookInfo struct {
	// Webhook URL, may be empty if webhook is not set up
	URL string `json:"url"`
	// True, if a custom certificate was provided for webhook certificate checks
	HasCustomCertificate bool `json:"has_custom_certificate"`
	// Number of updates awaiting delivery
	PendingUpdateCount int `json:"pending_update_count"`
	// Optional. Currently used webhook IP address
	IPAddress string `json:"ip_address,omitempty"`
	// Optional. Unix time for the most recent error that happened when trying
	// to deliver an update via webhook
	LastErrorDate int64 `json:"last_error_date,omitempty"`
	// Optional. Error message in human-readable format for the most recent
	// error that happened when trying to deliver an update via webhook
	LastErrorMessage string `json:"last_error_message,omitempty"`
	// Optional. Unix time of the most recent error that happened when trying to
	// synchronize available updates with Telegram datacenters
	LastSynchronizationErrorDate int64 `json:"last_synchronization_error_date,omitempty"`
	// Optional. Maximum allowed number of simultaneous HTTPS connections to the
	// webhook for update delivery
	MaxConnections int `json:"max_connections,omitempty"`
	// Optional. A list of update types the bot is subscribed to. Defaults to
	// all update types except chat_member
	AllowedUpdates []UpdateType `json:"allowed_updates,omitempty"`
}
