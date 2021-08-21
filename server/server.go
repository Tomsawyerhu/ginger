package server

import "net/http"

type Server struct {}

func (o *Server)HandleUpload(w http.ResponseWriter,r *http.Request){

}

func (o *Server)HandleDownload(w http.ResponseWriter,r *http.Request){

}

func InitServer(o *Server){}

func (o *Server)Start(){}