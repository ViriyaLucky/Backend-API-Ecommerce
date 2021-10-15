package handler

import (
	"fmt"
	"net/http"
)

func UserHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorHandler(w, r)
		return
	}
	fmt.Fprint(w, "welcome home")
}

func Loginhandler(w http.ResponseWriter, r *http.Request) {
	// if r.URL.Path != "/smth/" {
	// 	ErrorHandler(w, r, )
	// 	return
	// }
	fmt.Fprint(w, "welcome smth ", r.URL.Query().Get("id"))
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	// if r.URL.Path != "/smth/" {
	// 	ErrorHandler(w, r, )
	// 	return
	// }
	fmt.Fprint(w, "welcome smth ", r.URL.Query().Get("id"))
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	// if r.URL.Path != "/smth/" {
	// 	ErrorHandler(w, r, )
	// 	return
	// }
	fmt.Fprint(w, "welcome smth ", r.URL.Query().Get("id"))
}
func CartHandler(w http.ResponseWriter, r *http.Request) {
	// if r.URL.Path != "/smth/" {
	// 	ErrorHandler(w, r, )
	// 	return
	// }
	fmt.Fprint(w, "welcome smth ", r.URL.Query().Get("id"))
}

func ErrorHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "custom 404")
}
