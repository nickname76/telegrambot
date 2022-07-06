package telegrambot

// https://core.telegram.org/bots/api#available-types

import (
	"encoding/hex"
	"io"
	"math/rand"
)

// This object represents a Telegram user or bot.
//
// https://core.telegram.org/bots/api#user
type User struct {
	// Unique identifier for this user or bot.
	ID UserID `json:"id"`
	// True, if this user is a bot
	IsBot bool `json:"is_bot"`
	// User's or bot's first name
	FirstName string `json:"first_name"`
	// Optional. User's or bot's last name
	LastName string `json:"last_name,omitempty"`
	// Optional. User's or bot's username
	Username Username `json:"username,omitempty"`
	// Optional. IETF language tag of the user's language
	// https://en.wikipedia.org/wiki/IETF_language_tag
	LanguageCode LanguageCode `json:"language_code,omitempty"`
	// Optional. True, if this user is a Telegram Premium user
	IsPremium bool `json:"is_premium,omitempty"`
	// Optional. True, if this user added the bot to the attachment menu
	AddedToAttachmentMenu bool `json:"added_to_attachment_menu,omitempty"`
	// Optional. True, if the bot can be invited to groups. Returned only in
	// getMe. https://core.telegram.org/bots/api#getme
	CanJoinGroups bool `json:"can_join_groups,omitempty"`
	// Optional. True, if privacy mode is disabled for the bot. Returned only in
	// getMe. https://core.telegram.org/bots#privacy-mode
	// https://core.telegram.org/bots/api#getme
	CanReadAllGroupMessages bool `json:"can_read_all_group_messages,omitempty"`
	// Optional. True, if the bot supports inline queries. Returned only in
	// getMe. https://core.telegram.org/bots/api#getme
	SupportsInlineQueries bool `json:"supports_inline_queries,omitempty"`
}

// This object represents a chat.
//
// https://core.telegram.org/bots/api#chat
type Chat struct {
	// Unique identifier for this chat.
	ID ChatID `json:"id"`
	// Type of chat, can be either “private”, “group”, “supergroup” or “channel”
	Type ChatType `json:"type"`
	// Optional. Title, for supergroups, channels and group chats
	Title string `json:"title,omitempty"`
	// Optional. Username, for private chats, supergroups and channels if
	// available
	Username Username `json:"username,omitempty"`
	// Optional. First name of the other party in a private chat
	FirstName string `json:"first_name,omitempty"`
	// Optional. Last name of the other party in a private chat
	LastName string `json:"last_name,omitempty"`
	// Optional. Chat photo. Returned only in getChat.
	// https://core.telegram.org/bots/api#getchat
	Photo *ChatPhoto `json:"photo,omitempty"`
	// Optional. Bio of the other party in a private chat. Returned only in
	// getChat. https://core.telegram.org/bots/api#getchat
	Bio string `json:"bio,omitempty"`
	// Optional. True, if privacy settings of the other party in the private
	// chat allows to use tg://user?id=<user_id> links only in chats with the
	// user. Returned only in getChat.
	// https://core.telegram.org/bots/api#getchat
	HasPrivateForwards string `json:"has_private_forwards,omitempty"`
	// Optional. True, if users need to join the supergroup before they can send
	// messages. Returned only in getChat.
	// https://core.telegram.org/bots/api#getchat
	JoinToSendMessages bool `json:"join_to_send_messages,omitempty"`
	// Optional. True, if all users directly joining the supergroup need to be
	// approved by supergroup administrators. Returned only in getChat.
	// https://core.telegram.org/bots/api#getchat
	JoinByRequest bool `json:"join_by_request,omitempty"`
	// Optional. Description, for groups, supergroups and channel chats.
	// Returned only in getChat. https://core.telegram.org/bots/api#getchat
	Description string `json:"description,omitempty"`
	// Optional. Primary invite link, for groups, supergroups and channel chats.
	// Returned only in getChat. https://core.telegram.org/bots/api#getchat
	InviteLink string `json:"invite_link,omitempty"`
	// Optional. The most recent pinned message (by sending date). Returned only
	// in getChat. https://core.telegram.org/bots/api#getchat
	PinnedMessage *Message `json:"pinned_message,omitempty"`
	// Optional. Default chat member permissions, for groups and supergroups.
	// Returned only in getChat. https://core.telegram.org/bots/api#getchat
	Permissions *ChatPermissions `json:"permissions,omitempty"`
	// Optional. For supergroups, the minimum allowed delay between consecutive
	// messages sent by each unpriviledged user; in seconds. Returned only in
	// getChat. https://core.telegram.org/bots/api#getchat
	SlowModeDelay int `json:"slow_mode_delay,omitempty"`
	// Optional. The time after which all messages sent to the chat will be
	// automatically deleted; in seconds. Returned only in getChat.
	// https://core.telegram.org/bots/api#getchat
	MessageAutoDeleteTime int `json:"message_auto_delete_time,omitempty"`
	// Optional. True, if messages from the chat can't be forwarded to other
	// chats. Returned only in getChat.
	// https://core.telegram.org/bots/api#getchat
	HasProtectedContent bool `json:"has_protected_content,omitempty"`
	// Optional. For supergroups, name of group sticker set. Returned only in
	// getChat. https://core.telegram.org/bots/api#getchat
	StickerSetName string `json:"sticker_set_name,omitempty"`
	// Optional. True, if the bot can change the group sticker set. Returned
	// only in getChat. https://core.telegram.org/bots/api#getchat
	CanSetStickerSet bool `json:"can_set_sticker_set,omitempty"`
	// Optional. Unique identifier for the linked chat, i.e. the discussion
	// group identifier for a channel and vice versa; for supergroups and
	// channel chats. Returned only in getChat.
	// https://core.telegram.org/bots/api#getchat
	LinkedChatID ChatID `json:"linked_chat_id,omitempty"`
	// Optional. For supergroups, the location to which the supergroup is
	// connected. Returned only in getChat.
	// https://core.telegram.org/bots/api#getchat
	Location *ChatLocation `json:"location,omitempty"`
}

