package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sandisyd/schedule-mrt/modules/station"
)


func main(){
	initiateRouter()
}

func initiateRouter(){
	var (
		router = gin.Default() 
	api = router.Group("/v1/api"))

	station.Initial(api)
	router.Run(":8182")
}