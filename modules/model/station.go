package model

type StationsResponse struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Station struct {
	Id   string `json:"nid"`
	Name string `json:"title"`
}
