package common

type ItemUseData struct {
	UseOp    string   `json:"useOp,omitempty"` // enum ItemUseOp
	UseParam []string `json:"useParam,omitempty"`
}
