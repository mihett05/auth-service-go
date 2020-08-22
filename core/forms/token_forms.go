package forms

type TokenForm struct {
	AccessToken string `json:"access_token" binding:"required"`
}
