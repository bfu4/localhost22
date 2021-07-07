package main

import (
	"cdn/db"
	. "cdn/router"
	"cdn/structs"
	"cdn/util"
	. "os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	site, _ := LookupEnv("CDN_SITE_URL")
	sitePort, _ := LookupEnv("CDN_SITE_PORT")

	util.InitLogger("CDN")

	router := Router{
		Router: chi.NewRouter(),
	}

	mainSite := structs.Site{
		Name:             "cdn",
		RelativeLocation: ".",
		Port:             sitePort,
		Url:              site,
	}

	SetupRoutes(router, mainSite)

	db.OpenDatabase()

	// Start the server
	util.Info("Starting server on port {}!", sitePort)

	mainSite.Listen(router.Router)

	// Do not close program
	sc := make(chan Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, Interrupt, Kill)
	<-sc
}
