package ibpt

import (
	"errors"
)

var (
	ErrInvalidMessage         = errors.New("message is requerid")
	ErrInvalidAPIKey          = errors.New("API Key is requerid")
	ErrInvalidCnpj            = errors.New("CNPJ is requerid")
	ErrInvalidCode            = errors.New("Code is requerid")
	ErrInvalidUF              = errors.New("UF is requerid")
	ErrInvalidEX              = errors.New("EX is requerid")
	ErrInvalidDescription     = errors.New("Description is requerid")
	ErrInvalidUnitMeasurement = errors.New("Unit Measurement is requerid")
	ErrInvalidGtin            = errors.New("Gtin is requerid")
)

// Message represents list of targets, options, and payload for HTTP JSON
// messages.
type Request struct {
	Token           string
	CNPJ            string
	Code            string
	UF              string
	EX              int
	InternalCode    string
	Description     string
	UnitMeasurement string
	Value           float64
	Gtin            string
}

// Validate returns an error if the message is not well-formed.
func (msg *Request) Validate() error {
	if msg == nil {
		return ErrInvalidMessage
	}

	if msg.Token == "" {
		return ErrInvalidAPIKey
	}

	if msg.CNPJ == "" {
		return ErrInvalidCnpj
	}

	if msg.Code == "" {
		return ErrInvalidCode
	}

	if msg.UF == "" && len(msg.Code) > 2 {
		return ErrInvalidUF
	}

	if msg.Description == "" {
		return ErrInvalidDescription
	}

	if msg.UnitMeasurement == "" {
		return ErrInvalidUnitMeasurement
	}

	return nil
}
