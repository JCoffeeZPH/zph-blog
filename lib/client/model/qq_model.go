package model

type QQImageModel struct {
	Data []byte `json:"data"`
	Type string `json:"type"`
}

type QQDetails struct {
	QQAvatarCDNUrl string `json:"qq"`
	Nickname       string `json:"nickname"`
}
