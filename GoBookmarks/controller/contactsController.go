package controller

import (
	"bookmarks-api/models"
	"bookmarks-api/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var CreateContact = func(w http.ResponseWriter, r *http.Request) {
	// Grab the id of the user that send the request
	user := r.Context().Value("user").(uint)
	fmt.Println("Current contacts: ", user)
	contact := &models.Contact{}

	err := json.NewDecoder(r.Body).Decode(contact)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Error while decoding request body"))
		return
	}
	contact.UserId = user
	resp := contact.Create()
	utils.Respond(w, resp)
}

var GetContactsFor = func(w http.ResponseWriter, r *http.Request) {
	// pass our http request = r
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.Respond(w, utils.Message(false, "There was an error in your request"))
		return
	}
	data := models.GetContacts(uint(id))
	resp := utils.Message(true, "success")
	resp["data"] = data
	utils.Respond(w, resp)
}
