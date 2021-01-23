package main

import (
	"context"
	"fmt"
	"net/http"

	"bitbucket.org/oaroz/hai/app/config"
	"bitbucket.org/oaroz/hai/app/handler"
	"bitbucket.org/oaroz/hai/app/repository"
	"bitbucket.org/oaroz/hai/app/service"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {
	conf, err := config.Load("../config/app.yml")
	fmt.Println(conf)

	if err != nil {
		panic(fmt.Sprintf("Unable to load config file: %v\n", err))
	}
	db, err := pgxpool.Connect(context.Background(), conf.Database.Url)

	if err != nil {
		panic(fmt.Sprintf("Unable to connect to database: %v\n", err))
	}
	r := mux.NewRouter()
	mRepository := repository.NewMessageRepository(db)
	mService := service.NewMessageService(mRepository)
	handler.RegisterHandlers(r, mService)
	http.ListenAndServe(fmt.Sprintf("%s:%s", conf.Server.Host, conf.Server.Port), r)
}
