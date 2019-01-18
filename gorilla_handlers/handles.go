// Time        : 2019/01/11
// Description :

package main

import (
	"github.com/gorilla/handlers"
	"io"
	"net/http"
	"os"
)

func main() {
	http.Handle("/hello", handlers.CombinedLoggingHandler(os.Stdout, http.HandlerFunc(myHandler)))
	http.ListenAndServe(":1234", nil)
}

func myHandler(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(http.StatusAccepted)
	io.WriteString(rw, "hello world")
}
