package models

type Suggestion struct {
	IdSuggestion int    `json:"idSuggestion"`
	IdUsers      int    `json:"idUsers"`
	Suggestion   string `json:"suggestion"`
}
