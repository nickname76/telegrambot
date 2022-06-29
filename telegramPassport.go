package telegrambot

// https://core.telegram.org/bots/api#telegram-passport

import "fmt"

// Contains information about Telegram Passport data shared with the bot by the
// user.
//
// https://core.telegram.org/bots/api#passportdata
type PassportData struct {
	// Array with information about documents and other Telegram Passport
	// elements that was shared with the bot
	Data []*EncryptedPassportElement `json:"data"`
	// Encrypted credentials required to decrypt the data
	Credentials *EncryptedCredentials `json:"credentials"`
}

// This object represents a file uploaded to Telegram Passport. Currently all
// Telegram Passport files are in JPEG format when decrypted and don't exceed
// 10MB.
//
// https://core.telegram.org/bots/api#passportfile
type PassportFile struct {
	// Identifier for this file, which can be used to download or reuse the file
	FileID FileID `json:"file_id"`
	// Unique identifier for this file, which is supposed to be the same over
	// time and for different bots. Can't be used to download or reuse the file.
	FileUniqueID FileUniqueID `json:"file_unique_id"`
	// File size in bytes
	FileSize int64 `json:"file_size"`
	// Unix time when the file was uploaded
	FileDate int64 `json:"file_date"`
}

// Contains information about documents or other Telegram Passport elements
// shared with the bot by the user.
//
// https://core.telegram.org/bots/api#encryptedpassportelement
type EncryptedPassportElement struct {
	// Element type. One of "personal_details", "passport", "driver_license",
	// "identity_card", "internal_passport", "address", "utility_bill",
	// "bank_statement", "rental_agreement", "passport_registration",
	// "temporary_registration", "phone_number", "email".
	Type PassportElementType `json:"type"`
	// Optional. Base64-encoded encrypted Telegram Passport element data
	// provided by the user, available for "personal_details", "passport",
	// "driver_license", "identity_card", "internal_passport" and "address"
	// types. Can be decrypted and verified using the accompanying
	// EncryptedCredentials.
	// https://core.telegram.org/bots/api#encryptedcredentials
	Data string `json:"name,omitempty"`
	// Optional. User's verified phone number, available only for "phone_number"
	// type
	PhoneNumber string `json:"phone_number,omitempty"`
	// Optional. User's verified email address, available only for "email" type
	Email string `json:"email,omitempty"`
	// Optional. Array of encrypted files with documents provided by the user,
	// available for "utility_bill", "bank_statement", "rental_agreement",
	// "passport_registration" and "temporary_registration" types. Files can be
	// decrypted and verified using the accompanying EncryptedCredentials.
	// https://core.telegram.org/bots/api#encryptedcredentials
	Files []*PassportFile `json:"files,omitempty"`
	// Optional. Encrypted file with the front side of the document, provided by
	// the user. Available for "passport", "driver_license", "identity_card" and
	// "internal_passport". The file can be decrypted and verified using the
	// accompanying EncryptedCredentials.
	FrontSide *PassportFile `json:"front_side,omitempty"`
	// Optional. Encrypted file with the reverse side of the document, provided
	// by the user. Available for "driver_license" and "identity_card". The file
	// can be decrypted and verified using the accompanying
	// EncryptedCredentials.
	// https://core.telegram.org/bots/api#encryptedcredentials
	ReverseSide *PassportFile `json:"reverse_side,omitempty"`
	// Optional. Encrypted file with the selfie of the user holding a document,
	// provided by the user; available for "passport", "driver_license",
	// "identity_card" and "internal_passport". The file can be decrypted and
	// verified using the accompanying EncryptedCredentials.
	// https://core.telegram.org/bots/api#encryptedcredentials
	Selfie *PassportFile `json:"selfie,omitempty"`
	// Optional. Array of encrypted files with translated versions of documents
	// provided by the user. Available if requested for "passport",
	// "driver_license", "identity_card", "internal_passport", "utility_bill",
	// "bank_statement", "rental_agreement", "passport_registration" and
	// "temporary_registration" types. Files can be decrypted and verified using
	// the accompanying EncryptedCredentials.
	// https://core.telegram.org/bots/api#encryptedcredentials
	Translation []*PassportFile `json:"translation,omitempty"`
	// Base64-encoded element hash for using in PassportElementErrorUnspecified
	// https://core.telegram.org/bots/api#passportelementerrorunspecified
	Hash string `json:"hash"`
}

// Contains data required for decrypting and authenticating
// EncryptedPassportElement. See the Telegram Passport Documentation for a
// complete description of the data decryption and authentication processes.
// https://core.telegram.org/bots/api#encryptedpassportelement
// https://core.telegram.org/passport#receiving-information
//
// https://core.telegram.org/bots/api#encryptedcredentials
type EncryptedCredentials struct {
	// Base64-encoded encrypted JSON-serialized data with unique user's payload,
	// data hashes and secrets required for EncryptedPassportElement decryption
	// and authentication
	// https://core.telegram.org/bots/api#encryptedpassportelement
	Data string `json:"data"`
	// Base64-encoded data hash for data authentication
	Hash string `json:"hash"`
	// Base64-encoded secret, encrypted with the bot's public RSA key, required
	// for data decryption
	Secret string `json:"secret"`
}

