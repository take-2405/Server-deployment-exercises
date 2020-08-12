package timedata

import (
	"Attendance/pkg/server/view"
	"time"
)

func CreateTimeData() view.TimeData {
	//const layout = "2006-01-02 15:04"
	//const layout = "2006-01-02"
	var timeInformation view.TimeData
	time := time.Now()
	//yearDate:=time.Format(layout)
	timeInformation.Year = time.Year()
	timeInformation.Month = int(time.Month())
	timeInformation.Day = time.Day()
	timeInformation.Hour = time.Hour()
	timeInformation.Minute = time.Minute()
	return timeInformation
}
