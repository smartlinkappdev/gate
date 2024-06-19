package auth

type InterfaceAuth interface {
	Init(dsn string) error

	CreateToken(claims *Claims) (string, error)
	VerifyToken(token string) (*Claims, error)

	Signup(email, password string) (string, int, error)
	Login(email, password string) (string, error)
}
