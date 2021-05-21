package controllers

import (
	"net/http"

	"github.com/ryanadiputraa/gervichstore-server/config"
)

func Oauth2Controller(w http.ResponseWriter, r *http.Request) {
	url := config.GoogleOauthConfig.AuthCodeURL(config.RandomState)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

