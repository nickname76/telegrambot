// https://core.telegram.org/bots/api#stickers
package telegrambot

import "fmt"

// This object represents a sticker.
//
// https://core.telegram.org/bots/api#sticker
type Sticker struct {
	// Identifier for this file, which can be used to download or reuse the file
	FileID FileID `json:"file_id"`
	// Unique identifier for this file, which is supposed to be the same over
	// time and for different bots. Can't be used to download or reuse the file.
	FileUniqueID FileUniqueID `json:"file_unique_id"`
	// Sticker width
	Width int `json:"width"`
	// Sticker height
	Height int `json:"height"`
	// True, if the sticker is animated
	// https://telegram.org/blog/animated-stickers
	IsAnimated bool `json:"is_animated"`
	// True, if the sticker is a video sticker
	// https://telegram.org/blog/video-stickers-better-reactions
	IsVideo bool `json:"is_video"`
	// Optional. Sticker thumbnail in the .WEBP or .JPG format
	Thumb *PhotoSize `json:"thumb,omitempty"`
	// Optional. Emoji associated with the sticker
	Emoji string `json:"emoji,omitempty"`
	// Optional. Name of the sticker set to which the sticker belongs
	SetName StickerSetName `json:"set_name,omitempty"`
	// Optional. For mask stickers, the position where the mask should be placed
	MaskPosition *MaskPosition `json:"mask_position,omitempty"`
	// Optional. File size in bytes
	FileSize int64 `json:"file_size,omitempty"`
}

// This object represents a sticker set.
//
// https://core.telegram.org/bots/api#stickerset
type StickerSet struct {
	// asdfdsafadsfd
	Name StickerSetName `json:"name"`
	// asdfdsafadsfd
	Title string `json:"title"`
	// asdfdsafadsfd
	IsAnimated bool `json:"is_animated"`
	// asdfdsafadsfd
	IsVideo bool `json:"is_video"`
	// asdfdsafadsfd
	ContainsMasks bool `json:"contains_masks"`
	// asdfdsafadsfd
	Stickers []*Sticker `json:"stickers"`
	// sdafsdafasd
	Thumb *PhotoSize `json:"thumb,omitempty"`
}

// This object describes the position on faces where a mask should be placed by
// default.
//
// https://core.telegram.org/bots/api#maskposition
type MaskPosition struct {
	// The part of the face relative to which the mask should be placed. One of
	// "forehead", "eyes", "mouth", or "chin".
	Point MaskPositionPoint `json:"point"`
	// Shift by X-axis measured in widths of the mask scaled to the face size,
	// from left to right. For example, choosing -1.0 will place mask just to
	// the left of the default mask position.
	XShift float64 `json:"x_shift"`
	// Shift by Y-axis measured in heights of the mask scaled to the face size,
	// from top to bottom. For example, 1.0 will place the mask just below the
	// default mask position.
	YShift float64 `json:"y_shift"`
	// Mask scaling coefficient. For example, 2.0 means double size.
	Scale float64 `json:"scale"`
}

