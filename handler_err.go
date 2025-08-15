package main

import (
	"net/http"

	"go.uber.org/zap"
)

func handlerErr(w http.ResponseWriter, r * http.Request, log *zap.SugaredLogger){
	respondWithError(w, 400, "Something went wrong", log) 
}