package api

import (
	"Catalog/controler"
	"Catalog/niceErrors"
	"encoding/json"
	"net/http"
)

func NewUser(w http.ResponseWriter, r *http.Request) { //the uuid of the user should be for internal uses only! get the user by username

	/* for later reference
	var jsonBody map[string]interface{}
	json.NewDecoder(r.Body).Decode(&jsonBody)
	*/

	if r.Method == "POST" {
		var jsonBody map[string]interface{}
		decodeErr := json.NewDecoder(r.Body).Decode(&jsonBody)
		if decodeErr != nil {
			JsonRequestErrorResponder(w, niceErrors.FromErrorFull(decodeErr, "newUser decode json error", "Invalid json in request", niceErrors.INFO), 400)
			return
		}

		var username string
		if unRaw, ok := jsonBody["username"]; ok {
			if unString, ok := unRaw.(string); ok {
				username = unString
			}
		}

		var email string
		if eRaw, ok := jsonBody["email"]; ok {
			if eString, ok := eRaw.(string); ok {
				email = eString
			}
		}

		userCreated, nErr := controler.CreateUser(username, email)
		if nErr != nil {
			JsonRequestErrorResponder(w, nErr, 500)
			return
		}

		jsonConv, martialErr := json.Marshal(userCreated)
		if martialErr != nil {
			JsonRequestErrorResponder(w, niceErrors.New(martialErr.Error(), "User should be created but an error occurred", niceErrors.ERROR), 500)
			return
		}

		JsonRequestResponder(w, string(jsonConv), 200)
	} else {
		JsonRequestErrorResponder(w, niceErrors.New("user called getUser with "+r.Method, r.Method+" calls are not allowed", niceErrors.WARN), 405)
	}
}
