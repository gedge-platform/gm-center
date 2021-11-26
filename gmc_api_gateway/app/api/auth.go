package api

import (
	"encoding/json"
	"fmt"
	"gmc_api_gateway/app/db"
	"gmc_api_gateway/app/model"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

const (
	accessTokenCookieName  = "access-token"
	refreshTokenCookieName = "refresh-token"
	jwtSecretKey           = "some-secret-key"
)

func GetJWTSecret() string {
	return jwtSecretKey
}

type Claims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

func AuthenticateUser(id, password string) bool {
	db := db.DbManager()
	var user model.MemberWithPassword

	idCheck := strings.Compare(id, "") != 0
	passCheck := strings.Compare(password, "") != 0

	if idCheck && passCheck {

		if err := db.First(&user, model.MemberWithPassword{Member: model.Member{Id: id}, Password: password}).Error; err == nil {
			return true
		}

	}

	return false
}

func LoginUser(c echo.Context) (err error) {

	var user model.User

	Body := responseBody(c.Request().Body)
	fmt.Println("Body is : ", Body)
	// t, _ := ioutil.ReadAll(c.Request().Body)
	err = json.Unmarshal([]byte(Body), &user)
	if err != nil {
		c.String(http.StatusInternalServerError, "Invalid json provided")
		return
	}
	fmt.Println("Body Value is : ", user)
	fmt.Println("user email is : ", user.Id)
	fmt.Println("user password is : ", user.Password)

	loginResult := AuthenticateUser(user.Id, user.Password)

	fmt.Println("loginResult is : ", loginResult)

	if loginResult {
		accessToken, _, err := generateAccessToken(user.Id)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, err.Error())
		}
		fmt.Println("token is : ", accessToken)
		return c.JSON(http.StatusOK, echo.Map{
			"status":       200,
			"access-token": accessToken,
		})
	}
	return c.JSON(http.StatusUnauthorized, false)
}

func generateAccessToken(userid string) (string, time.Time, error) {

	expirationTime := time.Now().Add(time.Minute * 15)

	return generateToken(userid, expirationTime, []byte(GetJWTSecret()))
}

func generateToken(userid string, expirationTime time.Time, secret []byte) (string, time.Time, error) {
	claims := &Claims{
		Name: userid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", time.Now(), err
	}

	return tokenString, expirationTime, nil
}

// func VerifyAccessToken(c echo.Context) (err error) {

// }
