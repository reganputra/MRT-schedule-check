package model

type StationsResponse struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Station struct {
	Id   string `json:"nid"`
	Name string `json:"title"`
}

type Schedule struct {
	StationId   string `json:"nid"`
	StationName string `json:"title"`
	ScheduleHi  string `json:"jadwal_hi_biasa"`
	ScheduleLb  string `json:"jadwal_lb_biasa"`
}

type ScheduleResponse struct {
	StationName string `json:"station_name"`
	Time        string `json:"departure_time"`
}
