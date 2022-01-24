package handler

import (
	"fmt"
	"net/http"
	"webserver/helper"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	helper.LogRequest(r)

	fmt.Fprint(w, "OK")
}
