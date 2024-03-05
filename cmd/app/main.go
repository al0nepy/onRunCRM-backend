package main

import (
	"go.onRunCRM/config"
	"go.onRunCRM/internal/database/mongodb"
	"go.onRunCRM/internal/server"
)

func main() {
	config.LoadENV()
	mongodb.InitMongoDBConnection()
	server.RunGRPCServer()
}
