package oauth2utility

import (
	"bytes"
	"encoding/json"
	_ "encoding/json"
	"fmt"
	"github.com/config"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/models"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io"
	"log"
	"net/http"
	"os"
)

//type GoogleOauthToken struct {
//	AccessToken string
//	IdToken     string
//}
//
//func GetGoogleOauthToken() (*GoogleOauthToken, error) {
//	const rootURl = "https://oauth2.googleapis.com/token"
//
//	config, _ := LoadConfig("")
//
//	fmt.Println(config)
//	values := url.Values{}
//	values.Add("grant_type", "authorization_code")
//	values.Add("client_id", config.GoogleClientID)
//	values.Add("client_secret", config.GoogleClientSecret)
//	values.Add("redirect_uri", config.GoogleOAuthRedirectUrl)
//
//	query := values.Encode()
//
//	req, err := http.NewRequest("POST", rootURl, bytes.NewBufferString(query))
//	if err != nil {
//		return nil, err
//	}
//
//	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
//	client := http.Client{
//		Timeout: time.Second * 30,
//	}
//
//	res, err := client.Do(req)
//	if err != nil {
//		return nil, err
//	}
//
//	if res.StatusCode != http.StatusOK {
//		return nil, errors.New("could not retrieve token")
//	}
//
//	var resBody bytes.Buffer
//	_, err = io.Copy(&resBody, res.Body)
//	if err != nil {
//		return nil, err
//	}
//
//	var GoogleOauthTokenRes map[string]interface{}
//
//	if err := json.Unmarshal(resBody.Bytes(), &GoogleOauthTokenRes); err != nil {
//		return nil, err
//	}
//
//	tokenBody := &GoogleOauthToken{
//		AccessToken: GoogleOauthTokenRes["access_token"].(string),
//		IdToken:     GoogleOauthTokenRes["id_token"].(string),
//	}
//
//	return tokenBody, nil
//}
//
//func GetGoogleUser(accessToken string, idToken string) (*models.UserOtherEmailDescDB, error) {
//	rootUrl := fmt.Sprintf("https://www.googleapis.com/oauth2/v1/userinfo?alt=json&access_token=%s", accessToken)
//
//	req, err := http.NewRequest("GET", rootUrl, nil)
//	if err != nil {
//		return nil, err
//	}
//
//	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", idToken))
//
//	client := http.Client{
//		Timeout: time.Second * 30,
//	}
//	res, err := client.Do(req)
//	if err != nil {
//		return nil, err
//	}
//
//	if res.StatusCode != http.StatusOK {
//		return nil, errors.New("could not retrieve user")
//	}
//
//	var resBody bytes.Buffer
//	_, err = io.Copy(&resBody, res.Body)
//	if err != nil {
//		return nil, err
//	}
//
//	var GoogleUserRes map[string]interface{}
//
//	if err := json.Unmarshal(resBody.Bytes(), &GoogleUserRes); err != nil {
//		return nil, err
//	}
//
//	userBody := &models.UserOtherEmailDescDB{
//		Id:           GoogleUserRes["id"].(string),
//		Email:        GoogleUserRes["email"].(string),
//		UserName:     GoogleUserRes["name"].(string),
//		PhotoProfile: GoogleUserRes["picture"].(string),
//		PhoneNumber:  GoogleUserRes["phone_number"].(string),
//	}
//	return userBody, nil
//}
//
//func GoogleOAuthController(c *gin.Context) {
//	//var pathUrl string = "/"
//	//
//	//if c.Query("state") != "" {
//	//	pathUrl = c.Query("state")
//	//}
//
//	//if code == "" {
//	//	c.JSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "Authorization code not provided!"})
//	//	return
//	//}
//
//	tokenRes, err := GetGoogleOauthToken()
//
//	if err != nil {
//		c.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
//		return
//	}
//
//	google_user, err := GetGoogleUser(tokenRes.AccessToken, tokenRes.IdToken)
//	c.JSON(http.StatusOK, gin.H{
//		"user": google_user,
//	})
//}
//
//func LoadConfig(path string) (config Config, err error) {
//	viper.AddConfigPath(path)
//	viper.SetConfigType("env")
//	viper.SetConfigName("app")
//
//	viper.AutomaticEnv()
//
//	err = viper.ReadInConfig()
//	if err != nil {
//		return
//	}
//
//	err = viper.Unmarshal(&config)
//	return
//}
//
//type Config struct {
//	JWTTokenSecret string        `mapstructure:"JWT_SECRET"`
//	TokenExpiresIn time.Duration `mapstructure:"TOKEN_EXPIRED_IN"`
//	TokenMaxAge    int           `mapstructure:"TOKEN_MAXAGE"`
//
//	GoogleClientID         string `mapstructure:"GOOGLE_OAUTH_CLIENT_ID"`
//	GoogleClientSecret     string `mapstructure:"GOOGLE_OAUTH_CLIENT_SECRET"`
//	GoogleOAuthRedirectUrl string `mapstructure:"GOOGLE_OAUTH_REDIRECT_URL"`
//}

var cred config.GoogleCredentials
var conf *oauth2.Config

func init() {
	err := godotenv.Load("././handlers/env/google.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	clientID := os.Getenv("GOOGLE_OAUTH_CLIENT_ID")
	clientSecret := os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET")
	redirectURL := os.Getenv("GOOGLE_OAUTH_REDIRECT_URL")
	fmt.Println("haaaaaaaaaaaaa")
	fmt.Println(clientID)

	conf = &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.profile",
			"https://www.googleapis.com/auth/userinfo.email", // You have to select your own scope from here -> https://developers.google.com/identity/protocols/googlescopes#google_sign-in
		},
		Endpoint: google.Endpoint,
	}
}

func AuthGoogleHandler(c *gin.Context) {
	fmt.Println(conf)

	token, err := conf.Exchange(oauth2.NoContext, c.Query("code"))
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	client := conf.Client(oauth2.NoContext, token)

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

	userBody := &models.UserOtherEmailDescDB{
		Email:        GoogleUserRes["email"].(string),
		UserName:     GoogleUserRes["family_name"].(string),
		PhotoProfile: GoogleUserRes["picture"].(string),
		PhoneNumber:  "",
	}
	c.JSON(http.StatusOK, gin.H{
		"email":         userBody.Email,
		"user_name":     userBody.UserName,
		"photo_profile": userBody.PhotoProfile,
		"phone_number":  "",
	})
	// TODO make redirect to android app root
}

func GetLoginURL(state string) string {
	return conf.AuthCodeURL(state)
}
func GetConf() *oauth2.Config {
	return conf
}
