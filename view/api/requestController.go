package api

import (
	"net/http"
)

func StartRequestController() {
	http.HandleFunc("/info", GetDocumentation) //html api documentation
	http.HandleFunc("/api/check", Check)
	http.HandleFunc("/api/editRanking", EditRanking)
	http.HandleFunc("/api/getUser", GetUser)
	http.HandleFunc("/api/newUser", NewUser)
	http.HandleFunc("/api/editTopic", EditTopic)
	http.HandleFunc("/api/getRanking", GetRanking)
	http.HandleFunc("/api/getTopic", GetTopic)
	http.HandleFunc("/api/newTopic", NewTopic)
	http.HandleFunc("/api/searchTopic", SearchTopics)
	http.HandleFunc("/api/searchUsers", SearchUsers)
	http.HandleFunc("/api/getAdminTopics", SearchUsers)

	_ = http.ListenAndServe(":8080", nil)
}