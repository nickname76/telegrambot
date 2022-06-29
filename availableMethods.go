package telegrambot

// https://core.telegram.org/bots/api#available-methods

import "fmt"

// A simple method for testing your bot's authentication token. Requires no
// parameters. Returns basic information about the bot in form of a User object.
// https://core.telegram.org/bots/api#user
//
// https://core.telegram.org/bots/api#getme
func (api *API) GetMe() (*User, error) {
	user := &User{}

	_, err := api.makeAPICall("getMe", nil, nil, user)
	if err != nil {
		return nil, fmt.Errorf("GetMe: %w", err)
	}

	return user, nil
}

// Use this method to log out from the cloud Bot API server before launching the
// bot locally. You must log out the bot before running it locally, otherwise
// there is no guarantee that the bot will receive updates. After a successful
// call, you can immediately log in on a local server, but will not be able to
// log in back to the cloud Bot API server for 10 minutes. Returns True on
// success. Requires no parameters.
//
// https://core.telegram.org/bots/api#logout
func (api *API) LogOut() error {
	_, err := api.makeAPICall("logOut", nil, nil, nil)
	if err != nil {
		return fmt.Errorf("LogOut: %w", err)
	}

	return nil
}

// Use this method to close the bot instance before moving it from one local
// server to another. You need to delete the webhook before calling this method
// to ensure that the bot isn't launched again after server restart. The method
// will return error 429 in the first 10 minutes after the bot is launched.
// Returns True on success. Requires no parameters.
//
// https://core.telegram.org/bots/api#close
func (api *API) Close() error {
	_, err := api.makeAPICall("close", nil, nil, nil)
	if err != nil {
		return fmt.Errorf("Close: %w", err)
	}

	return nil
}

type SendMessageParams struct {
	// Unique identifier for the target chat or username of the target channel
	// (in the format @channelusername)
	ChatID ChatIDOrUsername `json:"chat_id"`
	// Text of the message to be sent, 1-4096 characters after entities parsin
	Text string `json:"text"`
	// Optional. Mode for parsing entities in the message text. See formatting
	// options for more details.
	ParseMode ParseMode `json:"parse_mode,omitempty"`
	// Optional. A JSON-serialized list of special entities that appear in
	// message text, which can be specified instead of parse_mode
	Entities []*MessageEntity `json:"entities,omitempty"`
	// Optional. Disables link previews for links in this message
	DisableWebPagePreview bool `json:"disable_web_page_preview,omitempty"`
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
	// Optional. Additional interface options. A JSON-serialized object for an
	// inline keyboard, custom reply keyboard, instructions to remove reply
	// keyboard or to force a reply from the user.
	// https://core.telegram.org/bots#inline-keyboards-and-on-the-fly-updating
	// https://core.telegram.org/bots#keyboards
	ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty"`
}

// Use this method to send text messages. On success, the sent Message is
// returned. https://core.telegram.org/bots/api#message
//
// https://core.telegram.org/bots/api#sendmessage
func (api *API) SendMessage(params *SendMessageParams) (*Message, error) {
	msg := &Message{}

	migrateToChatID, err := api.makeAPICall("sendMessage", params, nil, msg)
	if err != nil {
		if migrateToChatID != 0 {
			params.ChatID = migrateToChatID
			_, err = api.makeAPICall("sendMessage", params, nil, msg)
			if err != nil {
				return nil, fmt.Errorf("SendMessage: %w", err)
			}
		} else {
			return nil, fmt.Errorf("SendMessage: %w", err)
		}
	}

	return msg, nil
}

type ForwardMessageParams struct {
	// Unique identifier for the target chat or username of the target channel
	// (in the format @channelusername)
	ChatID ChatIDOrUsername `json:"chat_id"`
	// Unique identifier for the chat where the original message was sent (or
	// channel username in the format @channelusername)
	FromChatID ChatIDOrUsername `json:"from_chat_id"`
	// Optional. Sends the message silently. Users will receive a notification
	// with no sound. https://telegram.org/blog/channels-2-0#silent-messages
	DisableNotification bool `json:"disable_notification,omitempty"`
	// Optional. Protects the contents of the forwarded message from forwarding
	// and saving
	ProtectContent bool `json:"protect_content,omitempty"`
	// Message identifier in the chat specified in from_chat_id
	MessageID MessageID `json:"message_id"`
}

// Use this method to forward messages of any kind. Service messages can't be
// forwarded. On success, the sent Message is returned.
// https://core.telegram.org/bots/api#message
//
// https://core.telegram.org/bots/api#forwardmessage
func (api *API) ForwardMessage(params *ForwardMessageParams) (*Message, error) {
	msg := &Message{}

	migrateToChatID, err := api.makeAPICall("forwardMessage", params, nil, msg)
	if err != nil {
		if migrateToChatID != 0 {
			params.ChatID = migrateToChatID
			_, err = api.makeAPICall("forwardMessage", params, nil, msg)
			if err != nil {
				return nil, fmt.Errorf("ForwardMessage: %w", err)
			}
		} else {
			return nil, fmt.Errorf("ForwardMessage: %w", err)
		}
	}

	return msg, nil
}

type CopyMessageParams struct {
	// Unique identifier for the target chat or username of the target channel
	// (in the format @channelusername)
	ChatID ChatIDOrUsername `json:"chat_id"`
	// Unique identifier for the chat where the original message was sent (or
	// channel username in the format @channelusername)
	FromChatID ChatIDOrUsername `json:"from_chat_id"`
	// Message identifier in the chat specified in from_chat_id
	MessageID MessageID `json:"message_id"`
	// Optional. New caption for media, 0-1024 characters after entities
	// parsing. If not specified, the original caption is kept
	Caption string `json:"caption,omitempty"`
	// Optional. Mode for parsing entities in the new caption. See formatting
	// options for more details.
	// https://core.telegram.org/bots/api#formatting-options
	ParseMode ParseMode `json:"parse_mode,omitempty"`
	// Optional. A JSON-serialized list of special entities that appear in the
	// new caption, which can be specified instead of parse_mode
	CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"`
	// Optional. Sends the message silently. Users will receive a notification
	// with no sound. https://telegram.org/blog/channels-2-0#silent-messages
	DisableNotification bool `json:"disable_notification,omitempty"`
	// Optional. Protects the contents of the sent message from forwarding and
	// saving
	ProtectContent bool `json:"protect_content,omitempty"`
	// Optional. If the message is a reply, ID of the original message
	ReplyToMessageID bool `json:"reply_to_message_id,omitempty"`
	// Optional. Pass True, if the message should be sent even if the specified
	// replied-to message is not found
	AllowSendingWithoutReply bool `json:"allow_sending_without_reply,omitempty"`
	// Optional. Additional interface options. A JSON-serialized object for an
	// inline keyboard, custom reply keyboard, instructions to remove reply
	// keyboard or to force a reply from the user.
	// https://core.telegram.org/bots#inline-keyboards-and-on-the-fly-updating
	// https://core.telegram.org/bots#keyboards
	ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty"`
}

// Use this method to copy messages of any kind. Service messages and invoice
// messages can't be copied. The method is analogous to the method
// forwardMessage, but the copied message doesn't have a link to the original
// message. Returns the MessageId of the sent message on success.
// https://core.telegram.org/bots/api#forwardmessage
// https://core.telegram.org/bots/api#messageid
//
// https://core.telegram.org/bots/api#copymessage
func (api *API) CopyMessage(params *CopyMessageParams) (*MessageIDObject, error) {
	msgID := &MessageIDObject{}

	migrateToChatID, err := api.makeAPICall("copyMessage", params, nil, msgID)
	if err != nil {
		if migrateToChatID != 0 {
			params.ChatID = migrateToChatID
			_, err = api.makeAPICall("copyMessage", params, nil, msgID)
			if err != nil {
				return nil, fmt.Errorf("CopyMessage: %w", err)
			}
		} else {
			return nil, fmt.Errorf("CopyMessage: %w", err)
		}
	}

	return msgID, nil
}

type SendPhotoParams struct {
	// Unique identifier for the target chat or username of the target channel
	// (in the format @channelusername)
	ChatID ChatIDOrUsername `json:"chat_id"`
	// Photo to send. Pass a file_id as String to send a photo that exists on
	// the Telegram servers (recommended), pass an HTTP URL as a String for
	// Telegram to get a photo from the Internet, or upload a new photo using
	// multipart/form-data. The photo must be at most 10 MB in size. The photo's
	// width and height must not exceed 10000 in total. Width and height ratio
	// must be at most 20. More info on Sending Files »
	// https://core.telegram.org/bots/api#sending-files
	Photo InputFile `json:"photo"`
	// Optional. Photo caption (may also be used when resending photos by
	// file_id), 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`
	// Optional. Mode for parsing entities in the new caption. See formatting
	// options for more details.
	// https://core.telegram.org/bots/api#formatting-options
	ParseMode ParseMode `json:"parse_mode,omitempty"`
	// Optional. A JSON-serialized list of special entities that appear in the
	// new caption, which can be specified instead of parse_mode
	CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"`
	// Optional. Sends the message silently. Users will receive a notification
	// with no sound. https://telegram.org/blog/channels-2-0#silent-messages
	DisableNotification bool `json:"disable_notification,omitempty"`
	// Optional. Protects the contents of the sent message from forwarding and
	// saving
	ProtectContent bool `json:"protect_content,omitempty"`
	// Optional. If the message is a reply, ID of the original message
	ReplyToMessageID bool `json:"reply_to_message_id,omitempty"`
	// Optional. Pass True, if the message should be sent even if the specified
	// replied-to message is not found
	AllowSendingWithoutReply bool `json:"allow_sending_without_reply,omitempty"`
	// Optional. Additional interface options. A JSON-serialized object for an
	// inline keyboard, custom reply keyboard, instructions to remove reply
	// keyboard or to force a reply from the user.
	// https://core.telegram.org/bots#inline-keyboards-and-on-the-fly-updating
	// https://core.telegram.org/bots#keyboards
	ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty"`
}

// Use this method to send photos. On success, the sent Message is returned.
// https://core.telegram.org/bots/api#message
//
// https://core.telegram.org/bots/api#sendphoto
func (api *API) SendPhoto(params *SendPhotoParams) (*Message, error) {
	msg := &Message{}

	migrateToChatID, err := api.makeAPICall("sendPhoto", params, []InputFile{params.Photo}, msg)
	if err != nil {
		if migrateToChatID != 0 {
			params.ChatID = migrateToChatID
			_, err = api.makeAPICall("sendPhoto", params, []InputFile{params.Photo}, msg)
			if err != nil {
				return nil, fmt.Errorf("SendPhoto: %w", err)
			}
		} else {
			return nil, fmt.Errorf("SendPhoto: %w", err)
		}
	}

	return msg, nil
}

type SendAudioParams struct {
	// Unique identifier for the target chat or username of the target channel
	// (in the format @channelusername)
	ChatID ChatIDOrUsername `json:"chat_id"`
	// Audio file to send. Pass a file_id as String to send an audio file that
	// exists on the Telegram servers (recommended), pass an HTTP URL as a
	// String for Telegram to get an audio file from the Internet, or upload a
	// new one using multipart/form-data. More info on Sending Files »
	// https://core.telegram.org/bots/api#sending-files
	Audio InputFile `json:"audio"`
	// Optional. Audio caption, 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`
	// Optional. Mode for parsing entities in the new caption. See formatting
	// options for more details.
	// https://core.telegram.org/bots/api#formatting-options
	ParseMode ParseMode `json:"parse_mode,omitempty"`
	// Optional. A JSON-serialized list of special entities that appear in the
	// new caption, which can be specified instead of parse_mode
	CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"`
	// Optional. Duration of the audio in seconds
	Duration int `json:"duration,omitempty"`
	// Optional. Performer
	Performer string `json:"performer,omitempty"`
	// Track name
	Title string `json:"title,omitempty"`
	// Optional. Thumbnail of the file sent; can be ignored if thumbnail
	// generation for the file is supported server-side. The thumbnail should be
	// in JPEG format and less than 200 kB in size. A thumbnail's width and
	// height should not exceed 320. Ignored if the file is not uploaded using
	// multipart/form-data. Thumbnails can't be reused and can be only uploaded
	// as a new file, so you can pass "attach://<file_attach_name>" if the
	// thumbnail was uploaded using multipart/form-data under
	// <file_attach_name>. More info on Sending Files »
	// https://core.telegram.org/bots/api#sending-files
	Thumb InputFile `json:"thumb,omitempty"`
	// Optional. Sends the message silently. Users will receive a notification
	// with no sound. https://telegram.org/blog/channels-2-0#silent-messages
	DisableNotification bool `json:"disable_notification,omitempty"`
	// Optional. Protects the contents of the sent message from forwarding and
	// saving
	ProtectContent bool `json:"protect_content,omitempty"`
	// Optional. If the message is a reply, ID of the original message
	ReplyToMessageID bool `json:"reply_to_message_id,omitempty"`
	// Optional. Pass True, if the message should be sent even if the specified
	// replied-to message is not found
	AllowSendingWithoutReply bool `json:"allow_sending_without_reply,omitempty"`
	// Optional. Additional interface options. A JSON-serialized object for an
	// inline keyboard, custom reply keyboard, instructions to remove reply
	// keyboard or to force a reply from the user.
	// https://core.telegram.org/bots#inline-keyboards-and-on-the-fly-updating
	// https://core.telegram.org/bots#keyboards
	ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty"`
}

