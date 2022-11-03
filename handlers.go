package main

import "net/http"

type Question struct {
	OperationType string `json:"operation_type"`
	X             int    `json:"x"`
	Y             int    `json:"y"`
}

type Response struct {
	SlackUsername string `json:"slackUsername"`
	Result        int    `json:"result"`
	OperationType string `json:"operation_type"`
}

func arithmeticHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

}
