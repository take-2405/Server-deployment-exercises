package handler

import (
	"Attendance/pkg/server/model"
	"Attendance/pkg/server/view"
	"Attendance/pkg/timedata"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
)

func HandleResistTimeData(status string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var Response view.InsertTimeDataResponse
		var modelTimeData model.TimeData
		// Contextから認証済みのユーザIDを取得
		studentNumber := c.GetString("studentNumber")
		//fmt.Println(studentNumber)
		if len(studentNumber) == 0 {
			log.Println(errors.New("studentNumber is empty"))
			c.JSON(http.StatusInternalServerError, "Internal Server Error")
			return
		}

		//var timeInformation [5]int
		//timeInformation =timedata.CreateTimeData()
		//レスポンス用にviewの構造体に挿入
		Response.Timedata = timedata.CreateTimeData()
		//DBにも必要なので変換する
		modelTimeData = convertingViewAndModelStructure(Response.Timedata)

		//UUIDでユーザIDを生成する
		timeID, err := uuid.NewRandom()
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, "timeID is error")
			return
		}
		Response.TimeID = timeID.String()
		Response.Status = status
		if err := model.InsertAttendanceTime(timeID.String(), studentNumber, status, modelTimeData); err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, "Internal Server Error")
			return
		}
		c.JSON(http.StatusOK, view.InsertTimeDataResponse{TimeID: Response.TimeID, Status: Response.Status, Timedata: Response.Timedata})
		return
	}
}

func convertingViewAndModelStructure(viewTimedata view.TimeData) model.TimeData {
	var modelTimedata model.TimeData
	modelTimedata.Year = viewTimedata.Year
	modelTimedata.Month = viewTimedata.Month
	modelTimedata.Day = viewTimedata.Day
	modelTimedata.Hour = viewTimedata.Hour
	modelTimedata.Minute = viewTimedata.Minute
	return modelTimedata
}
