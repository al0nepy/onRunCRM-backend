package main

import (
	"go.onRunCRM/config"
	"go.onRunCRM/internal/server"
)

func main() {
	config.LoadENV()
	server.RunGRPCServer()
}
