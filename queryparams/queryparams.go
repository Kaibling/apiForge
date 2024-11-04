package queryparams

type QueryParams struct {
	Filter string `json:"filter"`
	Limit  int    `json:"limit"`
	Order  string `json:"order"`
}