type SetPassportDataErrorsParams struct {
	// User identifier
	UserID UserID `json:"user_id"`
	// A JSON-serialized array describing the errors
	Errors []*PassportElementError `json:"errors"`
}

// Informs a user that some of the Telegram Passport elements they provided
// contains errors. The user will not be able to re-submit their Passport to you
// until the errors are fixed (the contents of the field for which you returned
// the error must change). Returns True on success.
//
// Use this if the data submitted by the user doesn't satisfy the standards your
// service requires for any reason. For example, if a birthday date seems
// invalid, a submitted document is blurry, a scan shows evidence of tampering,
// etc. Supply some details in the error message to make sure the user knows how
// to correct the issues.
//
// https://core.telegram.org/bots/api#setpassportdataerrors
func (api *API) SetPassportDataErrors(params *SetPassportDataErrorsParams) error {
	_, err := api.makeAPICall("setPassportDataErrors", params, nil, nil)
	if err != nil {
		return fmt.Errorf("SetPassportDataErrors: %w", err)
	}

	return nil
}

// This object represents an error in the Telegram Passport element which was
// submitted that should be resolved by the user. It should be one of:
//   PassportElementErrorDataField - Represents an issue in one of the data fields that was provided by the user. The error is considered resolved when the field's value changes.
//   PassportElementErrorFrontSide - Represents an issue with the front side of a document. The error is considered resolved when the file with the front side of the document changes.
//   PassportElementErrorReverseSide - Represents an issue with the reverse side of a document. The error is considered resolved when the file with reverse side of the document changes.
//   PassportElementErrorSelfie - Represents an issue with the selfie with a document. The error is considered resolved when the file with the selfie changes.
//   PassportElementErrorFile - Represents an issue with a document scan. The error is considered resolved when the file with the document scan changes.
//   PassportElementErrorFiles - Represents an issue with a list of scans. The error is considered resolved when the list of files containing the scans changes.
//   PassportElementErrorTranslationFile - Represents an issue with one of the files that constitute the translation of a document. The error is considered resolved when the file changes.
//   PassportElementErrorTranslationFiles - Represents an issue with the translated version of a document. The error is considered resolved when a file with the document translation change.
//   PassportElementErrorUnspecified - Represents an issue in an unspecified place. The error is considered resolved when new data is added.
//
// https://core.telegram.org/bots/api#passportelementerror
// https://core.telegram.org/bots/api#passportelementerrordatafield
// https://core.telegram.org/bots/api#passportelementerrorfrontside
// https://core.telegram.org/bots/api#passportelementerrorreverseside
// https://core.telegram.org/bots/api#passportelementerrorselfie
// https://core.telegram.org/bots/api#passportelementerrorfile
// https://core.telegram.org/bots/api#passportelementerrorfiles
// https://core.telegram.org/bots/api#passportelementerrortranslationfile
// https://core.telegram.org/bots/api#passportelementerrortranslationfiles
// https://core.telegram.org/bots/api#passportelementerrorunspecified
type PassportElementError struct {
	// Error source
	//   PassportElementErrorDataField - must be data
	//   PassportElementErrorFrontSide - must be front_side
	//   PassportElementErrorReverseSide - must be reverse_side
	//   PassportElementErrorSelfie - must be selfie
	//   PassportElementErrorFile - must be file
	//   PassportElementErrorFiles - must be files
	//   PassportElementErrorTranslationFile - must be translation_file
	//   PassportElementErrorTranslationFiles - must be translation_files
	//   PassportElementErrorUnspecified - must be unspecified
	Source PassportElementErrorSource `json:"source"`
	// Error message
	Message string `json:"message"`

	// The section of the user's Telegram Passport which has the error
	//   PassportElementErrorDataField - one of "personal_details", "passport", "driver_license", "identity_card", "internal_passport", "address"
	//   PassportElementErrorFrontSide - one of "passport", "driver_license", "identity_card", "internal_passport"
	//   PassportElementErrorReverseSide - one of "driver_license", "identity_card"
	//   PassportElementErrorSelfie - one of "passport", "driver_license", "identity_card", "internal_passport"
	//   PassportElementErrorFile - one of "utility_bill", "bank_statement", "rental_agreement", "passport_registration", "temporary_registration"
	//   PassportElementErrorFiles - one of "utility_bill", "bank_statement", "rental_agreement", "passport_registration", "temporary_registration"
	//   PassportElementErrorTranslationFile - one of "passport", "driver_license", "identity_card", "internal_passport", "utility_bill", "bank_statement", "rental_agreement", "passport_registration", "temporary_registration"
	//   PassportElementErrorTranslationFiles - one of "passport", "driver_license", "identity_card", "internal_passport", "utility_bill", "bank_statement", "rental_agreement", "passport_registration", "temporary_registration"
	Type PassportElementType `json:"type"`

	// Name of the data field which has the error
	FieldName string `json:"field_name,omitempty"`
	// Base64-encoded data hash
	DataHash string `json:"data_hash,omitempty"`

	// Base64-encoded file hash
	FileHash string `json:"file_hash,omitempty"`

	// List of base64-encoded file hashes
	FileHashes []string `json:"file_hashes,omitempty"`

	// Base64-encoded element hash
	ElementHash string `json:"element_hash,omitempty"`
}
