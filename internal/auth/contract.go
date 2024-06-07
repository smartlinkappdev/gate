package auth

type InterfaceAuth interface {
	Init(dsn string) error

	CreateToken(claims *Claims) (string, error)
	VerifyToken(token string) (*Claims, error)

	Signup(username, password string) (string, int, error)
	Login(username, password string) (string, error)
}
