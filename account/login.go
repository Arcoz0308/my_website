package account

import (
	"github.com/arcoz0308/arcoz0308.tech/authentication"
	"io"
	"net/http"
	"text/template"
	"time"
)

type tem struct {
	Message        string
	User           string
	CustomRedirect string
}

func login(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./static/login/index.html"))
	if r.Method == http.MethodGet {
		if redirect := r.URL.Query().Get("redirect"); redirect != "" {
			t.Execute(w, tem{CustomRedirect: redirect})
			return
		}
		t.Execute(w, nil)
	} else {
		user := r.FormValue("name")
		password := r.FormValue("password")
		if user == "" {
			t.Execute(w, tem{Message: "please enter a username or a email", CustomRedirect: r.URL.Query().Get("redirect")})
			return
		}
		if len(password) < 6 {
			t.Execute(w, tem{Message: "password need 6 characters or more", User: user, CustomRedirect: r.URL.Query().Get("redirect")})
			return
		}
		a, err := authentication.GetAccount(user, password, true)
		if err != nil {
			if err == authentication.ErrUserDontExist || err == authentication.ErrPasswordAreInvalid {
				t.Execute(w, tem{Message: "invalid user or password", User: user, CustomRedirect: r.URL.Query().Get("redirect")})
				return
			}
			if err == authentication.ErrVerifyEmail {
				//TODO this
			} else {
				w.WriteHeader(http.StatusInternalServerError)
				io.WriteString(w, "internal server error")
				return
			}
		}
		c := authentication.GenerateCookie(a, authentication.CookieTypeSession)
		cookie := http.Cookie{
			Name:    "authentication",
			Value:   c,
			Domain:  ".arcoz0308.tech",
			Expires: time.Now().AddDate(0, 0, 30),
		}
		r.AddCookie(&cookie)
	}

}
