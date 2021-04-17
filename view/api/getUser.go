package api

import (
	"Catalog/model/sql/calls/queries"
	"Catalog/model/sql/connector"
	"Catalog/niceErrors"
	"encoding/json"
	"net/http"
)

func GetUser(w http.ResponseWriter, r *http.Request) { //the uuid of the user should be for internal uses only! get the user by username

	/* for later reference
	var jsonBody map[string]interface{}
	json.NewDecoder(r.Body).Decode(&jsonBody)
	*/

	if r.Method == "GET" {
		queryValues := r.URL.Query()
		userName := queryValues.Get("username")

		query := queries.UserQueryByUsername(userName)
		connErr := connector.SendQuery(query)
		if connErr != nil {
			nConnErr := niceErrors.FromError(connErr)
			JsonRequestErrorResponder(w, nConnErr, 500)
			return
		}

		if len(query.Results) == 0 {
			JsonRequestErrorResponder(w, niceErrors.New("userError", "Unable to find user with username: "+userName, niceErrors.INFO), 404)
			return
		} else if len(query.Results) < 1 {
			JsonRequestErrorResponder(w, niceErrors.New("Multiple users with username: "+userName, "Unable to find user with username: "+userName, niceErrors.ERROR), 500)
			return
		}

		jsonConv, martialErr := json.Marshal(query.Results[0])
		if martialErr != nil {
			JsonRequestErrorResponder(w, niceErrors.New(martialErr.Error(), "Unable to find user with username: "+userName, niceErrors.ERROR), 500)
		}

		JsonRequestResponder(w, string(jsonConv), 200)

	} else {
		JsonRequestErrorResponder(w, niceErrors.New("user called getUser with "+r.Method, r.Method+" calls are not allowed", niceErrors.WARN), 405)
	}
}
