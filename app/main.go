package main

import (
	"fmt"
	"log"
	"net/http"

	"bitbucket.org/oaroz/hai/app/config"
	"github.com/gorilla/mux"
)

func main() {
	conf, err := config.Load("../config/app.yml")
	fmt.Println(conf)
	if err != nil {
		log.Fatal(err)
	}
	r := mux.NewRouter()
	http.ListenAndServe(fmt.Sprintf("%s:%s", conf.Server.Host, conf.Server.Port), r)
}
