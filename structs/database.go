package structs

import (
	"cdn/auth"
	. "cdn/util"
	"database/sql"
	"database/sql/driver"
)

// Database redefines the internal SQL database for ease-of-access
type Database struct {
	*sql.DB
	Url  string
	Name string
}

// RegisterDriver register a sql driver with the given name
func RegisterDriver(name string, driver driver.Driver) {
	if !HasDriver(name) {
		sql.Register(name, driver)
	}
}

// HasDriver check if a driver is registered
func HasDriver(name string) bool {
	for _, d := range sql.Drivers() {
		if d == name {
			return true
		}
	}
	return false
}

// OpenDatabase Open a database
func OpenDatabase(protocol string, databaseUrl string, databaseName string) Database {
	db, err := sql.Open(protocol, databaseUrl+"/"+databaseName)

	// Fail immediately.
	if err != nil {
		Fatal("Failed because of {}!", err.Error())
	}

	return Database{db, databaseUrl, databaseName}
}

// OpenDatabaseWithCredentials open a database with the specified credentials
func OpenDatabaseWithCredentials(
	protocol string,
	databaseUrl string,
	databaseName string,
	credentials auth.Credentials,
) Database {
	url := Format("{}:{}@tcp({})", credentials.Username, credentials.Password, databaseUrl)
	return OpenDatabase(protocol, url, databaseName)
}

// OpenConnection open a connection from a database using the specified protocol and credentials
func (database Database) OpenConnection(protocol string, credentials auth.Credentials) Database {
	url := Format("{}:{}@tcp({})", credentials.Username, credentials.Password, database.Url)
	return OpenDatabase(protocol, url, database.Name)
}
