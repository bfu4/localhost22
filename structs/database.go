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

func RegisterDriver(name string, driver driver.Driver) {
	if !HasDriver(name) {
		sql.Register(name, driver)
	}
}

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

func OpenDatabaseWithCredentials(
	protocol string,
	databaseUrl string,
	databaseName string,
	credentials auth.Credentials,
) Database {
	url := Format("{}:{}@tcp({})", credentials.Username, credentials.Password, databaseUrl)
	return OpenDatabase(protocol, url, databaseName)
}

func (database Database) OpenConnection(protocol string, user string, password string) Database {
	url := Format("{}:{}@tcp({})", user, password, database.Url)
	return OpenDatabase(protocol, url, database.Name)
}
