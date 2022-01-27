package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func HealthController(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	totalRequests.WithLabelValues(r.RequestURI).Inc()

	LogRequest(200, r, l)
	fmt.Fprintln(w, "healthy")
}
