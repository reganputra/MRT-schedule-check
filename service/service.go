package service

import (
	"encoding/json"
	"errors"
	"mrt-schedule-checker/common/client"
	"mrt-schedule-checker/modules/model"
	"net/http"
	"strings"
	"time"
)

type Service interface {
	GetAllStations() (response []model.StationsResponse, err error)
	CheckSchedule(stationId string) (response []model.ScheduleResponse, err error)
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

func (s *ServiceImplementation) CheckSchedule(stationId string) (response []model.ScheduleResponse, err error) {
	url := "https://jakartamrt.co.id/id/val/stasiuns"

	body, err := client.GetRequest(s.client, url)
	if err != nil {
		return nil, err
	}
	var schedules []model.Schedule
	err = json.Unmarshal(body, &schedules)
	if err != nil {
		return nil, err
	}

	var selectedSchedules model.Schedule
	for _, item := range schedules {
		if item.StationId == stationId {
			selectedSchedules = item
			break
		}
	}

	if selectedSchedules.StationId == "" {
		err = errors.New("Station not found")
		return
	}

	resp, err := ConvertDataToResponse(selectedSchedules)
	if err != nil {
		return
	}

	return resp, nil
}

func ConvertDataToResponse(data model.Schedule) (response []model.ScheduleResponse, err error) {
	var (
		lebakBulusTripTime = "Station Lebak Bulus"
		bundaranHiTripTime = "Station Bundaran HI"
	)
	scheduleLebakBulus := data.ScheduleLb
	scheduleBundaranHi := data.ScheduleHi
	lebakBuluseParsed, err := ConvertScheduleTimeFormat(scheduleLebakBulus)
	if err != nil {
		return nil, err
	}
	bundaranHiParsed, err := ConvertScheduleTimeFormat(scheduleBundaranHi)
	if err != nil {
		return nil, err
	}
	for _, item := range lebakBuluseParsed {
		if item.Format("15:04") > time.Now().Format("15:04") {
			response = append(response, model.ScheduleResponse{
				StationName: lebakBulusTripTime,
				Time:        item.Format("15:04"),
			})
		}

	}

	for _, item := range bundaranHiParsed {
		if item.Format("15:04") > time.Now().Format("15:04") {
			response = append(response, model.ScheduleResponse{
				StationName: bundaranHiTripTime,
				Time:        item.Format("15:04"),
			})
		}

	}

	return response, nil

}

func ConvertScheduleTimeFormat(schedule string) (response []time.Time, err error) {
	var (
		parsedTime time.Time
		schedules  = strings.Split(schedule, ",")
	)
	for _, item := range schedules {
		trimedSpace := strings.TrimSpace(item)
		if trimedSpace == "" {
			continue
		}

		// Try to parse with the primary format
		parsedTime, err = time.Parse("15:04", trimedSpace)
		if err == nil {
			response = append(response, parsedTime)
			continue
		}

		// Try alternative formats if needed
		formats := []string{"3:04 PM", "3:04PM", "15.04"}
		parsed := false

		for _, format := range formats {
			parsedTime, err = time.Parse(format, trimedSpace)
			if err == nil {
				response = append(response, parsedTime)
				parsed = true
				break
			}
		}

		// If all parsing attempts fail, skip this time value and continue
		if !parsed {
			continue
		}
	}

	// If no valid time values were found, return an error
	if len(response) == 0 {
		return nil, errors.New("No valid time values found in schedule")
	}

	return response, nil
}
