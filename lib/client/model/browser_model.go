package model

type BrowserDetail struct {
	Code int `json:"code"`
	Message string `json:"message"`
	Data map[string]string `json:"data"`
}

