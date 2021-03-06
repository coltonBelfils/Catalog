package endpoints

import (
	"Catalog/model/sqlDatabase/calls/queries"
	"Catalog/model/sqlDatabase/connector"
	"Catalog/niceErrors"
	"Catalog/view/api/responder"
	"encoding/json"
	"net/http"
)

func GetAdminTopics(w http.ResponseWriter, r *http.Request) { //the uuid of the user should be for internal uses only! get the user by username

	/* for later reference
	var jsonBody map[string]interface{}
	json.NewDecoder(r.Body).Decode(&jsonBody)
	*/

	if r.Method == "GET" {
		queryValues := r.URL.Query()
		username := queryValues.Get("username")

		topicQuery := queries.TopicQueryByAdminUserUsername(username)
		nErr := connector.SendQuery(topicQuery)
		if nErr != nil {
			responder.JsonRequestErrorResponder(w, nErr, 500)
			return
		}

		jsonConv, martialErr := json.Marshal(topicQuery.Results)
		if martialErr != nil {
			responder.JsonRequestErrorResponder(w, niceErrors.New(martialErr.Error(), "Unable to get admin topics by username: "+username, niceErrors.JsonConvError, niceErrors.ERROR), 500)
			return
		}

		responder.JsonRequestResponder(w, string(jsonConv), 200)
	} else {
		responder.JsonRequestErrorResponder(w, niceErrors.New("user called getUser with "+r.Method, r.Method+" calls are not allowed", niceErrors.InvalidActionByUserError, niceErrors.WARN), 405)
	}
}
