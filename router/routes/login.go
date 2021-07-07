package routes

import (
	"cdn/db"
	"cdn/router/functions"
	"cdn/structs"
	"cdn/structs/models"
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

			var user models.User
			result := database.First(&user, "username = ?", username)

			if result.Error != nil {
				functions.SendError("User not found", 404, w)
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
