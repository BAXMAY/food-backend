package models

type SignUp struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
	UserRole string `json:"user_role" bson:"user_role"`
}

type SignIn struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}
