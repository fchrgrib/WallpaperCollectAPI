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

// EmailLoginGoogleController TODO make redirect to application
func EmailLoginGoogleController(c *gin.Context) {

	var userDesc models.UserOtherEmailDescDB

	db, err := database.ConnectDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": err.Error(),
		})
		panic(err)
		return
	}

	token, err := oauth2utility.GetGoogleConfRegis().Exchange(oauth2.NoContext, c.Query("code"))
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	client := oauth2utility.GetGoogleConfRegis().Client(oauth2.NoContext, token)

	userProfile, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo?alt=json&access_token=" + token.AccessToken)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	defer userProfile.Body.Close()

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
		c.AbortWithError(http.StatusBadRequest, err)
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

func RedirectGoogleLoginController(c *gin.Context) {
	state := data.RandToken()
	c.Redirect(http.StatusSeeOther, oauth2utility.GetGoogleLoginURL(state))
}
