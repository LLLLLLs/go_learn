// Time        : 2019/01/11
// Description :

package main

import (
	"github.com/gorilla/handlers"
	"golearn/util"
	"io"
	"net/http"
	"os"
)

func main() {
	http.Handle("/hello", handlers.CombinedLoggingHandler(os.Stdout, http.HandlerFunc(myHandler)))
	util.OkOrPanic(http.ListenAndServe(":1234", nil))
}

func myHandler(rw http.ResponseWriter, _ *http.Request) {
	rw.WriteHeader(http.StatusAccepted)
	_, err := io.WriteString(rw, "hello world")
	util.OkOrPanic(err)
}
