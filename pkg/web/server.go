package web

import (
	"html/template"
	"net/http"
)

// Serve serves a webserver
func Serve(plaidEnv string, plaidPublicKey string) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("pkg/web/public/index.html"))

		data := map[string]string{
			"PlaidEnv":       plaidEnv,
			"PlaidPublicKey": plaidPublicKey,
		}

		tmpl.Execute(w, data)
	})

	http.ListenAndServe(":80", nil)
}
