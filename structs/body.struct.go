package structs

type Body struct {
	Type        string `json:"type"`
	Destination int    `json:"destination"`
	Amount      int    `json:"amount"`
}
