package api

import (
	"net/http"
)

func StartRequestController() {
	http.HandleFunc("/info", GetDocumentation) //html api documentation
	http.HandleFunc("/api/check", Check)//Working, called to see if the api is online
	http.HandleFunc("/api/editRanking", EditRanking)
	http.HandleFunc("/api/getUser", GetUser)//Provisionally working. Auth not implemented yet
	http.HandleFunc("/api/newUser", NewUser)//Provisionally working. Auth not implemented yet, email verification not implemented yet
	http.HandleFunc("/api/editTopic", EditTopic)
	http.HandleFunc("/api/getRanking", GetRanking)
	http.HandleFunc("/api/getTopic", GetTopic)
	http.HandleFunc("/api/newTopic", NewTopic)
	http.HandleFunc("/api/searchTopic", SearchTopics)
	http.HandleFunc("/api/searchUsers", SearchUsers)
	http.HandleFunc("/api/getAdminTopics", GetAdminTopics)//Provisionally working. Auth not implemented yet

	_ = http.ListenAndServe(":80", nil)
}