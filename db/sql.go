package db

import (
	"cdn/auth"
	"cdn/structs"
	"cdn/util"
	"github.com/go-sql-driver/mysql"
)

// SqlDatabase type-alias the database locally for this file
type SqlDatabase struct {
	structs.Database
}

func InitSql() {
	structs.RegisterDriver("mysql", &mysql.MySQLDriver{})
}

// OpenDatabase open a given database
func OpenDatabase(databaseUrl string, databaseName string, user string, password string) SqlDatabase {
	creds := auth.Credentials{
		Username: user,
		Password: password,
	}
	return SqlDatabase{
		structs.OpenDatabaseWithCredentials(
			"mysql", databaseUrl,
			databaseName,
			creds,
		),
	}
}

// Login open a connection to the database with the specified credentials
func (database SqlDatabase) Login(user string, password string) {
	database.OpenConnection("mysql", user, password)
}

// Query make a query and get the scan
func (database SqlDatabase) Query(query string) int {
	rows, err := database.DB.Query(query)
	if err != nil {
		util.Info("Query errored because of {}!", err.Error())
		return 0
	} else {
		var count int
		rows.Scan(&count)
		// Log
		util.Info(
			"Successfully executed statement for [{}]! Found {} matches.",
			query,
			string(rune(count)),
		)
		return count
	}
}

// Execute execute a statement
func (database SqlDatabase) Execute(statement string) {
	_, err := database.DB.Exec(statement)
	if err != nil {
		util.Info("Statement execution failed because of {}!", err.Error())
	} else {
		util.Info("Successfully executed statement for [{}]!", statement)
	}
}
