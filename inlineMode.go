package telegrambot

// https://core.telegram.org/bots/api#inline-mode

import "fmt"

// This object represents an incoming inline query. When the user sends an empty
// query, your bot could return some default or trending results.
//
// https://core.telegram.org/bots/api#inlinequery
type InlineQuery struct {
	// Unique identifier for this query
	ID InlineQueryID `json:"id"`
	// Sender
	From *User `json:"from"`
	// Text of the query (up to 256 characters)
	Query string `json:"query"`
	// Offset of the results to be returned, can be controlled by the bot
	Offset string `json:"offset"`
	// Optional. Type of the chat, from which the inline query was sent. Can be
	// either "sender" for a private chat with the inline query sender,
	// "private", "group", "supergroup", or "channel". The chat type should be
	// always known for requests sent from official clients and most third-party
	// clients, unless the request was sent from a secret chat
	ChatType InlineQueryChatType `json:"chat_type,omitempty"`
	// Optional. Sender location, only for bots that request user location
	Location *Location `json:"location,omitempty"`
}

type AnswerInlineQueryParams struct {
	// Unique identifier for the answered query
	InlineQueryID InlineQueryID `json:"inline_query_id"`
	// A JSON-serialized array of results for the inline query
	Results []*InlineQueryResult `json:"results"`
	// The maximum amount of time in seconds that the result of the inline query
	// may be cached on the server. Defaults to 300.
	CacheTime int `json:"cache_time,omitempty"`
	// Pass True, if results may be cached on the server side only for the user
	// that sent the query. By default, results may be returned to any user who
	// sends the same query
	IsPersonal bool `json:"is_personal,omitempty"`
	// Pass the offset that a client should send in the next query with the same
	// text to receive more results. Pass an empty string if there are no more
	// results or if you don't support pagination. Offset length can't exceed 64
	// bytes.
	NextOffset string `json:"next_offset,omitempty"`
	// If passed, clients will display a button with specified text that
	// switches the user to a private chat with the bot and sends the bot a
	// start message with the parameter switch_pm_parameter
	SwitchPMText string `json:"switch_pm_text,omitempty"`
	// Deep-linking parameter for the /start message sent to the bot when user
	// presses the switch button. 1-64 characters, only `A-Z`, `a-z`, `0-9`, `_`
	// and `-` are allowed. https://core.telegram.org/bots#deep-linking
	// https://core.telegram.org/bots/api#inlinekeyboardmarkup
	SwitchPMParameter string `json:"switch_pm_parameter,omitempty"`
}

// Use this method to send answers to an inline query. On success, True is
// returned. No more than *50* results per query are allowed.
//
// https://core.telegram.org/bots/api#answerinlinequery
func (api *API) AnswerInlineQuery(params *AnswerInlineQueryParams) error {
	_, err := api.makeAPICall("getGameHighScores", params, nil, nil)
	if err != nil {
		return fmt.Errorf("GetGameHighScores: %w", err)
	}

	return nil
}