// This object represents a message.
//
// https://core.telegram.org/bots/api#message
type Message struct {
	// Unique message identifier inside this chat
	MessageID MessageID `json:"message_id"`
	// Optional. Sender of the message; empty for messages sent to channels. For
	// backward compatibility, the field contains a fake sender user in
	// non-channel chats, if the message was sent on behalf of a chat.
	From *User `json:"from,omitempty"`
	// Optional. Sender of the message, sent on behalf of a chat. For example,
	// the channel itself for channel posts, the supergroup itself for messages
	// from anonymous group administrators, the linked channel for messages
	// automatically forwarded to the discussion group. For backward
	// compatibility, the field from contains a fake sender user in non-channel
	// chats, if the message was sent on behalf of a chat.
	SenderChat *Chat `json:"sender_chat,omitempty"`
	// Date the message was sent in Unix time
	Date int64 `json:"date"`
	// Conversation the message belongs to
	Chat *Chat `json:"chat"`
	// Optional. For forwarded messages, sender of the original message
	ForwardFrom *User `json:"forward_from,omitempty"`
	// Optional. For messages forwarded from channels or from anonymous
	// administrators, information about the original sender chat
	ForwardFromChat *Chat `json:"forward_from_chat,omitempty"`
	// Optional. For messages forwarded from channels, identifier of the
	// original message in the channel
	ForwardFromMessageID MessageID `json:"forward_from_message_id,omitempty"`
	// Optional. For forwarded messages that were originally sent in channels or
	// by an anonymous chat administrator, signature of the message sender if
	// present
	ForwardSignature string `json:"forward_signature,omitempty"`
	// Optional. Sender's name for messages forwarded from users who disallow
	// adding a link to their account in forwarded messages
	ForwardSenderName string `json:"forward_sender_name,omitempty"`
	// Optional. For forwarded messages, date the original message was sent in
	// Unix time
	ForwardDate int64 `json:"forward_date,omitempty"`
	// Optional. True, if the message is a channel post that was automatically
	// forwarded to the connected discussion group
	IsAutomaticForward bool `json:"is_automatic_forward,omitempty"`
	// Optional. For replies, the original message. Note that the Message object
	// in this field will not contain further reply_to_message fields even if it
	// itself is a reply.
	ReplyToMessage *Message `json:"reply_to_message,omitempty"`
	// Optional. Bot through which the message was sent
	ViaBot *User `json:"via_bot,omitempty"`
	// Optional. Date the message was last edited in Unix time
	EditDate int64 `json:"edit_date,omitempty"`
	// Optional. True, if the message can't be forwarded
	HasProtectedContent bool `json:"has_protected_content,omitempty"`
	// Optional. The unique identifier of a media message group this message
	// belongs to
	MediaGroupID string `json:"media_group_id,omitempty"`
	// Optional. Signature of the post author for messages in channels, or the
	// custom title of an anonymous group administrator
	AuthorSignature string `json:"author_signature,omitempty"`
	// Optional. For text messages, the actual UTF-8 text of the message, 0-4096
	// characters
	Text string `json:"text,omitempty"`
	// Optional. For text messages, special entities like usernames, URLs, bot
	// commands, etc. that appear in the text
	Entities []*MessageEntity `json:"entities,omitempty"`
	// Optional. Message is an animation, information about the animation. For
	// backward compatibility, when this field is set, the document field will
	// also be set
	Animation *Animation `json:"animation,omitempty"`
	// Optional. Message is an audio file, information about the file
	Audio *Audio `json:"audio,omitempty"`
	// Optional. Message is a general file, information about the file
	Document *Document `json:"document,omitempty"`
	// Optional. Message is a photo, available sizes of the photo
	Photo []*PhotoSize `json:"photo,omitempty"`
	// Optional. Message is a sticker, information about the sticker
	Sticker *Sticker `json:"sticker,omitempty"`
	// Optional. Message is a video, information about the video
	Video *Video `json:"video,omitempty"`
	// Optional. Message is a video note, information about the video message
	// https://telegram.org/blog/video-messages-and-telescope
	VideoNote *VideoNote `json:"video_note,omitempty"`
	// Optional. Message is a voice message, information about the file
	Voice *Voice `json:"voice,omitempty"`
	// Optional. Caption for the animation, audio, document, photo, video or
	// voice, 0-1024 characters
	Caption string `json:"caption,omitempty"`
	// Optional. For messages with a caption, special entities like usernames,
	// URLs, bot commands, etc. that appear in the caption
	CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"`
	// Optional. Message is a shared contact, information about the contact
	Contact *Contact `json:"contact,omitempty"`
	// Optional. Message is a dice with random value
	Dice *Dice `json:"dice,omitempty"`
	// Optional. Message is a game, information about the game.
	// https://core.telegram.org/bots/api#games
	Game *Game `json:"game,omitempty"`
	// Optional. Message is a native poll, information about the poll
	Poll *Poll `json:"poll,omitempty"`
	// Optional. Message is a venue, information about the venue. For backward
	// compatibility, when this field is set, the location field will also be
	// set
	Venue *Venue `json:"venue,omitempty"`
	// Optional. Message is a shared location, information about the location
	Location *Location `json:"location,omitempty"`
	// Optional. New members that were added to the group or supergroup and
	// information about them (the bot itself may be one of these members)
	NewChatMembers []*User `json:"new_chat_members,omitempty"`
	// Optional. A member was removed from the group, information about them
	// (this member may be the bot itself)
	LeftChatMember *User `json:"left_chat_member,omitempty"`
	// Optional. A chat title was changed to this value
	NewChatTitle string `json:"new_chat_title,omitempty"`
	// Optional. A chat photo was change to this value
	NewChatPhoto []*PhotoSize `json:"new_chat_photo,omitempty"`
	// Optional. Service message: the chat photo was deleted
	DeleteChatPhoto bool `json:"delete_chat_photo,omitempty"`
	// Optional. Service message: the group has been created
	GroupChatCreated bool `json:"group_chat_created,omitempty"`
	// Optional. Service message: the supergroup has been created. This field
	// can't be received in a message coming through updates, because bot can't
	// be a member of a supergroup when it is created. It can only be found in
	// reply_to_message if someone replies to a very first message in a directly
	// created supergroup.
	SupergroupChatCreated bool `json:"supergroup_chat_created,omitempty"`
	// Optional. Service message: the channel has been created. This field can't
	// be received in a message coming through updates, because bot can't be a
	// member of a channel when it is created. It can only be found in
	// reply_to_message if someone replies to a very first message in a channel.
	ChannelChatCreated bool `json:"channel_chat_created,omitempty"`
	// Optional. Service message: auto-delete timer settings changed in the chat
	MessageAutoDeleteTimerChanged *MessageAutoDeleteTimerChanged `json:"message_auto_delete_timer_changed,omitempty"`
	// Optional. The group has been migrated to a supergroup with the specified
	// identifier. This number may have more than 32 significant bits and some
	// programming languages may have difficulty/silent defects in interpreting
	// it. But it has at most 52 significant bits, so a signed 64-bit integer or
	// double-precision float type are safe for storing this identifier.
	MigrateToChatID ChatID `json:"migrate_to_chat_id,omitempty"`
	// Optional. The supergroup has been migrated from a group with the
	// specified identifier. This number may have more than 32 significant bits
	// and some programming languages may have difficulty/silent defects in
	// interpreting it. But it has at most 52 significant bits, so a signed
	// 64-bit integer or double-precision float type are safe for storing this
	// identifier.
	MigrateFromChatID ChatID `json:"migrate_from_chat_id,omitempty"`
	// Optional. Specified message was pinned. Note that the Message object in
	// this field will not contain further reply_to_message fields even if it is
	// itself a reply.
	PinnedMessage *Message `json:"pinned_message,omitempty"`
	// Optional. Message is an invoice for a payment, information about the
	// invoice. https://core.telegram.org/bots/api#payments
	Invoice *Invoice `json:"invoice,omitempty"`
	// Optional. Message is a service message about a successful payment,
	// information about the payment.
	// https://core.telegram.org/bots/api#payments
	SuccessfulPayment *SuccessfulPayment `json:"successful_payment,omitempty"`
	// Optional. The domain name of the website on which the user has logged in.
	// https://core.telegram.org/widgets/login
	ConnectedWebsite string `json:"connected_website,omitempty"`
	// Optional. Telegram Passport data
	PassportData *PassportData `json:"passport_data,omitempty"`
	// Optional. Service message. A user in the chat triggered another user's
	// proximity alert while sharing Live Location.
	ProximityAlertTriggered *ProximityAlertTriggered `json:"proximity_alert_triggered,omitempty"`
	// Optional. Service message: video chat scheduled
	VideoChatScheduled *VideoChatScheduled `json:"video_chat_scheduled,omitempty"`
	// Optional. Service message: video chat started
	VideoChatStarted *VideoChatStarted `json:"video_chat_started,omitempty"`
	// Optional. Service message: video chat ended
	VideoChatEnded *VideoChatEnded `json:"video_chat_ended,omitempty"`
	// Optional. Service message: new participants invited to a video chat
	VideoChatParticipantsInvited *VideoChatParticipantsInvited `json:"video_chat_participants_invited,omitempty"`
	// Optional. Service message: data sent by a Web App
	WebAppData *WebAppData `json:"web_app_data,omitempty"`
	// Optional. Inline keyboard attached to the message. `login_url` buttons
	// are represented as ordinary `url` buttons.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// This object represents a unique message identifier.
//
// https://core.telegram.org/bots/api#messageid
type MessageIDObject struct {
	// Unique message identifier
	MessageID MessageID `json:"message_id"`
}

// This object represents one special entity in a text message. For example,
// hashtags, usernames, URLs, etc.
//
// https://core.telegram.org/bots/api#messageentity
type MessageEntity struct {
	// Type of the entity. Currently, can be “mention” (@username), “hashtag”
	// (#hashtag), “cashtag” ($USD), “bot_command” (/start@jobs_bot), “url”
	// (https://telegram.org), “email” (do-not-reply@telegram.org),
	// “phone_number” (+1-212-555-0123), “bold” (bold text), “italic” (italic
	// text), “underline” (underlined text), “strikethrough” (strikethrough
	// text), “spoiler” (spoiler message), “code” (monowidth string), “pre”
	// (monowidth block), “text_link” (for clickable text URLs), “text_mention”
	// (for users without usernames) https://telegram.org/blog/edit#new-mentions
	Type MessageEntityType `json:"type"`
	// Offset in UTF-16 code units to the start of the entity
	Offset int `json:"offset"`
	// Length of the entity in UTF-16 code units
	Length int `json:"length"`
	// Optional. For “text_link” only, url that will be opened after user taps
	// on the text
	URL string `json:"url,omitempty"`
	// Optional. For “text_mention” only, the mentioned user
	User *User `json:"user,omitempty"`
	// Optional. For “pre” only, the programming language of the entity text
	Language string `json:"language,omitempty"`
}

// This object represents one size of a photo or a file / sticker thumbnail.
// https://core.telegram.org/bots/api#document
// https://core.telegram.org/bots/api#sticker
//
// https://core.telegram.org/bots/api#photosize
type PhotoSize struct {
	// Identifier for this file, which can be used to download or reuse the file
	FileID FileID `json:"file_id"`
	// Unique identifier for this file, which is supposed to be the same over
	// time and for different bots. Can't be used to download or reuse the file.
	FileUniqueID FileUniqueID `json:"file_unique_id"`
	// Photo width
	Width int `json:"width"`
	// Photo height
	Height int `json:"height"`
	// Optional. File size in bytes
	FileSize int64 `json:"file_size,omitempty"`
}

// This object represents an animation file (GIF or H.264/MPEG-4 AVC video
// without sound).
//
// https://core.telegram.org/bots/api#animation
type Animation struct {
	// Identifier for this file, which can be used to download or reuse the file
	FileID FileID `json:"file_id"`
	// Unique identifier for this file, which is supposed to be the same over
	// time and for different bots. Can't be used to download or reuse the file.
	FileUniqueID FileUniqueID `json:"file_unique_id"`
	// Video width as defined by sender
	Width int `json:"width"`
	// Video height as defined by sender
	Height int `json:"height"`
	// Duration of the video in seconds as defined by sender
	Duration int `json:"duration"`
	// Optional. Animation thumbnail as defined by sender
	Thumb *PhotoSize `json:"thumb,omitempty"`
	// Optional. Original animation filename as defined by sender
	FileName string `json:"file_name,omitempty"`
	// Optional. MIME type of the file as defined by sender
	MimeType string `json:"mime_type,omitempty"`
	// Optional. File size in bytes
	FileSize int64 `json:"file_size,omitempty"`
}

// This object represents an audio file to be treated as music by the Telegram
// clients. https://core.telegram.org/bots/api#photosize
// https://core.telegram.org/bots/api#voice
// https://core.telegram.org/bots/api#audio
//
// https://core.telegram.org/bots/api#audio
type Audio struct {
	// Identifier for this file, which can be used to download or reuse the file
	FileID FileID `json:"file_id"`
	// Unique identifier for this file, which is supposed to be the same over
	// time and for different bots. Can't be used to download or reuse the file.
	FileUniqueID FileUniqueID `json:"file_unique_id"`
	// Duration of the audio in seconds as defined by sender
	Duration int `json:"duration"`
	// Optional. Performer of the audio as defined by sender or by audio tags
	Performer string `json:"performer,omitempty"`
	// Optional. Title of the audio as defined by sender or by audio tags
	Title string `json:"title,omitempty"`
	// Optional. Original filename as defined by sender
	FileName string `json:"file_name,omitempty"`
	// Optional. MIME type of the file as defined by sender
	MimeType string `json:"mime_type,omitempty"`
	// Optional. File size in bytes
	FileSize int64 `json:"file_size,omitempty"`
	// Optional. Thumbnail of the album cover to which the music file belongs
	Thumb *PhotoSize `json:"thumb,omitempty"`
}

// This object represents a general file (as opposed to photos, voice messages
// and audio files).
//
// https://core.telegram.org/bots/api#document
type Document struct {
	// Identifier for this file, which can be used to download or reuse the file
	FileID FileID `json:"file_id"`
	// Unique identifier for this file, which is supposed to be the same over
	// time and for different bots. Can't be used to download or reuse the file.
	FileUniqueID FileUniqueID `json:"file_unique_id"`
	// Optional. Document thumbnail as defined by sender
	Thumb *PhotoSize `json:"thumb,omitempty"`
	// Optional. Original filename as defined by sender
	FileName string `json:"file_name,omitempty"`
	// Optional. MIME type of the file as defined by sender
	MimeType string `json:"mime_type,omitempty"`
	// Optional. File size in bytes
	FileSize int64 `json:"file_size,omitempty"`
}

// This object represents a video file.
//
// https://core.telegram.org/bots/api#video
type Video struct {
	// Identifier for this file, which can be used to download or reuse the file
	FileID FileID `json:"file_id"`
	// Unique identifier for this file, which is supposed to be the same over
	// time and for different bots. Can't be used to download or reuse the file.
	FileUniqueID FileUniqueID `json:"file_unique_id"`
	// Video width as defined by sender
	Width int `json:"width"`
	// Video height as defined by sender
	Height int `json:"height"`
	// Duration of the video in seconds as defined by sender
	Duration int `json:"duration"`
	// Optional. Video thumbnail
	Thumb *PhotoSize `json:"thumb,omitempty"`
	// Optional. Original filename as defined by sender
	FileName string `json:"file_name,omitempty"`
	// Optional. Mime type of a file as defined by sender
	MimeType string `json:"mime_type,omitempty"`
	// Optional. File size in bytes
	FileSize int64 `json:"file_size,omitempty"`
}

// This object represents a video message (available in Telegram apps as of
// v.4.0). https://telegram.org/blog/video-messages-and-telescope
//
// https://core.telegram.org/bots/api#videonote
type VideoNote struct {
	// Identifier for this file, which can be used to download or reuse the file
	FileID FileID `json:"file_id"`
	// Unique identifier for this file, which is supposed to be the same over
	// time and for different bots. Can't be used to download or reuse the file.
	FileUniqueID FileUniqueID `json:"file_unique_id"`
	// Video width and height (diameter of the video message) as defined by
	// sender
	Length int `json:"length"`
	// Duration of the video in seconds as defined by sender
	Duration int `json:"duration"`
	// Optional. Video thumbnail
	Thumb *PhotoSize `json:"thumb,omitempty"`
	// Optional. File size in bytes
	FileSize int64 `json:"file_size,omitempty"`
}

// This object represents a voice note.
//
// https://core.telegram.org/bots/api#voice
type Voice struct {
	// Identifier for this file, which can be used to download or reuse the file
	FileID FileID `json:"file_id"`
	// Unique identifier for this file, which is supposed to be the same over
	// time and for different bots. Can't be used to download or reuse the file.
	FileUniqueID FileUniqueID `json:"file_unique_id"`
	// Duration of the audio in seconds as defined by sender
	Duration int `json:"duration"`
	// Optional. MIME type of the file as defined by sender
	MimeType string `json:"mime_type,omitempty"`
	// Optional. File size in bytes
	FileSize int64 `json:"file_size,omitempty"`
}

// This object represents a phone contact.
//
// https://core.telegram.org/bots/api#contact
type Contact struct {
	// Contact's phone number
	PhoneNumber string `json:"phone_number"`
	// Contact's first name
	FirstName string `json:"first_name"`
	// Optional. Contact's last name
	LastName string `json:"last_name,omitempty"`
	// Optional. Contact's user identifier in Telegram.
	UserID UserID `json:"user_id,omitempty"`
	// Optional. Additional data about the contact in the form of a vCard
	// https://en.wikipedia.org/wiki/VCard
	VCard string `json:"vcard,omitempty"`
}

// This object represents an animated emoji that displays a random value.
//
// https://core.telegram.org/bots/api#dice
type Dice struct {
	// Emoji on which the dice throw animation is based
	Emoji string `json:"emoji"`
	// Value of the dice, 1-6 for “🎲”, “🎯” and “🎳” base emoji, 1-5 for “🏀”
	// and “⚽” base emoji, 1-64 for “🎰” base emoji
	Value int `json:"value"`
}

// This object contains information about one answer option in a poll.
//
// https://core.telegram.org/bots/api#polloption
type PollOption struct {
	// Option text, 1-100 characters
	Text string `json:"text"`
	// Number of users that voted for this option
	VoterCount int `json:"voter_count"`
}

// This object represents an answer of a user in a non-anonymous poll.
//
// https://core.telegram.org/bots/api#pollanswer
type PollAnswer struct {
	// Unique poll identifier
	PollID PollID `json:"poll_id"`
	// The user, who changed the answer to the poll
	User *User `json:"user"`
	// 0-based identifiers of answer options, chosen by the user. May be empty
	// if the user retracted their vote.
	OptionIDs []int `json:"option_ids"`
}

// This object contains information about a poll.
//
// https://core.telegram.org/bots/api#poll
type Poll struct {
	// Unique poll identifier
	ID PollID `json:"id"`
	// Poll question, 1-300 characters
	Question string `json:"question"`
	// List of poll options
	Options []*PollOption `json:"options"`
	// Total number of users that voted in the poll
	TotalVoterCount int `json:"total_voter_count"`
	// True, if the poll is closed
	IsClosed bool `json:"is_closed"`
	// True, if the poll is anonymous
	IsAnonymous bool `json:"is_anonymous"`
	// Poll type, currently can be “regular” or “quiz”
	Type PollType `json:"type"`
	// True, if the poll allows multiple answers
	AllowsMultipleAnswers bool `json:"allows_multiple_answers"`
	// Optional. 0-based identifier of the correct answer option. Available only
	// for polls in the quiz mode, which are closed, or was sent (not forwarded)
	// by the bot or to the private chat with the bot.
	CorrectOptionID int `json:"correct_option_id,omitempty"`
	// Optional. Text that is shown when a user chooses an incorrect answer or
	// taps on the lamp icon in a quiz-style poll, 0-200 characters
	Explanation string `json:"explanation,omitempty"`
	// Optional. Special entities like usernames, URLs, bot commands, etc. that
	// appear in the explanation
	ExplanationEntities []*MessageEntity `json:"explanation_entities,omitempty"`
	// Optional. Amount of time in seconds the poll will be active after
	// creation
	OpenPeriod int `json:"open_period,omitempty"`
	// Optional. Point in time (Unix timestamp) when the poll will be
	// automatically closed
	CloseDate int64 `json:"close_date,omitempty"`
}

// This object represents a point on the map.
//
// https://core.telegram.org/bots/api#location
type Location struct {
	// Longitude as defined by sender
	Longitude float64 `json:"longitude"`
	// Latitude as defined by sender
	Latitude float64 `json:"latitude"`
	// Optional. The radius of uncertainty for the location, measured in meters;
	// 0-1500
	HorizontalAccuracy float64 `json:"horizontal_accuracy,omitempty"`
	// Optional. Time relative to the message sending date, during which the
	// location can be updated; in seconds. For active live locations only.
	LivePeriod int `json:"live_period,omitempty"`
	// Optional. The direction in which user is moving, in degrees; 1-360. For
	// active live locations only.
	Heading int `json:"heading,omitempty"`
	// Optional. Maximum distance for proximity alerts about approaching another
	// chat member, in meters. For sent live locations only.
	ProximityAlertRadius int `json:"proximity_alert_radius,omitempty"`
}

// This object represents a venue.
//
// https://core.telegram.org/bots/api#venue
type Venue struct {
	// Venue location. Can't be a live location
	Location *Location `json:"location"`
	// Name of the venue
	Title string `json:"title"`
	// Address of the venue
	Address string `json:"address"`
	// Optional. Foursquare identifier of the venue
	Foursquare_id string `json:"foursquare_id,omitempty"`
	// Optional. Foursquare type of the venue. (For example,
	// “arts_entertainment/default”, “arts_entertainment/aquarium” or
	// “food/icecream”.)
	Foursquare_type string `json:"foursquare_type,omitempty"`
	// Optional. Google Places identifier of the venue
	GooglePlaceID string `json:"google_place_id,omitempty"`
	// Optional. Google Places type of the venue. (See supported types.)
	// https://developers.google.com/maps/documentation/places/web-service/supported_types
	GooglePlaceType string `json:"google_place_type,omitempty"`
}

// Contains data sent from a Web App to the bot.
// https://core.telegram.org/bots/webapps
//
// https://core.telegram.org/bots/api#webappdata
type WebAppData struct {
	// The data. Be aware that a bad client can send arbitrary data in this
	// field.
	Data string `json:"data"`
	// Text of the web_app keyboard button, from which the Web App was opened.
	// Be aware that a bad client can send arbitrary data in this field.
	ButtonText string `json:"button_text"`
}

// This object represents the content of a service message, sent whenever a user
// in the chat triggers a proximity alert set by another user.
//
// https://core.telegram.org/bots/api#proximityalerttriggered
type ProximityAlertTriggered struct {
	// User that triggered the alert
	Traveler *User `json:"traveler"`
	// User that set the alert
	Watcher *User `json:"watcher"`
	// The distance between the users
	Distance int `json:"distance"`
}

// This object represents a service message about a change in auto-delete timer
// settings.
//
// https://core.telegram.org/bots/api#messageautodeletetimerchanged
type MessageAutoDeleteTimerChanged struct {
	// New auto-delete time for messages in the chat; in seconds
	MessageAutoDeleteTime int `json:"message_auto_delete_time"`
}

// This object represents a service message about a video chat scheduled in the
// chat.
//
// https://core.telegram.org/bots/api#videochatscheduled
type VideoChatScheduled struct {
	// Point in time (Unix timestamp) when the video chat is supposed to be
	// started by a chat administrator
	StartDate int64 `json:"start_date"`
}

// This object represents a service message about a voice chat started in the
// chat. Currently holds no information.
//
// https://core.telegram.org/bots/api#videochatstarted
type VideoChatStarted struct{}

// This object represents a service message about a video chat ended in the
// chat.
//
// https://core.telegram.org/bots/api#videochatended
type VideoChatEnded struct {
	// Video chat duration in seconds
	Duration int `json:"duration"`
}

// This object represents a service message about new members invited to a video
// chat.
//
// https://core.telegram.org/bots/api#videochatparticipantsinvited
type VideoChatParticipantsInvited struct {
	// New members that were invited to the video chat
	Users []*User `json:"users,omitempty"`
}

// This object represent a user's profile pictures.
//
// https://core.telegram.org/bots/api#userprofilephotos
type UserProfilePhotos struct {
	// Total number of profile pictures the target user has
	TotalCount int `json:"total_count"`
	// Requested profile pictures (in up to 4 sizes each)
	Photos [][]*PhotoSize `json:"photos"`
}

// This object represents a file ready to be downloaded. The file can be
// downloaded via the link https://api.telegram.org/file/bot<token>/<file_path>.
// It is guaranteed that the link will be valid for at least 1 hour. When the
// link expires, a new one can be requested by calling getFile.
// https://core.telegram.org/bots/api#getfile
//
// The maximum file size to download is 20 MB
//
// https://core.telegram.org/bots/api#file
type File struct {
	// Identifier for this file, which can be used to download or reuse the file
	FileID FileID `json:"file_id"`
	// Unique identifier for this file, which is supposed to be the same over
	// time and for different bots. Can't be used to download or reuse the file.
	FileUniqueID FileUniqueID `json:"file_unique_id"`
	// Optional. File size in bytes, if known
	FileSize int64 `json:"file_size,omitempty"`
	// Optional. File path. Use
	// `https://api.telegram.org/file/bot<token>/<file_path>` to get the file.
	FilePath string `json:"file_path,omitempty"`
}

// Contains information about a Web App. https://core.telegram.org/bots/webapps
//
// https://core.telegram.org/bots/api#webappinfo
type WebAppInfo struct {
	// An HTTPS URL of a Web App to be opened with additional data as specified
	// in Initializing Web Apps
	// https://core.telegram.org/bots/webapps#initializing-web-apps
	URL string `json:"url"`
}

// This object represents a custom keyboard with reply options (see Introduction
// to bots for details and examples). https://core.telegram.org/bots#keyboards
//
// https://core.telegram.org/bots/api#replykeyboardmarkup
type ReplyKeyboardMarkup struct {
	// Array of button rows, each represented by an Array of KeyboardButton
	// objects
	Keyboard [][]*KeyboardButton `json:"keyboard"`
	// Optional. Requests clients to resize the keyboard vertically for optimal
	// fit (e.g., make the keyboard smaller if there are just two rows of
	// buttons). Defaults to false, in which case the custom keyboard is always
	// of the same height as the app's standard keyboard.
	ResizeKeyboard bool `json:"resize_keyboard,omitempty"`
	// Optional. Requests clients to hide the keyboard as soon as it's been
	// used. The keyboard will still be available, but clients will
	// automatically display the usual letter-keyboard in the chat – the user
	// can press a special button in the input field to see the custom keyboard
	// again. Defaults to false.
	OneTimeKeyboard bool `json:"one_time_keyboard,omitempty"`
	// Optional. The placeholder to be shown in the input field when the
	// keyboard is active; 1-64 characters
	InputFieldPlaceholder string `json:"input_field_placeholder,omitempty"`
	// Optional. Use this parameter if you want to show the keyboard to specific
	// users only. Targets: 1) users that are @mentioned in the text of the
	// Message object; 2) if the bot's message is a reply (has
	// reply_to_message_id), sender of the original message.
	//
	// Example: A user requests to change the bot's language, bot replies to the
	// request with a keyboard to select the new language. Other users in the
	// group don't see the keyboard.
	Selective bool `json:"selective,omitempty"`
}

// This object represents one button of the reply keyboard. For simple text
// buttons String can be used instead of this object to specify text of the
// button. Optional fields web_app, request_contact, request_location, and
// request_poll are mutually exclusive.
//
//   Note: request_contact and request_location options will only work in Telegram versions released after 9 April, 2016. Older clients will display unsupported message.
//   Note: request_poll option will only work in Telegram versions released after 23 January, 2020. Older clients will display unsupported message.
//   Note: web_app option will only work in Telegram versions released after 16 April, 2022. Older clients will display unsupported message.
//
// https://core.telegram.org/bots/api#keyboardbutton
type KeyboardButton struct {
	// Text of the button. If none of the optional fields are used, it will be
	// sent as a message when the button is pressed
	Text string `json:"text"`
	// Optional. If True, the user's phone number will be sent as a contact when
	// the button is pressed. Available in private chats only
	RequestContact bool `json:"request_contact,omitempty"`
	// Optional. If True, the user's current location will be sent when the
	// button is pressed. Available in private chats only
	RequestLocation bool `json:"request_location,omitempty"`
	// Optional. If specified, the user will be asked to create a poll and send
	// it to the bot when the button is pressed. Available in private chats only
	RequestPoll *KeyboardButtonPollType `json:"request_poll,omitempty"`
	// Optional. If specified, the described Web App will be launched when the
	// button is pressed. The Web App will be able to send a “web_app_data”
	// service message. Available in private chats only.
	// https://core.telegram.org/bots/webapps
	WebApp *WebAppInfo `json:"web_app,omitempty"`
}

// This object represents type of a poll, which is allowed to be created and
// sent when the corresponding button is pressed.
//
// https://core.telegram.org/bots/api#keyboardbuttonpolltype
type KeyboardButtonPollType struct {
	// Optional. If quiz is passed, the user will be allowed to create only
	// polls in the quiz mode. If regular is passed, only regular polls will be
	// allowed. Otherwise, the user will be allowed to create a poll of any
	// type.
	Type PollType `json:"type,omitempty"`
}

// Upon receiving a message with this object, Telegram clients will remove the
// current custom keyboard and display the default letter-keyboard. By default,
// custom keyboards are displayed until a new keyboard is sent by a bot. An
// exception is made for one-time keyboards that are hidden immediately after
// the user presses a button (see ReplyKeyboardMarkup).
// https://core.telegram.org/bots/api#replykeyboardmarkup
//
// https://core.telegram.org/bots/api#replykeyboardremove
type ReplyKeyboardRemove struct {
	// Requests clients to remove the custom keyboard (user will not be able to
	// summon this keyboard; if you want to hide the keyboard from sight but
	// keep it accessible, use one_time_keyboard in ReplyKeyboardMarkup)
	RemoveKeyboard bool `json:"remove_keyboard"`
	// Optional. Use this parameter if you want to remove the keyboard for
	// specific users only. Targets: 1) users that are @mentioned in the text of
	// the Message object; 2) if the bot's message is a reply (has
	// reply_to_message_id), sender of the original message.
	//
	// Example: A user votes in a poll, bot returns confirmation message in
	// reply to the vote and removes the keyboard for that user, while still
	// showing the keyboard with poll options to users who haven't voted yet.
	Selective bool `json:"selective,omitempty"`
}

// This object represents an inline keyboard that appears right next to the
// message it belongs to.
// https://core.telegram.org/bots#inline-keyboards-and-on-the-fly-updating
//
// Note: This will only work in Telegram versions released after 9 April, 2016.
// Older clients will display unsupported message.
//
// https://core.telegram.org/bots/api#inlinekeyboardmarkup
type InlineKeyboardMarkup struct {
	// Array of button rows, each represented by an Array of
	// InlineKeyboardButton objects
	// https://core.telegram.org/bots/api#inlinekeyboardbutton
	InlineKeyboard [][]*InlineKeyboardButton `json:"inline_keyboard"`
}

// This object represents one button of an inline keyboard. You *must* use
// exactly one of the optional fields.
//
// https://core.telegram.org/bots/api#inlinekeyboardbutton
type InlineKeyboardButton struct {
	// Label text on the button
	Text string `json:"text"`
	// Optional. HTTP or tg:// url to be opened when the button is pressed.
	// Links `tg://user?id=<user_id>` can be used to mention a user by their ID
	// without using a username, if this is allowed by their privacy settings.
	URL string `json:"url,omitempty"`
	// Optional. Data to be sent in a callback query to the bot when button is
	// pressed, 1-64 bytes https://core.telegram.org/bots/api#callbackquery
	CallbackData string `json:"callback_data,omitempty"`
	// Optional. Description of the Web App that will be launched when the user
	// presses the button. The Web App will be able to send an arbitrary message
	// on behalf of the user using the method answerWebAppQuery. Available only
	// in private chats between a user and the bot.
	// https://core.telegram.org/bots/webapps
	// https://core.telegram.org/bots/api#answerwebappquery
	WebApp *WebAppInfo `json:"web_app,omitempty"`
	// Optional. An HTTP URL used to automatically authorize the user. Can be
	// used as a replacement for the Telegram Login Widget.
	// https://core.telegram.org/widgets/login
	LoginURL *LoginURL `json:"login_url,omitempty"`
	// Optional. If set, pressing the button will prompt the user to select one
	// of their chats, open that chat and insert the bot's username and the
	// specified inline query in the input field. Can be empty, in which case
	// just the bot's username will be inserted.
	//
	// Note: This offers an easy way for users to start using your bot in inline
	// mode when they are currently in a private chat with it. Especially useful
	// when combined with switch_pm... actions – in this case the user will be
	// automatically returned to the chat they switched from, skipping the chat
	// selection screen.
	//
	// https://core.telegram.org/bots/inline
	// https://core.telegram.org/bots/api#answerinlinequery
	SwitchInlineQuery string `json:"switch_inline_query,omitempty"`
	// Optional. If set, pressing the button will insert the bot's username and
	// the specified inline query in the current chat's input field. Can be
	// empty, in which case only the bot's username will be inserted. This
	// offers a quick way for the user to open your bot in inline mode in the
	// same chat – good for selecting something from multiple options.
	SwitchInlineQueryCurrentChat string `json:"switch_inline_query_current_chat,omitempty"`
	// Optional. Description of the game that will be launched when the user
	// presses the button. *NOTE*: This type of button *must* always be the
	// first button in the first row.
	CallbackGame *CallbackGame `json:"callback_game,omitempty"`
	// Optional. Specify True, to send a Pay button. *NOTE*: This type of button
	// *must* always be the first button in the first row and can only be used
	// in invoice messages.
	Pay bool `json:"pay,omitempty"`
}

// This object represents a parameter of the inline keyboard button used to
// automatically authorize a user. Serves as a great replacement for the
// Telegram Login Widget when the user is coming from Telegram. All the user
// needs to do is tap/click a button and confirm that they want to log in:
// https://core.telegram.org/file/811140015/1734/8VZFkwWXalM.97872/6127fa62d8a0bf2b3c
// https://core.telegram.org/widgets/login
//
// Telegram apps support these buttons as of version 5.7.
// https://telegram.org/blog/privacy-discussions-web-bots#meet-seamless-web-bots
//
// Sample bot: @discussbot
//
// https://core.telegram.org/bots/api#loginurl
type LoginURL struct {
	// An HTTP URL to be opened with user authorization data added to the query
	// string when the button is pressed. If the user refuses to provide
	// authorization data, the original URL without information about the user
	// will be opened. The data added is the same as described in Receiving
	// authorization data.
	// https://core.telegram.org/widgets/login#receiving-authorization-data
	// *NOTE*: You *must* always check the hash of the received data to verify
	// the authentication and the integrity of the data as described in Checking
	// authorization.
	// https://core.telegram.org/widgets/login#checking-authorization
	URL string `json:"url"`
	// Optional. New text of the button in forwarded messages.
	ForwardText string `json:"forward_text,omitempty"`
	// Optional. Username of a bot, which will be used for user authorization.
	// See Setting up a bot for more details. If not specified, the current
	// bot's username will be assumed. The url's domain must be the same as the
	// domain linked with the bot. See Linking your domain to the bot for more
	// details. https://core.telegram.org/widgets/login#setting-up-a-bot
	// https://core.telegram.org/widgets/login#linking-your-domain-to-the-bot
	BotUsername Username `json:"bot_username,omitempty"`
	// Optional. Pass True to request the permission for your bot to send
	// messages to the user.
	RequestWriteAccess bool `json:"request_write_access,omitempty"`
}

// This object represents an incoming callback query from a callback button in
// an inline keyboard. If the button that originated the query was attached to a
// message sent by the bot, the field message will be present. If the button was
// attached to a message sent via the bot (in inline mode), the field
// inline_message_id will be present. Exactly one of the fields data or
// game_short_name will be present.
//
// *NOTE*: After the user presses a callback button, Telegram clients will
// display a progress bar until you call answerCallbackQuery. It is, therefore,
// necessary to react by calling answerCallbackQuery even if no notification to
// the user is needed (e.g., without specifying any of the optional parameters).
// https://core.telegram.org/bots/api#answercallbackquery
//
// https://core.telegram.org/bots/api#callbackquery
type CallbackQuery struct {
	// Unique identifier for this query
	ID CallbackQueryID `json:"id"`
	// Sender
	From *User `json:"from"`
	// Optional. Message with the callback button that originated the query.
	// Note that message content and message date will not be available if the
	// message is too old
	Message *Message `json:"message,omitempty"`
	// Optional. Identifier of the message sent via the bot in inline mode, that
	// originated the query.
	InlineMessageID InlineMessageID `json:"inline_message_id,omitempty"`
	// Global identifier, uniquely corresponding to the chat to which the
	// message with the callback button was sent. Useful for high scores in
	// games. https://core.telegram.org/bots/api#games
	ChatInstance ChatInstance `json:"chat_instance"`
	// Optional. Data associated with the callback button. Be aware that a bad
	// client can send arbitrary data in this field.
	Data string `json:"data,omitempty"`
	// Optional. Short name of a Game to be returned, serves as the unique
	// identifier for the game https://core.telegram.org/bots/api#games
	GameShortName GameShortName `json:"game_short_name,omitempty"`
}

// Upon receiving a message with this object, Telegram clients will display a
// reply interface to the user (act as if the user has selected the bot's
// message and tapped 'Reply'). This can be extremely useful if you want to
// create user-friendly step-by-step interfaces without having to sacrifice
// privacy mode. https://core.telegram.org/bots#privacy-mode
//
//
// Example: A poll bot for groups runs in privacy mode (only receives commands,
// replies to its messages and mentions). There could be two ways to create a
// new poll: - Explain the user how to send a command with parameters (e.g.
// /newpoll question answer1 answer2). May be appealing for hardcore users but
// lacks modern day polish. - Guide the user through a step-by-step process.
// 'Please send me your question', 'Cool, now let's add the first answer
// option', 'Great. Keep adding answer options, then send /done when you're
// ready'. The last option is definitely more attractive. And if you use
// ForceReply in your bot's questions, it will receive the user's answers even
// if it only receives replies, commands and mentions — without any extra work
// for the user.
//
// https://t.me/PollBot
//
// https://core.telegram.org/bots/api#forcereply
type ForceReply struct {
	// Shows reply interface to the user, as if they manually selected the bot's
	// message and tapped 'Reply'
	ForceReply bool `json:"force_reply"`
	// Optional. The placeholder to be shown in the input field when the reply
	// is active; 1-64 characters
	InputFieldPlaceholder string `json:"input_field_placeholder,omitempty"`
	// Optional. Use this parameter if you want to force reply from specific
	// users only. Targets: 1) users that are @mentioned in the text of the
	// Message object; 2) if the bot's message is a reply (has
	// reply_to_message_id), sender of the original message.
	Selective bool `json:"selective,omitempty"`
}

// This object represents a chat photo.
//
// https://core.telegram.org/bots/api#chatphoto
type ChatPhoto struct {
	// File identifier of small (160x160) chat photo. This file_id can be used
	// only for photo download and only for as long as the photo is not changed.
	SmallFileID FileID `json:"small_file_id"`
	// Unique file identifier of small (160x160) chat photo, which is supposed
	// to be the same over time and for different bots. Can't be used to
	// download or reuse the file.
	SmallFileUniqueID FileUniqueID `json:"small_file_unique_id"`
	// File identifier of big (640x640) chat photo. This file_id can be used
	// only for photo download and only for as long as the photo is not changed.
	BigFileID FileID `json:"big_file_id"`
	// Unique file identifier of big (640x640) chat photo, which is supposed to
	// be the same over time and for different bots. Can't be used to download
	// or reuse the file.
	BigFileUniqueID FileUniqueID `json:"big_file_unique_id"`
}

// Represents an invite link for a chat.
//
// https://core.telegram.org/bots/api#chatinvitelink
type ChatInviteLink struct {
	// The invite link. If the link was created by another chat administrator,
	// then the second part of the link will be replaced with "...".
	InviteLink string `json:"invite_link"`
	// Creator of the link
	Creator *User `json:"creator"`
	// True, if users joining the chat via the link need to be approved by chat
	// administrators
	CreatesJoinRequest bool `json:"creates_join_request"`
	// True, if the link is primary
	IsPrimary bool `json:"is_primary"`
	// True, if the link is revoked
	IsRevoked bool `json:"is_revoked"`
	// Optional. Invite link name
	Name string `json:"name,omitempty"`
	// Optional. Point in time (Unix timestamp) when the link will expire or has
	// been expired
	ExpireDate int64 `json:"expire_date,omitempty"`
	// Optional. Maximum number of users that can be members of the chat
	// simultaneously after joining the chat via this invite link; 1-99999
	MemberLimit int `json:"member_limit,omitempty"`
	// Optional. Number of pending join requests created using this link
	PendingJoinRequestCount int `json:"pending_join_request_count,omitempty"`
}

// Represents the rights of an administrator in a chat.
//
// https://core.telegram.org/bots/api#chatadministratorrights
type ChatAdministratorRights struct {
	// True, if the user's presence in the chat is hidden
	IsAnonymous bool `json:"is_anonymous,omitempty"`
	// True, if the administrator can access the chat event log, chat
	// statistics, message statistics in channels, see channel members, see
	// anonymous administrators in supergroups and ignore slow mode. Implied by
	// any other administrator privilege
	CanManageChat bool `json:"can_manage_chat,omitempty"`
	// True, if the administrator can delete messages of other users
	CanDeleteMessages bool `json:"can_delete_messages,omitempty"`
	// True, if the administrator can manage video chats
	CanManageVideoChats bool `json:"can_manage_video_chats,omitempty"`
	// True, if the administrator can restrict, ban or unban chat members
	CanRestrictMembers bool `json:"can_restrict_members,omitempty"`
	// True, if the administrator can add new administrators with a subset of
	// their own privileges or demote administrators that he has promoted,
	// directly or indirectly (promoted by administrators that were appointed by
	// the user)
	CanPromoteMembers bool `json:"can_promote_members,omitempty"`
	// True, if the user is allowed to change the chat title, photo and other
	// settings
	CanChangeInfo bool `json:"can_change_info,omitempty"`
	// True, if the user is allowed to invite new users to the chat
	CanInviteUsers bool `json:"can_invite_users,omitempty"`
	// Optional. True, if the administrator can post in the channel; channels
	// only
	CanPostMessages bool `json:"can_post_messages,omitempty"`
	// Optional. True, if the administrator can edit messages of other users and
	// can pin messages; channels only
	CanEditMessages bool `json:"can_edit_messages,omitempty"`
	// Optional. True, if the user is allowed to pin messages; groups and
	// supergroups only
	CanPinMessages bool `json:"can_pin_messages,omitempty"`
}

// This object contains information about one member of a chat. Currently, the
// following 6 types of chat members are supported:
//   ChatMemberOwner - Represents a chat member that owns the chat and has all administrator privileges.
//   ChatMemberAdministrator - Represents a chat member that has some additional privileges.
//   ChatMemberMember - Represents a chat member that has no additional privileges or restrictions.
//   ChatMemberRestricted - Represents a chat member that is under certain restrictions in the chat. Supergroups only.
//   ChatMemberLeft - Represents a chat member that isn't currently a member of the chat, but may join it themselves.
//   ChatMemberBanned - Represents a chat member that was banned in the chat and can't return to the chat or view chat messages.
//
// https://core.telegram.org/bots/api#chatmember
// https://core.telegram.org/bots/api#chatmemberowner
// https://core.telegram.org/bots/api#chatmemberadministrator
// https://core.telegram.org/bots/api#chatmembermember
// https://core.telegram.org/bots/api#chatmemberrestricted
// https://core.telegram.org/bots/api#chatmemberleft
// https://core.telegram.org/bots/api#chatmemberbanned
type ChatMember struct {
	// ChatMemberOwner, ChatMemberAdministrator, ChatMemberMember,
	// ChatMemberRestricted, ChatMemberLeft, ChatMemberBanned

	// The member's status in the chat
	//   ChatMemberOwner - always "creator"
	//   ChatMemberAdministrator - always "administrator"
	//   ChatMemberMember - always "member"
	//   ChatMemberRestricted - always "restricted"
	//   ChatMemberLeft - always "left"
	//   ChatMemberBanned - always "kicked"
	Status ChatMemberStatus `json:"status"`
	// Information about the user
	User *User `json:"user"`

	// ChatMemberOwner, ChatMemberAdministrator

	// True, if the user's presence in the chat is hidden
	IsAnonymous bool `json:"is_anonymous,omitempty"`
	// Optional. Custom title for this user
	CustomTitle string `json:"custom_title,omitempty"`

	// ChatMemberAdministrator

	// True, if the bot is allowed to edit administrator privileges of that user
	CanBeEdited bool `json:"can_be_edited,omitempty"`
	// True, if the administrator can access the chat event log, chat
	// statistics, message statistics in channels, see channel members, see
	// anonymous administrators in supergroups and ignore slow mode. Implied by
	// any other administrator privilege
	CanManageChat bool `json:"can_manage_chat,omitempty"`
	// True, if the administrator can delete messages of other users
	CanDeleteMessages bool `json:"can_delete_messages,omitempty"`
	// True, if the administrator can manage video chats
	CanManageVideoChats bool `json:"can_manage_video_chats,omitempty"`
	// True, if the administrator can restrict, ban or unban chat members
	CanRestrictMembers bool `json:"can_restrict_members,omitempty"`
	// True, if the administrator can add new administrators with a subset of
	// their own privileges or demote administrators that he has promoted,
	// directly or indirectly (promoted by administrators that were appointed by
	// the user)
	CanPromoteMembers bool `json:"can_promote_members,omitempty"`
	// Optional. True, if the administrator can post in the channel; channels
	// only
	CanPostMessages bool `json:"can_post_messages,omitempty"`
	// Optional. True, if the administrator can edit messages of other users and
	// can pin messages; channels only
	CanEditMessages bool `json:"can_edit_messages,omitempty"`

	// ChatMemberAdministrator, ChatMemberRestricted

	// True, if the user is allowed to change the chat title, photo and other
	// settings
	CanChangeInfo bool `json:"can_change_info,omitempty"`
	// True, if the user is allowed to invite new users to the chat
	CanInviteUsers bool `json:"can_invite_users,omitempty"`
	// Optional. True, if the user is allowed to pin messages; groups and
	// supergroups only
	CanPinMessages bool `json:"can_pin_messages,omitempty"`

	// ChatMemberRestricted

	// True, if the user is a member of the chat at the moment of the request
	IsMember bool `json:"is_member,omitempty"`
	// True, if the user is allowed to send text messages, contacts, locations
	// and venues
	CanSendMessages bool `json:"can_send_messages,omitempty"`
	// True, if the user is allowed to send audios, documents, photos, videos,
	// video notes and voice notes
	CanSendMediaMessages bool `json:"can_send_media_messages,omitempty"`
	// True, if the user is allowed to send polls
	CanSendPolls bool `json:"can_send_polls,omitempty"`
	// True, if the user is allowed to send animations, games, stickers and use
	// inline bots
	CanSendOtherMessages bool `json:"can_send_other_messages,omitempty"`
	// True, if the user is allowed to add web page previews to their messages
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews,omitempty"`

	// ChatMemberRestricted, ChatMemberBanned

	// Date when restrictions will be lifted for this user; unix time. If 0,
	// then the user is restricted forever
	UntilDate int64 `json:"until_date,omitempty"`
}

// This object represents changes in the status of a chat member.
//
// https://core.telegram.org/bots/api#chatmemberupdated
type ChatMemberUpdated struct {
	// Chat the user belongs to
	Chat *Chat `json:"chat"`
	// Performer of the action, which resulted in the change
	From *User `json:"from"`
	// Date the change was done in Unix time
	Date int64 `json:"date"`
	// Previous information about the chat member
	OldChatMember *ChatMember `json:"old_chat_member"`
	// New information about the chat member
	NewChatMember *ChatMember `json:"new_chat_member"`
	// Optional. Chat invite link, which was used by the user to join the chat;
	// for joining by invite link events only.
	InviteLink *ChatInviteLink `json:"invite_link,omitempty"`
}

// Represents a join request sent to a chat.
//
// https://core.telegram.org/bots/api#chatjoinrequest
type ChatJoinRequest struct {
	// Chat to which the request was sent
	Chat *Chat `json:"chat"`
	// User that sent the join request
	From *User `json:"from"`
	// Date the request was sent in Unix time
	Date int64 `json:"date"`
	// Optional. Bio of the user.
	Bio string `json:"bio,omitempty"`
	// Optional. Chat invite link that was used by the user to send the join
	// request
	InviteLink *ChatInviteLink `json:"invite_link,omitempty"`
}

// Describes actions that a non-administrator user is allowed to take in a chat.
//
// https://core.telegram.org/bots/api#chatpermissions
type ChatPermissions struct {
	// Optional. True, if the user is allowed to send text messages, contacts,
	// locations and venues
	CanSendMessages bool `json:"can_send_messages,omitempty"`
	// Optional. True, if the user is allowed to send audios, documents, photos,
	// videos, video notes and voice notes, implies can_send_messages
	CanSendMediaMessages bool `json:"can_send_media_messages,omitempty"`
	// Optional. True, if the user is allowed to send polls, implies
	// can_send_messages
	CanSendPolls bool `json:"can_send_polls,omitempty"`
	// Optional. True, if the user is allowed to send animations, games,
	// stickers and use inline bots, implies can_send_media_messages
	CanSendOtherMessages bool `json:"can_send_other_messages,omitempty"`
	// Optional. True, if the user is allowed to add web page previews to their
	// messages, implies can_send_media_messages
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews,omitempty"`
	// Optional. True, if the user is allowed to change the chat title, photo
	// and other settings. Ignored in public supergroups
	CanChangeInfo bool `json:"can_change_info,omitempty"`
	// Optional. True, if the user is allowed to invite new users to the chat
	CanInviteUsers bool `json:"can_invite_users,omitempty"`
	// Optional. True, if the user is allowed to pin messages. Ignored in public
	// supergroups
	CanPinMessages bool `json:"can_pin_messages,omitempty"`
}

// Represents a location to which a chat is connected.
//
// https://core.telegram.org/bots/api#chatlocation
type ChatLocation struct {
	// The location to which the supergroup is connected. Can't be a live
	// location.
	Location *Location `json:"location"`
	// Location address; 1-64 characters, as defined by the chat owner
	Address string `json:"address"`
}

// This object represents a bot command.
//
// https://core.telegram.org/bots/api#botcommand
type BotCommand struct {
	// Text of the command; 1-32 characters. Can contain only lowercase English
	// letters, digits and underscores.
	Command string `json:"command"`
	// Description of the command; 1-256 characters.
	Description string `json:"description"`
}

// This object represents the scope to which bot commands are applied.
//
// Currently, the following 7 scopes are supported:
//   BotCommandScopeDefault - Represents the default scope of bot commands. Default commands are used if no commands with a narrower scope are specified for the user.
//   BotCommandScopeAllPrivateChats - Represents the scope of bot commands, covering all private chats.
//   BotCommandScopeAllGroupChats - Represents the scope of bot commands, covering all group and supergroup chats.
//   BotCommandScopeAllChatAdministrators - Represents the scope of bot commands, covering all group and supergroup chat administrators.
//   BotCommandScopeChat - Represents the scope of bot commands, covering a specific chat.
//   BotCommandScopeChatAdministrators - Represents the scope of bot commands, covering all administrators of a specific group or supergroup chat.
//   BotCommandScopeChatMember - Represents the scope of bot commands, covering a specific member of a group or supergroup chat.
//
// Determining list of commands
//
// The following algorithm is used to determine the list of commands for a
// particular user viewing the bot menu. The first list of commands which is set
// is returned: Commands in the chat with the bot:
//   botCommandScopeChat + language_code
//   botCommandScopeChat
//   botCommandScopeAllPrivateChats + language_code
//   botCommandScopeAllPrivateChats
//   botCommandScopeDefault + language_code
//   botCommandScopeDefault
//
// Commands in group and supergroup chats:
//   botCommandScopeChatMember + language_code
//   botCommandScopeChatMember
//   botCommandScopeChatAdministrators + language_code (administrators only)
//   botCommandScopeChatAdministrators (administrators only)
//   botCommandScopeChat + language_code
//   botCommandScopeChat
//   botCommandScopeAllChatAdministrators + language_code (administrators only)
//   botCommandScopeAllChatAdministrators (administrators only)
//   botCommandScopeAllGroupChats + language_code
//   botCommandScopeAllGroupChats
//   botCommandScopeDefault + language_code
//   botCommandScopeDefault
//
// https://core.telegram.org/bots/api#botcommandscope
// https://core.telegram.org/bots/api#botcommandscopedefault
// https://core.telegram.org/bots/api#botcommandscopeallprivatechats
// https://core.telegram.org/bots/api#botcommandscopeallgroupchats
// https://core.telegram.org/bots/api#botcommandscopeallchatadministrators
// https://core.telegram.org/bots/api#botcommandscopechat
// https://core.telegram.org/bots/api#botcommandscopechatadministrators
// https://core.telegram.org/bots/api#botcommandscopechatmember
type BotCommandScope struct {
	// Scope type
	//     BotCommandScopeDefault - must be default
	//     BotCommandScopeAllPrivateChats - must be all_private_chats
	//     BotCommandScopeAllGroupChats - must be all_group_chats
	//     BotCommandScopeAllChatAdministrators - must be all_chat_administrators
	//     BotCommandScopeChat - must be chat
	//     BotCommandScopeChatAdministrators - must be chat_administrators
	//     BotCommandScopeChatMember - must be chat_member
	Type BotCommandScopeType `json:"type"`

	// Unique identifier for the target chat or username of the target
	// supergroup (in the format @supergroupusername)
	ChatID ChatIDOrUsername `json:"chat_id,omitempty"`

	// Unique identifier of the target user
	UserID UserID `json:"user_id,omitempty"`
}

// This object describes the bot's menu button in a private chat. It should be
// one of
//   MenuButtonCommands - Represents a menu button, which opens the bot's list of commands.
//   MenuButtonWebApp - Represents a menu button, which launches a Web App. https://core.telegram.org/bots/webapps
//   MenuButtonDefault - Describes that no specific value for the menu button was set.
//
// If a menu button other than MenuButtonDefault is set for a private chat, then
// it is applied in the chat. Otherwise the default menu button is applied. By
// default, the menu button opens the list of bot commands.
// https://core.telegram.org/bots/api#menubuttondefault
//
// https://core.telegram.org/bots/api#menubutton
// https://core.telegram.org/bots/api#menubuttoncommands
// https://core.telegram.org/bots/api#menubuttonwebapp
// https://core.telegram.org/bots/api#menubuttondefault
type MenuButton struct {
	// MenuButtonCommands, MenuButtonWebApp, MenuButtonDefault

	// Type of the button
	//   MenuButtonCommands - must be commands
	//   MenuButtonWebApp - must be web_app
	//   MenuButtonDefault - must be default
	Type MenuButtonType `json:"type"`

	// MenuButtonWebApp

	// Text on the button
	Text string `json:"text,omitempty"`
	// Description of the Web App that will be launched when the user presses
	// the button. The Web App will be able to send an arbitrary message on
	// behalf of the user using the method answerWebAppQuery.
	// https://core.telegram.org/bots/api#answerwebappquery
	WebApp *WebAppInfo `json:"web_app,omitempty"`
}

// Describes why a request was unsuccessful.
//
// *Used internally by this library*
// https://core.telegram.org/bots/api#responseparameters
type ResponseParameters struct {
	// Optional. The group has been migrated to a supergroup with the specified
	// identifier.
	MigrateToChatID ChatID `json:"migrate_to_chat_id,omitempty"`
	// Optional. In case of exceeding flood control, the number of seconds left
	// to wait before the request can be repeated
	RetryAfter int `json:"retry_after,omitempty"`
}

// This object represents the content of a media message to be sent. It should
// be one of
//   InputMediaAnimation - Represents an animation file (GIF or H.264/MPEG-4 AVC video without sound) to be sent.
//   InputMediaDocument - Represents a general file to be sent.
//   InputMediaAudio - Represents an audio file to be treated as music to be sent.
//   InputMediaPhoto - Represents a photo to be sent.
//   InputMediaVideo - Represents a video to be sent.
//
// https://core.telegram.org/bots/api#inputmedia
// https://core.telegram.org/bots/api#inputmediaanimation
// https://core.telegram.org/bots/api#inputmediadocument
// https://core.telegram.org/bots/api#inputmediaaudio
// https://core.telegram.org/bots/api#inputmediaphoto
// https://core.telegram.org/bots/api#inputmediavideo
type InputMedia struct {
	// Type of the result
	//   InputMediaPhoto - must be photo
	//   InputMediaVideo - must be video
	//   InputMediaAnimation - must be animation
	//   InputMediaAudio - must be audio
	//   InputMediaDocument - must be document
	Type InputMediaType `json:"type"`
	// File to send
	Media InputFile `json:"media"`
	// Optional. Caption of the file to be sent, 0-1024 characters after
	// entities parsing
	Caption string `json:"caption,omitempty"`
	// Optional. Mode for parsing entities in the animation caption. See
	// formatting options for more details.
	// https://core.telegram.org/bots/api#formatting-options
	ParseMode ParseMode `json:"parse_mode,omitempty"`
	// Optional. List of special entities that appear in the caption, which can
	// be specified instead of parse_mode
	CaptionEntities []*MessageEntity `json:"caption_entities,omitempty"`

	// Optional. Thumbnail of the file sent; can be ignored if thumbnail
	// generation for the file is supported server-side. The thumbnail should be
	// in JPEG format and less than 200 kB in size. A thumbnail's width and
	// height should not exceed 320.
	Thumb InputFile `json:"thumb,omitempty"`

	// Optional. Video or Animation width
	Width int `json:"width,omitempty"`
	// Optional. Video or Animation height
	Height int `json:"height,omitempty"`

	// Optional. Video, animation or audio duration in seconds
	Duration int `json:"duration,omitempty"`

	// Optional. Pass True, if the uploaded video is suitable for streaming
	SupportsStreaming bool `json:"supports_streaming,omitempty"`

	// Optional. Performer of the audio
	Performer int `json:"performer,omitempty"`
	// Optional. Title of the audio
	Title int `json:"title,omitempty"`

	// Optional. Disables automatic server-side content type detection for files
	// uploaded using multipart/form-data. Always True, if the document is sent
	// as part of an album.
	DisableContentTypeDetection bool `json:"disable_content_type_detection,omitempty"`
}

// INPUT FILE IS NOT A VANILLA TYPE FROM TELEGRAM BOT API DOCUMENTATION THERE IS
// NO VANILLA TYPE SPECIFICATION FOR INPUT FILE

// This object represents the contents of a file to be uploaded. Must be posted
// using multipart/form-data in the usual way that files are uploaded via the
// browser.
//
// *NOTE FROM THIS LIBRARY DEVELOPER*: InputFile can be either FileID, FileURL,
// or FileReader. If you want to upload large files, use self-hosted telegram
// bot api server, and upload using FileURL with the file URL scheme
// https://github.com/tdlib/telegram-bot-api
// https://hub.docker.com/r/rmuhamedgaliev/telegram-bot-api
// https://github.com/rmuhamedgaliev/telegram-bot-api
// https://en.wikipedia.org/wiki/File_URI_scheme
//
// https://core.telegram.org/bots/api#inputfile
// https://core.telegram.org/bots/api#sending-files
type InputFile interface {
	multipartFormFile() (fieldname string, filename string, reader io.Reader)
}

// Identifier for a file, which can be used to download or reuse the file And
// file id for InputFile fields
//
// https://core.telegram.org/bots/api#sending-files
type FileID string

func (FileID) multipartFormFile() (fieldname string, filename string, reader io.Reader) {
	return "", "", nil
}

// File URL for InputFile fields
//
// https://core.telegram.org/bots/api#sending-files
type FileURL string

func (FileURL) multipartFormFile() (fieldname string, filename string, reader io.Reader) {
	return "", "", nil
}

// File reader for InputFile fields
//
// https://core.telegram.org/bots/api#sending-files
type FileReader struct {
	// Name of the file
	Name   string
	Reader io.Reader

	fieldname string
}

func (fr *FileReader) multipartFormFile() (fieldname string, filename string, reader io.Reader) {
	fr.checkFieldname()

	return fr.fieldname, fr.Name, fr.Reader
}

func (fr *FileReader) MarshalJSON() ([]byte, error) {
	fr.checkFieldname()

	return []byte(`"` + "attach://" + fr.fieldname + `"`), nil
}

func (fr *FileReader) checkFieldname() {
	if fr.fieldname != "" {
		return
	}

	b := make([]byte, 6)
	rand.Read(b)

	fr.fieldname = hex.EncodeToString(b)
}
