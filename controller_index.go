package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func IndexController(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	totalRequests.WithLabelValues(r.RequestURI).Inc()

	LogRequest(200, r, l)
	fmt.Fprintln(w, "This API has a number of endpoints to look up fiat currency spot prices for Bitcoin.")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "API endpoints are:")
	fmt.Fprintln(w, "  - /health:         Health status endpoint")
	fmt.Fprintln(w, "  - /healthz:        Health status endpoint using standard Kubernetes naming convention")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "  - /usd or /USD:    Spot price lookup for US Dollars")
	fmt.Fprintln(w, "  - /eur or /EUR:    Spot price lookup for Euros")
	fmt.Fprintln(w, "  - /jpy or /JPY:    Spot price lookup for Japan Yen")
	fmt.Fprintln(w, "  - /gbp or /GBP:    Spot price lookup for Great Britain Pounds")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "  - /api/{currency}: Dyanmic endpoint to lookup the spot price for any standard 3 character currency abbreviation")
	fmt.Fprintln(w, "")
	fmt.Fprintln(w, "  - /metrics:        Prometheus metrics endpoint")
}