// Use this method to send audio files, if you want Telegram clients to display
// them in the music player. Your audio must be in the .MP3 or .M4A format. On
// success, the sent Message is returned. Bots can currently send audio files of
// up to 50 MB in size, this limit may be changed in the future.
//
// For sending voice messages, use the sendVoice method instead.
// https://core.telegram.org/bots/api#sendvoice
// https://core.telegram.org/bots/api#message
//
// https://core.telegram.org/bots/api#sendaudio
func (api *API) SendAudio(params *SendAudioParams) (*Message, error) {
	msg := &Message{}

	migrateToChatID, err := api.makeAPICall("sendAudio", params, []InputFile{params.Audio, params.Thumb}, msg)
	if err != nil {
		if migrateToChatID != 0 {
			params.ChatID = migrateToChatID
			_, err = api.makeAPICall("sendAudio", params, []InputFile{params.Audio, params.Thumb}, msg)
			if err != nil {
				return nil, fmt.Errorf("SendAudio: %w", err)
			}
		} else {
			return nil, fmt.Errorf("SendAudio: %w", err)
		}
	}

	return msg, nil
}

type SendDocumentParams struct {
	// Unique identifier for the target chat or username of the target channel
	// (in the format @channelusername)
	ChatID ChatIDOrUsername `json:"chat_id"`
	// File to send. Pass a file_id as String to send a file that exists on the
	// Telegram servers (recommended), pass an HTTP URL as a String for Telegram
	// to get a file from the Internet, or upload a new one using
	// multipart/form-data. More info on Sending Files »  //
	// https://core.telegram.org/bots/api#sending-files
	Document InputFile `json:"document"`
	// Optional. Thumbnail of the file sent; can be ignored if thumbnail
	// generation for the file is supported server-side. The thumbnail should be
	// in JPEG format and less than 200 kB in size. A thumbnail's width and
	// height should not exceed 320. Ignored if the file is not uploaded using
	// multipart/form-data. Thumbnails can't be reused and can be only uploaded
	// as a new file, so you can pass "attach://<file_attach_name>" if the
	// thumbnail was uploaded using multipart/form-data under
	// <file_attach_name>. More info on Sending Files »
	// https://core.telegram.org/bots/api#sending-files
	Thumb InputFile `json:"thumb,omitempty"`
	// Optional. Document caption (may also be used when resending documents by
	// file_id), 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`
	// Optional. Mode for parsing entities in the new caption. See formatting
	// options for more details.
	// https://core.telegram.org/bots/api#formatting-options
	ParseMode ParseMode `json:"parse_mode,omitempty"`
	// Optional. A JSON-serialized list of special entities that appear in the
	// new caption, which can be specified instead of parse_mode
	CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"`
	// Disables automatic server-side content type detection for files uploaded
	// using multipart/form-data
	DisableContentTypeDetection bool `json:"disable_content_type_detection,omitempty"`
	// Optional. Sends the message silently. Users will receive a notification
	// with no sound. https://telegram.org/blog/channels-2-0#silent-messages
	DisableNotification bool `json:"disable_notification,omitempty"`
	// Optional. Protects the contents of the sent message from forwarding and
	// saving
	ProtectContent bool `json:"protect_content,omitempty"`
	// Optional. If the message is a reply, ID of the original message
	ReplyToMessageID bool `json:"reply_to_message_id,omitempty"`
	// Optional. Pass True, if the message should be sent even if the specified
	// replied-to message is not found
	AllowSendingWithoutReply bool `json:"allow_sending_without_reply,omitempty"`
	// Optional. Additional interface options. A JSON-serialized object for an
	// inline keyboard, custom reply keyboard, instructions to remove reply
	// keyboard or to force a reply from the user.
	// https://core.telegram.org/bots#inline-keyboards-and-on-the-fly-updating
	// https://core.telegram.org/bots#keyboards
	ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty"`
}

// Use this method to send general files. On success, the sent Message is
// returned. Bots can currently send files of any type of up to 50 MB in size,
// this limit may be changed in the future.
// https://core.telegram.org/bots/api#message
//
// https://core.telegram.org/bots/api#senddocument
func (api *API) SendDocument(params *SendDocumentParams) (*Message, error) {
	msg := &Message{}

	migrateToChatID, err := api.makeAPICall("sendDocument", params, []InputFile{params.Document, params.Thumb}, msg)
	if err != nil {
		if migrateToChatID != 0 {
			params.ChatID = migrateToChatID
			_, err = api.makeAPICall("sendDocument", params, []InputFile{params.Document, params.Thumb}, msg)
			if err != nil {
				return nil, fmt.Errorf("SendDocument: %w", err)
			}
		} else {
			return nil, fmt.Errorf("SendDocument: %w", err)
		}
	}

	return msg, nil
}

type SendVideoParams struct {
	// Unique identifier for the target chat or username of the target channel
	// (in the format @channelusername)
	ChatID ChatIDOrUsername `json:"chat_id"`
	// Video to send. Pass a file_id as String to send a video that exists on
	// the Telegram servers (recommended), pass an HTTP URL as a String for
	// Telegram to get a video from the Internet, or upload a new video using
	// multipart/form-data. More info on Sending Files »
	// https://core.telegram.org/bots/api#sending-files
	Video InputFile `json:"video"`
	//Optional. Duration of sent video in seconds
	Duration int `json:"duration,omitempty"`
	// Optional. Video width
	Width int `json:"width,omitempty"`
	// Optional. Video height
	Height int `json:"height,omitempty"`
	// Optional. Thumbnail of the file sent; can be ignored if thumbnail
	// generation for the file is supported server-side. The thumbnail should be
	// in JPEG format and less than 200 kB in size. A thumbnail's width and
	// height should not exceed 320. Ignored if the file is not uploaded using
	// multipart/form-data. Thumbnails can't be reused and can be only uploaded
	// as a new file, so you can pass "attach://<file_attach_name>" if the
	// thumbnail was uploaded using multipart/form-data under
	// <file_attach_name>. More info on Sending Files »
	// https://core.telegram.org/bots/api#sending-files
	Thumb InputFile `json:"thumb,omitempty"`
	// Optional. Video caption (may also be used when resending videos by
	// file_id), 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`
	// Optional. Mode for parsing entities in the new caption. See formatting
	// options for more details.
	// https://core.telegram.org/bots/api#formatting-options
	ParseMode ParseMode `json:"parse_mode,omitempty"`
	// Optional. A JSON-serialized list of special entities that appear in the
	// new caption, which can be specified instead of parse_mode
	CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"`
	// Disables automatic server-side content type detection for files uploaded
	// using multipart/form-data
	DisableContentTypeDetection bool `json:"disable_content_type_detection,omitempty"`
	// Optional. Sends the message silently. Users will receive a notification
	// with no sound. https://telegram.org/blog/channels-2-0#silent-messages
	DisableNotification bool `json:"disable_notification,omitempty"`
	// Optional. Protects the contents of the sent message from forwarding and
	// saving
	ProtectContent bool `json:"protect_content,omitempty"`
	// Optional. If the message is a reply, ID of the original message
	ReplyToMessageID bool `json:"reply_to_message_id,omitempty"`
	// Optional. Pass True, if the message should be sent even if the specified
	// replied-to message is not found
	AllowSendingWithoutReply bool `json:"allow_sending_without_reply,omitempty"`
	// Optional. Additional interface options. A JSON-serialized object for an
	// inline keyboard, custom reply keyboard, instructions to remove reply
	// keyboard or to force a reply from the user.
	// https://core.telegram.org/bots#inline-keyboards-and-on-the-fly-updating
	// https://core.telegram.org/bots#keyboards
	ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty"`
}

// Use this method to send video files, Telegram clients support mp4 videos
// (other formats may be sent as Document). On success, the sent Message is
// returned. Bots can currently send video files of up to 50 MB in size, this
// limit may be changed in the future.
// https://core.telegram.org/bots/api#document
// https://core.telegram.org/bots/api#message
//
// https://core.telegram.org/bots/api#sendvideo
func (api *API) SendVideo(params *SendVideoParams) (*Message, error) {
	msg := &Message{}

	migrateToChatID, err := api.makeAPICall("sendVideo", params, []InputFile{params.Video, params.Thumb}, msg)
	if err != nil {
		if migrateToChatID != 0 {
			params.ChatID = migrateToChatID
			_, err = api.makeAPICall("sendVideo", params, []InputFile{params.Video, params.Thumb}, msg)
			if err != nil {
				return nil, fmt.Errorf("SendVideo: %w", err)
			}
		} else {
			return nil, fmt.Errorf("SendVideo: %w", err)
		}
	}

	return msg, nil
}

type SendAnimationParams struct {
	// Unique identifier for the target chat or username of the target channel
	// (in the format @channelusername)
	ChatID ChatIDOrUsername `json:"chat_id"`
	// Animation to send. Pass a file_id as String to send an animation that
	// exists on the Telegram servers (recommended), pass an HTTP URL as a
	// String for Telegram to get an animation from the Internet, or upload a
	// new animation using multipart/form-data. More info on Sending Files »
	// https://core.telegram.org/bots/api#sending-files
	Animation InputFile `json:"animation"`
	//Optional. Duration of sent animation in seconds
	Duration int `json:"duration,omitempty"`
	// Optional. Animation width
	Width int `json:"width,omitempty"`
	// Optional. Animation height
	Height int `json:"height,omitempty"`
	// Thumbnail of the file sent; can be ignored if thumbnail generation for
	// the file is supported server-side. The thumbnail should be in JPEG format
	// and less than 200 kB in size. A thumbnail's width and height should not
	// exceed 320. Ignored if the file is not uploaded using
	// multipart/form-data. Thumbnails can't be reused and can be only uploaded
	// as a new file, so you can pass “attach://<file_attach_name>” if the
	// thumbnail was uploaded using multipart/form-data under
	// <file_attach_name>. More info on Sending Files »
	// https://core.telegram.org/bots/api#sending-files
	Thumb InputFile `json:"thumb,omitempty"`
	// Optional. Animation caption (may also be used when resending animation by
	// file_id), 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`
	// Optional. Mode for parsing entities in the new caption. See formatting
	// options for more details.
	// https://core.telegram.org/bots/api#formatting-options
	ParseMode ParseMode `json:"parse_mode,omitempty"`
	// Optional. A JSON-serialized list of special entities that appear in the
	// new caption, which can be specified instead of parse_mode
	CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"`
	// Disables automatic server-side content type detection for files uploaded
	// using multipart/form-data
	DisableContentTypeDetection bool `json:"disable_content_type_detection,omitempty"`
	// Optional. Sends the message silently. Users will receive a notification
	// with no sound. https://telegram.org/blog/channels-2-0#silent-messages
	DisableNotification bool `json:"disable_notification,omitempty"`
	// Optional. Protects the contents of the sent message from forwarding and
	// saving
	ProtectContent bool `json:"protect_content,omitempty"`
	// Optional. If the message is a reply, ID of the original message
	ReplyToMessageID bool `json:"reply_to_message_id,omitempty"`
	// Optional. Pass True, if the message should be sent even if the specified
	// replied-to message is not found
	AllowSendingWithoutReply bool `json:"allow_sending_without_reply,omitempty"`
	// Optional. Additional interface options. A JSON-serialized object for an
	// inline keyboard, custom reply keyboard, instructions to remove reply
	// keyboard or to force a reply from the user.
	// https://core.telegram.org/bots#inline-keyboards-and-on-the-fly-updating
	// https://core.telegram.org/bots#keyboards
	ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty"`
}

// Use this method to send animation files (GIF or H.264/MPEG-4 AVC video
// without sound). On success, the sent Message is returned. Bots can currently
// send animation files of up to 50 MB in size, this limit may be changed in the
// future. https://core.telegram.org/bots/api#message
//
// https://core.telegram.org/bots/api#sendanimation
func (api *API) SendAnimation(params *SendAnimationParams) (*Message, error) {
	msg := &Message{}

	migrateToChatID, err := api.makeAPICall("sendAnimation", params, []InputFile{params.Animation, params.Thumb}, msg)
	if err != nil {
		if migrateToChatID != 0 {
			params.ChatID = migrateToChatID
			_, err = api.makeAPICall("sendAnimation", params, []InputFile{params.Animation, params.Thumb}, msg)
			if err != nil {
				return nil, fmt.Errorf("SendAnimation: %w", err)
			}
		} else {
			return nil, fmt.Errorf("SendAnimation: %w", err)
		}
	}

	return msg, nil
}

