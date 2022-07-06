package tbtools

import "github.com/nickname76/telegrambot"

// Stores keyboard with handlers for each buttons.
// See methods for this type for more information.
type ReplyKeyboardHandler [][]ReplyKeyboardHandlerButton

// Button for ReplyKeyboardHandler
type ReplyKeyboardHandlerButton struct {
	// Button text
	Text string
	// Used by HandleMessageKeyboardButton.
	// Must not be nil.
	Handler func() error
}

// Returns reply markup with pre-defined options composed from ReplyKeyboardHandler
func (rkh ReplyKeyboardHandler) ReplyMarkup() telegrambot.ReplyMarkup {
	return rkh.ReplyMarkupWithOptions(telegrambot.ReplyKeyboardMarkup{
		ResizeKeyboard: true,
		Selective:      true,
	})
}

// Returns reply markup with passed options composed from ReplyKeyboardHandler.
// You should not pass keyboard field in options.
func (rkh ReplyKeyboardHandler) ReplyMarkupWithOptions(options telegrambot.ReplyKeyboardMarkup) telegrambot.ReplyMarkup {
	keyboard := [][]*telegrambot.KeyboardButton{}

	for _, row := range rkh {
		keyboardRow := []*telegrambot.KeyboardButton{}
		for _, button := range row {
			keyboardRow = append(keyboardRow, &telegrambot.KeyboardButton{
				Text: button.Text,
			})
		}

		keyboard = append(keyboard, keyboardRow)
	}

	return &telegrambot.ReplyKeyboardMarkup{
		Keyboard:              keyboard,
		ResizeKeyboard:        options.ResizeKeyboard,
		OneTimeKeyboard:       options.OneTimeKeyboard,
		InputFieldPlaceholder: options.InputFieldPlaceholder,
		Selective:             options.Selective,
	}
}

// Runs handler for a button stored in ReplyKeyboardHandler.
// If no handler found, returns handled == false
func (rkh ReplyKeyboardHandler) HandleMessageKeyboardButton(msg *telegrambot.Message) (handled bool, err error) {
	if msg == nil || msg.Text == "" {
		return false, nil
	}

	text := msg.Text

	for _, row := range rkh {
		for _, button := range row {
			if button.Text == text {
				return true, button.Handler()
			}
		}
	}

	return false, nil
}
