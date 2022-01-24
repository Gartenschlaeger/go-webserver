package handlers

import (
	"fmt"
	"net/http"

	"github.com/KaiGartenschlaeger/go-webserver/pkg/logging"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	logging.LogRequest(r)

	fmt.Fprint(w, "OK")
}
