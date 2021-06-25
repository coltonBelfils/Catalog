package controller

import (
	"Catalog/view/api/endpoints"
	m "Catalog/view/api/middleware"
	"net/http"
)

func RequestController() {
	//path scheme: -- /api/<kind of action, verb:read, write, delete, assign, revoke, execute>/<what is being acted on, noun: ownUserRankings, otherUserInfo, ...> --
	http.Handle("/", m.CorsMiddleware(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.URL.Path != "/" && req.URL.Path != "" { //Checking to make sure the pattern matches / and isn't something else
			http.NotFound(w, req)
			return
		}
		endpoints.GetHomepage(w, req)
	})))                                                                                          //html home page
	http.Handle("/documentation", m.CorsMiddleware(http.HandlerFunc(endpoints.GetDocumentation))) //html api documentation
	http.Handle("/api/check", m.CorsMiddleware(http.HandlerFunc(endpoints.Check)))                //Working, called to see if the api is online
	http.Handle("/api/editRanking", m.CorsMiddleware(http.HandlerFunc(endpoints.EditRanking)))
	http.Handle("/api/getUser", m.CorsMiddleware(m.AuthMiddleware(http.HandlerFunc(endpoints.GetUser)))) //Provisionally working.
	http.Handle("/api/newUser", m.CorsMiddleware(http.HandlerFunc(endpoints.NewUser)))                   //Provisionally working.
	http.Handle("/api/editTopic", m.CorsMiddleware(m.AuthMiddleware(http.HandlerFunc(endpoints.EditTopic))))
	http.Handle("/api/getRanking", m.CorsMiddleware(http.HandlerFunc(endpoints.GetRanking)))
	http.Handle("/api/getTopic", m.CorsMiddleware(http.HandlerFunc(endpoints.GetTopic)))
	http.Handle("/api/newTopic", m.CorsMiddleware(m.AuthMiddleware(http.HandlerFunc(endpoints.NewTopic))))
	http.Handle("/api/searchTopic", m.CorsMiddleware(http.HandlerFunc(endpoints.SearchTopics)))
	http.Handle("/api/searchUsers", m.CorsMiddleware(http.HandlerFunc(endpoints.SearchUsers)))
	http.Handle("/api/getAdminTopics", m.CorsMiddleware(m.AuthMiddleware(http.HandlerFunc(endpoints.GetAdminTopics)))) //Provisionally working.

	_ = http.ListenAndServe(":80", nil)
}
