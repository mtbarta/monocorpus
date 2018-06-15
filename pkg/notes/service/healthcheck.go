package service

import (
	"encoding/json"
	"net/http"
)

type Alive struct {
	Alive bool
}

func HealthCheckHandlerFunc(w http.ResponseWriter, r *http.Request) {
	// A very simple health check.

	w.Header().Set("Content-Type", "application/json")

	// In the future we could report back on the status of our DB, or our cache
	// (e.g. Redis) by performing a simple PING, and include them in the response.
	alive := Alive{true}

	json.NewEncoder(w).Encode(alive)

}
