  
package view

//AuthCreateRequest リクエスト
type AuthCreateRequest struct {
	StudentNumber string `form:"studentNumber" json:"studentNumber" xml:"studentNumber"`
	Name string `form:"name" json:"name" xml:"name"`
}

//AuthCreateResponse レスポンス
type AuthCreateResponse struct {
	Token string `json:"token"`
}