package main

import (
	"os"

	"github.com/astaxie/beego"
	_ "github.com/oreuta/easytrip/routers"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	err := beego.AppConfig.Set("httpport", port)
	if err != nil {
		panic(err)
	}

	beego.Run()
}
