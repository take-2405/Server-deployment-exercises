package model

import (
	"Attendance/pkg/db"
	"database/sql"
	"log"
)

// User userテーブルデータ
type User struct {
	StudentNumber string
	AuthToken string
	Name      string
}

func InsertUser(AuthToken,Name,StudentNumber string)error{
	stmt, err := db.Conn.Prepare("INSERT INTO users VALUES (?,?,?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(StudentNumber,AuthToken,Name)
	return err
}

// SelectUserByAuthToken auth_tokenを条件にレコードを取得する
func SelectUserByAuthToken(authToken string) (*User, error) {
	// TODO: auth_tokenを条件にSELECTを行うSQLを第1引数に入力する
	row := db.Conn.QueryRow("SELECT * FROM users WHERE auth_token=?", authToken)
	return convertToUser(row)
}

// convertToUser rowデータをUserデータへ変換する
func convertToUser(row *sql.Row) (*User, error) {
	user := User{}
	if err := row.Scan(&user.StudentNumber, &user.AuthToken, &user.Name); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Println(err)
		return nil, err
	}
	return &user, nil
}