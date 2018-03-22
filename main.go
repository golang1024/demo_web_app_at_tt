package main

import (
	"demo_web_app/router"
	"demo_web_app/db"

	"net/http"
	"log"
)

func StartApp() {
	//db load config
	db.LoadConfig()

	//router init
	r := router.InitRouter()
	//http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":80", r))
}


func main() {
	StartApp()

}