type SendVoiceParams struct {
	// Unique identifier for the target chat or username of the target channel
	// (in the format @channelusername)
	ChatID ChatIDOrUsername `json:"chat_id"`
	// Audio file to send. Pass a file_id as String to send a file that exists
	// on the Telegram servers (recommended), pass an HTTP URL as a String for
	// Telegram to get a file from the Internet, or upload a new one using
	// multipart/form-data. More info on Sending Files »
	// https://core.telegram.org/bots/api#sending-files
	Voice InputFile `json:"voice"`
	// Optional. Voice message caption, 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`
	// Optional. Mode for parsing entities in the new caption. See formatting
	// options for more details.
	// https://core.telegram.org/bots/api#formatting-options
	ParseMode ParseMode `json:"parse_mode,omitempty"`
	// Optional. Duration of the voice message in seconds
	Duration int `json:"duration,omitempty"`
	// Optional. A JSON-serialized list of special entities that appear in the
	// new caption, which can be specified instead of parse_mode
	CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"`
	// Disables automatic server-side content type detection for files uploaded
	// using multipart/form-data
	DisableContentTypeDetection bool `json:"disable_content_type_detection,omitempty"`
	// Optional. Sends the message silently. Users will receive a notification
	// with no sound. https://telegram.org/blog/channels-2-0#silent-messages
	DisableNotification bool `json:"disable_notification,omitempty"`
	// Optional. Protects the contents of the sent message from forwarding and
	// saving
	ProtectContent bool `json:"protect_content,omitempty"`
	// Optional. If the message is a reply, ID of the original message
	ReplyToMessageID bool `json:"reply_to_message_id,omitempty"`
	// Optional. Pass True, if the message should be sent even if the specified
	// replied-to message is not found
	AllowSendingWithoutReply bool `json:"allow_sending_without_reply,omitempty"`
	// Optional. Additional interface options. A JSON-serialized object for an
	// inline keyboard, custom reply keyboard, instructions to remove reply
	// keyboard or to force a reply from the user.
	// https://core.telegram.org/bots#inline-keyboards-and-on-the-fly-updating
	// https://core.telegram.org/bots#keyboards
	ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty"`
}

// Use this method to send audio files, if you want Telegram clients to display
// the file as a playable voice message. For this to work, your audio must be in
// an .OGG file encoded with OPUS (other formats may be sent as Audio or
// Document). On success, the sent Message is returned. Bots can currently send
// voice messages of up to 50 MB in size, this limit may be changed in the
// future. https://core.telegram.org/bots/api#audio
// https://core.telegram.org/bots/api#document
// https://core.telegram.org/bots/api#message
//
// https://core.telegram.org/bots/api#sendvoice
func (api *API) SendVoice(params *SendVoiceParams) (*Message, error) {
	msg := &Message{}

	migrateToChatID, err := api.makeAPICall("sendVoice", params, []InputFile{params.Voice}, msg)
	if err != nil {
		if migrateToChatID != 0 {
			params.ChatID = migrateToChatID
			_, err = api.makeAPICall("sendVoice", params, []InputFile{params.Voice}, msg)
			if err != nil {
				return nil, fmt.Errorf("SendVoice: %w", err)
			}
		} else {
			return nil, fmt.Errorf("SendVoice: %w", err)
		}
	}

	return msg, nil
}

type SendVideoNoteParams struct {
	// Unique identifier for the target chat or username of the target channel
	// (in the format @channelusername)
	ChatID ChatIDOrUsername `json:"chat_id"`
	// Video note to send. Pass a file_id as String to send a video note that
	// exists on the Telegram servers (recommended) or upload a new video using
	// multipart/form-data. More info on Sending Files ». Sending video notes by
	// a URL is currently unsupported
	// https://core.telegram.org/bots/api#sending-files
	VideoNote InputFile `json:"video_note"`
	// Optional. Duration of sent video in seconds
	Duration int `json:"duration,omitempty"`
	// Optional. Video width and height, i.e. diameter of the video message
	Length int `json:"length,omitempty"`
	// Thumbnail of the file sent; can be ignored if thumbnail generation for
	// the file is supported server-side. The thumbnail should be in JPEG format
	// and less than 200 kB in size. A thumbnail's width and height should not
	// exceed 320. Ignored if the file is not uploaded using
	// multipart/form-data. Thumbnails can't be reused and can be only uploaded
	// as a new file, so you can pass “attach://<file_attach_name>” if the
	// thumbnail was uploaded using multipart/form-data under
	// <file_attach_name>. More info on Sending Files »
	// https://core.telegram.org/bots/api#sending-files
	Thumb InputFile `json:"thumb,omitempty"`
	// Optional. Animation caption (may also be used when resending animation by
	// file_id), 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`
	// Optional. Mode for parsing entities in the new caption. See formatting
	// options for more details.
	// https://core.telegram.org/bots/api#formatting-options
	ParseMode ParseMode `json:"parse_mode,omitempty"`
	// Optional. A JSON-serialized list of special entities that appear in the
	// new caption, which can be specified instead of parse_mode
	CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"`
	// Disables automatic server-side content type detection for files uploaded
	// using multipart/form-data
	DisableContentTypeDetection bool `json:"disable_content_type_detection,omitempty"`
	// Optional. Sends the message silently. Users will receive a notification
	// with no sound. https://telegram.org/blog/channels-2-0#silent-messages
	DisableNotification bool `json:"disable_notification,omitempty"`
	// Optional. Protects the contents of the sent message from forwarding and
	// saving
	ProtectContent bool `json:"protect_content,omitempty"`
	// Optional. If the message is a reply, ID of the original message
	ReplyToMessageID bool `json:"reply_to_message_id,omitempty"`
	// Optional. Pass True, if the message should be sent even if the specified
	// replied-to message is not found
	AllowSendingWithoutReply bool `json:"allow_sending_without_reply,omitempty"`
	// Optional. Additional interface options. A JSON-serialized object for an
	// inline keyboard, custom reply keyboard, instructions to remove reply
	// keyboard or to force a reply from the user.
	// https://core.telegram.org/bots#inline-keyboards-and-on-the-fly-updating
	// https://core.telegram.org/bots#keyboards
	ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty"`
}

// As of v.4.0, Telegram clients support rounded square mp4 videos of up to 1
// minute long. Use this method to send video messages. On success, the sent
// Message is returned. https://telegram.org/blog/video-messages-and-telescope
// https://core.telegram.org/bots/api#message
//
// https://core.telegram.org/bots/api#sendvideonote
func (api *API) SendVideoNote(params *SendVideoNoteParams) (*Message, error) {
	msg := &Message{}

	migrateToChatID, err := api.makeAPICall("sendVideoNote", params, []InputFile{params.VideoNote, params.Thumb}, msg)
	if err != nil {
		if migrateToChatID != 0 {
			params.ChatID = migrateToChatID
			_, err = api.makeAPICall("sendVideoNote", params, []InputFile{params.VideoNote, params.Thumb}, msg)
			if err != nil {
				return nil, fmt.Errorf("SendVideoNote: %w", err)
			}
		} else {
			return nil, fmt.Errorf("SendVideoNote: %w", err)
		}
	}

	return msg, nil
}

type SendMediaGroupParams struct {
	// Unique identifier for the target chat or username of the target channel
	// (in the format @channelusername)
	ChatID ChatIDOrUsername `json:"chat_id"`
	// A JSON-serialized array describing messages to be sent, must include 2-10
	// items
	Media []*InputMedia `json:"media"`
	// Optional. Sends messages silently. Users will receive a notification with
	// no sound. https://telegram.org/blog/channels-2-0#silent-messages
	DisableNotification bool `json:"disable_notification,omitempty"`
	// Optional. Protects the contents of the sent message from forwarding and
	// saving
	ProtectContent bool `json:"protect_content,omitempty"`
	// Optional. If the message is a reply, ID of the original message
	ReplyToMessageID bool `json:"reply_to_message_id,omitempty"`
	// Optional. Pass True, if the message should be sent even if the specified
	// replied-to message is not found
	AllowSendingWithoutReply bool `json:"allow_sending_without_reply,omitempty"`
}

// Use this method to send a group of photos, videos, documents or audios as an
// album. Documents and audio files can be only grouped in an album with
// messages of the same type. On success, an array of Messages that were sent is
// returned. https://core.telegram.org/bots/api#message
//
// https://core.telegram.org/bots/api#sendmediagroup
func (api *API) SendMediaGroup(params *SendMediaGroupParams) ([]*Message, error) {
	inputFiles := []InputFile{}
	for _, inputMedia := range params.Media {
		inputFiles = append(inputFiles, inputMedia.Media, inputMedia.Thumb)
	}

	msgs := []*Message{}

	migrateToChatID, err := api.makeAPICall("sendMediaGroup", params, inputFiles, &msgs)
	if err != nil {

		if migrateToChatID != 0 {
			params.ChatID = migrateToChatID
			_, err = api.makeAPICall("sendMediaGroup", params, inputFiles, &msgs)
			if err != nil {
				return nil, fmt.Errorf("SendMediaGroup: %w", err)
			}
		} else {
			return nil, fmt.Errorf("SendMediaGroup: %w", err)
		}
	}

	return msgs, nil
}

type SendLocationParams struct {
	// Unique identifier for the target chat or username of the target channel
	// (in the format @channelusername)
	ChatID ChatIDOrUsername `json:"chat_id"`
	// Latitude of the location
	Latitude float64 `json:"latitude"`
	// Longitude of the location
	Longitude float64 `json:"longitude"`
	// Optional. The radius of uncertainty for the location, measured in meters;
	// 0-1500
	HorizontalAccuracy float64 `json:"horizontal_accuracy,omitempty"`
	// Optional. Period in seconds for which the location will be updated (see
	// Live Locations, should be between 60 and 86400.
	// https://telegram.org/blog/live-locations
	LivePeriod int `json:"live_period,omitempty"`
	// Optional. For live locations, a direction in which the user is moving, in
	// degrees. Must be between 1 and 360 if specified.
	Heading int `json:"heading,omitempty"`
	// Optional. For live locations, a maximum distance for proximity alerts
	// about approaching another chat member, in meters. Must be between 1 and
	// 100000 if specified.
	ProximityAlertRadius int `json:"proximity_alert_radius,omitempty"`
	// Optional. Sends the message silently. Users will receive a notification
	// with no sound. https://telegram.org/blog/channels-2-0#silent-messages
	DisableNotification bool `json:"disable_notification,omitempty"`
	// Optional. Protects the contents of the sent message from forwarding and
	// saving
	ProtectContent bool `json:"protect_content,omitempty"`
	// Optional. If the message is a reply, ID of the original message
	ReplyToMessageID bool `json:"reply_to_message_id,omitempty"`
	// Optional. Pass True, if the message should be sent even if the specified
	// replied-to message is not found
	AllowSendingWithoutReply bool `json:"allow_sending_without_reply,omitempty"`
	// Optional. Additional interface options. A JSON-serialized object for an
	// inline keyboard, custom reply keyboard, instructions to remove reply
	// keyboard or to force a reply from the user.
	// https://core.telegram.org/bots#inline-keyboards-and-on-the-fly-updating
	// https://core.telegram.org/bots#keyboards
	ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty"`
}

// Use this method to send point on the map. On success, the sent Message is
// returned. https://core.telegram.org/bots/api#message
//
// https://core.telegram.org/bots/api#sendlocation
func (api *API) SendLocation(params *SendLocationParams) (*Message, error) {
	msg := &Message{}

	migrateToChatID, err := api.makeAPICall("sendLocation", params, nil, msg)
	if err != nil {
		if migrateToChatID != 0 {
			params.ChatID = migrateToChatID
			_, err = api.makeAPICall("sendLocation", params, nil, msg)
			if err != nil {
				return nil, fmt.Errorf("SendLocation: %w", err)
			}
		} else {
			return nil, fmt.Errorf("SendLocation: %w", err)
		}
	}

	return msg, nil
}

