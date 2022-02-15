package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type JsonRequest struct {
	Id int `json:"id"`
}

type JsonResponse struct {
	Message string `json:"message"`
}

func ReceiveJson(w http.ResponseWriter, r *http.Request) {
	var request JsonRequest
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &request); err != nil {
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	response := JsonResponse{Message: fmt.Sprintf("Received JSON Message with id: %d", request.Id)}
	json.NewEncoder(w).Encode(response)
}
