package auth

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Auth struct {
	conn *gorm.DB
}

func New() InterfaceAuth {
	return &Auth{}
}

func (a *Auth) Init(dsn string) error {
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	//conn.AutoMigrate(&UserAuth{})
	a.conn = conn
	return nil
}

var mySigningKey = []byte("secret")

// CreateToken создает JWT
func (a *Auth) CreateToken(claims *Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": claims.Email,
		"id":    claims.ID,
		"exp":   time.Now().Add(time.Hour * 24).Unix(), // Токен, действительный сутки
	})
	return token.SignedString(mySigningKey)
}

// VerifyToken проверяет JWT
func (a *Auth) VerifyToken(tokenString string) (*Claims, error) {
	fmt.Println("tokenString:", tokenString)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return mySigningKey, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims)

		c := Claims{
			Email: claims["email"].(string),
			ID:    int(claims["id"].(float64)),
		}

		fmt.Println(c)
		return &c, nil
	} else {
		return nil, err
	}
}

func (a *Auth) Signup(email, password string) (string, int, error) {
	salt, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", 0, fmt.Errorf("failed to generate salt: %v", err)
	}

	// Хэширование пароля с солью
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", 0, fmt.Errorf("failed to hash password: %v", err)
	}

	// Создание пользователя с хэшированным паролем и солью
	user := UserAuth{
		Hash:  string(hashedPassword),
		Salt:  string(salt),
		Email: email,
	}

	a.conn.Create(&user)

	token, err := a.CreateToken(&Claims{
		ID:    user.ID,
		Email: email,
	})
	if err != nil {
		return "", 0, err
	}

	return token, user.ID, err
}
func (a *Auth) Login(email, password string) (string, error) {
	user := UserAuth{}
	a.conn.Model(&UserAuth{}).Where("email = ?", email).Find(&user)
	err := bcrypt.CompareHashAndPassword([]byte(user.Hash), []byte(password))
	if err != nil {
		return "", err
	}

	fmt.Println(user)

	return a.CreateToken(&Claims{
		ID:    user.ID,
		Email: email,
	})
}
