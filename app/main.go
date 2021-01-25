package main

import (
	"context"
	"errors"
	"flag"
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

type App struct {
	Router *mux.Router
	Db     *pgxpool.Pool
	Conf   *config.Config
}

func main() {
	configPathPtr := flag.String("configPath", "../config/app.yml", "Path to config file.")
	flag.Parse()
	app := App{}
	app.Init(*configPathPtr)
	app.Start()
}

func (app *App) Init(path string) {
	var err error
	app.Conf, err = config.Load(path)

	if err != nil {
		log.Fatalf("Unable to load config file: %v\n", err)
	}
	app.Db, err = app.tryConnectDb(5, app.Conf.Database.Url)

	if err != nil {
		log.Fatal(err)
	}
	app.Router = mux.NewRouter()
	mRepository := repository.NewMessageRepository(app.Db)
	mService := service.NewMessageService(mRepository)
	handler.RegisterHandlers(app.Router, mService)
}

func (app *App) Start() {
	log.Println("Server started...")
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", app.Conf.Server.Host, app.Conf.Server.Port), app.Router))
}

func (a *App) tryConnectDb(tries int32, url string) (*pgxpool.Pool, error) {
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
