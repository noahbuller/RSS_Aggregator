package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func main(){
	fmt.Printf("")
	//Logger Config
	logger := zap.NewExample().Sugar()
	defer logger.Sync()
	logger.Infof("Logger Configured")

	godotenv.Load(".env")
	portString:= os.Getenv("PORT");
	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))


	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", func(w http.ResponseWriter, r *http.Request) {
		handlerReadiness(w, r, logger)
	})
	v1Router.Get("/err", func(w http.ResponseWriter, r* http.Request){
		handlerErr(w, r, logger)
	})


	v1Router.Mount("/v1", v1Router)
	srv:= &http.Server{
		Handler: v1Router,
		Addr:	":" + portString,
	}


	logger.Infof("Server starting on port %v", portString)
	err := srv.ListenAndServe()
	if err != nil{
		logger.Fatal(err)
	}
	
}