package models

type Renew struct {
	RefreshToken string `json:"refresh_token" bson:"refresh_token"`
}
