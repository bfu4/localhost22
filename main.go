package main

import (
	"cdn/db"
	"cdn/util"
	"github.com/joho/godotenv"
)

func main() {
	// This does not want to load unless we do it manually
	_ = godotenv.Load(".env", ".env.example")
	// host, _ := LookupEnv("CDN_HOST_ADDRESS")
	// port, _ := LookupEnv("CDN_DATABASE_PORT")
	// user, _ := LookupEnv("CDN_DATABASE_USER")
	// pass, _ := LookupEnv("CDN_DATABASE_USER_PASSWORD")
	// url, _ := LookupEnv("CDN_HOST_URL")

	util.InitLogger("CDN")
	db.InitSql()

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
