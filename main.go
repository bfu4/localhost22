package main

import (
	"cdn/db"
	. "cdn/router"
	"cdn/structs"
	"cdn/util"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	"net/http"
	. "os"
)

func main() {
	// This does not want to load unless we do it manually
	_ = godotenv.Load(".env", ".env.example")
	site, _ := LookupEnv("CDN_SITE_URL")
	sitePort, _ := LookupEnv("CDN_SITE_PORT")
	// port, _ := LookupEnv("CDN_DATABASE_PORT")
	user, _ := LookupEnv("CDN_DATABASE_USER")
	pass, _ := LookupEnv("CDN_DATABASE_USER_PASSWORD")
	url, _ := LookupEnv("CDN_HOST_URL")

	util.InitLogger("CDN")
	db.InitSql()

	router := Router{
		Router: chi.NewRouter(),
	}

	mainSite := structs.Site{
		Name:             "cdn",
		RelativeLocation: ".",
		Url:              site,
	}

	SetupRoutes(router, mainSite)

	database := db.OpenDatabase(
		url,
		"test",
		user,
		pass,
	)

	db.SetGlobalDatabase(database)
	// create a table
	database.Query("create table if not exists uploaded (name VARCHAR(255), extension VARCHAR(255), site VARCHAR(255));")
	util.Info("Finished sql setup.")

	// Start the server
	util.Info("Starting server on port {}!", sitePort)
	err := http.ListenAndServe(":"+sitePort, router)

	if err != nil {
		util.Fatal("Failed to start server on port {}! {}", sitePort, err.Error())
	}

	// Do not close program
	// sc := make(chan Signal, 1)
	// signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, Interrupt, Kill)
	// <- sc
}
