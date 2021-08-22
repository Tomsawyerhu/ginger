package server

import "net/http"

type HttpHandler struct{
	serveMux *http.ServeMux
}

func (o HttpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	o.serveMux.ServeHTTP(w,r)
}
