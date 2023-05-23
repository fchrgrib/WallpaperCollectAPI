package config

import (
	"github.com/golang-jwt/jwt/v4"
)

var JwtKey = []byte("skdvsjknvdhve788yhghvysuyzgc6tzzt67tc6zUTAVFGsgfuyUv6svtGVSTV^VSvt&RV%S&TVYUSGVS&T(S&&(9S&VTsVgSVgS*V*SvFVs8VfY")

type Claims struct {
	Id    string
	Email string
	jwt.RegisteredClaims
}

type GoogleCredentials struct {
	Cid         string `json:"c_id"`
	CSecret     string `json:"c_secret"`
	RedirectURL string `json:"redirect_url"`
}
