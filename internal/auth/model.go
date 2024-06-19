package auth

type UserAuth struct {
	ID    int    `json:"id" gorm:"primary_key"`
	Hash  string `json:"hash"`
	Salt  string `json:"salt"`
	Email string `json:"email"`
}

type Claims struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}
