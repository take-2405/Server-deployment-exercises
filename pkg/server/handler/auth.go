package handler

import (
	"Attendance/pkg/server/model"
	"Attendance/pkg/server/view"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
)

func HandleAuthCreate() gin.HandlerFunc {
	return func(c *gin.Context) {

		// リクエストBodyから情報を取得
		var requestBody view.AuthCreateRequest
		if err := c.ShouldBindJSON(&requestBody); err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, "Request is error")
			return
		}

		// UUIDで認証トークンを生成する
		AuthToken, err := uuid.NewRandom()
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, "AuthToken is error")
			return
		}

		// データベースにユーザデータを登録する
		// TODO: ユーザデータの登録クエリを入力する
		err = model.InsertUser(AuthToken.String(), requestBody.Name, requestBody.StudentNumber)
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusInternalServerError, "DB is error")
			return
		}
		// // 生成した認証トークンを返却
		c.JSON(http.StatusOK, view.AuthCreateResponse{Token: AuthToken.String()})
	}
}
