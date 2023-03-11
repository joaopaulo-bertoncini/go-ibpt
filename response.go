package ibpt

import (
	"encoding/json"
	"errors"
)

var (
	ErrMissingRegistration = errors.New("missing registration token")
)

// Response represents the FCM server's response to the application
// server's sent message.
type Response struct {
	Code          string  `json:"Codigo"`
	UF            string  `json:"UF"`
	EX            int     `json:"EX"`
	Description   string  `json:"Descricao"`
	National      float64 `json:"Nacional"`
	State         float64 `json:"Estadual"`
	Imported      float64 `json:"Importado"`
	Municipal     float64 `json:"Municipal"`
	Type          string  `json:"Tipo"`
	BeginningTerm string  `json:"VigenciaInicio"`
	TermEnd       string  `json:"VigenciaFim"`
	Key           string  `json:"Chave"`
	Version       string  `json:"Versao"`
	Source        string  `json:"Fonte"`
}

// UnmarshalJSON implements json.Unmarshaler interface.
func (r *Response) unmarshalJSON(data []byte) error {
	var response struct {
		Code          string  `json:"Codigo"`
		UF            string  `json:"UF"`
		EX            int     `json:"EX"`
		Description   string  `json:"Descricao"`
		National      float64 `json:"Nacional"`
		State         float64 `json:"Estadual"`
		Imported      float64 `json:"Importado"`
		Municipal     float64 `json:"Municipal"`
		Type          string  `json:"Tipo"`
		BeginningTerm string  `json:"VigenciaInicio"`
		TermEnd       string  `json:"VigenciaFim"`
		Key           string  `json:"Chave"`
		Version       string  `json:"Versao"`
		Source        string  `json:"Fonte"`
	}

	if err := json.Unmarshal(data, &response); err != nil {
		return err
	}

	r.Code = response.Code
	r.UF = response.UF
	r.EX = response.EX
	r.Description = response.Description
	r.National = response.National
	r.State = response.State
	r.Imported = response.Imported
	r.Municipal = response.Municipal
	r.Type = response.Type
	r.BeginningTerm = response.BeginningTerm
	r.TermEnd = response.TermEnd
	r.Key = response.Key
	r.Version = response.Version
	r.Source = response.Source
	return nil
}
