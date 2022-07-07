package model

type IPDetails struct {
	IP       string `json:"ip"`
	IpSource string `json:"ip_source"`
}

type RegionData struct {
	Status    string `json:"status"`
	Info      string `json:"info"`
	Infocode  string `json:"infocode"`
	Province  string `json:"province"`
	City      string `json:"city"`
	Adcode    string `json:"adcode"`
	Rectangle string `json:"rectangle"`
}
