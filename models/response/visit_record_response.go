package response

type VisitRecord struct {
	Date []string `json:"date"`
	PV []int `json:"pv"`
	UV []int `json:"uv"`
}