type EditMessageLiveLocationParams struct {
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
	// Latitude of new location
	Latitude float64 `json:"latitude"`
	// Longitude of new location
	Longitude float64 `json:"longitude"`
	// Optional. The radius of uncertainty for the location, measured in meters;
	// 0-1500
	HorizontalAccuracy float64 `json:"horizontal_accuracy,omitempty"`
	// Optional. Direction in which the user is moving, in degrees. Must be
	// between 1 and 360 if specified.
	Heading int `json:"heading,omitempty"`
	// Optional. Maximum distance for proximity alerts about approaching another
	// chat member, in meters. Must be between 1 and 100000 if specified.
	ProximityAlertRadius int `json:"proximity_alert_radius,omitempty"`
	// Optional. A JSON-serialized object for a new inline keyboard.
	// https://core.telegram.org/bots#inline-keyboards-and-on-the-fly-updating
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// Use this method to edit live location messages. A location can be edited
// until its live_period expires or editing is explicitly disabled by a call to
// stopMessageLiveLocation. On success, if the edited message is not an inline
// message, the edited Message is returned, otherwise True is returned.
// https://core.telegram.org/bots/api#stopmessagelivelocation
// https://core.telegram.org/bots/api#message
//
// https://core.telegram.org/bots/api#editmessagelivelocation
func (api *API) EditMessageLiveLocation(params *EditMessageLiveLocationParams) (*Message, error) {
	var msg *Message

	if params.InlineMessageID != "" {
		msg = &Message{}
	}

	migrateToChatID, err := api.makeAPICall("editMessageLiveLocation", params, nil, msg)
	if err != nil {
		if migrateToChatID != 0 {
			params.ChatID = migrateToChatID
			_, err = api.makeAPICall("editMessageLiveLocation", params, nil, msg)
			if err != nil {
				return nil, fmt.Errorf("EditMessageLiveLocation: %w", err)
			}
		} else {
			return nil, fmt.Errorf("EditMessageLiveLocation: %w", err)
		}
	}

	return msg, nil
}

type StopMessageLiveLocationParams struct {
	// Optional. Required if inline_message_id is not specified. Unique
	// identifier for the target chat or username of the target channel (in the
	// format @channelusername)
	ChatID ChatIDOrUsername `json:"chat_id,omitempty"`
	// Optional. Required if inline_message_id is not specified. Identifier of
	// the message with live location to stop
	MessageID MessageID `json:"message_id,omitempty"`
	// Optional. Required if chat_id and message_id are not specified.
	// Identifier of the inline message
	InlineMessageID InlineMessageID `json:"inline_message_id,omitempty"`
	// Optional. A JSON-serialized object for a new inline keyboard.
	// https://core.telegram.org/bots#inline-keyboards-and-on-the-fly-updating
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// Use this method to stop updating a live location message before live_period
// expires. On success, if the message is not an inline message, the edited
// Message is returned, otherwise True is returned.
// https://core.telegram.org/bots/api#message
//
// https://core.telegram.org/bots/api#stopmessagelivelocation
func (api *API) StopMessageLiveLocation(params *StopMessageLiveLocationParams) (*Message, error) {
	var msg *Message

	if params.InlineMessageID != "" {
		msg = &Message{}
	}

	migrateToChatID, err := api.makeAPICall("stopMessageLiveLocation", params, nil, msg)
	if err != nil {
		if migrateToChatID != 0 {
			params.ChatID = migrateToChatID
			_, err = api.makeAPICall("stopMessageLiveLocation", params, nil, msg)
			if err != nil {
				return nil, fmt.Errorf("StopMessageLiveLocation: %w", err)
			}
		} else {
			return nil, fmt.Errorf("StopMessageLiveLocation: %w", err)
		}
	}

	return msg, nil
}

type SendVenueParams struct {
	// Unique identifier for the target chat or username of the target channel
	// (in the format @channelusername)
	ChatID ChatIDOrUsername `json:"chat_id"`
	// Latitude of the venue
	Latitude float64 `json:"latitude"`
	// Longitude of the venue
	Longitude float64 `json:"longitude"`
	// Name of the venue
	Title string `json:"title"`
	// Address of the venue
	Address string `json:"address"`
	// Optional. Foursquare identifier of the venue
	FoursquareID string `json:"foursquare_id,omitempty"`
	// Optional. Foursquare type of the venue, if known. (For example,
	// “arts_entertainment/default”, “arts_entertainment/aquarium” or
	// “food/icecream”.)
	FoursquareType string `json:"foursquare_type,omitempty"`
	// Optional. Google Places identifier of the venue
	GooglePlaceID string `json:"google_place_id,omitempty"`
	// Optional. Google Places type of the venue. (See supported types.)
	// https://developers.google.com/places/web-service/supported_types
	GooglePlaceType string `json:"google_place_type,omitempty"`
	// Optional. Sends the message silently. Users will receive a notification
	// with no sound. https://telegram.org/blog/channels-2-0#silent-messages
	DisableNotification bool `json:"disable_notification,omitempty"`
	// Optional. Protects the contents of the sent message from forwarding and
	// saving
	ProtectContent bool `json:"protect_content,omitempty"`
	// Optional. If the message is a reply, ID of the original message
	ReplyToMessageID bool `json:"reply_to_message_id,omitempty"`
	// Optional. Pass True, if the message should be sent even if the specified
	// replied-to message is not found
	AllowSendingWithoutReply bool `json:"allow_sending_without_reply,omitempty"`
	// Optional. Additional interface options. A JSON-serialized object for an
	// inline keyboard, custom reply keyboard, instructions to remove reply
	// keyboard or to force a reply from the user.
	// https://core.telegram.org/bots#inline-keyboards-and-on-the-fly-updating
	// https://core.telegram.org/bots#keyboards
	ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty"`
}

// Use this method to send information about a venue. On success, the sent
// Message is returned. https://core.telegram.org/bots/api#message
//
// https://core.telegram.org/bots/api#sendvenue
func (api *API) SendVenue(params *SendVenueParams) (*Message, error) {
	msg := &Message{}

	migrateToChatID, err := api.makeAPICall("sendVenue", params, nil, msg)
	if err != nil {
		if migrateToChatID != 0 {
			params.ChatID = migrateToChatID
			_, err = api.makeAPICall("sendVenue", params, nil, msg)
			if err != nil {
				return nil, fmt.Errorf("SendVenue: %w", err)
			}
		} else {
			return nil, fmt.Errorf("SendVenue: %w", err)
		}
	}

	return msg, nil
}

type SendContactParams struct {
	// Unique identifier for the target chat or username of the target channel
	// (in the format @channelusername)
	ChatID ChatIDOrUsername `json:"chat_id"`
	// Contact's phone number
	PhoneNumber string `json:"phone_number"`
	// Contact's first name
	FirstName string `json:"first_name"`
	// Optional. Contact's last name
	LastName string `json:"last_name,omitempty"`
	// Optional. Additional data about the contact in the form of a vCard,
	// 0-2048 bytes https://en.wikipedia.org/wiki/VCard
	VCard string `json:"vcard,omitempty"`
	// Optional. Sends the message silently. Users will receive a notification
	// with no sound. https://telegram.org/blog/channels-2-0#silent-messages
	DisableNotification bool `json:"disable_notification,omitempty"`
	// Optional. Protects the contents of the sent message from forwarding and
	// saving
	ProtectContent bool `json:"protect_content,omitempty"`
	// Optional. If the message is a reply, ID of the original message
	ReplyToMessageID bool `json:"reply_to_message_id,omitempty"`
	// Optional. Pass True, if the message should be sent even if the specified
	// replied-to message is not found
	AllowSendingWithoutReply bool `json:"allow_sending_without_reply,omitempty"`
	// Optional. Additional interface options. A JSON-serialized object for an
	// inline keyboard, custom reply keyboard, instructions to remove reply
	// keyboard or to force a reply from the user.
	// https://core.telegram.org/bots#inline-keyboards-and-on-the-fly-updating
	// https://core.telegram.org/bots#keyboards
	ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty"`
}

// Use this method to send phone contacts. On success, the sent Message is
// returned. https://core.telegram.org/bots/api#message
//
// https://core.telegram.org/bots/api#sendcontact
func (api *API) SendContact(params *SendContactParams) (*Message, error) {
	msg := &Message{}

	migrateToChatID, err := api.makeAPICall("sendContact", params, nil, msg)
	if err != nil {
		if migrateToChatID != 0 {
			params.ChatID = migrateToChatID
			_, err = api.makeAPICall("sendContact", params, nil, msg)
			if err != nil {
				return nil, fmt.Errorf("SendContact: %w", err)
			}
		} else {
			return nil, fmt.Errorf("SendContact: %w", err)
		}
	}

	return msg, nil
}

type SendPollParams struct {
	// Unique identifier for the target chat or username of the target channel
	// (in the format @channelusername)
	ChatID ChatIDOrUsername `json:"chat_id"`
	// Poll question, 1-300 characters
	Question string `json:"question"`
	// A JSON-serialized list of answer options, 2-10 strings 1-100 characters
	// each
	Options []string `json:"options"`
	// Optional. True, if the poll needs to be anonymous, defaults to True
	IsAnonymous bool `json:"is_anonymous,omitempty"`
	// Optional. Poll type, “quiz” or “regular”, defaults to “regular”
	Type PollType `json:"type,omitempty"`
	// Optional. True, if the poll allows multiple answers, ignored for polls in
	// quiz mode, defaults to False
	AllowsMultipleAnswers bool `json:"allows_multiple_answers,omitempty"`
	// Optional. 0-based identifier of the correct answer option, required for
	// polls in quiz mode
	CorrectOptionID int `json:"correct_option_id,omitempty"`
	// Optional. Text that is shown when a user chooses an incorrect answer or
	// taps on the lamp icon in a quiz-style poll, 0-200 characters with at most
	// 2 line feeds after entities parsing
	Explanation string `json:"explanation,omitempty"`
	// Optional. Mode for parsing entities in the explanation. See formatting
	// options for more details.
	// https://core.telegram.org/bots/api#formatting-options
	ExplanationParseMode ParseMode `json:"explanation_parse_mode,omitempty"`
	// Optional. A JSON-serialized list of special entities that appear in the
	// poll explanation, which can be specified instead of parse_mode
	ExplanationEntities []*MessageEntity `json:"explanation_entities,omitempty"`
	// Optional. Amount of time in seconds the poll will be active after
	// creation, 5-600. Can't be used together with close_date.
	OpenPeriod int `json:"open_period,omitempty"`
	// Optional. Point in time (Unix timestamp) when the poll will be
	// automatically closed. Must be at least 5 and no more than 600 seconds in
	// the future. Can't be used together with open_period.
	CloseDate int64 `json:"close_date,omitempty"`
	// Optional. Pass True, if the poll needs to be immediately closed. This can
	// be useful for poll preview.
	IsClosed bool `json:"is_closed,omitempty"`
	// Optional. Sends the message silently. Users will receive a notification
	// with no sound. https://telegram.org/blog/channels-2-0#silent-messages
	DisableNotification bool `json:"disable_notification,omitempty"`
	// Optional. Protects the contents of the sent message from forwarding and
	// saving
	ProtectContent bool `json:"protect_content,omitempty"`
	// Optional. If the message is a reply, ID of the original message
	ReplyToMessageID bool `json:"reply_to_message_id,omitempty"`
	// Optional. Pass True, if the message should be sent even if the specified
	// replied-to message is not found
	AllowSendingWithoutReply bool `json:"allow_sending_without_reply,omitempty"`
	// Optional. Additional interface options. A JSON-serialized object for an
	// inline keyboard, custom reply keyboard, instructions to remove reply
	// keyboard or to force a reply from the user.
	// https://core.telegram.org/bots#inline-keyboards-and-on-the-fly-updating
	// https://core.telegram.org/bots#keyboards
	ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty"`
}

// Use this method to send a native poll. On success, the sent Message is
// returned. https://core.telegram.org/bots/api#message
//
// https://core.telegram.org/bots/api#sendpoll
func (api *API) SendPoll(params *SendPollParams) (*Message, error) {
	msg := &Message{}

	migrateToChatID, err := api.makeAPICall("sendPoll", params, nil, msg)
	if err != nil {
		if migrateToChatID != 0 {
			params.ChatID = migrateToChatID
			_, err = api.makeAPICall("sendPoll", params, nil, msg)
			if err != nil {
				return nil, fmt.Errorf("SendPoll: %w", err)
			}
		} else {
			return nil, fmt.Errorf("SendPoll: %w", err)
		}
	}

	return msg, nil
}

type SendDiceParams struct {
	// Unique identifier for the target chat or username of the target channel
	// (in the format @channelusername)
	ChatID ChatIDOrUsername `json:"chat_id"`
	// Optional. Emoji on which the dice throw animation is based. Currently,
	// must be one of “🎲”, “🎯”, “🏀”, “⚽”, “🎳”, or “🎰”. Dice can have values
	// 1-6 for “🎲”, “🎯” and “🎳”, values 1-5 for “🏀” and “⚽”, and values 1-64
	// for “🎰”. Defaults to “🎲”
	Emoji DiceEmoji `json:"emoji,omitempty"`
	// Optional. Sends the message silently. Users will receive a notification
	// with no sound. https://telegram.org/blog/channels-2-0#silent-messages
	DisableNotification bool `json:"disable_notification,omitempty"`
	// Optional. Protects the contents of the sent message from forwarding and
	// saving
	ProtectContent bool `json:"protect_content,omitempty"`
	// Optional. If the message is a reply, ID of the original message
	ReplyToMessageID bool `json:"reply_to_message_id,omitempty"`
	// Optional. Pass True, if the message should be sent even if the specified
	// replied-to message is not found
	AllowSendingWithoutReply bool `json:"allow_sending_without_reply,omitempty"`
	// Optional. Additional interface options. A JSON-serialized object for an
	// inline keyboard, custom reply keyboard, instructions to remove reply
	// keyboard or to force a reply from the user.
	// https://core.telegram.org/bots#inline-keyboards-and-on-the-fly-updating
	// https://core.telegram.org/bots#keyboards
	ReplyMarkup ReplyMarkup `json:"reply_markup,omitempty"`
}

// Use this method to send an animated emoji that will display a random value.
// On success, the sent Message is returned.
// https://core.telegram.org/bots/api#message
//
// https://core.telegram.org/bots/api#senddice
func (api *API) SendDice(params *SendDiceParams) (*Message, error) {
	msg := &Message{}

	migrateToChatID, err := api.makeAPICall("sendDice", params, nil, msg)
	if err != nil {
		if migrateToChatID != 0 {
			params.ChatID = migrateToChatID
			_, err = api.makeAPICall("sendDice", params, nil, msg)
			if err != nil {
				return nil, fmt.Errorf("SendDice: %w", err)
			}
		} else {
			return nil, fmt.Errorf("SendDice: %w", err)
		}
	}

	return msg, nil
}

type SendChatActionParams struct {
	// Unique identifier for the target chat or username of the target channel
	// (in the format @channelusername)
	ChatID ChatIDOrUsername `json:"chat_id"`
	// Type of action to broadcast. Choose one, depending on what the user is
	// about to receive: typing for text messages, upload_photo for photos,
	// record_video or upload_video for videos, record_voice or upload_voice for
	// voice notes, upload_document for general files, choose_sticker for
	// stickers, find_location for location data, record_video_note or
	// upload_video_note for video notes.
	// https://core.telegram.org/bots/api#sendmessage
	// https://core.telegram.org/bots/api#sendphoto
	// https://core.telegram.org/bots/api#sendvideo
	// https://core.telegram.org/bots/api#sendvoice
	// https://core.telegram.org/bots/api#senddocument
	// https://core.telegram.org/bots/api#sendsticker
	// https://core.telegram.org/bots/api#sendlocation
	// https://core.telegram.org/bots/api#sendvideonote
	Action ChatAction `json:"action"`
}

// Use this method when you need to tell the user that something is happening on
// the bot's side. The status is set for 5 seconds or less (when a message
// arrives from your bot, Telegram clients clear its typing status). Returns
// True on success.
//
// Example: The ImageBot needs some time to process a request and upload the
// image. Instead of sending a text message along the lines of “Retrieving
// image, please wait…”, the bot may use sendChatAction with action =
// upload_photo. The user will see a “sending photo” status for the bot.
// https://t.me/imagebot https://core.telegram.org/bots/api#sendchataction
//
// We only recommend using this method when a response from the bot will take a
// *noticeable* amount of time to arrive.
//
// https://core.telegram.org/bots/api#sendchataction
func (api *API) SendChatAction(params *SendChatActionParams) error {
	migrateToChatID, err := api.makeAPICall("sendChatAction", params, nil, nil)
	if err != nil {
		if migrateToChatID != 0 {
			params.ChatID = migrateToChatID
			_, err = api.makeAPICall("sendChatAction", params, nil, nil)
			if err != nil {
				return fmt.Errorf("SendChatAction: %w", err)
			}
		} else {
			return fmt.Errorf("SendChatAction: %w", err)
		}
	}

	return nil
}

type GetUserProfilePhotosParams struct {
	// Unique identifier of the target user
	UserID UserID `json:"user_id"`
	// Optional. Sequential number of the first photo to be returned. By
	// default, all photos are returned.
	Offset int `json:"offset,omitempty"`
	// Optional. Limits the number of photos to be retrieved. Values between
	// 1-100 are accepted. Defaults to 100.
	Limit int `json:"limit,omitempty"`
}

// Use this method to get a list of profile pictures for a user. Returns a
// UserProfilePhotos object.
// https://core.telegram.org/bots/api#userprofilephotos
//
// https://core.telegram.org/bots/api#getuserprofilephotos
func (api *API) GetUserProfilePhotos(params *GetUserProfilePhotosParams) (*UserProfilePhotos, error) {
	userProfilePhotos := &UserProfilePhotos{}

	_, err := api.makeAPICall("getUserProfilePhotos", params, nil, userProfilePhotos)
	if err != nil {
		return nil, fmt.Errorf("GetUserProfilePhotos: %w", err)
	}

	return userProfilePhotos, nil
}

type GetFileParams struct {
	// File identifier to get info about
	FileID FileID `json:"file_id"`
}

// Use this method to get basic information about a file and prepare it for
// downloading. For the moment, bots can download files of up to 20MB in size.
// On success, a File object is returned. The file can then be downloaded via
// the link https://api.telegram.org/file/bot<token>/<file_path>, where
// <file_path> is taken from the response. It is guaranteed that the link will
// be valid for at least 1 hour. When the link expires, a new one can be
// requested by calling getFile again. https://core.telegram.org/bots/api#file
//
// Note: This function may not preserve the original file name and MIME type.
// You should save the file's MIME type and name (if available) when the File
// object is received.
//
// https://core.telegram.org/bots/api#getfile
func (api *API) GetFile(params *GetFileParams) (*File, error) {
	file := &File{}

	_, err := api.makeAPICall("getFile", params, nil, file)
	if err != nil {
		return nil, fmt.Errorf("GetFile: %w", err)
	}

	return file, nil
}

type BanChatMemberParams struct {
	// Unique identifier for the target group or username of the target
	// supergroup or channel (in the format @channelusername)
	ChatID ChatIDOrUsername `json:"chat_id"`
	// Unique identifier of the target user
	UserID UserID `json:"user_id"`
	// Optional. Date when the user will be unbanned, unix time. If user is
	// banned for more than 366 days or less than 30 seconds from the current
	// time they are considered to be banned forever. Applied for supergroups
	// and channels only.
	UntilDate int64 `json:"until_date,omitempty"`
	// Optional. Pass True to delete all messages from the chat for the user
	// that is being removed. If False, the user will be able to see messages in
	// the group that were sent before the user was removed. Always True for
	// supergroups and channels.
	RevokeMessages bool `json:"revoke_messages,omitempty"`
}

// Use this method to ban a user in a group, a supergroup or a channel. In the
// case of supergroups and channels, the user will not be able to return to the
// chat on their own using invite links, etc., unless unbanned first. The bot
// must be an administrator in the chat for this to work and must have the
// appropriate administrator rights. Returns True on success.
// https://core.telegram.org/bots/api#unbanchatmember
//
// https://core.telegram.org/bots/api#banchatmember
func (api *API) BanChatMember(params *BanChatMemberParams) error {
	migrateToChatID, err := api.makeAPICall("banChatMember", params, nil, nil)
	if err != nil {
		if migrateToChatID != 0 {
			params.ChatID = migrateToChatID
			_, err = api.makeAPICall("banChatMember", params, nil, nil)
			if err != nil {
				return fmt.Errorf("BanChatMember: %w", err)
			}
		} else {
			return fmt.Errorf("BanChatMember: %w", err)
		}
	}

	return nil
}

type UnbanChatMemberParams struct {
	// Unique identifier for the target group or username of the target
	// supergroup or channel (in the format @channelusername)
	ChatID ChatIDOrUsername `json:"chat_id"`
	// Unique identifier of the target user
	UserID UserID `json:"user_id"`
	// Optional. Do nothing if the user is not banned
	OnlyIfBanned bool `json:"only_if_banned,omitempty"`
}

// Use this method to unban a previously banned user in a supergroup or channel.
// The user will not return to the group or channel automatically, but will be
// able to join via link, etc. The bot must be an administrator for this to
// work. By default, this method guarantees that after the call the user is not
// a member of the chat, but will be able to join it. So if the user is a member
// of the chat they will also be removed from the chat. If you don't want this,
// use the parameter only_if_banned. Returns True on success.
// https://core.telegram.org/bots/api#unbanchatmember
//
// https://core.telegram.org/bots/api#unbanchatmember
func (api *API) UnbanChatMember(params *UnbanChatMemberParams) error {
	migrateToChatID, err := api.makeAPICall("unbanChatMember", params, nil, nil)
	if err != nil {
		if migrateToChatID != 0 {
			params.ChatID = migrateToChatID
			_, err = api.makeAPICall("unbanChatMember", params, nil, nil)
			if err != nil {
				return fmt.Errorf("UnbanChatMember: %w", err)
			}
		} else {
			return fmt.Errorf("UnbanChatMember: %w", err)
		}
	}

	return nil
}

type RestrictChatMemberParams struct {
	// Unique identifier for the target chat or username of the target
	// supergroup (in the format @supergroupusername)
	ChatID ChatIDOrUsername `json:"chat_id"`
	// Unique identifier of the target user
	UserID UserID `json:"user_id"`
	// A JSON-serialized object for new user permissions
	Permissions *ChatPermissions `json:"permissions"`
	// Optional. Date when restrictions will be lifted for the user, unix time.
	// If user is restricted for more than 366 days or less than 30 seconds from
	// the current time, they are considered to be restricted forever
	UntilDate int64 `json:"until_date,omitempty"`
}

// Use this method to restrict a user in a supergroup. The bot must be an
// administrator in the supergroup for this to work and must have the
// appropriate administrator rights. Pass True for all permissions to lift
// restrictions from a user. Returns True on success.
//
// https://core.telegram.org/bots/api#restrictchatmember
func (api *API) RestrictChatMember(params *RestrictChatMemberParams) error {
	migrateToChatID, err := api.makeAPICall("restrictChatMember", params, nil, nil)
	if err != nil {
		if migrateToChatID != 0 {
			params.ChatID = migrateToChatID
			_, err = api.makeAPICall("restrictChatMember", params, nil, nil)
			if err != nil {
				return fmt.Errorf("RestrictChatMember: %w", err)
			}
		} else {
			return fmt.Errorf("RestrictChatMember: %w", err)
		}
	}

	return nil
}

type PromoteChatMemberParams struct {
	// Unique identifier for the target chat or username of the target channel
	// (in the format @channelusername)
	ChatID ChatIDOrUsername `json:"chat_id"`
	// Unique identifier of the target user
	UserID UserID `json:"user_id"`
	// Optional. Pass True, if the administrator's presence in the chat is
	// hidden
	IsAnonymous bool `json:"is_anonymous,omitempty"`
	// Optional. Pass True, if the administrator can access the chat event log,
	// chat statistics, message statistics in channels, see channel members, see
	// anonymous administrators in supergroups and ignore slow mode. Implied by
	// any other administrator privilege
	CanManageChat bool `json:"can_manage_chat,omitempty"`
	// Optional. Pass True, if the administrator can create channel posts,
	// channels only
	CanPostMessages bool `json:"can_post_messages,omitempty"`
	// Optional. Pass True, if the administrator can edit messages of other
	// users and can pin messages, channels only
	CanEditMessages bool `json:"can_edit_messages,omitempty"`
	// Optional. Pass True, if the administrator can delete messages of other
	// users
	CanDeleteMessages bool `json:"can_delete_messages,omitempty"`
	// Optional. Pass True, if the administrator can manage video chats
	CanManageVideoChats bool `json:"can_manage_video_chats,omitempty"`
	// Optional. Pass True, if the administrator can restrict, ban or unban chat
	// members
	CanRestrictMembers bool `json:"can_restrict_members,omitempty"`
	// Optional. Pass True, if the administrator can add new administrators with
	// a subset of their own privileges or demote administrators that he has
	// promoted, directly or indirectly (promoted by administrators that were
	// appointed by him)
	CanPromoteMembers bool `json:"can_promote_members,omitempty"`
	// Optional. Pass True, if the administrator can change chat title, photo
	// and other settings
	CanChangeInfo bool `json:"can_change_info,omitempty"`
	// Optional. Pass True, if the administrator can invite new users to the
	// chat
	CanInviteUsers bool `json:"can_invite_users,omitempty"`
	// Optional. Pass True, if the administrator can pin messages, supergroups
	// only
	CanPinMessages bool `json:"can_pin_messages,omitempty"`
}

// Use this method to promote or demote a user in a supergroup or a channel. The
// bot must be an administrator in the chat for this to work and must have the
// appropriate administrator rights. Pass False for all boolean parameters to
// demote a user. Returns True on success.
//
// https://core.telegram.org/bots/api#promotechatmember
func (api *API) PromoteChatMember(params *PromoteChatMemberParams) error {
	migrateToChatID, err := api.makeAPICall("promoteChatMember", params, nil, nil)
	if err != nil {
		if migrateToChatID != 0 {
			params.ChatID = migrateToChatID
			_, err = api.makeAPICall("promoteChatMember", params, nil, nil)
			if err != nil {
				return fmt.Errorf("PromoteChatMember: %w", err)
			}
		} else {
			return fmt.Errorf("PromoteChatMember: %w", err)
		}
	}

	return nil
}

type SetChatAdministratorCustomTitleParams struct {
	// Unique identifier for the target chat or username of the target
	// supergroup (in the format @supergroupusername)
	ChatID ChatIDOrUsername `json:"chat_id"`
	// Unique identifier of the target user
	UserID UserID `json:"user_id"`
	// New custom title for the administrator; 0-16 characters, emoji are not
	// allowed
	CustomTitle string `json:"custom_title"`
}

// Use this method to set a custom title for an administrator in a supergroup
// promoted by the bot. Returns True on success.
//
// https://core.telegram.org/bots/api#setchatadministratorcustomtitle
func (api *API) SetChatAdministratorCustomTitle(params *SetChatAdministratorCustomTitleParams) error {
	migrateToChatID, err := api.makeAPICall("setChatAdministratorCustomTitle", params, nil, nil)
	if err != nil {
		if migrateToChatID != 0 {
			params.ChatID = migrateToChatID
			_, err = api.makeAPICall("setChatAdministratorCustomTitle", params, nil, nil)
			if err != nil {
				return fmt.Errorf("SetChatAdministratorCustomTitle: %w", err)
			}
		} else {
			return fmt.Errorf("SetChatAdministratorCustomTitle: %w", err)
		}
	}

	return nil
}

type BanChatSenderChatParams struct {
	// Unique identifier for the target chat or username of the target channel
	// (in the format @channelusername)
	ChatID ChatIDOrUsername `json:"chat_id"`
	// Unique identifier of the target sender chat
	SenderChatID ChatID `json:"sender_chat_id"`
}

// Use this method to ban a channel chat in a supergroup or a channel. Until the
// chat is unbanned, the owner of the banned chat won't be able to send messages
// on behalf of any of their channels. The bot must be an administrator in the
// supergroup or channel for this to work and must have the appropriate
// administrator rights. Returns True on success.
// https://core.telegram.org/bots/api#unbanchatsenderchat
//
// https://core.telegram.org/bots/api#banchatsenderchat
func (api *API) BanChatSenderChat(params *BanChatSenderChatParams) error {
	migrateToChatID, err := api.makeAPICall("banChatSenderChat", params, nil, nil)
	if err != nil {
		if migrateToChatID != 0 {
			params.ChatID = migrateToChatID
			_, err = api.makeAPICall("banChatSenderChat", params, nil, nil)
			if err != nil {
				return fmt.Errorf("BanChatSenderChat: %w", err)
			}
		} else {
			return fmt.Errorf("BanChatSenderChat: %w", err)
		}
	}

	return nil
}

type UnbanChatSenderChatParams struct {
	// Unique identifier for the target chat or username of the target channel
	// (in the format @channelusername)
	ChatID ChatIDOrUsername `json:"chat_id"`
	// Unique identifier of the target sender chat
	SenderChatID ChatID `json:"sender_chat_id"`
}

// Use this method to unban a previously banned channel chat in a supergroup or
// channel. The bot must be an administrator for this to work and must have the
// appropriate administrator rights. Returns True on success.
//
// https://core.telegram.org/bots/api#unbanchatsenderchat
func (api *API) UnbanChatSenderChat(params *UnbanChatSenderChatParams) error {
	migrateToChatID, err := api.makeAPICall("unbanChatSenderChat", params, nil, nil)
	if err != nil {
		if migrateToChatID != 0 {
			params.ChatID = migrateToChatID
			_, err = api.makeAPICall("unbanChatSenderChat", params, nil, nil)
			if err != nil {
				return fmt.Errorf("UnbanChatSenderChat: %w", err)
			}
		} else {
			return fmt.Errorf("UnbanChatSenderChat: %w", err)
		}
	}

	return nil
}

type SetChatPermissionsParams struct {
	// Unique identifier for the target chat or username of the target
	// supergroup (in the format @supergroupusername)
	ChatID ChatIDOrUsername `json:"chat_id"`
	// A JSON-serialized object for new default chat permissions
	Permissions *ChatPermissions `json:"permissions"`
}

// Use this method to set default chat permissions for all members. The bot must
// be an administrator in the group or a supergroup for this to work and must
// have the can_restrict_members administrator rights. Returns True on success.
//
// https://core.telegram.org/bots/api#setchatpermissions
func (api *API) SetChatPermissions(params *SetChatPermissionsParams) error {
	migrateToChatID, err := api.makeAPICall("setChatPermissions", params, nil, nil)
	if err != nil {
		if migrateToChatID != 0 {
			params.ChatID = migrateToChatID
			_, err = api.makeAPICall("setChatPermissions", params, nil, nil)
			if err != nil {
				return fmt.Errorf("SetChatPermissions: %w", err)
			}
		} else {
			return fmt.Errorf("SetChatPermissions: %w", err)
		}
	}

	return nil
}

type ExportChatInviteLinkParams struct {
	// Unique identifier for the target chat or username of the target channel
	// (in the format @channelusername)
	ChatID ChatIDOrUsername `json:"chat_id"`
}

// Use this method to generate a new primary invite link for a chat; any
// previously generated primary link is revoked. The bot must be an
// administrator in the chat for this to work and must have the appropriate
// administrator rights. Returns the new invite link as String on success.
//
// Note: Each administrator in a chat generates their own invite links. Bots
// can't use invite links generated by other administrators. If you want your
// bot to work with invite links, it will need to generate its own link using
// exportChatInviteLink or by calling the getChat method. If your bot needs to
// generate a new primary invite link replacing its previous one, use
// exportChatInviteLink again.
// https://core.telegram.org/bots/api#exportchatinvitelink
// https://core.telegram.org/bots/api#getchat
//
// https://core.telegram.org/bots/api#exportchatinvitelink
func (api *API) ExportChatInviteLink(params *ExportChatInviteLinkParams) (string, error) {
	link := ""

	migrateToChatID, err := api.makeAPICall("exportChatInviteLink", params, nil, &link)
	if err != nil {
		if migrateToChatID != 0 {
			params.ChatID = migrateToChatID
			_, err = api.makeAPICall("exportChatInviteLink", params, nil, &link)
			if err != nil {
				return "", fmt.Errorf("ExportChatInviteLink: %w", err)
			}
		} else {
			return "", fmt.Errorf("ExportChatInviteLink: %w", err)
		}
	}

	return link, nil
}

type CreateChatInviteLinkParams struct {
	// Unique identifier for the target chat or username of the target channel
	// (in the format @channelusername)
	ChatID ChatIDOrUsername `json:"chat_id"`
	// Optional. Invite link name; 0-32 characters
	Name string `json:"name,omitempty"`
	// Optional. Point in time (Unix timestamp) when the link will expire
	ExpireDate int64 `json:"expire_date,omitempty"`
	// Optional. Maximum number of users that can be members of the chat
	// simultaneously after joining the chat via this invite link; 1-99999
	MemberLimit int `json:"member_limit,omitempty"`
	// Optional. True, if users joining the chat via the link need to be
	// approved by chat administrators. If True, member_limit can't be specified
	CreatesJoinRequest bool `json:"creates_join_request,omitempty"`
}

// Use this method to create an additional invite link for a chat. The bot must
// be an administrator in the chat for this to work and must have the
// appropriate administrator rights. The link can be revoked using the method
// revokeChatInviteLink. Returns the new invite link as ChatInviteLink object.
// https://core.telegram.org/bots/api#revokechatinvitelink
// https://core.telegram.org/bots/api#chatinvitelink
//
// https://core.telegram.org/bots/api#createchatinvitelink
func (api *API) CreateChatInviteLink(params *CreateChatInviteLinkParams) (*ChatInviteLink, error) {
	chatInviteLink := &ChatInviteLink{}

	migrateToChatID, err := api.makeAPICall("exportChatInviteLink", params, nil, chatInviteLink)
	if err != nil {
		if migrateToChatID != 0 {
			params.ChatID = migrateToChatID
			_, err = api.makeAPICall("exportChatInviteLink", params, nil, chatInviteLink)
			if err != nil {
				return nil, fmt.Errorf("ExportChatInviteLink: %w", err)
			}
		} else {
			return nil, fmt.Errorf("ExportChatInviteLink: %w", err)
		}
	}

	return chatInviteLink, nil
}

type EditChatInviteLinkParams struct {
	// Unique identifier for the target chat or username of the target channel
	// (in the format @channelusername)
	ChatID ChatIDOrUsername `json:"chat_id"`
	// The invite link to edit
	InviteLink string `json:"invite_link"`
	// Optional. Invite link name; 0-32 characters
	Name string `json:"name,omitempty"`
	// Optional. Point in time (Unix timestamp) when the link will expire
	ExpireDate int64 `json:"expire_date,omitempty"`
	// Optional. Maximum number of users that can be members of the chat
	// simultaneously after joining the chat via this invite link; 1-99999
	MemberLimit int `json:"member_limit,omitempty"`
	// Optional. True, if users joining the chat via the link need to be
	// approved by chat administrators. If True, member_limit can't be specified
	CreatesJoinRequest bool `json:"creates_join_request,omitempty"`
}

// Use this method to edit a non-primary invite link created by the bot. The bot
// must be an administrator in the chat for this to work and must have the
// appropriate administrator rights. Returns the edited invite link as a
// ChatInviteLink object. https://core.telegram.org/bots/api#chatinvitelink
//
// https://core.telegram.org/bots/api#editchatinvitelink
func (api *API) EditChatInviteLink(params *EditChatInviteLinkParams) (*ChatInviteLink, error) {
	chatInviteLink := &ChatInviteLink{}

	migrateToChatID, err := api.makeAPICall("editChatInviteLink", params, nil, chatInviteLink)
	if err != nil {
		if migrateToChatID != 0 {
			params.ChatID = migrateToChatID
			_, err = api.makeAPICall("editChatInviteLink", params, nil, chatInviteLink)
			if err != nil {
				return nil, fmt.Errorf("EditChatInviteLink: %w", err)
			}
		} else {
			return nil, fmt.Errorf("EditChatInviteLink: %w", err)
		}
	}

	return chatInviteLink, nil
}

type RevokeChatInviteLinkParams struct {
	// Unique identifier of the target chat or username of the target channel
	// (in the format @channelusername)
	ChatID ChatIDOrUsername `json:"chat_id"`
	// The invite link to revoke
	InviteLink string `json:"invite_link"`
}

// Use this method to revoke an invite link created by the bot. If the primary
// link is revoked, a new link is automatically generated. The bot must be an
// administrator in the chat for this to work and must have the appropriate
// administrator rights. Returns the revoked invite link as ChatInviteLink
// object. https://core.telegram.org/bots/api#chatinvitelink
//
// https://core.telegram.org/bots/api#revokechatinvitelink
func (api *API) RevokeChatInviteLink(params *RevokeChatInviteLinkParams) (*ChatInviteLink, error) {
	chatInviteLink := &ChatInviteLink{}

	migrateToChatID, err := api.makeAPICall("revokeChatInviteLink", params, nil, chatInviteLink)
	if err != nil {
		if migrateToChatID != 0 {
			params.ChatID = migrateToChatID
			_, err = api.makeAPICall("revokeChatInviteLink", params, nil, chatInviteLink)
			if err != nil {
				return nil, fmt.Errorf("RevokeChatInviteLink: %w", err)
			}
		} else {
			return nil, fmt.Errorf("RevokeChatInviteLink: %w", err)
		}
	}

	return chatInviteLink, nil
}

type ApproveChatJoinRequestParams struct {
	// Unique identifier for the target chat or username of the target channel
	// (in the format @channelusername)
	ChatID ChatIDOrUsername `json:"chat_id"`
	// Unique identifier of the target user
	UserID UserID `json:"user_id"`
}

// Use this method to approve a chat join request. The bot must be an
// administrator in the chat for this to work and must have the can_invite_users
// administrator right. Returns True on success.
//
// https://core.telegram.org/bots/api#approvechatjoinrequest
func (api *API) ApproveChatJoinRequest(params *ApproveChatJoinRequestParams) error {
	migrateToChatID, err := api.makeAPICall("approveChatJoinRequest", params, nil, nil)
	if err != nil {
		if migrateToChatID != 0 {
			params.ChatID = migrateToChatID
			_, err = api.makeAPICall("approveChatJoinRequest", params, nil, nil)
			if err != nil {
				return fmt.Errorf("ApproveChatJoinRequest: %w", err)
			}
		} else {
			return fmt.Errorf("ApproveChatJoinRequest: %w", err)
		}
	}

	return nil
}

type DeclineChatJoinRequestParams struct {
	// Unique identifier for the target chat or username of the target channel
	// (in the format @channelusername)
	ChatID ChatIDOrUsername `json:"chat_id"`
	// Unique identifier of the target user
	UserID UserID `json:"user_id"`
}

// Use this method to decline a chat join request. The bot must be an
// administrator in the chat for this to work and must have the can_invite_users
// administrator right. Returns True on success.
//
// https://core.telegram.org/bots/api#declinechatjoinrequest
func (api *API) DeclineChatJoinRequest(params *DeclineChatJoinRequestParams) error {
	migrateToChatID, err := api.makeAPICall("declineChatJoinRequest", params, nil, nil)
	if err != nil {
		if migrateToChatID != 0 {
			params.ChatID = migrateToChatID
			_, err = api.makeAPICall("declineChatJoinRequest", params, nil, nil)
			if err != nil {
				return fmt.Errorf("DeclineChatJoinRequest: %w", err)
			}
		} else {
			return fmt.Errorf("DeclineChatJoinRequest: %w", err)
		}
	}

	return nil
}

type SetChatPhotoParams struct {
	// Unique identifier for the target chat or username of the target channel
	// (in the format @channelusername)
	ChatID ChatIDOrUsername `json:"chat_id"`
	// New chat photo, uploaded using multipart/form-data
	Photo InputFile `json:"user_id"`
}

// Use this method to set a new profile photo for the chat. Photos can't be
// changed for private chats. The bot must be an administrator in the chat for
// this to work and must have the appropriate administrator rights. Returns True
// on success.
//
// https://core.telegram.org/bots/api#setchatphoto
func (api *API) SetChatPhoto(params *SetChatPhotoParams) error {
	migrateToChatID, err := api.makeAPICall("setChatPhoto", params, []InputFile{params.Photo}, nil)
	if err != nil {
		if migrateToChatID != 0 {
			params.ChatID = migrateToChatID
			_, err = api.makeAPICall("setChatPhoto", params, []InputFile{params.Photo}, nil)
			if err != nil {
				return fmt.Errorf("SetChatPhoto: %w", err)
			}
		} else {
			return fmt.Errorf("SetChatPhoto: %w", err)
		}
	}

	return nil
}

type DeleteChatPhotoParams struct {
	// Unique identifier for the target chat or username of the target channel
	// (in the format @channelusername)
	ChatID ChatIDOrUsername `json:"chat_id"`
}

// Use this method to delete a chat photo. Photos can't be changed for private
// chats. The bot must be an administrator in the chat for this to work and must
// have the appropriate administrator rights. Returns True on success.
//
// https://core.telegram.org/bots/api#deletechatphoto
func (api *API) DeleteChatPhoto(params *DeleteChatPhotoParams) error {
	migrateToChatID, err := api.makeAPICall("deleteChatPhoto", params, nil, nil)
	if err != nil {
		if migrateToChatID != 0 {
			params.ChatID = migrateToChatID
			_, err = api.makeAPICall("deleteChatPhoto", params, nil, nil)
			if err != nil {
				return fmt.Errorf("DeleteChatPhoto: %w", err)
			}
		} else {
			return fmt.Errorf("DeleteChatPhoto: %w", err)
		}
	}

	return nil
}

type SetChatTitleParams struct {
	// Unique identifier for the target chat or username of the target channel
	// (in the format @channelusername)
	ChatID ChatIDOrUsername `json:"chat_id"`
	// New chat title, 1-255 characters
	Title string `json:"title"`
}

// Use this method to change the title of a chat. Titles can't be changed for
// private chats. The bot must be an administrator in the chat for this to work
// and must have the appropriate administrator rights. Returns True on success.
//
// https://core.telegram.org/bots/api#setchattitle
func (api *API) SetChatTitle(params *SetChatTitleParams) error {
	migrateToChatID, err := api.makeAPICall("setChatTitle", params, nil, nil)
	if err != nil {
		if migrateToChatID != 0 {
			params.ChatID = migrateToChatID
			_, err = api.makeAPICall("setChatTitle", params, nil, nil)
			if err != nil {
				return fmt.Errorf("SetChatTitle: %w", err)
			}
		} else {
			return fmt.Errorf("SetChatTitle: %w", err)
		}
	}

	return nil
}

type SetChatDescriptionParams struct {
	// Unique identifier for the target chat or username of the target channel
	// (in the format @channelusername)
	ChatID ChatIDOrUsername `json:"chat_id"`
	// Optional. New chat description, 0-255 characters
	Description string `json:"description,omitempty"`
}

// Use this method to change the description of a group, a supergroup or a
// channel. The bot must be an administrator in the chat for this to work and
// must have the appropriate administrator rights. Returns True on success.
//
// https://core.telegram.org/bots/api#setchatdescription
func (api *API) SetChatDescription(params *SetChatDescriptionParams) error {
	migrateToChatID, err := api.makeAPICall("setChatDescription", params, nil, nil)
	if err != nil {
		if migrateToChatID != 0 {
			params.ChatID = migrateToChatID
			_, err = api.makeAPICall("setChatDescription", params, nil, nil)
			if err != nil {
				return fmt.Errorf("SetChatDescription: %w", err)
			}
		} else {
			return fmt.Errorf("SetChatDescription: %w", err)
		}
	}

	return nil
}

type PinChatMessageParams struct {
	// Unique identifier for the target chat or username of the target channel
	// (in the format @channelusername)
	ChatID ChatIDOrUsername `json:"chat_id"`
	// Identifier of a message to pin
	MessageID MessageID `json:"message_id"`
	// Optinal. Pass True, if it is not necessary to send a notification to all
	// chat members about the new pinned message. Notifications are always
	// disabled in channels and private chats.
	DisableNotification bool `json:"disable_notification,omitempty"`
}

// Use this method to add a message to the list of pinned messages in a chat. If
// the chat is not a private chat, the bot must be an administrator in the chat
// for this to work and must have the 'can_pin_messages' administrator right in
// a supergroup or 'can_edit_messages' administrator right in a channel. Returns
// True on success.
//
// https://core.telegram.org/bots/api#pinchatmessage
func (api *API) PinChatMessage(params *PinChatMessageParams) error {
	migrateToChatID, err := api.makeAPICall("pinChatMessage", params, nil, nil)
	if err != nil {
		if migrateToChatID != 0 {
			params.ChatID = migrateToChatID
			_, err = api.makeAPICall("pinChatMessage", params, nil, nil)
			if err != nil {
				return fmt.Errorf("PinChatMessage: %w", err)
			}
		} else {
			return fmt.Errorf("PinChatMessage: %w", err)
		}
	}

	return nil
}

type UnpinChatMessageParams struct {
	// Unique identifier for the target chat or username of the target channel
	// (in the format @channelusername)
	ChatID ChatIDOrUsername `json:"chat_id"`
	// Optional. Identifier of a message to unpin. If not specified, the most
	// recent pinned message (by sending date) will be unpinned.
	MessageID MessageID `json:"message_id,omitempty"`
}

// Use this method to remove a message from the list of pinned messages in a
// chat. If the chat is not a private chat, the bot must be an administrator in
// the chat for this to work and must have the 'can_pin_messages' administrator
// right in a supergroup or 'can_edit_messages' administrator right in a
// channel. Returns True on success.
//
// https://core.telegram.org/bots/api#unpinchatmessage
func (api *API) UnpinChatMessage(params *UnpinChatMessageParams) error {
	migrateToChatID, err := api.makeAPICall("unpinChatMessage", params, nil, nil)
	if err != nil {
		if migrateToChatID != 0 {
			params.ChatID = migrateToChatID
			_, err = api.makeAPICall("unpinChatMessage", params, nil, nil)
			if err != nil {
				return fmt.Errorf("UnpinChatMessage: %w", err)
			}
		} else {
			return fmt.Errorf("UnpinChatMessage: %w", err)
		}
	}

	return nil
}

type UnpinAllChatMessagesParams struct {
	// Unique identifier for the target chat or username of the target channel
	// (in the format @channelusername)
	ChatID ChatIDOrUsername `json:"chat_id"`
}

// Use this method to clear the list of pinned messages in a chat. If the chat
// is not a private chat, the bot must be an administrator in the chat for this
// to work and must have the 'can_pin_messages' administrator right in a
// supergroup or 'can_edit_messages' administrator right in a channel. Returns
// True on success.
//
// https://core.telegram.org/bots/api#unpinallchatmessages
func (api *API) UnpinAllChatMessages(params *UnpinAllChatMessagesParams) error {
	migrateToChatID, err := api.makeAPICall("unpinAllChatMessages", params, nil, nil)
	if err != nil {
		if migrateToChatID != 0 {
			params.ChatID = migrateToChatID
			_, err = api.makeAPICall("unpinAllChatMessages", params, nil, nil)
			if err != nil {
				return fmt.Errorf("UnpinAllChatMessages: %w", err)
			}
		} else {
			return fmt.Errorf("UnpinAllChatMessages: %w", err)
		}
	}

	return nil
}

type LeaveChatParams struct {
	// Unique identifier for the target chat or username of the target
	// supergroup or channel (in the format @channelusername)
	ChatID ChatIDOrUsername `json:"chat_id"`
}

// Use this method for your bot to leave a group, supergroup or channel. Returns
// True on success.
//
// https://core.telegram.org/bots/api#leavechat
func (api *API) LeaveChat(params *LeaveChatParams) error {
	migrateToChatID, err := api.makeAPICall("leaveChat", params, nil, nil)
	if err != nil {
		if migrateToChatID != 0 {
			params.ChatID = migrateToChatID
			_, err = api.makeAPICall("leaveChat", params, nil, nil)
			if err != nil {
				return fmt.Errorf("LeaveChat: %w", err)
			}
		} else {
			return fmt.Errorf("LeaveChat: %w", err)
		}
	}

	return nil
}

type GetChatParams struct {
	// Unique identifier for the target chat or username of the target
	// supergroup or channel (in the format @channelusername)
	ChatID ChatIDOrUsername `json:"chat_id"`
}

// Use this method to get up to date information about the chat (current name of
// the user for one-on-one conversations, current username of a user, group or
// channel, etc.). Returns a Chat object on success.
// https://core.telegram.org/bots/api#chat
//
// https://core.telegram.org/bots/api#getchat
func (api *API) GetChat(params *GetChatParams) (*UserProfilePhotos, error) {
	userProfilePhotos := &UserProfilePhotos{}

	migrateToChatID, err := api.makeAPICall("getChat", params, nil, userProfilePhotos)
	if err != nil {
		if migrateToChatID != 0 {
			params.ChatID = migrateToChatID
			_, err = api.makeAPICall("getChat", params, nil, userProfilePhotos)
			if err != nil {
				return nil, fmt.Errorf("GetChat: %w", err)
			}
		} else {
			return nil, fmt.Errorf("GetChat: %w", err)
		}
	}

	return userProfilePhotos, nil
}

type GetChatAdministratorsParams struct {
	// Unique identifier for the target chat or username of the target
	// supergroup or channel (in the format @channelusername)
	ChatID ChatIDOrUsername `json:"chat_id"`
}

// Use this method to get a list of administrators in a chat. On success,
// returns an Array of ChatMember objects that contains information about all
// chat administrators except other bots. If the chat is a group or a supergroup
// and no administrators were appointed, only the creator will be returned.
// https://core.telegram.org/bots/api#chatmember
//
// https://core.telegram.org/bots/api#getchatadministrators
func (api *API) GetChatAdministrators(params *GetChatAdministratorsParams) ([]*ChatMember, error) {
	chatMembers := []*ChatMember{}

	migrateToChatID, err := api.makeAPICall("getChatAdministrators", params, nil, &chatMembers)
	if err != nil {
		if migrateToChatID != 0 {
			params.ChatID = migrateToChatID
			_, err = api.makeAPICall("getChatAdministrators", params, nil, &chatMembers)
			if err != nil {
				return nil, fmt.Errorf("GetChatAdministrators: %w", err)
			}
		} else {
			return nil, fmt.Errorf("GetChatAdministrators: %w", err)
		}
	}

	return chatMembers, nil
}

type GetChatMemberCountParams struct {
	// Unique identifier for the target chat or username of the target
	// supergroup or channel (in the format @channelusername)
	ChatID ChatIDOrUsername `json:"chat_id"`
}

// Use this method to get the number of members in a chat. Returns Int on
// success.
//
// https://core.telegram.org/bots/api#getchatmembercount
func (api *API) GetChatMemberCount(params *GetChatMemberCountParams) (int, error) {
	memberCount := 0

	migrateToChatID, err := api.makeAPICall("getChatMemberCount", params, nil, &memberCount)
	if err != nil {
		if migrateToChatID != 0 {
			params.ChatID = migrateToChatID
			_, err = api.makeAPICall("getChatMemberCount", params, nil, &memberCount)
			if err != nil {
				return 0, fmt.Errorf("GetChatMemberCount: %w", err)
			}
		} else {
			return 0, fmt.Errorf("GetChatMemberCount: %w", err)
		}
	}

	return memberCount, nil
}

type GetChatMemberParams struct {
	// Unique identifier for the target chat or username of the target
	// supergroup or channel (in the format @channelusername)
	ChatID ChatIDOrUsername `json:"chat_id"`
	// Unique identifier of the target user
	UserID UserID `json:"user_id"`
}

// Use this method to get information about a member of a chat. Returns a
// ChatMember object on success.
//
// https://core.telegram.org/bots/api#getchatmember
func (api *API) GetChatMember(params *GetChatMemberParams) (*ChatMember, error) {
	chatMember := &ChatMember{}

	migrateToChatID, err := api.makeAPICall("getChatMember", params, nil, chatMember)
	if err != nil {
		if migrateToChatID != 0 {
			params.ChatID = migrateToChatID
			_, err = api.makeAPICall("getChatMember", params, nil, chatMember)
			if err != nil {
				return nil, fmt.Errorf("GetChatMember: %w", err)
			}
		} else {
			return nil, fmt.Errorf("GetChatMember: %w", err)
		}
	}

	return chatMember, nil
}

type SetChatStickerSetParams struct {
	// Unique identifier for the target chat or username of the target
	// supergroup or channel (in the format @channelusername)
	ChatID ChatIDOrUsername `json:"chat_id"`
	// Name of the sticker set to be set as the group sticker set
	StickerSetName StickerSetName `json:"sticker_set_name"`
}

// Use this method to set a new group sticker set for a supergroup. The bot must
// be an administrator in the chat for this to work and must have the
// appropriate administrator rights. Use the field can_set_sticker_set
// optionally returned in getChat requests to check if the bot can use this
// method. Returns True on success. https://core.telegram.org/bots/api#getchat
//
// https://core.telegram.org/bots/api#setchatstickerset
func (api *API) SetChatStickerSet(params *SetChatStickerSetParams) error {
	migrateToChatID, err := api.makeAPICall("setChatStickerSet", params, nil, nil)
	if err != nil {
		if migrateToChatID != 0 {
			params.ChatID = migrateToChatID
			_, err = api.makeAPICall("setChatStickerSet", params, nil, nil)
			if err != nil {
				return fmt.Errorf("SetChatStickerSet: %w", err)
			}
		} else {
			return fmt.Errorf("SetChatStickerSet: %w", err)
		}
	}

	return nil
}

type DeleteChatStickerSetParams struct {
	// Unique identifier for the target chat or username of the target
	// supergroup or channel (in the format @channelusername)
	ChatID ChatIDOrUsername `json:"chat_id"`
}

// Use this method to delete a group sticker set from a supergroup. The bot must
// be an administrator in the chat for this to work and must have the
// appropriate administrator rights. Use the field can_set_sticker_set
// optionally returned in getChat requests to check if the bot can use this
// method. Returns True on success. https://core.telegram.org/bots/api#getchat
//
// https://core.telegram.org/bots/api#deletechatstickerset
func (api *API) DeleteChatStickerSet(params *DeleteChatStickerSetParams) error {
	migrateToChatID, err := api.makeAPICall("deleteChatStickerSet", params, nil, nil)
	if err != nil {
		if migrateToChatID != 0 {
			params.ChatID = migrateToChatID
			_, err = api.makeAPICall("deleteChatStickerSet", params, nil, nil)
			if err != nil {
				return fmt.Errorf("DeleteChatStickerSet: %w", err)
			}
		} else {
			return fmt.Errorf("DeleteChatStickerSet: %w", err)
		}
	}

	return nil
}

type AnswerCallbackQueryParams struct {
	// Unique identifier for the query to be answered
	CallbackQueryID CallbackQueryID `json:"callback_query_id"`
	// Optional. Text of the notification. If not specified, nothing will be
	// shown to the user, 0-200 characters
	Text string `json:"text,omitempty"`
	// Optional. If True, an alert will be shown by the client instead of a
	// notification at the top of the chat screen. Defaults to false.
	ShowAlert bool `json:"show_alert,omitempty"`
	// Optional. URL that will be opened by the user's client. If you have
	// created a Game and accepted the conditions via @Botfather, specify the
	// URL that opens your game — note that this will only work if the query
	// comes from a callback_game button.
	// https://core.telegram.org/bots/api#game https://t.me/botfather
	// https://core.telegram.org/bots/api#inlinekeyboardbutton
	//
	// Otherwise, you may use links like t.me/your_bot?start=XXXX that open your
	// bot with a parameter.
	URL string `json:"url,omitempty"`
	// Optional. The maximum amount of time in seconds that the result of the
	// callback query may be cached client-side. Telegram apps will support
	// caching starting in version 3.14. Defaults to 0.
	CacheTime int `json:"cache_time,omitempty"`
}

// Use this method to send answers to callback queries sent from inline
// keyboards. The answer will be displayed to the user as a notification at the
// top of the chat screen or as an alert. On success, True is returned.
// https://core.telegram.org/bots#inline-keyboards-and-on-the-fly-updating
//
// Alternatively, the user can be redirected to the specified Game URL. For this
// option to work, you must first create a game for your bot via @Botfather and
// accept the terms. Otherwise, you may use links like
// `t.me/your_bot?start=XXXX` that open your bot with a parameter.
// https://t.me/botfather
//
// https://core.telegram.org/bots/api#answercallbackquery
func (api *API) AnswerCallbackQuery(params *AnswerCallbackQueryParams) error {
	_, err := api.makeAPICall("answerCallbackQuery", params, nil, nil)
	if err != nil {
		return fmt.Errorf("AnswerCallbackQuery: %w", err)
	}

	return nil
}

type SetMyCommandsParams struct {
	// A JSON-serialized list of bot commands to be set as the list of the bot's
	// commands. At most 100 commands can be specified.
	Commands []*BotCommand `json:"commands"`
	// Optional. A JSON-serialized object, describing scope of users for which
	// the commands are relevant. Defaults to BotCommandScopeDefault.
	// https://core.telegram.org/bots/api#botcommandscopedefault
	Scope *BotCommandScope `json:"scope,omitempty"`
	// Optional. A two-letter ISO 639-1 language code. If empty, commands will
	// be applied to all users from the given scope, for whose language there
	// are no dedicated commands
	LanguageCode LanguageCode `json:"language_code,omitempty"`
}

// Use this method to change the list of the bot's commands. See
// https://core.telegram.org/bots#commands for more details about bot commands.
// Returns True on success.
//
// https://core.telegram.org/bots/api#setmycommands
func (api *API) SetMyCommands(params *SetMyCommandsParams) error {
	_, err := api.makeAPICall("setMyCommands", params, nil, nil)
	if err != nil {
		return fmt.Errorf("SetMyCommands: %w", err)
	}

	return nil
}

type DeleteMyCommandsParams struct {
	// Optional. A JSON-serialized object, describing scope of users for which
	// the commands are relevant. Defaults to BotCommandScopeDefault.
	// https://core.telegram.org/bots/api#botcommandscopedefault
	Scope *BotCommandScope `json:"scope,omitempty"`
	// Optional. A two-letter ISO 639-1 language code. If empty, commands will
	// be applied to all users from the given scope, for whose language there
	// are no dedicated commands
	LanguageCode LanguageCode `json:"language_code,omitempty"`
}

// Use this method to delete the list of the bot's commands for the given scope
// and user language. After deletion, higher level commands will be shown to
// affected users. Returns True on success.
// https://core.telegram.org/bots/api#determining-list-of-commands
//
// https://core.telegram.org/bots/api#deletemycommands
func (api *API) DeleteMyCommands(params *DeleteMyCommandsParams) error {
	_, err := api.makeAPICall("deleteMyCommands", params, nil, nil)
	if err != nil {
		return fmt.Errorf("DeleteMyCommands: %w", err)
	}

	return nil
}

type GetMyCommandsParams struct {
	// Optional. A JSON-serialized object, describing scope of users. Defaults
	// to BotCommandScopeDefault.
	// https://core.telegram.org/bots/api#botcommandscopedefault
	Scope *BotCommandScope `json:"scope,omitempty"`
	// Optional. A two-letter ISO 639-1 language code or an empty string
	LanguageCode LanguageCode `json:"language_code,omitempty"`
}

// Use this method to get the current list of the bot's commands for the given
// scope and user language. Returns Array of BotCommand on success. If commands
// aren't set, an empty list is returned.
// https://core.telegram.org/bots/api#botcommand
//
// https://core.telegram.org/bots/api#getmycommands
func (api *API) GetMyCommands(params *GetMyCommandsParams) ([]*BotCommand, error) {
	commands := []*BotCommand{}

	_, err := api.makeAPICall("getMyCommands", params, nil, &commands)
	if err != nil {
		return nil, fmt.Errorf("GetMyCommands: %w", err)
	}

	return commands, nil
}

type SetChatMenuButtonParams struct {
	// Optional. Unique identifier for the target private chat. If not
	// specified, default bot's menu button will be changed
	ChatID ChatID `json:"chat_id,omitempty"`
	// Optional. A JSON-serialized object for the new bot's menu button.
	// Defaults to MenuButtonDefault
	// https://core.telegram.org/bots/api#menubuttondefault
	MenuButton *MenuButton `json:"menu_button,omitempty"`
}

// Use this method to change the bot's menu button in a private chat, or the
// default menu button. Returns True on success.
//
// https://core.telegram.org/bots/api#setchatmenubutton
func (api *API) SetChatMenuButton(params *SetChatMenuButtonParams) error {
	migrateToChatID, err := api.makeAPICall("setChatMenuButton", params, nil, nil)
	if err != nil {
		if migrateToChatID != 0 {
			params.ChatID = migrateToChatID
			_, err = api.makeAPICall("setChatMenuButton", params, nil, nil)
			if err != nil {
				return fmt.Errorf("SetChatMenuButton: %w", err)
			}
		} else {
			return fmt.Errorf("SetChatMenuButton: %w", err)
		}
	}

	return nil
}

type GetChatMenuButtonParams struct {
	// Optional. Unique identifier for the target private chat. If not
	// specified, default bot's menu button will be returned
	ChatID ChatID `json:"chat_id,omitempty"`
}

// Use this method to get the current value of the bot's menu button in a
// private chat, or the default menu button. Returns MenuButton on success.
// https://core.telegram.org/bots/api#menubutton
//
// https://core.telegram.org/bots/api#getchatmenubutton
func (api *API) GetChatMenuButton(params *GetChatMenuButtonParams) (*MenuButton, error) {
	var mb *MenuButton

	migrateToChatID, err := api.makeAPICall("getChatMenuButton", params, nil, mb)
	if err != nil {
		if migrateToChatID != 0 {
			params.ChatID = migrateToChatID
			_, err = api.makeAPICall("getChatMenuButton", params, nil, mb)
			if err != nil {
				return nil, fmt.Errorf("GetChatMenuButton: %w", err)
			}
		} else {
			return nil, fmt.Errorf("GetChatMenuButton: %w", err)
		}
	}

	return mb, nil
}

type SetMyDefaultAdministratorRightsParams struct {
	// Optional. A JSON-serialized object describing new default administrator
	// rights. If not specified, the default administrator rights will be
	// cleared.
	Rights *ChatAdministratorRights `json:"rights,omitempty"`
	// Optional. Pass True to change the default administrator rights of the bot
	// in channels. Otherwise, the default administrator rights of the bot for
	// groups and supergroups will be changed.
	ForChannels bool `json:"for_channels,omitempty"`
}

// Use this method to change the default administrator rights requested by the
// bot when it's added as an administrator to groups or channels. These rights
// will be suggested to users, but they are are free to modify the list before
// adding the bot. Returns True on success.
//
// https://core.telegram.org/bots/api#setmydefaultadministratorrights
func (api *API) SetMyDefaultAdministratorRights(params *SetMyDefaultAdministratorRightsParams) error {
	_, err := api.makeAPICall("setMyDefaultAdministratorRights", params, nil, nil)
	if err != nil {
		return fmt.Errorf("SetMyDefaultAdministratorRights: %w", err)
	}

	return nil
}

type GetMyDefaultAdministratorRightsParams struct {
	// Optional. Pass True to change the default administrator rights of the bot
	// in channels. Otherwise, the default administrator rights of the bot for
	// groups and supergroups will be changed.
	ForChannels bool `json:"for_channels,omitempty"`
}

// Use this method to get the current default administrator rights of the bot.
// Returns ChatAdministratorRights on success.
// https://core.telegram.org/bots/api#chatadministratorrights
//
// https://core.telegram.org/bots/api#getmydefaultadministratorrights
func (api *API) GetMyDefaultAdministratorRights(params *GetMyDefaultAdministratorRightsParams) (*ChatAdministratorRights, error) {
	var car *ChatAdministratorRights

	_, err := api.makeAPICall("getMyDefaultAdministratorRights", params, nil, car)
	if err != nil {
		return nil, fmt.Errorf("GetMyDefaultAdministratorRights: %w", err)
	}

	return car, nil
}