// This object represents one result of an inline query. Telegram clients
// currently support results of the following 20 types:
//   InlineQueryResultCachedAudio - Represents a link to an MP3 audio file stored on the Telegram servers. By default, this audio file will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the audio.
//   InlineQueryResultCachedDocument - Represents a link to a file stored on the Telegram servers. By default, this file will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the file.
//   InlineQueryResultCachedGif - Represents a link to an animated GIF file stored on the Telegram servers. By default, this animated GIF file will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with specified content instead of the animation.
//   InlineQueryResultCachedMpeg4Gif - Represents a link to a video animation (H.264/MPEG-4 AVC video without sound) stored on the Telegram servers. By default, this animated MPEG-4 file will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the animation.
//   InlineQueryResultCachedPhoto - Represents a link to a photo stored on the Telegram servers. By default, this photo will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the photo.
//   InlineQueryResultCachedSticker - Represents a link to a sticker stored on the Telegram servers. By default, this sticker will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the sticker.
//   InlineQueryResultCachedVideo - Represents a link to a video file stored on the Telegram servers. By default, this video file will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the video.
//   InlineQueryResultCachedVoice - Represents a link to a voice message stored on the Telegram servers. By default, this voice message will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the voice message.
//   InlineQueryResultArticle - Represents a link to an article or web page.
//   InlineQueryResultAudio - Represents a link to an MP3 audio file. By default, this audio file will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the audio.
//   InlineQueryResultContact - Represents a contact with a phone number. By default, this contact will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the contact.
//   InlineQueryResultGame - Represents a Game. https://core.telegram.org/bots/api#games
//   InlineQueryResultDocument - Represents a link to a file. By default, this file will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the file. Currently, only *.PDF* and *.ZIP* files can be sent using this method.
//   InlineQueryResultGif - Represents a link to an animated GIF file. By default, this animated GIF file will be sent by the user with optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the animation.
//   InlineQueryResultLocation - Represents a location on a map. By default, the location will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the location.
//   InlineQueryResultMpeg4Gif - Represents a link to a video animation (H.264/MPEG-4 AVC video without sound). By default, this animated MPEG-4 file will be sent by the user with optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the animation.
//   InlineQueryResultPhoto - Represents a link to a photo. By default, this photo will be sent by the user with optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the photo.
//   InlineQueryResultVenue - Represents a venue. By default, the venue will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the venue.
//   InlineQueryResultVideo - Represents a link to a page containing an embedded video player or a video file. By default, this video file will be sent by the user with an optional caption. Alternatively, you can use input_message_content to send a message with the specified content instead of the video. (If an InlineQueryResultVideo message contains an embedded video (e.g., YouTube), you *must* replace its content using input_message_content.)
//   InlineQueryResultVoice - Represents a link to a voice recording in an .OGG container encoded with OPUS. By default, this voice recording will be sent by the user. Alternatively, you can use input_message_content to send a message with the specified content instead of the the voice message.
//
// Note: All URLs passed in inline query results will be available to end users
// and therefore must be assumed to be *public*.
//
// https://core.telegram.org/bots/api#inlinequeryresult
// https://core.telegram.org/bots/api#inlinequeryresultcachedaudio
// https://core.telegram.org/bots/api#inlinequeryresultcacheddocument
// https://core.telegram.org/bots/api#inlinequeryresultcachedgif
// https://core.telegram.org/bots/api#inlinequeryresultcachedmpeg4gif
// https://core.telegram.org/bots/api#inlinequeryresultcachedphoto
// https://core.telegram.org/bots/api#inlinequeryresultcachedsticker
// https://core.telegram.org/bots/api#inlinequeryresultcachedvideo
// https://core.telegram.org/bots/api#inlinequeryresultcachedvoice
// https://core.telegram.org/bots/api#inlinequeryresultarticle
// https://core.telegram.org/bots/api#inlinequeryresultaudio
// https://core.telegram.org/bots/api#inlinequeryresultcontact
// https://core.telegram.org/bots/api#inlinequeryresultgame
// https://core.telegram.org/bots/api#inlinequeryresultdocument
// https://core.telegram.org/bots/api#inlinequeryresultgif
// https://core.telegram.org/bots/api#inlinequeryresultlocation
// https://core.telegram.org/bots/api#inlinequeryresultmpeg4gif
// https://core.telegram.org/bots/api#inlinequeryresultphoto
// https://core.telegram.org/bots/api#inlinequeryresultvenue
// https://core.telegram.org/bots/api#inlinequeryresultvideo
// https://core.telegram.org/bots/api#inlinequeryresultvoice
type InlineQueryResult struct {
	// Type of the result
	//   InlineQueryResultArticle - must be article
	//   InlineQueryResultPhoto - must be photo
	//   InlineQueryResultGif - must be gif
	//   InlineQueryResultMpeg4Gif - must be mpeg4_gif
	//   InlineQueryResultVideo - must be video
	//   InlineQueryResultAudio - must be audio
	//   InlineQueryResultVoice - must be voice
	//   InlineQueryResultDocument - must be document
	//   InlineQueryResultLocation - must be location
	//   InlineQueryResultVenue - must be venue
	//   InlineQueryResultContact - must be contact
	//   InlineQueryResultGame - must be game
	//   InlineQueryResultCachedPhoto - must be photo
	//   InlineQueryResultCachedGif - must be gif
	//   InlineQueryResultCachedMpeg4Gif - must be mpeg4_gif
	//   InlineQueryResultCachedSticker - must be sticker
	//   InlineQueryResultCachedDocument - must be document
	//   InlineQueryResultCachedVideo - must be video
	//   InlineQueryResultCachedVoice - must be voice
	//   InlineQueryResultCachedAudio - must be audio
	Type InlineQueryResultType `json:"type"`
	// Unique identifier for this result, 1-64 Bytes
	ID InlineQueryResultID `json:"id"`

	// Title of the result
	Title string `json:"title,omitempty"`
	// Content of the message to be sent
	InputMessageContent *InputMessageContent `json:"input_message_content,omitempty"`
	// Optional. Inline keyboard attached to the message
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	// Optional. Short description of the result
	Description string `json:"description,omitempty"`
	// Optional. Url of the thumbnail for the result
	ThumbURL string `json:"thumb_url,omitempty"`
	// Optional. Thumbnail width
	ThumbWidth int `json:"thumb_width,omitempty"`
	// Optional. Thumbnail height
	ThumbHeight int `json:"thumb_height,omitempty"`

	// Optional. Caption of the result to be sent, 0-1024 characters after
	// entities parsing
	Caption string `json:"caption,omitempty"`
	// Optional. Mode for parsing entities in the result caption. See formatting
	// options for more details.
	// https://core.telegram.org/bots/api#formatting-options
	ParseMode ParseMode `json:"parse_mode,omitempty"`
	// Optional. List of special entities that appear in the caption, which can
	// be specified instead of parse_mode
	CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"`

	// Optional. URL of the result
	URL string `json:"url,omitempty"`
	// Optional. Pass True, if you don't want the URL to be shown in the message
	HideURL bool `json:"hide_url,omitempty"`

	// A valid URL of the photo. Photo must be in JPEG format. Photo size must
	// not exceed 5MB
	PhotoURL string `json:"photo_url,omitempty"`

	// A valid URL for the GIF file. File size must not exceed 1MB
	GifURL string `json:"gif_url,omitempty"`
	// Optional. Width of the GIF
	GifWidth int `json:"gif_width,omitempty"`
	// Optional. Height of the GIF
	GifHeight int `json:"gif_height,omitempty"`
	// Optional. Duration of the GIF in seconds
	GifDuration int `json:"gif_duration,omitempty"`

	// Optional. MIME type of the thumbnail
	ThumbMimeType string `json:"thumb_mime_type,omitempty"`

	// A valid URL for the MP4 file. File size must not exceed 1MB
	Mpeg4URL string `json:"mpeg4_url,omitempty"`
	// Optional. Video width
	Mpeg4Width int `json:"mpeg4_width,omitempty"`
	// Optional. Video height
	Mpeg4Height int `json:"mpeg4_height,omitempty"`
	// Optional. Video duration in seconds
	Mpeg4Duration int `json:"mpeg4_duration,omitempty"`

	// Mime type of the content of result url
	MimeType string `json:"mime_type,omitempty"`

	// A valid URL for the embedded video player or video file
	VideoURL string `json:"video_url,omitempty"`
	// Optional. Video width
	VideoWidth int `json:"video_width,omitempty"`
	// Optional. Video height
	VideoHeight int `json:"video_height,omitempty"`
	// Optional. Video duration in seconds
	VideoDuration int `json:"video_duration,omitempty"`

	// A valid URL for the audio file
	AudioURL string `json:"audio_url,omitempty"`
	// Optional. Performer
	Performer string `json:"performer,omitempty"`
	// Optional. Audio duration in seconds
	AudioDuration int `json:"audio_duration,omitempty"`

	// A valid URL for the voice recording
	VoiceURL string `json:"voice_url,omitempty"`

	// A valid URL for the file
	DocumentURL string `json:"document_url,omitempty"`

	// Location latitude in degrees
	Latitude float64 `json:"latitude,omitempty"`
	// Location longitude in degrees
	Longitude float64 `json:"longitude,omitempty"`

	// Optional. The radius of uncertainty for the location, measured in meters;
	// 0-1500
	HorizontalAccuracy float64 `json:"horizontal_accuracy,omitempty"`
	// Optional. Period in seconds for which the location can be updated, should
	// be between 60 and 86400.
	LivePeriod int `json:"live_period,omitempty"`
	// Optional. For live locations, a direction in which the user is moving, in
	// degrees. Must be between 1 and 360 if specified.
	Heading int `json:"heading,omitempty"`
	// Optional. For live locations, a maximum distance for proximity alerts
	// about approaching another chat member, in meters. Must be between 1 and
	// 100000 if specified.
	ProximityAlertRadius int `json:"proximity_alert_radius,omitempty"`

	// Address of the venue
	Address string `json:"address,omitempty"`
	// Optional. Foursquare identifier of the venue if known
	FoursquareID string `json:"foursquare_id,omitempty"`
	// Optional. Foursquare type of the venue, if known. (For example,
	// "arts_entertainment/default", "arts_entertainment/aquarium" or
	// "food/icecream".)
	FoursquareType string `json:"foursquare_type,omitempty"`
	// Optional. Google Places identifier of the venue
	GooglePlaceID string `json:"google_place_id,omitempty"`
	// Optional. Google Places type of the venue. (See supported types.)
	// https://developers.google.com/places/web-service/supported_types
	GooglePlaceType string `json:"google_place_type,omitempty"`

	// Contact's phone number
	PhoneNumber string `json:"phone_number,omitempty"`
	// Contact's first name
	FirstName string `json:"first_name,omitempty"`
	// Optional. Contact's last name
	LastName string `json:"last_name,omitempty"`
	// Optional. Additional data about the contact in the form of a vCard,
	// 0-2048 bytes https://en.wikipedia.org/wiki/VCard
	VCard string `json:"vcard,omitempty"`

	// Short name of the game
	GameShortName GameShortName `json:"game_short_name,omitempty"`

	// A valid file identifier of the photo
	PhotoFileID FileID `json:"photo_file_id,omitempty"`

	// A valid file identifier for the GIF file
	GifFileID FileID `json:"gif_file_id,omitempty"`

	// A valid file identifier for the MP4 file
	Mpeg4FileID FileID `json:"mpeg4_file_id,omitempty"`

	// A valid file identifier of the sticker
	StickerFileID FileID `json:"sticker_file_id,omitempty"`

	// A valid file identifier for the file
	DocumentFileID FileID `json:"document_file_id,omitempty"`

	// A valid file identifier for the video file
	VideoFileID FileID `json:"video_file_id,omitempty"`

	// A valid file identifier for the voice message
	VoiceFileID FileID `json:"voice_file_id,omitempty"`

	// A valid file identifier for the audio file
	AudioFileID FileID `json:"audio_file_id,omitempty"`
}

