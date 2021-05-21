package config

import (
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	GoogleOauthConfig = &oauth2.Config{
		RedirectURL: "http://localhost:8080/api/login/callback",
		Scopes: []string{"https://www.googleapis.com/auth/userinfo.profile"},
		Endpoint: google.Endpoint,
	}
	RandomState = String(16)
)

func init() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
		return
	}
	GoogleOauthConfig.ClientID = os.Getenv("GOOGLE_CLIENT_ID")
	GoogleOauthConfig.ClientSecret = os.Getenv("GOOGLE_CLIENT_SECRET")
}

// generate random state
const charset = "abcdefghijklmnopqrstuvwxyz" +
  "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
  rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
  b := make([]byte, length)
  for i := range b {
    b[i] = charset[seededRand.Intn(len(charset))]
  }
  return string(b)
}

func String(length int) string {
  return StringWithCharset(length, charset)
}