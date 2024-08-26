package main

import (
	_ "personal-api/docs"
	Server "personal-api/src"
	Env "personal-api/src/utils"
)

// @title API using Gin
// @version 1.0
// @description Documentation
// @host localhost:3000
// @BasePath /
// @schemes http https
func main() {

	//Env;
	Env.Load()

	//Server;
	Server.Run()
}
