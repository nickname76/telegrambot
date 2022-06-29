package telegrambot

// Types which are not exactly specified in the official Bot API documentation

// Unique user identifier
type UserID ChatID

// Unique chat identifier
type ChatID int64

// Unique user specified string identifier
type Username string

// Literally ChatID or Username typed value
type ChatIDOrUsername interface {
	chatIDOrUsername()
}

func (ChatID) chatIDOrUsername()   {}
func (Username) chatIDOrUsername() {}

// Unique identifier for file, which is supposed to be the same over time and
// for different bots. Can't be used to download or reuse the file.
type FileUniqueID string

// A two-letter ISO 639-1 language code.
type LanguageCode string

const (
	LanguageCodeRussian    LanguageCode = "ru"
	LanguageCodeEnglish    LanguageCode = "en"
	LanguageCodeBelarusian LanguageCode = "be"
	LanguageCodeUkrainian  LanguageCode = "uk"
	LanguageCodeKorean     LanguageCode = "ko"
	LanguageCodeCatalan    LanguageCode = "ca"
	LanguageCodeDutch      LanguageCode = "nl"
	LanguageCodeFrench     LanguageCode = "fr"
	LanguageCodeGerman     LanguageCode = "de"
	LanguageCodeItalian    LanguageCode = "it"
	LanguageCodeMalay      LanguageCode = "ms"
	LanguageCodePolish     LanguageCode = "pl"
	LanguageCodePortuguese LanguageCode = "pt"
	LanguageCodeSpanish    LanguageCode = "es"
	LanguageCodeTurkish    LanguageCode = "tr"
)

// Type of the chat. Can be either “sender” for a private chat with the inline
// query sender, “private”, “group”, “supergroup”, or “channel”.
type ChatType string

const (
	ChatTypePrivate    ChatType = "private"
	ChatTypeGroup      ChatType = "group"
	ChatTypeSupergroup ChatType = "supergroup"
	ChatTypeChannel    ChatType = "channel"
)

// Unique message identifier inside chat
type MessageID int

// Type of the entity.
type MessageEntityType string

const (
	MessageEntityTypeMention       MessageEntityType = "mention"
	MessageEntityTypeHashtag       MessageEntityType = "hashtag"
	MessageEntityTypeCashtag       MessageEntityType = "cashtag"
	MessageEntityTypeBotCommand    MessageEntityType = "bot_command"
	MessageEntityTypeURL           MessageEntityType = "url"
	MessageEntityTypeEmail         MessageEntityType = "email"
	MessageEntityTypePhoneNumber   MessageEntityType = "phone_number"
	MessageEntityTypeBold          MessageEntityType = "bold"
	MessageEntityTypeItalic        MessageEntityType = "italic"
	MessageEntityTypeUnderline     MessageEntityType = "underline"
	MessageEntityTypeStrikethrough MessageEntityType = "strikethrough"
	MessageEntityTypeSpoiler       MessageEntityType = "spoiler"
	MessageEntityTypeCode          MessageEntityType = "code"
	MessageEntityTypePre           MessageEntityType = "pre"
	MessageEntityTypeTextLink      MessageEntityType = "text_link"
	MessageEntityTypeTextMention   MessageEntityType = "text_mention"
)

// Short name of a Game, serves as the unique identifier for the game
// https://core.telegram.org/bots/api#games
type GameShortName string

// Unique poll identifier
type PollID string

// Poll type, currently can be “regular” or “quiz”
type PollType string

const (
	PollTypeRegular PollType = "regular"
	PollTypeQuiz    PollType = "quiz"
)

// Unique identifier for callback query
type CallbackQueryID string

// Global identifier, uniquely corresponding to the chat to which a message with
// the callback button was sent. Useful for high scores in games.
// https://core.telegram.org/bots/api#games
type ChatInstance string

// The member's status in a chat
type ChatMemberStatus string

const (
	ChatMemberStatusCreator       ChatMemberStatus = "creator"
	ChatMemberStatusAdministrator ChatMemberStatus = "administrator"
	ChatMemberStatusMember        ChatMemberStatus = "member"
	ChatMemberStatusRestricted    ChatMemberStatus = "restricted"
	ChatMemberStatusLeft          ChatMemberStatus = "left"
	ChatMemberStatusKicked        ChatMemberStatus = "kicked"
)

type BotCommandScopeType string

