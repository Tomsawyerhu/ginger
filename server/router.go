package server

func (o *Server)InitRouter(){
	o.serveMux.HandleFunc("/upload",o.HandleUpload)
	o.serveMux.HandleFunc("/download",o.HandleDownload)
}
