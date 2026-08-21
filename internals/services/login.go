package services

import (
	"context"
	"net/http"
	"office-expense-management-backend/database"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte(os.Getenv("JWT_SECRET"))

func generateToken(username string, departement string) (string, error) {
	claims := jwt.MapClaims{
		"sub":         "123",
		"username":    username,
		"departement": departement,
		"exp":         time.Now().Add(24 * time.Hour).Unix(),
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
	var req UserReq

	c.ShouldBindJSON(&req)

	pool, err := database.Connect()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"errorInternalDatabase": err})
		return
	}

	defer pool.Close()

	row := pool.QueryRow(context.Background(), "SELECT username, password FROM users WHERE username = $1", req.Username)

	var username, password, departement string

	if err := row.Scan(&username, &password, &departement); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	if req.Password != password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid credentials"})
		return
	}

	token, err := generateToken(username, departement)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{"token": token})
}
