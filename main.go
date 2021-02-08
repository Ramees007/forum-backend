package main

import "github.com/rameesThattarath/qaForum/server"

func main() {
	s := &server.Server{}
	s.Initialize(server.GetConfig())
	s.Run(":8080")
}
