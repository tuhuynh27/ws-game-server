package main

import (
	"github.com/oddx-team/odd-game-server/configs"
	"github.com/oddx-team/odd-game-server/pkg/server"
)

func main() {
	configs.InitMongo()
	server.Serve()
}
