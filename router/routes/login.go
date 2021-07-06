package routes

import (
	"cdn/auth"
	"cdn/db"
	"cdn/router/functions"
	"cdn/structs"
	"cdn/util"
	"github.com/matthewhartstonge/argon2"
	"net/http"
	"strings"
)

func Login(site structs.Site) structs.Route {
	endpoint := structs.Endpoint{
		Name:    "/login",
		HostUrl: site.Url,
	}

	return structs.Route{
		Endpoint:      endpoint,
		Authenticated: false,
		Methods:       []string{"POST"},
		Callback: func(w http.ResponseWriter, r *http.Request, userId int) {
			_ = r.ParseMultipartForm(util.DefaultFormMaxMem)

			username := r.FormValue("username")
			password := r.FormValue("password")

			if len(username) == 0 || len(password) == 0 {
				functions.SendError("No username or password specified", 400, w)
				return
			}

			database := db.GetGlobalDatabase()

			rows, err := database.DB.Query(
				"SELECT id, password FROM users WHERE username = ?",
				username,
			)

			if rows == nil {
				message := util.ErrorOrMessage(err, "Something went wrong.")
				functions.SendError(message, 500, w)
				return
			}

			user := auth.Credentials{}

			if !rows.Next() {
				functions.SendError("User not found", 404, w)
				return
			}

			err = rows.Scan(&user.Id, &user.Password)

			if err != nil {
				functions.SendError(err.Error(), 500, w)
				return
			}

			encoded, err := argon2.VerifyEncoded([]byte(password), []byte(user.Password))
			success := encoded && (err == nil)

			if !success {
				functions.SendError("Incorrect password", 403, w)
				return
			}

			if err != nil {
				functions.SendError(err.Error(), 500, w)
				return
			}

			jwtWrapper := util.GetJWTWrapper()

			expiry := jwtWrapper.GetExpiry()
			tokenString, err := jwtWrapper.GenerateToken(user.Id, expiry)

			if err != nil {
				functions.SendError(err.Error(), 403, w)
				return
			}

			cookie := http.Cookie{
				HttpOnly: true,
				SameSite: http.SameSiteLaxMode,
				Secure:   !strings.Contains(site.Url, "localhost"),
				Path:     "/",
				Expires:  expiry,
				Value:    tokenString,
				Name:     "token",
			}

			http.SetCookie(w, &cookie)

			_, _ = w.Write([]byte("OK"))
		},
	}
}
