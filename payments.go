// https://core.telegram.org/bots/api#payments
package telegrambot

import "fmt"

type SendInvoiceParams struct {
	// Unique identifier for the target chat or username of the target channel
	// (in the format @channelusername)
	ChatID ChatIDOrUsername `json:"chat_id"`
	// Product name, 1-32 characters
	Title string `json:"title"`
	// Product description, 1-255 characters
	Description string `json:"description"`
	// Bot-defined invoice payload, 1-128 bytes. This will not be displayed to
	// the user, use for your internal processes.
	Payload string `json:"payload"`
	// Payments provider token, obtained via Botfather https://t.me/botfather
	ProviderToken string `json:"provider_token"`
	// Three-letter ISO 4217 currency code, see more on currencies
	// https://core.telegram.org/bots/payments#supported-currencies
	Currency string `json:"currency"`
	// Price breakdown, a JSON-serialized list of components (e.g. product
	// price, tax, discount, delivery cost, delivery tax, bonus, etc.)
	Prices []*LabeledPrice `json:"prices"`
	// Optional. The maximum accepted amount for tips in the smallest units of
	// the currency (integer, not float/double). For example, for a maximum tip
	// of US$ 1.45 pass max_tip_amount = 145. See the exp parameter in
	// currencies.json, it shows the number of digits past the decimal point for
	// each currency (2 for the majority of currencies). Defaults to 0
	// https://core.telegram.org/bots/payments/currencies.json
	MaxTipAmount int `json:"max_tip_amount,omitempty"`
	// Optional. A JSON-serialized array of suggested amounts of tips in the
	// smallest units of the currency (integer, not float/double). At most 4
	// suggested tip amounts can be specified. The suggested tip amounts must be
	// positive, passed in a strictly increased order and must not exceed
	// max_tip_amount.
	SuggestedTipAmounts []int `json:"suggested_tip_amounts,omitempty"`
	// Optional. Unique deep-linking parameter. If left empty, forwarded copies
	// of the sent message will have a Pay button, allowing multiple users to
	// pay directly from the forwarded message, using the same invoice. If
	// non-empty, forwarded copies of the sent message will have a URL button
	// with a deep link to the bot (instead of a Pay button), with the value
	// used as the start parameter
	StartParameter string `json:"start_parameter,omitempty"`
	// Optional. A JSON-serialized data about the invoice, which will be shared
	// with the payment provider. A detailed description of required fields
	// should be provided by the payment provider.
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
	// Optional. Sends the message silently. Users will receive a notification
	// with no sound. https://telegram.org/blog/channels-2-0#silent-messages
	DisableNotification bool `json:"disable_notification,omitempty"`
	// Optional. Protects the contents of the sent message from forwarding and
	// saving
	ProtectContent bool `json:"protect_content,omitempty"`
	// If the message is a reply, ID of the original message
	ReplyToMessageID MessageID `json:"reply_to_message_id,omitempty"`
	// Optional. Pass True, if the message should be sent even if the specified
	// replied-to message is not found
	AllowSendingWithoutReply bool `json:"allow_sending_without_reply,omitempty"`
	// Optional. A JSON-serialized object for an inline keyboard. If empty, one
	// 'Pay `total price`' button will be shown. If not empty, the first button
	// must be a Pay button.
	// https://core.telegram.org/bots#inline-keyboards-and-on-the-fly-updating
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// Use this method to send invoices. On success, the sent Message is returned.
// https://core.telegram.org/bots/api#message
//
// https://core.telegram.org/bots/api#sendinvoice
func (api *API) SendInvoice(params *SendInvoiceParams) (*Message, error) {
	msg := &Message{}

	migrateToChatID, err := api.makeAPICall("sendInvoice", params, nil, msg)
	if err != nil {
		if migrateToChatID != 0 {
			params.ChatID = migrateToChatID
			_, err = api.makeAPICall("sendInvoice", params, nil, msg)
			if err != nil {
				return nil, fmt.Errorf("SendInvoice: %w", err)
			}
		} else {
			return nil, fmt.Errorf("SendInvoice: %w", err)
		}
	}

	return msg, nil
}

type CreateInvoiceLinkParams struct {
	// Product name, 1-32 characters
	Title string `json:"title"`
	// Product description, 1-255 characters
	Description string `json:"description"`
	// Bot-defined invoice payload, 1-128 bytes. This will not be displayed to
	// the user, use for your internal processes.
	Payload string `json:"payload"`
	// Payments provider token, obtained via Botfather https://t.me/botfather
	ProviderToken string `json:"provider_token"`
	// Three-letter ISO 4217 currency code, see more on currencies
	// https://core.telegram.org/bots/payments#supported-currencies
	Currency string `json:"currency"`
	// Price breakdown, a JSON-serialized list of components (e.g. product
	// price, tax, discount, delivery cost, delivery tax, bonus, etc.)
	Prices []*LabeledPrice `json:"prices"`
	// Optional. The maximum accepted amount for tips in the smallest units of
	// the currency (integer, not float/double). For example, for a maximum tip
	// of US$ 1.45 pass max_tip_amount = 145. See the exp parameter in
	// currencies.json, it shows the number of digits past the decimal point for
	// each currency (2 for the majority of currencies). Defaults to 0
	// https://core.telegram.org/bots/payments/currencies.json
	MaxTipAmount int `json:"max_tip_amount,omitempty"`
	// Optional. A JSON-serialized array of suggested amounts of tips in the
	// smallest units of the currency (integer, not float/double). At most 4
	// suggested tip amounts can be specified. The suggested tip amounts must be
	// positive, passed in a strictly increased order and must not exceed
	// max_tip_amount.
	SuggestedTipAmounts []int `json:"suggested_tip_amounts,omitempty"`
	// Optional. A JSON-serialized data about the invoice, which will be shared
	// with the payment provider. A detailed description of required fields
	// should be provided by the payment provider.
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

// Use this method to create a link for an invoice. Returns the created invoice
// link as String on success.
//
// https://core.telegram.org/bots/api#createinvoicelink
func (api *API) CreateInvoiceLink(params *CreateInvoiceLinkParams) (string, error) {
	link := ""

	_, err := api.makeAPICall("createInvoiceLink", params, nil, &link)
	if err != nil {
		return "", fmt.Errorf("createInvoiceLink: %w", err)
	}

	return link, nil
}

type AnswerShippingQueryParams struct {
	// Unique identifier for the query to be answered
	ShippingQueryID ShippingQueryID `json:"shipping_query_id"`
	// Specify True if delivery to the specified address is possible and False
	// if there are any problems (for example, if delivery to the specified
	// address is not possible)
	OK bool `json:"ok"`
	// Optional. Required if ok is True. A JSON-serialized array of available
	// shipping options.
	ShippingOptions []*ShippingOption `json:"shipping_options,omitempty"`
	// Optional. Required if ok is False. Error message in human readable form
	// that explains why it is impossible to complete the order (e.g. "Sorry,
	// delivery to your desired address is unavailable'). Telegram will display
	// this message to the user.
	ErrorMessage string `json:"error_message,omitempty"`
}

// If you sent an invoice requesting a shipping address and the parameter
// is_flexible was specified, the Bot API will send an Update with a
// shipping_query field to the bot. Use this method to reply to shipping
// queries. On success, True is returned.
// https://core.telegram.org/bots/api#update
//
// https://core.telegram.org/bots/api#answershippingquery
func (api *API) AnswerShippingQuery(params *AnswerShippingQueryParams) error {
	_, err := api.makeAPICall("answerShippingQuery", params, nil, nil)
	if err != nil {
		return fmt.Errorf("AnswerShippingQuery: %w", err)
	}

	return nil
}

type AnswerPreCheckoutQueryParams struct {
	// Unique identifier for the query to be answered
	PreCheckoutQueryID PreCheckoutQueryID `json:"pre_checkout_query_id"`
	// Specify True if everything is alright (goods are available, etc.) and the
	// bot is ready to proceed with the order. Use False if there are any
	// problems.
	OK bool `json:"ok"`
	// Optional. Required if ok is False. Error message in human readable form
	// that explains the reason for failure to proceed with the checkout (e.g.
	// "Sorry, somebody just bought the last of our amazing black T-shirts while
	// you were busy filling out your payment details. Please choose a different
	// color or garment!"). Telegram will display this message to the user.
	ErrorMessage string `json:"error_message,omitempty"`
}

// Once the user has confirmed their payment and shipping details, the Bot API
// sends the final confirmation in the form of an Update with the field
// pre_checkout_query. Use this method to respond to such pre-checkout queries.
// On success, True is returned. *Note*: The Bot API must receive an answer
// within 10 seconds after the pre-checkout query was sent.
// https://core.telegram.org/bots/api#update
//
// https://core.telegram.org/bots/api#answerprecheckoutquery
func (api *API) AnswerPreCheckoutQuery(params *AnswerPreCheckoutQueryParams) error {
	_, err := api.makeAPICall("answerPreCheckoutQuery", params, nil, nil)
	if err != nil {
		return fmt.Errorf("AnswerPreCheckoutQuery: %w", err)
	}

	return nil
}

// This object represents a portion of the price for goods or services.
//
// https://core.telegram.org/bots/api#labeledprice
type LabeledPrice struct {
	// Portion label
	Label string `json:"label"`
	// Price of the product in the smallest units of the currency (integer, not
	// float/double). For example, for a price of US$ 1.45 pass amount = 145.
	// See the exp parameter in currencies.json, it shows the number of digits
	// past the decimal point for each currency (2 for the majority of
	// currencies).
	Amount int `json:"amount"`
}

// This object contains basic information about an invoice.
//
// https://core.telegram.org/bots/api#invoice
type Invoice struct {
	// Product name
	Title string `json:"title"`
	// Product description
	Description string `json:"description"`
	// Unique bot deep-linking parameter that can be used to generate this
	// invoice
	StartParameter string `json:"start_parameter"`
	// Three-letter ISO 4217 currency code
	// https://core.telegram.org/bots/payments#supported-currencies
	Currency string `json:"currency"`
	// Total price in the smallest units of the currency (integer, not
	// float/double). For example, for a price of US$ 1.45 pass amount = 145.
	// See the exp parameter in currencies.json, it shows the number of digits
	// past the decimal point for each currency (2 for the majority of
	// currencies). https://core.telegram.org/bots/payments/currencies.json
	// https://core.telegram.org/bots/payments/currencies.json
	TotalAmount int `json:"total_amount"`
}

// This object represents a shipping address.
//
// https://core.telegram.org/bots/api#shippingaddress
type ShippingAddress struct {
	// ISO 3166-1 alpha-2 country code
	Country_code string `json:"country_code"`
	// State, if applicable
	State string `json:"state"`
	// City
	City string `json:"city"`
	// First line for the address
	StreetLine1 string `json:"street_line1"`
	// Second line for the address
	StreetLine2 string `json:"street_line2"`
	// Address
	PostCode string `json:"post_code"`
}

// This object represents information about an order.
//
// https://core.telegram.org/bots/api#orderinfo
type OrderInfo struct {
	// Optional. User name
	Name string `json:"name,omitempty"`
	// Optional. User's phone number
	PhoneNumber string `json:"phone_number,omitempty"`
	// Optional. User email
	Email string `json:"email,omitempty"`
	// Optional. User shipping address
	ShippingAddress *ShippingAddress `json:"shipping_address,omitempty"`
}

// This object represents one shipping option.
//
// https://core.telegram.org/bots/api#shippingoption
type ShippingOption struct {
	// Shipping option identifier
	ID ShippingOptionID `json:"id"`
	// Option title
	Title string `json:"title"`
	// List of price portions
	Prices []*LabeledPrice `json:"prices"`
}

// This object contains basic information about a successful payment.
//
// https://core.telegram.org/bots/api#successfulpayment
type SuccessfulPayment struct {
	// Three-letter ISO 4217 currency code
	// https://core.telegram.org/bots/payments#supported-currencies
	Currency string `json:"currency"`
	// Total price in the smallest units of the currency (integer, not
	// float/double). For example, for a price of US$ 1.45 pass amount = 145.
	// See the exp parameter in currencies.json, it shows the number of digits
	// past the decimal point for each currency (2 for the majority of
	// currencies). https://core.telegram.org/bots/payments/currencies.json
	TotalAmount int `json:"total_amount"`
	// Bot specified invoice payload
	InvoicePayload string `json:"invoice_payload"`
	// Optional. Identifier of the shipping option chosen by the user
	ShippingOptionID ShippingOptionID `json:"shipping_option_id,omitempty"`
	// Optional. Order info provided by the user
	OrderInfo *OrderInfo `json:"order_info,omitempty"`
	// Telegram payment identifier
	TelegramPaymentChargeID string `json:"telegram_payment_charge_id"`
	// Provider payment identifier
	ProviderPaymentChargeID string `json:"provider_payment_charge_id"`
}

// This object contains information about an incoming shipping query.
//
// https://core.telegram.org/bots/api#shippingquery
type ShippingQuery struct {
	// Unique query identifier
	ID ShippingQueryID `json:"id"`
	// User who sent the query
	From *User `json:"from"`
	// Bot specified invoice payload
	InvoicePayload string `json:"invoice_payload"`
	// User specified shipping address
	ShippingAddress *ShippingAddress `json:"shipping_address"`
}

// This object contains information about an incoming pre-checkout query.
//
// https://core.telegram.org/bots/api#precheckoutquery
type PreCheckoutQuery struct {
	// Unique query identifier
	ID PreCheckoutQueryID `json:"id"`
	// User who sent the query
	From *User `json:"from"`
	// Three-letter ISO 4217 currency code
	// https://core.telegram.org/bots/payments#supported-currencies
	Currency string `json:"currency"`
	// Total price in the smallest units of the currency (integer, not
	// float/double). For example, for a price of US$ 1.45 pass amount = 145.
	// See the exp parameter in currencies.json, it shows the number of digits
	// past the decimal point for each currency (2 for the majority of
	// currencies). https://core.telegram.org/bots/payments/currencies.json
	TotalAmount int `json:"total_amount"`
	// Bot specified invoice payload
	InvoicePayload string `json:"invoice_payload"`
	// Optional. Identifier of the shipping option chosen by the user
	ShippingOptionID ShippingOptionID `json:"shipping_option_id,omitempty"`
	// Optional. Order info provided by the user
	OrderInfo *OrderInfo `json:"order_info,omitempty"`
}
