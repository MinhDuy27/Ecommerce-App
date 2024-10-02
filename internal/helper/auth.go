package helper

import (
	"errors"
	
	"go-app/domain"
	"net/http"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	secret string
}

func GetAuth(s string) Auth{
	return Auth{
		secret: s,
	}
}

func (a Auth) HashPassword(Pp string) (string, error) {
	if len(Pp) < 6 {
		return "", errors.New("the length of password must be greater than 6")
	}
	hashP, err := bcrypt.GenerateFromPassword([]byte(Pp), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("error hashing password")
	}
	return string(hashP), nil
}
func (a Auth) VerifyPassword(Hp string, Pp string) error {

	if len(Hp) < 6 || len(Pp) < 6 {
		return errors.New("the length of password must be greater than 6")
	}
	err := bcrypt.CompareHashAndPassword([]byte(Hp),[]byte(Pp))
	if err != nil {
		return errors.New("incorrect email or password, pls try again")
	}
	return nil
}
func (a Auth) GenerateToken(id uint, email string, role string) (string,error) {
	if id == 0 || email == "" || role == "" {
		return "",errors.New("missing input to generate token")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": id,
		"email": email,
		"role": role,
		"exp": time.Now().Add(time.Hour * 24 ).Unix(),
	})
	tokenString, err := token.SignedString([]byte(a.secret))
	if err != nil {
		return "", errors.New("error generating token")
	}
	return tokenString,nil
}

func (a Auth) VerifyToken(t string) (domain.User,error) {
	tokenarr := strings.Split(t, " ")

	if tokenarr[0] != "Bearer" {
		return domain.User{},errors.New("invalid token")
	}
	token, err := jwt.Parse(tokenarr[1], func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok{
			return nil, errors.New("error parsing token")
		}
		return []byte(a.secret), nil
	})
	if err != nil {
		return domain.User{}, errors.New("error signing method")
	}

	clamis, ok := token.Claims.(jwt.MapClaims)
	if  ok && token.Valid {
		if clamis["exp"].(float64) < float64(time.Now().Unix()) {
			return domain.User{}, errors.New("token expired, pls login again")
		}
		return domain.User{
			ID: uint(clamis["id"].(float64)),
			Email: clamis["email"].(string),
			UserType: clamis["role"].(string),
		},nil
	}
	return domain.User{}, errors.New("error verify token")
}
func (a Auth) Authorize (ctx *fiber.Ctx) error{
	authHeader := ctx.Get("Authorization")
	user ,err := a.VerifyToken(authHeader)
	if err == nil || user.ID != 0 {
		ctx.Locals("user", user)
		return ctx.Next()
	}
	return ctx.Status(http.StatusUnauthorized).JSON(&fiber.Map{
		"message": "Authorization failed",
	})

}
func (a Auth) GetUser(ctx *fiber.Ctx) domain.User {
	user := ctx.Locals("user").(domain.User)
	return user
}