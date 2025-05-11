package dto

type Request struct {
	Urls    []string `json:"urls"`
	Audio   bool     `json:"audio"`
	Quality bool     `json:"quality"`
}