const (
	BotCommandScopeTypeDefault               BotCommandScopeType = "default"
	BotCommandScopeTypeAllPrivateChats       BotCommandScopeType = "all_private_chats"
	BotCommandScopeTypeAllGroupChats         BotCommandScopeType = "all_group_chats"
	BotCommandScopeTypeAllChatAdministrators BotCommandScopeType = "all_chat_administrators"
	BotCommandScopeTypeChat                  BotCommandScopeType = "chat"
	BotCommandScopeTypeChatAdministrator     BotCommandScopeType = "chat_administrators"
	BotCommandScopeTypeChatMember            BotCommandScopeType = "chat_member"
)

// https://core.telegram.org/bots/api#formatting-options
type ParseMode string

const (
	// https://core.telegram.org/bots/api#markdownv2-style
	ParseModeMarkdownV2 ParseMode = "MarkdownV2"
	// https://core.telegram.org/bots/api#html-style
	ParseModeHTML ParseMode = "HTML"
	// https://core.telegram.org/bots/api#markdown-style
	ParseModeMarkdown ParseMode = "Markdown"
)

type InputMediaType string

const (
	InputMediaTypePhoto     InputMediaType = "photo"
	InputMediaTypeVideo     InputMediaType = "video"
	InputMediaTypeAnimation InputMediaType = "animation"
	InputMediaTypeAudio     InputMediaType = "audio"
	InputMediaTypeDocument  InputMediaType = "document"
)

type StickerSetName string

type MaskPositionPoint string

const (
	MaskPositionPointForehead MaskPositionPoint = "forehead"
	MaskPositionPointEyes     MaskPositionPoint = "eyes"
	MaskPositionPointMounth   MaskPositionPoint = "mouth"
	MaskPositionPointChin     MaskPositionPoint = "chin"
)

// Unique identifier for an inline query
type InlineQueryID string

type InlineQueryChatType ChatType

const (
	InlineQueryChatTypeSender     InlineQueryChatType = "sender"
	InlineQueryChatTypePrivate    InlineQueryChatType = InlineQueryChatType(ChatTypePrivate)
	InlineQueryChatTypeGroup      InlineQueryChatType = InlineQueryChatType(ChatTypeGroup)
	InlineQueryChatTypeSupergroup InlineQueryChatType = InlineQueryChatType(ChatTypeSupergroup)
	InlineQueryChatTypeChannel    InlineQueryChatType = InlineQueryChatType(ChatTypeChannel)
)

// Type of an inline query result
type InlineQueryResultType string

const (
	InlineQueryResultTypeArticle  InlineQueryResultType = "article"
	InlineQueryResultTypePhoto    InlineQueryResultType = "photo"
	InlineQueryResultTypeGif      InlineQueryResultType = "gif"
	InlineQueryResultTypeMpeg4Gif InlineQueryResultType = "mpeg4_gif"
	InlineQueryResultTypeVideo    InlineQueryResultType = "video"
	InlineQueryResultTypeAudio    InlineQueryResultType = "audio"
	InlineQueryResultTypeVoice    InlineQueryResultType = "voice"
	InlineQueryResultTypeDocument InlineQueryResultType = "document"
	InlineQueryResultTypeLocation InlineQueryResultType = "location"
	InlineQueryResultTypeVenue    InlineQueryResultType = "venue"
	InlineQueryResultTypeContact  InlineQueryResultType = "contact"
	InlineQueryResultTypeGame     InlineQueryResultType = "game"
	InlineQueryResultTypeSticker  InlineQueryResultType = "sticker"
)

type InlineQueryResultID string

type UpdateID int

type UpdateType string

const (
	UpdateTypeMessage            UpdateType = "message"
	UpdateTypeEditedMessage      UpdateType = "edited_message"
	UpdateTypeChannelPost        UpdateType = "channel_post"
	UpdateTypeEditedChannelPost  UpdateType = "edited_channel_post"
	UpdateTypeInlineQuery        UpdateType = "inline_query"
	UpdateTypeChosenInlineResult UpdateType = "chosen_inline_result"
	UpdateTypeCallbackQuery      UpdateType = "callback_query"
	UpdateTypeShippingQuery      UpdateType = "shipping_query"
	UpdateTypePreCheckoutQuery   UpdateType = "pre_checkout_query"
	UpdateTypePoll               UpdateType = "poll"
	UpdateTypePollAnswer         UpdateType = "poll_answer"
	UpdateTypeMyChatMember       UpdateType = "my_chat_member"
	UpdateTypeChatMember         UpdateType = "chat_member"
	UpdateTypeChatJoinRequest    UpdateType = "chat_join_request"
)

