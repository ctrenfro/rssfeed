package main

import "net/http"

func handlerHealthzGet(w http.ResponseWriter, r *http.Request) {

	type response struct {
		Status string `json:"Status"`
	}

	respondWithJSON(w, http.StatusOK, response{
		Status: "ok",
	})
}
