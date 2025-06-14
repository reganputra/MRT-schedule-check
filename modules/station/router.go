package station

import (
	"github.com/gin-gonic/gin"
	"mrt-schedule-checker/common/response"
	"mrt-schedule-checker/service"
	"net/http"
)

func Intialize(router *gin.RouterGroup) {

	stationService := service.NewServiceImpl()

	station := router.Group("/station")
	station.GET("", func(c *gin.Context) {
		GetAllStations(c, stationService)
	})
}

func GetAllStations(c *gin.Context, implementation *service.ServiceImplementation) {
	resp, err := implementation.GetAllStations()
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.ApiResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		})
		return
	}

	c.JSON(http.StatusOK, response.ApiResponse{
		Success: true,
		Message: "Successfully retrieved all stations",
		Data:    resp,
	})
}
