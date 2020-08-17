package server

import (
	"Attendance/pkg/middleware"
	"Attendance/pkg/server/handler"

	"github.com/gin-gonic/gin"
)

var (
	//Server gin flameworkのserver
	Server *gin.Engine
)

func init() {
	Server = gin.Default()
	Server.GET("/", handler.HandlePing())
	Server.GET("/Auth", handler.HandleAuthCreate())
	Server.GET("/Active", middleware.Authenticate(handler.HandleResistTimeData("active")))
	Server.GET("/BreakIN", middleware.Authenticate(handler.HandleResistTimeData("breakIN")))
	Server.GET("/BreakOUT", middleware.Authenticate(handler.HandleResistTimeData("breakOUT")))
	Server.GET("/Leave", middleware.Authenticate(handler.HandleResistTimeData("leave")))
	//Server.GET("/Attendance", middleware.Authenticate(handler.HandleResistAttend()))
	//Server.GET("/BreakIN", middleware.Authenticate(handler.HandleResistBreakIN()))
	//Server.GET("/BreakOUT", middleware.Authenticate(handler.HandleResistBreakOUT()))
	//Server.GET("/Retired", middleware.Authenticate(handler.HandleResistRetired()))
}
