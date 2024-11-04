package queryparams

type QueryParams struct {
	Filter string `json:"filter"`
	Limit  int    `json:"limit"`
	Order  string `json:"order"`
	After  string `json:"after"`
	Before string `json:"before"`
}
