package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"bitbucket.org/oaroz/hai/app/config"
	"bitbucket.org/oaroz/hai/app/handler"
	"bitbucket.org/oaroz/hai/app/repository"
	"bitbucket.org/oaroz/hai/app/service"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {
	conf, err := config.Load("../config/app.yml")

	if err != nil {
		log.Fatalf("Unable to load config file: %v\n", err)
	}
	db, err := tryConnectDb(5, conf.Database.Url)

	if err != nil {
		log.Fatalf(fmt.Sprintf(err.Error()))
	}
	r := mux.NewRouter()
	mRepository := repository.NewMessageRepository(db)
	mService := service.NewMessageService(mRepository)
	handler.RegisterHandlers(r, mService)
	log.Println("Server started...")
	http.ListenAndServe(fmt.Sprintf("%s:%s", conf.Server.Host, conf.Server.Port), r)
}

func tryConnectDb(tries int32, url string) (*pgxpool.Pool, error) {
	for tries > 0 {
		db, err := pgxpool.Connect(context.Background(), url)
		if err != nil {
			log.Printf("Unable to connect to database: %v\n Left tries: %d\n", err, tries)
			tries--
			time.Sleep(5 * time.Second)
		} else {
			log.Println("Connection to database has been established.")
			return db, err
		}
	}
	return nil, errors.New("Unable to connect to database")
}
