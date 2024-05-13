package structs

type Body struct {
	Type        string `json:"type,omitempty"`
	Destination string `json:"destination,omitempty"`
	Origin      string `json:"origin,omitempty"`
	Amount      int    `json:"amount,omitempty"`
}
