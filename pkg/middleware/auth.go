package middleware

import (
	"log"
	"net/http"

	"Attendance/pkg/server/model"

	"github.com/gin-gonic/gin"
)

// Authenticate ユーザ認証を行ってContextへユーザID情報を保存する
func Authenticate(ginNextMethod gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		// リクエストヘッダからX-API-TOKEN(認証トークン)を取得
		key := "X-API-TOKEN"
		token := c.GetHeader(key)
		if len(token) == 0 {
			log.Println("X-API-TOKEN is empty")
			return
		}

		// データベースから認証トークンに紐づくユーザの情報を取得
		user, err := model.SelectUserByAuthToken(token)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, "User select is error.")
			return
		}
		if user == nil {
			log.Printf("user not found. token=%s", token)
			c.JSON(http.StatusBadRequest, "User is nil")
			return
		}
		// ユーザIDをContextへ保存して以降の処理に利用する
		c.Set("studentNumber", user.StudentNumber)
		ginNextMethod(c)
	}
}
