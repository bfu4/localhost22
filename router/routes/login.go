package routes

import (
	"cdn/auth"
	"cdn/db"
	"cdn/router/functions"
	"cdn/structs"
	"github.com/matthewhartstonge/argon2"
	"net/http"
)

func Login(site structs.Site) structs.Route {
	endpoint := structs.Endpoint{
		Name:    "/login",
		HostUrl: site.Url,
	}

	return structs.Route{
		Endpoint:      endpoint,
		Authenticated: false,
		Callback: func(w http.ResponseWriter, r *http.Request) {
			username := r.FormValue("username")
			password := r.FormValue("password")

			database := db.GetGlobalDatabase()

			rows, err := database.DB.Query(
				"SELECT password FROM users WHERE username = ?",
				username,
			)

			if rows == nil {
				var message string

				if err == nil {
					message = "Something went wrong"
				} else {
					message = err.Error()
				}

				functions.SendError(message, 500, w)
				return
			}

			user := auth.Credentials{}

			err = rows.Scan(&user.Password)

			if err != nil {
				functions.SendError(err.Error(), 500, w)
				return
			}

			encoded, err := argon2.VerifyEncoded([]byte(password), []byte(user.Password))
			success := encoded && err != nil

			if !success {
				functions.SendError("incorrect password", 403, w)
				return
			}

			_, _ = w.Write([]byte("OK"))
		},
	}
}
