package file

import (
	"cdn/db"
	"cdn/structs"
	"cdn/util"
	"database/sql"
	"os"
)

func Remove(fileName string, site structs.Site, database db.SqlDatabase) {
	doRemove(fileName, site, database)
}

func doRemove(fileName string, site structs.Site, database db.SqlDatabase) {
	retrieveQueryString := util.Format("select * from uploaded where name = '{}' and site = '{}';", fileName, site.Name)
	rows := getRows(database.DB.Query(retrieveQueryString))
	if rows != nil {
		err := os.Remove(site.RelativeLocation + "/content/" + fileName + (*rows).FileExtension)
		if err != nil {
			util.Info("Failed to remove file {} from {}!", fileName, site.Name)
			return
		}
		queryString := util.Format("delete from uploaded where name = '{}' and site = '{}';", fileName, site.Name)
		database.Query(queryString)
	}
}

func getRows(rows *sql.Rows, err error) *structs.DatabaseEntry {
	if err != nil {
		println(err.Error())
		return nil
	}
	for rows.Next() {
		entry := structs.DatabaseEntry{}
		_ = rows.Scan(&entry.OriginalName, &entry.FileName, &entry.FileExtension, &entry.Site)
		return &entry
	}
	return nil
}