// This object represents the content of a message to be sent as a result of an
// inline query. Telegram clients currently support the following 5 types:
//   InputTextMessageContent - Represents the content of a text message to be sent as the result of an inline query.
//   InputLocationMessageContent - Represents the content of a location message to be sent as the result of an inline query.
//   InputVenueMessageContent - Represents the content of a venue message to be sent as the result of an inline query.
//   InputContactMessageContent - Represents the content of a contact message to be sent as the result of an inline query.
//   InputInvoiceMessageContent - Represents the content of an invoice message to be sent as the result of an inline query.
//
// https://core.telegram.org/bots/api#inputmessagecontent
// https://core.telegram.org/bots/api#inputtextmessagecontent
// https://core.telegram.org/bots/api#inputlocationmessagecontent
// https://core.telegram.org/bots/api#inputvenuemessagecontent
// https://core.telegram.org/bots/api#inputcontactmessagecontent
// https://core.telegram.org/bots/api#inputinvoicemessagecontent
type InputMessageContent struct {
	// Text of the message to be sent, 1-4096 characters
	MessageText string `json:"message_text,omitempty"`
	// Optional. Mode for parsing entities in the message text. See formatting
	// options for more details.
	// https://core.telegram.org/bots/api#formatting-options
	ParseMode ParseMode `json:"parse_mode,omitempty"`
	// Optional. List of special entities that appear in message text, which can
	// be specified instead of parse_mode
	Entities []*MessageEntity `json:"entities,omitempty"`
	// Optional. Disables link previews for links in the sent message
	DisableWebPagePreview bool `json:"disable_web_page_preview,omitempty"`

	// Latitude of the location in degrees
	Latitude float64 `json:"latitude,omitempty"`
	// Longitude of the location in degrees
	Longitude float64 `json:"longitude,omitempty"`

	// Optional. The radius of uncertainty for the location, measured in meters;
	// 0-1500
	HorizontalAccuracy float64 `json:"horizontal_accuracy,omitempty"`
	// Optional. Period in seconds for which the location can be updated, should
	// be between 60 and 86400.
	LivePeriod int `json:"live_period,omitempty"`
	// Optional. For live locations, a direction in which the user is moving, in
	// degrees. Must be between 1 and 360 if specified.
	Heading int `json:"heading,omitempty"`
	// Optional. For live locations, a maximum distance for proximity alerts
	// about approaching another chat member, in meters. Must be between 1 and
	// 100000 if specified.
	ProximityAlertRadius int `json:"proximity_alert_radius,omitempty"`

	// Name of the venue / Product name, 1-32 characters
	Title string `json:"title,omitempty"`
	// Address of the venue
	Address string `json:"address,omitempty"`
	// Optional. Foursquare identifier of the venue, if known
	FoursquareID string `json:"foursquare_id,omitempty"`
	// Optional. Foursquare type of the venue, if known. (For example,
	// "arts_entertainment/default", "arts_entertainment/aquarium" or
	// "food/icecream".)
	FoursquareType string `json:"foursquare_type,omitempty"`
	// Optional. Google Places identifier of the venue
	GooglePlaceID string `json:"google_place_id,omitempty"`
	// Optional. Google Places type of the venue. (See supported types.)
	// https://developers.google.com/places/web-service/supported_types
	GooglePlaceType string `json:"google_place_type,omitempty"`

	// Contact's phone number
	PhoneNumber string `json:"phone_number,omitempty"`
	// Contact's first name
	FirstName string `json:"first_name,omitempty"`
	// Optional. Contact's last name
	LastName string `json:"last_name,omitempty"`
	// Optional. Additional data about the contact in the form of a vCard,
	// 0-2048 bytes https://en.wikipedia.org/wiki/VCard
	VCard string `json:"vcard,omitempty"`

	// Product description, 1-255 characters
	Description string `json:"description,omitempty"`
	// Bot-defined invoice payload, 1-128 bytes. This will not be displayed to
	// the user, use for your internal processes.
	Payload string `json:"payload,omitempty"`
	// Payments provider token, obtained via Botfather https://t.me/botfather
	ProviderToken string `json:"provider_token,omitempty"`
	// Three-letter ISO 4217 currency code, see more on currencies
	// https://core.telegram.org/bots/payments#supported-currencies
	Currency string `json:"currency,omitempty"`
	// Price breakdown, a JSON-serialized list of components (e.g. product
	// price, tax, discount, delivery cost, delivery tax, bonus, etc.)
	Prices []*LabeledPrice `json:"prices,omitempty"`
	// Optional. The maximum accepted amount for tips in the smallest units of
	// the currency (integer, not float/double). For example, for a maximum tip
	// of US$ 1.45 pass max_tip_amount = 145. See the exp parameter in
	// currencies.json, it shows the number of digits past the decimal point for
	// each currency (2 for the majority of currencies). Defaults to 0
	// https://core.telegram.org/bots/payments/currencies.json
	MaxTipAmount int `json:"max_tip_amount,omitempty"`
	// Optional. A JSON-serialized array of suggested amounts of tip in the
	// smallest units of the currency (integer, not float/double). At most 4
	// suggested tip amounts can be specified. The suggested tip amounts must be
	// positive, passed in a strictly increased order and must not exceed
	// max_tip_amount.
	SuggestedTipAmounts []int `json:"suggested_tip_amounts,omitempty"`
	// Optional. A JSON-serialized object for data about the invoice, which will
	// be shared with the payment provider. A detailed description of the
	// required fields should be provided by the payment provider.
	ProviderData string `json:"provider_data,omitempty"`
	// Optional. URL of the product photo for the invoice. Can be a photo of the
	// goods or a marketing image for a service. People like it better when they
	// see what they are paying for.
	PhotoURL string `json:"photo_url,omitempty"`
	// Optional. Photo size
	PhotoSize int `json:"photo_size,omitempty"`
	// Optional. Photo width
	PhotoWidth int `json:"photo_width,omitempty"`
	// Optional. Photo height
	PhotoHeight int `json:"photo_height,omitempty"`
	// Optional. Pass True, if you require the user's full name to complete the
	// order
	NeedName bool `json:"need_name,omitempty"`
	// Optional. Pass True, if you require the user's phone number to complete
	// the order
	NeedPhoneNumber bool `json:"need_phone_number,omitempty"`
	// Optional. Pass True, if you require the user's email address to complete
	// the order
	NeedEmail bool `json:"need_email,omitempty"`
	// Optional. Pass True, if you require the user's shipping address to
	// complete the order
	NeedShippingAddress bool `json:"need_shipping_address,omitempty"`
	// Optional. Pass True, if user's phone number should be sent to provider
	SendPhoneNumberToProvider bool `json:"send_phone_number_to_provider,omitempty"`
	// Optional. Pass True, if user's email address should be sent to provider
	SendEmailToProvider bool `json:"send_email_to_provider,omitempty"`
	// Optional. Pass True, if the final price depends on the shipping method
	IsFlexible bool `json:"is_flexible,omitempty"`
}

