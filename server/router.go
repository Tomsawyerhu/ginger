package server

import (
	"net/http"
)

func (o *Server)InitRouter(){
	http.HandleFunc("/upload",o.HandleUpload)
	http.HandleFunc("/download",o.HandleDownload)
}
