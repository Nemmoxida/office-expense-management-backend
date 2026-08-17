package services

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

var secretKey = []byte(os.Getenv("JWT_SECRET"))

func generateToken(username string) (string, error) {
	claims := jwt.MapClaims{
		"sub":      "123",
		"username": username,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

type UserReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	godotenv.Load()
	var req UserReq
	fmt.Print(secretKey)

	c.ShouldBindJSON(&req)

	if req.Username != "kermit" && req.Password != "123" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Wrong username or password"})
		return
	}

	token, err := generateToken(req.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.Data(http.StatusAccepted, "application/json", []byte(token))
}