type SendStickerParams struct {
	// Unique identifier for the target chat or username of the target channel
	// (in the format @channelusername)
	ChatID ChatIDOrUsername `json:"chat_id"`
	// Sticker to send. Pass a file_id as String to send a file that exists on
	// the Telegram servers (recommended), pass an HTTP URL as a String for
	// Telegram to get a .WEBP file from the Internet, or upload a new one using
	// multipart/form-data. More info on Sending Files »
	// https://core.telegram.org/bots/api#sending-files
	Sticker InputFile `json:"sticker"`
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

// Use this method to send static .WEBP, animated .TGS, or video .WEBM stickers.
// On success, the sent Message is returned.
// https://telegram.org/blog/animated-stickers
// https://telegram.org/blog/video-stickers-better-reactions
// https://core.telegram.org/bots/api#message
//
// https://core.telegram.org/bots/api#sendsticker
func (api *API) SendSticker(params *SendStickerParams) (*Message, error) {
	msg := &Message{}

	migrateToChatID, err := api.makeAPICall("sendSticker", params, []InputFile{params.Sticker}, msg)
	if err != nil {
		if migrateToChatID != 0 {
			params.ChatID = migrateToChatID
			_, err = api.makeAPICall("sendSticker", params, []InputFile{params.Sticker}, msg)
			if err != nil {
				return nil, fmt.Errorf("SendSticker: %w", err)
			}
		} else {
			return nil, fmt.Errorf("SendSticker: %w", err)
		}
	}

	return msg, nil
}

type GetStickerSetParams struct {
	// Name of the sticker set
	Name StickerSetName `json:"name"`
}

// Use this method to get a sticker set. On success, a StickerSet object is
// returned. https://core.telegram.org/bots/api#stickerset
//
// https://core.telegram.org/bots/api#getstickerset
func (api *API) GetStickerSet(params *GetStickerSetParams) (*StickerSet, error) {
	stickerSet := &StickerSet{}

	_, err := api.makeAPICall("getStickerSet", params, nil, stickerSet)
	if err != nil {
		return nil, fmt.Errorf("GetStickerSet: %w", err)
	}

	return stickerSet, nil
}

type UploadStickerFileParams struct {
	// User identifier of sticker file owner
	UserID UserID `json:"user_id"`
	// PNG image with the sticker, must be up to 512 kilobytes in size,
	// dimensions must not exceed 512px, and either width or height must be
	// exactly 512px. More info on Sending Files »
	// https://core.telegram.org/bots/api#sending-files
	PNGSticker InputFile `json:"png_sticker"`
}

// Use this method to upload a .PNG file with a sticker for later use in
// createNewStickerSet and addStickerToSet methods (can be used multiple times).
// Returns the uploaded File on success. https://core.telegram.org/bots/api#file
//
// https://core.telegram.org/bots/api#uploadstickerfile
func (api *API) UploadStickerFile(params *UploadStickerFileParams) (*File, error) {
	file := &File{}

	_, err := api.makeAPICall("answerPreCheckoutQuery", params, []InputFile{params.PNGSticker}, file)
	if err != nil {
		return nil, fmt.Errorf("AnswerPreCheckoutQuery: %w", err)
	}

	return file, nil
}

type CreateNewStickerSetParams struct {
	// User identifier of created sticker set owner
	UserID UserID `json:"user_id"`
	// Short name of sticker set, to be used in t.me/addstickers/ URLs (e.g.,
	// animals). Can contain only english letters, digits and underscores. Must
	// begin with a letter, can't contain consecutive underscores and must end
	// in "_by_<bot_username>". <bot_username> is case insensitive. 1-64
	// characters.
	Name StickerSetName `json:"name"`
	// Sticker set title, 1-64 characters
	Title string `json:"title"`
	// Optional. *PNG* image with the sticker, must be up to 512 kilobytes in
	// size, dimensions must not exceed 512px, and either width or height must
	// be exactly 512px. Pass a file_id as a String to send a file that already
	// exists on the Telegram servers, pass an HTTP URL as a String for Telegram
	// to get a file from the Internet, or upload a new one using
	// multipart/form-data. More info on Sending Files »
	PNGSticker InputFile `json:"png_sticker,omitempty"`
	// Optional. *TGS* animation with the sticker, uploaded using
	// multipart/form-data. See
	// https://core.telegram.org/stickers#animated-sticker-requirements for
	// technical requirements
	TGSSticker InputFile `json:"tgs_sticker,omitempty"`
	// Optional. *WEBM* video with the sticker, uploaded using
	// multipart/form-data. See
	// https://core.telegram.org/stickers#video-sticker-requirements for
	// technical requirements
	WEBMSticker InputFile `json:"webm_sticker,omitempty"`
	// One or more emoji corresponding to the sticker
	Emojis string `json:"emojis"`
	// Optional. Pass True, if a set of mask stickers should be created
	ContainsMasks bool `json:"contains_masks,omitempty"`
	// Optional. A JSON-serialized object for position where the mask should be
	// placed on faces
	MaskPosition *MaskPosition `json:"mask_position,omitempty"`
}

// Use this method to create a new sticker set owned by a user. The bot will be
// able to edit the sticker set thus created. You *must* use exactly one of the
// fields png_sticker, tgs_sticker, or webm_sticker. Returns True on success.
//
// https://core.telegram.org/bots/api#createnewstickerset
func (api *API) CreateNewStickerSet(params *CreateNewStickerSetParams) error {
	_, err := api.makeAPICall("createNewStickerSet", params, []InputFile{params.PNGSticker, params.TGSSticker, params.WEBMSticker}, nil)
	if err != nil {
		return fmt.Errorf("CreateNewStickerSet: %w", err)
	}

	return nil
}

type AddStickerToSetParams struct {
	// User identifier of sticker set owner
	UserID UserID `json:"user_id"`
	// Sticker set name
	Name StickerSetName `json:"name"`
	// Optional. *PNG* image with the sticker, must be up to 512 kilobytes in
	// size, dimensions must not exceed 512px, and either width or height must
	// be exactly 512px. Pass a file_id as a String to send a file that already
	// exists on the Telegram servers, pass an HTTP URL as a String for Telegram
	// to get a file from the Internet, or upload a new one using
	// multipart/form-data. More info on Sending Files »
	// https://core.telegram.org/bots/api#sending-files
	PNGSticker InputFile `json:"png_sticker,omitempty"`
	// Optional. *TGS* animation with the sticker, uploaded using
	// multipart/form-data. See
	// https://core.telegram.org/stickers#animated-sticker-requirements for
	// technical requirements
	TGSSticker InputFile `json:"tgs_sticker,omitempty"`
	// Optional. *WEBM* video with the sticker, uploaded using
	// multipart/form-data. See
	// https://core.telegram.org/stickers#video-sticker-requirements for
	// technical requirements
	WEBMSticker InputFile `json:"webm_sticker,omitempty"`
	// One or more emoji corresponding to the sticker
	Emojis string `json:"emojis"`
	// Optional. A JSON-serialized object for position where the mask should be
	// placed on faces
	MaskPosition *MaskPosition `json:"mask_position,omitempty"`
}

// Use this method to add a new sticker to a set created by the bot. You *must*
// use exactly one of the fields png_sticker, tgs_sticker, or webm_sticker.
// Animated stickers can be added to animated sticker sets and only to them.
// Animated sticker sets can have up to 50 stickers. Static sticker sets can
// have up to 120 stickers. Returns True on success.
//
// https://core.telegram.org/bots/api#addstickertoset
func (api *API) AddStickerToSet(params *AddStickerToSetParams) error {
	_, err := api.makeAPICall("addStickerToSet", params, []InputFile{params.PNGSticker, params.TGSSticker, params.WEBMSticker}, nil)
	if err != nil {
		return fmt.Errorf("AddStickerToSet: %w", err)
	}

	return nil
}

type SetStickerPositionInSetParams struct {
	// File identifier of the sticker
	Sticker FileID `json:"sticker"`
	// New sticker position in the set, zero-based
	Position int `json:"position"`
}

// Use this method to move a sticker in a set created by the bot to a specific
// position. Returns True on success.
//
// https://core.telegram.org/bots/api#setstickerpositioninset
func (api *API) SetStickerPositionInSet(params *SetStickerPositionInSetParams) error {
	_, err := api.makeAPICall("setStickerPositionInSet", params, nil, nil)
	if err != nil {
		return fmt.Errorf("SetStickerPositionInSet: %w", err)
	}

	return nil
}

type DeleteStickerFromSetParams struct {
	// File identifier of the sticker
	Sticker FileID `json:"sticker"`
}

// Use this method to delete a sticker from a set created by the bot. Returns
// True on success.
//
// https://core.telegram.org/bots/api#deletestickerfromset
func (api *API) DeleteStickerFromSet(params *DeleteStickerFromSetParams) error {
	_, err := api.makeAPICall("deleteStickerFromSet", params, nil, nil)
	if err != nil {
		return fmt.Errorf("DeleteStickerFromSet: %w", err)
	}

	return nil
}

type SetStickerSetThumbParams struct {
	// Sticker set name
	Name StickerSetName `json:"name"`
	// User identifier of the sticker set owner
	UserID UserID `json:"user_id"`
	// Optional. A PNG image with the thumbnail, must be up to 128 kilobytes in
	// size and have width and height exactly 100px, or a TGS animation with the
	// thumbnail up to 32 kilobytes in size; see
	// https://core.telegram.org/stickers#animated-sticker-requirements for
	// animated sticker technical requirements, or a WEBM video with the
	// thumbnail up to 32 kilobytes in size; see
	// https://core.telegram.org/stickers#video-sticker-requirements for video
	// sticker technical requirements. Pass a file_id as a String to send a file
	// that already exists on the Telegram servers, pass an HTTP URL as a String
	// for Telegram to get a file from the Internet, or upload a new one using
	// multipart/form-data. More info on Sending Files ». Animated sticker set
	// thumbnails can't be uploaded via HTTP URL.
	// https://core.telegram.org/bots/api#sending-files
	Thumb InputFile `json:"thumb,omitempty"`
}

// Use this method to set the thumbnail of a sticker set. Animated thumbnails
// can be set for animated sticker sets only. Video thumbnails can be set only
// for video sticker sets only. Returns True on success.
//
// https://core.telegram.org/bots/api#setstickersetthumb
func (api *API) SetStickerSetThumb(params *SetStickerSetThumbParams) error {
	_, err := api.makeAPICall("setStickerSetThumb", params, []InputFile{params.Thumb}, nil)
	if err != nil {
		return fmt.Errorf("SetStickerSetThumb: %w", err)
	}

	return nil
}
