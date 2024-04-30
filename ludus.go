package main

import (
	"github.com/GustavoFreitas22/ludus_sys/config"
	"github.com/GustavoFreitas22/ludus_sys/controller"
)

func main() {
	config.InitDatabase()
	controller.Init()
}
