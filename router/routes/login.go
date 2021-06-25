package routes

import (
	"cdn/auth"
	"cdn/db"
	"cdn/structs"
	"cdn/util"
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
				"SELECT password FROM user WHERE username = ?",
				username,
			)

			if rows == nil {
				w.Write([]byte("lol"))

				return
			}

			user := auth.Credentials{}

			err = rows.Scan(&user.Password)

			if err != nil {
				// todo: (@alii) `functions.SendError(err, errCode, w)`
				w.Write(
					[]byte(util.Stringify(util.JsonObject{
						Key:   "error",
						Value: err.Error(),
					})),
				)

				return
			}

			encoded, err := argon2.VerifyEncoded([]byte(password), []byte(user.Password))
			success := encoded && err != nil

			if !success {
				_, _ = w.Write(
					[]byte(util.Stringify(util.JsonObject{
						Key:   "error",
						Value: "incorrect password",
					})),
				)

				return
			}

			w.Write([]byte("cool"))
		},
	}
}