// Represents a result of an inline query that was chosen by the user and sent
// to their chat partner. https://core.telegram.org/bots/api#inlinequeryresult
//
// Note: It is necessary to enable inline feedback via @BotFather in order to
// receive these objects in updates.
// https://core.telegram.org/bots/inline#collecting-feedback
// https://t.me/botfather
//
// https://core.telegram.org/bots/api#choseninlineresult
type ChosenInlineResult struct {
	// The unique identifier for the result that was chosen
	ResultID InlineQueryResultID `json:"result_id"`
	// The user that chose the result
	From *User `json:"from"`
	// Optional. Sender location, only for bots that require user location
	Location *Location `json:"location,omitempty"`
	// Optional. Identifier of the sent inline message. Available only if there
	// is an inline keyboard attached to the message. Will be also received in
	// callback queries and can be used to edit the message.
	// https://core.telegram.org/bots/api#inlinekeyboardmarkup
	// https://core.telegram.org/bots/api#callbackquery
	// https://core.telegram.org/bots/api#updating-messages
	InlineMessageID InlineMessageID `json:"inline_message_id,omitempty"`
	// The query that was used to obtain the result
	Query string `json:"query"`
}

type AnswerWebAppQueryParams struct {
	// Unique identifier for the query to be answered
	WebAppQueryID WebAppQueryID `json:"web_app_query_id"`
	// A JSON-serialized object describing the message to be sent
	Result *InlineQueryResult `json:"result"`
}

// Use this method to set the result of an interaction with a Web App and send a
// corresponding message on behalf of the user to the chat from which the query
// originated. On success, a SentWebAppMessage object is returned.
// https://core.telegram.org/bots/webapps
// https://core.telegram.org/bots/api#sentwebappmessage
//
// https://core.telegram.org/bots/api#answerwebappquery
func (api *API) AnswerWebAppQuery(params *AnswerWebAppQueryParams) (*SentWebAppMessage, error) {
	var swamsg *SentWebAppMessage

	_, err := api.makeAPICall("answerWebAppQuery", params, nil, swamsg)
	if err != nil {
		return nil, fmt.Errorf("AnswerWebAppQuery: %w", err)
	}

	return swamsg, nil
}

// Contains information about an inline message sent by a Web App on behalf of a
// user. https://core.telegram.org/bots/webapps
//
// https://core.telegram.org/bots/api#sentwebappmessage
type SentWebAppMessage struct {
	// Optional. Identifier of the sent inline message. Available only if there
	// is an inline keyboard attached to the message.
	// https://core.telegram.org/bots/api#inlinekeyboardmarkup
	InlineMessageID InlineMessageID `json:"inline_message_id,omitempty"`
}
