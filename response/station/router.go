package station

import "github.com/gin-gonic/gin"

func Intialize(router *gin.RouterGroup) {
	station := router.Group("/station")
	station.GET("")
}
