package handler

import "net/http"

type HandlerInterface interface {
	Serve(w http.ResponseWriter, r *http.Request)
}
