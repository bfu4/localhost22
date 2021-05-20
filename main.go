package main

import (
	"cdn/db"
	. "cdn/router"
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
	// user, _ := LookupEnv("CDN_DATABASE_USER")
	// pass, _ := LookupEnv("CDN_DATABASE_USER_PASSWORD")
	// url, _ := LookupEnv("CDN_HOST_URL")

	util.InitLogger("CDN")
	db.InitSql()

	router := Router{
		Router:    chi.NewRouter(),
	}

	SetupRoutes(router, site)

	util.Info("Starting server on port {}!", sitePort)
	err := http.ListenAndServe(":" + sitePort, router)

	if err != nil {
		util.Fatal("Failed to start server on port {}! {}", sitePort, err.Error())
	}

	// database := db.OpenDatabase(
	//	url,
	//	"test",
	//	user,
	//	pass,
	// )
	// database.Execute("SELECT \"Hello World!\";")

	// Do not close program
	// sc := make(chan Signal, 1)
	// signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, Interrupt, Kill)
	// <- sc
}
