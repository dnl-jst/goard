package routeHandler

import (
	"fmt"
	"time"
	"net/http"
	"go.iondynamics.net/webapp"
	"github.com/parnurzeal/gorequest"
)

func CheckHttp(w http.ResponseWriter, req *http.Request) *webapp.Error {

	method   := req.FormValue("method")
	protocol := req.FormValue("protocol")
	host     := req.FormValue("host")
	port     := req.FormValue("port")
	path     := req.FormValue("path")

	requestUrl := protocol + "://" + host

	if port != "" {
		requestUrl += ":" + port
	}

	requestUrl += path

	request := gorequest.New()

	var response *http.Response
	var body string
	var err []error

	start := time.Now()

	switch method {
		case "get":
			response, body, err = request.Get(requestUrl).End()
		case "post":
			response, body, err = request.Post(requestUrl).End()
		case "delete":
			response, body, err = request.Delete(requestUrl).End()
		case "head":
			response, body, err = request.Head(requestUrl).End()
		case "put":
			response, body, err = request.Put(requestUrl).End()
		case "patch":
			response, body, err = request.Patch(requestUrl).End()
	}

	if err != nil {
		writeError(w, "request failed")
	} else {

		elapsed := time.Since(start)

		writeSuccess(w, fmt.Sprintf("%s, %d bytes in %s", response.Status, len(body), elapsed))
	}

	return nil
}

func writeError(w http.ResponseWriter, status string) {
	fmt.Fprintf(w, status)
}

func writeWarning(w http.ResponseWriter, status string) {
	fmt.Fprintf(w, status)
}

func writeSuccess(w http.ResponseWriter, status string) {
	fmt.Fprintf(w, status)
}