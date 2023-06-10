package login

import (
	"bytes"
	"encoding/json"
	"github.com/config"
	"github.com/database"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/libs/utils/data"
	"github.com/libs/utils/oauth2utility"
	"github.com/models"
	"golang.org/x/oauth2"
	"io"
	"net/http"
	"time"
)

// EmailLoginFacebookController TODO make redirect to application
func EmailLoginFacebookController(c *gin.Context) {
	var (
		userProfiles models.UserProfileFacebook
		user         models.UserDescDB
	)

	db, err := database.ConnectDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": err.Error(),
		})
		panic(err)
		return
	}

	token, err := oauth2utility.GetFacebookConfRegis().Exchange(oauth2.NoContext, c.Query("code"))
	client := oauth2utility.GetFacebookConfRegis().Client(oauth2.NoContext, token)

	userProfile, err := client.Get("https://graph.facebook.com/v13.0/me?fields=id,name,email,picture")
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	defer userProfile.Body.Close()

	var resBody bytes.Buffer
	_, err = io.Copy(&resBody, userProfile.Body)

	if err := json.Unmarshal(resBody.Bytes(), &userProfiles); err != nil {
		return
	}

	if err := db.Table("user").Where("email = ?", userProfiles.Email).First(&user).Error; err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	expTime := time.Now().Local().Add(time.Hour * 3)
	claims := &config.Claims{
		Id:    user.Id,
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "go-jwt-mux",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenJWT, err := tokenAlgo.SignedString(config.JwtKey)

	if err != nil {
		panic(err)
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

func RedirectFacebookLoginController(c *gin.Context) {
	state := data.RandToken()
	c.JSON(http.StatusOK, gin.H{
		"url":    oauth2utility.GetFacebookLoginURL(state),
		"status": "ok",
	})
	return
}
