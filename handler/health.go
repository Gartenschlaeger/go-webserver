package handler

import (
	"fmt"
	"net/http"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	logRequest(r)
	fmt.Fprint(w, "OK")
}
