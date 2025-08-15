package main
import(
	"net/http"
	"encoding/json"
	"go.uber.org/zap"
)

func respondWithJSON(w http.ResponseWriter, code int, payload any, log *zap.SugaredLogger) {
	logger := *log
	defer logger.Sync()

	data, er := json.Marshal(payload)
	if er != nil{
		w.WriteHeader(500)
		logger.Errorf("Failed to Marshal Json Payload:%f", payload)
	}
	w.Header().Add("Content-Type", "applicaiton/json")
	w.WriteHeader(code)
	w.Write(data)
}

func respondWithError(w http.ResponseWriter, code int, msg string, log *zap.SugaredLogger){
	if code > 499{ //server side error, 
		log.Infof("Responding with 5XX err:", msg)
	}
	type errResponse struct{
		Error string `json:"error"`
	}

	respondWithJSON(w, code, errResponse{
		Error: msg, 
	}, log)
}