package controller

import (
	"bookmarks-api/models"
	"bookmarks-api/utils"
	"encoding/json"
	"net/http"
)

// {"email" : "testApi@gmail.com", "password" : "1234qwerty"}
var CreateAccount = func(w http.ResponseWriter, r *http.Request) {
	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Invalid request"))
		return
	}
	resp := account.Create()
	utils.Respond(w, resp)
}

var Authenticate = func(w http.ResponseWriter, r *http.Request) {
	account := &models.Account{}
	err := json.NewDecoder(r.Body).Decode(account)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Invalid request"))
		return
	}
	resp := models.Login(account.Email, account.Password)
	utils.Respond(w, resp)
}
