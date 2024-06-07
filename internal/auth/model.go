package auth

type UserAuth struct {
	ID       int    `json:"id" gorm:"primary_key"`
	Hash     string `json:"hash"`
	Salt     string `json:"salt"`
	Username string `json:"username"`
}

type Claims struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}
