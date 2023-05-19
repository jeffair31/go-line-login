package authen

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type USER_PASSWORD struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var memberx = map[string]string{
	"najaemin": hashPassword("najaemin1308"),
	"renjun":   hashPassword("2303renjun"),
	"jamren":   hashPassword("1323jamren"),
}

// git
func LoginHandler(c *gin.Context) {

	account := USER_PASSWORD{}
	if err := c.ShouldBindJSON(&account); err != nil {
		c.AbortWithError(http.StatusBadRequest, errors.New(`invalid request`))
		return
	}

	// check user
	if val, ok := memberx[account.Username]; !ok || !checkPasswordHash(account.Password, val) {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	// implement login logic here
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(5 * time.Minute).Unix(),
	})

	ss, err := token.SignedString([]byte("MySignature"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"account_info": account,
		"token":        ss,
	})
}

func hashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func validateToken(token string) error {
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte("MySignature"), nil
	})

	return err
}

func Authorizationx(c *gin.Context) {
	s := c.Request.Header.Get("Authorization")

	token := strings.TrimPrefix(s, "Bearer ")

	if err := validateToken(token); err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

}
