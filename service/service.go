package service

import (
	"encoding/json"
	"mrt-schedule-checker/common/client"
	"mrt-schedule-checker/modules/model"
	"net/http"
	"time"
)

type Service interface {
	GetAllStations() (response []model.StationsResponse, err error)
}

type ServiceImplementation struct {
	client *http.Client
}

func NewServiceImpl() *ServiceImplementation {
	return &ServiceImplementation{
		client: &http.Client{
			Timeout: 10 * time.Second, // Set a suitable timeout
		},
	}
}

func (s *ServiceImplementation) GetAllStations() (response []model.StationsResponse, err error) {
	url := "https://jakartamrt.co.id/id/val/stasiuns"

	body, err := client.GetRequest(s.client, url)
	if err != nil {
		return nil, err
	}
	var stations []model.Station
	err = json.Unmarshal(body, &stations)
	if err != nil {
		return nil, err
	}

	response = make([]model.StationsResponse, len(stations))
	for i, item := range stations {
		response[i] = model.StationsResponse{
			Id:   item.Id,
			Name: item.Name,
		}
	}

	return response, nil
}
