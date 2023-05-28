package oauth2utility

import (
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
	_ "golang.org/x/oauth2/facebook"
	"log"
	"os"
)

var facebookConf *oauth2.Config

func init() {

}

// GetFacebookRegisterURL func AuthFacebookHandler(c *gin.Context) {
//
//	token, err := facebookConf.Exchange(oauth2.NoContext, c.Query("code"))
//
//	client := facebookConf.Client(oauth2.NoContext, token)
//
//	userProfile, err := client.Get("https://graph.facebook.com/v13.0/me?fields=id,name,email,picture")
//	if err != nil {
//		c.AbortWithError(http.StatusBadRequest, err)
//		return
//	}
//	defer userProfile.Body.Close()
//
//	var resBody bytes.Buffer
//	_, err = io.Copy(&resBody, userProfile.Body)
//
//	var FacebookUserRes map[string]interface{}
//	if err := json.Unmarshal(resBody.Bytes(), &FacebookUserRes); err != nil {
//		return
//	}
//
//	c.JSON(http.StatusOK, gin.H{
//		"body": FacebookUserRes,
//	})
//
// }
func GetFacebookRegisterURL(state string) string {
	return GetFacebookConfRegis().AuthCodeURL(state)
}
func GetFacebookLoginURL(state string) string {
	return GetFacebookConfLogin().AuthCodeURL(state)
}
func GetFacebookConfRegis() *oauth2.Config {
	err := godotenv.Load("././handlers/env/facebook.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	clientID := os.Getenv("FACEBOOK_OAUTH_CLIENT_ID")
	clientSecret := os.Getenv("FACEBOOK_OAUTH_CLIENT_SECRET")
	redirectURL := os.Getenv("FACEBOOK_OAUTH_REDIRECT_URL")

	return &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Scopes: []string{
			"public_profile",
			"email",
		},
		Endpoint: facebook.Endpoint,
	}
}
func GetFacebookConfLogin() *oauth2.Config {
	err := godotenv.Load("././handlers/env/facebook.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	clientID := os.Getenv("FACEBOOK_OAUTH_CLIENT_ID")
	clientSecret := os.Getenv("FACEBOOK_OAUTH_CLIENT_SECRET")
	redirectURL := os.Getenv("FACEBOOK_LOGIN_REDIRECT_URL")

	return &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Scopes: []string{
			"public_profile",
			"email",
		},
		Endpoint: facebook.Endpoint,
	}
}

// TODO make redirect to android app root
