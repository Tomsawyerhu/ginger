package main

import "ginger/server"

func main() {
	s:=new(server.Server)
	s.InitServer()
	s.Start()
}
