package main

import "net/http"

func errGet(w http.ResponseWriter, r *http.Request) {

	respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
}
