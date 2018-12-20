package main

import (
	"os"
	"strconv"

	"github.com/oreuta/easytrip/sql1"

	"github.com/astaxie/beego"
	_ "github.com/oreuta/easytrip/clients"
	_ "github.com/oreuta/easytrip/routers"
)

func main() {
	var err error
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	beego.BConfig.Listen.HTTPPort, err = strconv.Atoi(port)
	beego.BConfig.WebConfig.Session.SessionOn = true
	if err != nil {
		panic(err)
	}
	go sql1.Update()
	beego.Run()
}
