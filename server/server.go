package server

import (
	"net/http"
)

type Server struct {
	serveMux *http.ServeMux
	http.Server
}

func (o *Server)HandleUpload(w http.ResponseWriter,r *http.Request){
	_, _ = w.Write([]byte("upload received"))
}

func (o *Server)HandleDownload(w http.ResponseWriter,r *http.Request){

}

func (o *Server) InitServer(){
	o.Addr="127.0.0.1:8080"
	o.serveMux=http.NewServeMux()

	handler:=new(HttpHandler)
	handler.serveMux=o.serveMux
	o.Handler=handler
	o.InitRouter()
}

func (o *Server)Start(){
	_ = o.ListenAndServe()
}