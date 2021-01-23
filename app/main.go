package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"bitbucket.org/oaroz/hai/app/config"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"
)

func main() {
	conf, err := config.Load("../config/app.yml")
	fmt.Println(conf)
	if err != nil {
		log.Fatal(err)
	}

	conn, err := pgx.Connect(context.Background(), conf.Database.Url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	} else {
		fmt.Println("Connection established")
	}
	defer conn.Close(context.Background())

	r := mux.NewRouter()
	http.ListenAndServe(fmt.Sprintf("%s:%s", conf.Server.Host, conf.Server.Port), r)
}
