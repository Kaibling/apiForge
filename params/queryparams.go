package params

type Pagination struct {
	Filter string  `json:"filter,omitempty"`
	Limit  int     `json:"limit"`
	Order  string  `json:"order"`
	After  *string `json:"after"`
	Before *string `json:"before"`
}