type InlineMessageID string

type ShippingOptionID string

type ShippingQueryID string

type PreCheckoutQueryID string

type PassportElementType string

const (
	PassportElementTypePersonalDetails       PassportElementType = "personal_details"
	PassportElementTypePassport              PassportElementType = "passport"
	PassportElementTypeDriverLicense         PassportElementType = "driver_license"
	PassportElementTypeIdentityCard          PassportElementType = "identity_card"
	PassportElementTypeInternalPassport      PassportElementType = "internal_passport"
	PassportElementTypeAddress               PassportElementType = "address"
	PassportElementTypeUtilityBill           PassportElementType = "utility_bill"
	PassportElementTypeBankStatement         PassportElementType = "bank_statement"
	PassportElementTypeRentalAgreement       PassportElementType = "rental_agreement"
	PassportElementTypePassportRegistration  PassportElementType = "passport_registration"
	PassportElementTypeTemporaryRegistration PassportElementType = "temporary_registration"
	PassportElementTypePhoneNumber           PassportElementType = "phone_number"
	PassportElementTypeEmail                 PassportElementType = "email"
)

type PassportElementErrorSource string

const (
	PassportElementErrorSourceData             PassportElementErrorSource = "data"
	PassportElementErrorSourceFrontSide        PassportElementErrorSource = "front_side"
	PassportElementErrorSourceReverseSide      PassportElementErrorSource = "reverse_side"
	PassportElementErrorSourceSelfie           PassportElementErrorSource = "selfie"
	PassportElementErrorSourceFile             PassportElementErrorSource = "file"
	PassportElementErrorSourceFiles            PassportElementErrorSource = "files"
	PassportElementErrorSourceTranslationFile  PassportElementErrorSource = "translation_file"
	PassportElementErrorSourceTranslationFiles PassportElementErrorSource = "translation_files"
	PassportElementErrorSourceUnspecified      PassportElementErrorSource = "unspecified"
)

type ReplyMarkup interface {
	replyMarkup()
}

func (*InlineKeyboardMarkup) replyMarkup() {}
func (*ReplyKeyboardMarkup) replyMarkup()  {}
func (*ReplyKeyboardRemove) replyMarkup()  {}
func (*ForceReply) replyMarkup()           {}

type DiceEmoji string

const (
	DiceEmojiGameDie     DiceEmoji = "🎲"
	DiceEmojiBullseye    DiceEmoji = "🎯"
	DiceEmojiBasketball  DiceEmoji = "🏀"
	DiceEmojiSoccerBall  DiceEmoji = "⚽"
	DiceEmojiBowling     DiceEmoji = "🎳"
	DiceEmojiSlotMachine DiceEmoji = "🎰"
)

type ChatAction string

const (
	// https://core.telegram.org/bots/api#sendmessage
	ChatActionTyping ChatAction = "typing"
	// https://core.telegram.org/bots/api#sendphoto
	ChatActionUploadPhoto ChatAction = "upload_photo"
	// https://core.telegram.org/bots/api#sendvideo
	ChatActionRecordVideo ChatAction = "record_video"
	// https://core.telegram.org/bots/api#sendvideo
	ChatActionUploadVideo ChatAction = "upload_video"
	// https://core.telegram.org/bots/api#sendvoice
	ChatActionRecordVoice ChatAction = "record_voice"
	// https://core.telegram.org/bots/api#sendvoice
	ChatActionUploadVoice ChatAction = "upload_voice"
	// https://core.telegram.org/bots/api#senddocument
	ChatActionUploadDocument ChatAction = "upload_document"
	// https://core.telegram.org/bots/api#sendsticker
	ChatActionChooseSticker ChatAction = "choose_sticker"
	// https://core.telegram.org/bots/api#sendlocation
	ChatActionFindLocation ChatAction = "find_location"
	// https://core.telegram.org/bots/api#sendvideonote
	ChatActionRecordVideoNote ChatAction = "record_video_note"
	// https://core.telegram.org/bots/api#sendvideonote
	ChatActionUploadVideoNote ChatAction = "upload_video_note"
)

type WebAppQueryID string

type MenuButtonType string

const (
	MenuButtonTypeCommands MenuButtonType = "commands"
	MenuButtonTypeWebApp   MenuButtonType = "web_app"
	MenuButtonTypeDefault  MenuButtonType = "default"
)
