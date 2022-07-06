package telegrambot

// https://core.telegram.org/bots/api

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"sort"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/nickname76/repeater"
	"github.com/valyala/fasthttp"
)

// URL of official Telegram Bot API endpoint
const DefaultAPIEndpointURL = "https://api.telegram.org/bot"

var defaultFasthttpClient = &fasthttp.Client{
	NoDefaultUserAgentHeader:      true,
	DisableHeaderNamesNormalizing: true,
	DisablePathNormalizing:        true,
}

// Default function for performing http requests by API
func DefaultHttpDoRequest(method string, url string, headers map[string]string, body []byte) (respBody []byte, err error) {
	req := &fasthttp.Request{}

	req.Header.SetMethod(method)
	req.SetRequestURI(url)

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	req.SetBody(body)

	resp := &fasthttp.Response{}

	err = defaultFasthttpClient.Do(req, resp)
	if err != nil {
		return nil, fmt.Errorf("DefaultHttpDoRequest: %w", err)
	}

	respBody = resp.Body()

	return respBody, nil
}

// Main object in this library, for performing Telegram Bot API requests
type API struct {
	Token         string
	EndpointURL   string
	HttpDoRequest func(method string, url string, headers map[string]string, body []byte) (respBody []byte, err error)
}

// Creates Telegram Bot API interface instance. If you want to customize http
// requests behavior or api endpoint url (e.x. use local instance), then
// instance API struct directly
func NewAPI(token string) (*API, *User, error) {
	api := &API{
		Token:         token,
		EndpointURL:   DefaultAPIEndpointURL,
		HttpDoRequest: DefaultHttpDoRequest,
	}

	user, err := api.GetMe()
	if err != nil {
		return nil, nil, fmt.Errorf("NewAPI: %w", err)
	}

	return api, user, nil
}

// Response on API request. Used internally by this library.
type Response struct {
	OK          bool   `json:"ok"`
	ErrorCode   int    `json:"error_code,omitempty"`
	Description string `json:"description,omitempty"`

	Parameters *ResponseParameters `json:"parameters,omitempty"`

	Result any `json:"result,omitempty"`
}

func (api *API) makeAPICall(method string, requestData any, inputFiles []InputFile, resultDest any) (migrateToChatID ChatID, err error) {
	var (
		reqURL         = api.EndpointURL + api.Token + "/" + method
		reqContentType string
		reqBody        []byte
	)

	jsoniterCfg := jsoniter.Config{
		OnlyTaggedField:               true,
		ObjectFieldMustBeSimpleString: true,
		CaseSensitive:                 true,
	}.Froze()

	requestDataJSON, err := jsoniterCfg.Marshal(requestData)
	if err != nil {
		return 0, fmt.Errorf("makeAPICall: %w", err)
	}

	if inputFilesToUpload := filterInputFilesNeedingUpload(inputFiles); len(inputFilesToUpload) == 0 {
		reqContentType = "application/json"
		reqBody = requestDataJSON
	} else {
		reqBodyBuf := bytes.NewBuffer(nil)

		mw := multipart.NewWriter(reqBodyBuf)

		var err error
		iter := jsoniterCfg.BorrowIterator(requestDataJSON)
		iter.ReadMapCB(func(i *jsoniter.Iterator, s string) bool {
			err = mw.WriteField(s, i.ReadAny().ToString())
			return err == nil
		})
		if err != nil {
			return 0, fmt.Errorf("makeAPICall: %w", err)
		}

		for _, inputFile := range inputFilesToUpload {
			fieldname, filename, reader := inputFile.multipartFormFile()
			filew, err := mw.CreateFormFile(fieldname, filename)
			if err != nil {
				return 0, fmt.Errorf("makeAPICall: %w", err)
			}

			_, err = io.Copy(filew, reader)
			if err != nil {
				return 0, fmt.Errorf("makeAPICall: %w", err)
			}
		}

		err = mw.Close()
		if err != nil {
			return 0, fmt.Errorf("makeAPICall: %w", err)
		}

		reqContentType = mw.FormDataContentType()
		reqBody = reqBodyBuf.Bytes()
	}

loop:
	for {
		respBody, err := api.HttpDoRequest("POST", reqURL, map[string]string{
			"Content-Type": reqContentType,
		}, reqBody)
		if err != nil {
			return 0, fmt.Errorf("makeAPICall: %w", err)
		}

		apiResp := &Response{
			Result: resultDest,
		}

		err = jsoniterCfg.Unmarshal(respBody, apiResp)
		if err != nil {
			return 0, fmt.Errorf("makeAPICall: %w", err)
		}

		if !apiResp.OK {
			if apiRespParams := apiResp.Parameters; apiRespParams != nil {
				switch {
				case apiRespParams.MigrateToChatID != 0:
					return apiRespParams.MigrateToChatID, nil
				case apiRespParams.RetryAfter != 0:
					time.Sleep(time.Second * time.Duration(apiRespParams.RetryAfter))
					continue loop
				}
			}

			return 0, fmt.Errorf("makeAPICall - telegram bot api error: %w", errors.New(apiResp.Description))
		}

		return 0, nil
	}
}

func filterInputFilesNeedingUpload(inputFiles []InputFile) []InputFile {
	if inputFiles == nil {
		return nil
	}

	filteredInputFiles := []InputFile{}

	for _, inputFile := range inputFiles {
		if inputFile == nil {
			continue
		}

		if _, _, reader := inputFile.multipartFormFile(); reader != nil {
			filteredInputFiles = append(filteredInputFiles, inputFile)
		}
	}

	return filteredInputFiles
}

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
