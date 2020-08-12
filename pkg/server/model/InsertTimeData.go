package model

import (
	"Attendance/pkg/db"
)

type TimeData struct {
	Year   int
	Month  int
	Day    int
	Hour   int
	Minute int
}

func InsertAttendanceTime(timeID, studentNumber, state string, timeInformation TimeData) error {
	stmt, err := db.Conn.Prepare("INSERT INTO timeManagement VALUES (?,?,?,?,?,?,?,?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(timeID, studentNumber, state, timeInformation.Year, timeInformation.Month, timeInformation.Day, timeInformation.Hour, timeInformation.Minute)
	return err
}
