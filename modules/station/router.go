package station

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sandisyd/schedule-mrt/common/response"
)

func Initial(router *gin.RouterGroup){
	stationService := NewService()
station := router.Group("/stations")
station.GET("",func(c  *gin.Context){
	GetAllStation(c, stationService)
})

station.GET("/:id", func(c *gin.Context) {
	CheckScheduleByStation(c, stationService)
})
}

func GetAllStation(c *gin.Context, service Service){
	datas, err := service.GetAllStation()
	if err != nil {
		c.JSON(
			http.StatusBadRequest, response.APIResponse{
				Success: false,
				Message: "Not found " + err.Error(),
				Data: nil,
			},
		)
		return 
	}

	//handle response 
	c.JSON(
		http.StatusOK, response.APIResponse{
			Success: true,
			Message: "Success get response stations",
			Data: datas,
		},
	)
}

func CheckScheduleByStation(c *gin.Context, service Service){
	datas, err := service.CheckScheduleByStation()

	if err != nil {
		c.JSON(
			http.StatusBadRequest, response.APIResponse{
				Success: false,
				Message: "Not found " + err.Error(),
				Data: nil,
			},
		)
		return 
	}
	//handle response 
	c.JSON(
		http.StatusOK, response.APIResponse{
			Success: true,
			Message: "Success get response schedule by station",
			Data: datas,
		},
	)
}