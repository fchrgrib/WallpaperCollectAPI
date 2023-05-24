package oauth2utility

import (
	_ "encoding/json"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"log"
	"os"
)

//
//func AuthGoogleHandler(c *gin.Context) {
//
//	token, err := googleConf.Exchange(oauth2.NoContext, c.Query("code"))
//	if err != nil {
//		c.AbortWithError(http.StatusUnauthorized, err)
//		return
//	}
//
//	client := googleConf.Client(oauth2.NoContext, token)
//
//	userProfile, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo?alt=json&access_token=" + token.AccessToken)
//	if err != nil {
//		c.AbortWithError(http.StatusBadRequest, err)
//		return
//	}
//	defer userProfile.Body.Close()
//
//	var resBody bytes.Buffer
//	_, err = io.Copy(&resBody, userProfile.Body)
//
//	var GoogleUserRes map[string]interface{}
//	if err := json.Unmarshal(resBody.Bytes(), &GoogleUserRes); err != nil {
//		return
//	}
//
//	userBody := &models.UserOtherEmailDescDB{
//		Email:        GoogleUserRes["email"].(string),
//		UserName:     GoogleUserRes["family_name"].(string),
//		PhotoProfile: GoogleUserRes["picture"].(string),
//		PhoneNumber:  "",
//	}
//	c.JSON(http.StatusOK, gin.H{
//		"email":         userBody.Email,
//		"user_name":     userBody.UserName,
//		"photo_profile": userBody.PhotoProfile,
//		"phone_number":  "",
//	})
//	// TODO make redirect to android app root
//}

func GetGoogleRegisterURL(state string) string {
	return GetGoogleConfRegis().AuthCodeURL(state)
}
func GetGoogleLoginURL(state string) string {
	return GetGoogleConfLogin().AuthCodeURL(state)
}
func GetGoogleConfRegis() *oauth2.Config {
	err := godotenv.Load("././handlers/env/google.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	clientID := os.Getenv("GOOGLE_OAUTH_CLIENT_ID")
	clientSecret := os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET")
	redirectURL := os.Getenv("GOOGLE_OAUTH_REDIRECT_URL")

	return &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.profile",
			"https://www.googleapis.com/auth/userinfo.email",
		},
		Endpoint: google.Endpoint,
	}
}
func GetGoogleConfLogin() *oauth2.Config {
	err := godotenv.Load("././handlers/env/google.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	clientID := os.Getenv("GOOGLE_OAUTH_CLIENT_ID")
	clientSecret := os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET")
	redirectURL := os.Getenv("GOOGLE_LOGIN_REDIRECT_URL")

	return &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.profile",
			"https://www.googleapis.com/auth/userinfo.email",
		},
		Endpoint: google.Endpoint,
	}
}
