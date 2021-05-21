package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ryanadiputraa/gervichstore-server/config"
	"github.com/ryanadiputraa/gervichstore-server/helpers"
	"golang.org/x/oauth2"
)

func LoginCallbackController(w http.ResponseWriter, r *http.Request) {
	if r.FormValue("state") != config.RandomState {
		helpers.WriteErrorResponse(w, r, http.StatusUnauthorized, "state is not valid")
		return
	}

	token, err := config.GoogleOauthConfig.Exchange(oauth2.NoContext, r.FormValue("code"))
	fmt.Println(token)
	if err != nil {
		helpers.WriteErrorResponse(w, r, http.StatusBadGateway, "could not get token")
		return
	}

	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token="+token.AccessToken)
	if err != nil {
		helpers.WriteErrorResponse(w, r, http.StatusBadGateway, "cant request token to google api")
		return
	}

	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		helpers.WriteInternalServerError(w, r,)
		return
	}

	helpers.WriteResponse(w, r, "application/json", http.StatusOK, data)
}