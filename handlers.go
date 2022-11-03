package main

import (
	"encoding/json"
	"net/http"
)

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

	var question Question

	err := json.NewDecoder(r.Body).Decode(&question)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, operation_type := check(question)

	res := Response{
		SlackUsername: "noornee",
		Result:        result,
		OperationType: operation_type,
	}

	resByte, _ := json.Marshal(res)

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write(resByte)

}
