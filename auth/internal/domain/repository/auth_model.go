package repository

type AuthModel struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (AuthModel) TableName() string {
	return "auths"
}
