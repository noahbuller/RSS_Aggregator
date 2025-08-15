package main

import (
	"net/http"

	"go.uber.org/zap"
)

func handlerReadiness(w http.ResponseWriter, r *http.Request, log *zap.SugaredLogger){
	respondWithJSON(w, 200, struct {}{}, log)
	//only care about response code
}