package main

import (
	"os"
	"strconv"

	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/config"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/oreuta/easytrip/routers"
	"github.com/oreuta/easytrip/sql1"
)

func main() {
	var err error
	port := os.Getenv("PORT")
	if port == "" {
		port = beego.AppConfig.String("HTTPPort")
	}

	beego.BConfig.Listen.HTTPPort, err = strconv.Atoi(port)
	if err != nil {
		panic(err)
	}
	sql1.Db = sql1.CreateConnect(beego.AppConfig.String("connect"))

	beego.Run()
}
