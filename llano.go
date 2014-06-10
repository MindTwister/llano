/*
Llano is a mock server, useful when testing other libraries during development.
*/
package llano

import (
	"io"
	"net/http"
	"time"
)

/*
Returns a http.HandlerFunc that always respons with 200: standardResponse
*/
func Code200(standardResponse string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(standardResponse))
	}
}

/*
More or less a wrapper around http.Redirect (http://golang.org/pkg/net/http/#Redirect)
*/
func Code301(to string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, to, 301)
	}
}

/*
More or less a wrapper around http.Redirect (http://golang.org/pkg/net/http/#Redirect)
*/
func Code302(to string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, to, 302)
	}
}

/*
Guess, always returns 500
*/
func Code500() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
	}
}

/*
Holds the connection for indefinite time
*/
func Timeout(w http.ResponseWriter, r *http.Request) {
	time.Sleep(24 * time.Hour)
}

/*
Returns the request body with the original Content-Type forwarded
*/
func Echo(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Add("Content-Type", r.Header.Get("Content-Type"))
	io.Copy(w, r.Body)
}

/*
Standalone llano server, listens for all http.Handlers specified in the package with their response code as url (/200, /301) (echo for... echo)
*/
func Standalone(address, default200 string) {
	http.HandleFunc("/200", Code200(default200))
	http.HandleFunc("/301", Code301("/200"))
	http.HandleFunc("/302", Code302("/200"))
	http.HandleFunc("/500", Code500())
	http.HandleFunc("/echo", Echo)
	http.HandleFunc("/timeout", Timeout)
	http.ListenAndServe(address, nil)
}
