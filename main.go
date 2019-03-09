package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/gorilla/mux"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/twitter"
	"github.com/na-bot-o/mydicthub/controllers"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("can't load .env file")
	}

	goth.UseProviders(
		twitter.New(os.Getenv("TWITTER_KEY"), os.Getenv("TWITTER_SECRET"), os.Getenv("CALLBACK_URL")),
	)
}

func main() {

	controller := controllers.Controller{}

	router := mux.NewRouter()

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	router.HandleFunc("/", controller.IndexHandler).Methods("GET")
	router.HandleFunc("/login", controller.LoginHandler).Methods("GET")
	router.HandleFunc("/auth/{provider}/login", gothic.BeginAuthHandler)
	router.HandleFunc("/auth/{provider}/callback", controller.CallbackHandler)

	log.Fatal(http.ListenAndServe(":8085", router))
}
