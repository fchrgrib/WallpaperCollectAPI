package login

import (
	"bytes"
	"encoding/json"
	"github.com/config"
	"github.com/database"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/models"
	"io"
	"net/http"
	"time"
)

func EmailLoginGoogleController(c *gin.Context) {

	var (
		userDesc    models.UserOtherEmailDescDB
		googleToken models.GoogleToken
	)

	db, err := database.ConnectDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": err.Error(),
		})
		panic(err)
		return
	}

	if err := c.ShouldBindJSON(&googleToken); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err,
		})
		return
	}

	userProfile, err := http.Get("https://oauth2.googleapis.com/tokeninfo?id_token=" + googleToken.Token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": err,
		})
		return
	}
	defer func(Body io.ReadCloser) {
		if err := Body.Close(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status": err,
			})
			return
		}
	}(userProfile.Body)

	var resBody bytes.Buffer
	_, err = io.Copy(&resBody, userProfile.Body)

	var GoogleUserRes map[string]interface{}
	if err := json.Unmarshal(resBody.Bytes(), &GoogleUserRes); err != nil {
		return
	}

	user := &models.UserOtherEmailDescDB{
		Email:        GoogleUserRes["email"].(string),
		UserName:     GoogleUserRes["family_name"].(string),
		PhotoProfile: GoogleUserRes["picture"].(string),
		PhoneNumber:  "",
	}

	if err := db.Table("user").Where("email = ?", user.Email).First(&userDesc).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": err,
		})
		return
	}

	expTime := time.Now().Local().Add(time.Hour * 3)
	claims := &config.Claims{
		Id:    userDesc.Id,
		Email: userDesc.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "go-jwt-mux",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenJWT, err := tokenAlgo.SignedString(config.JwtKey)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": err.Error(),
		})
		return
	}

	c.SetCookie("token", tokenJWT, 4*3600, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
	return
}
