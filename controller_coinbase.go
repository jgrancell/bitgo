package main

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
)

func CoinbaseController(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	totalRequests.WithLabelValues(r.RequestURI).Inc()

	var currency string
	// Perform lookup from Coinbase
	if len(p.ByName("currency")) == 0 {
		currency = strings.Trim(r.RequestURI, "/")
	} else {
		currency = p.ByName("currency")
	}

	data, err := CallCoinbase(strings.ToUpper(currency))
	if err != nil {
		LogRequest(500, r, l)
		l.Error(err)
	} else {
		LogRequest(200, r, l)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusFound)
	json.NewEncoder(w).Encode(data)
}
