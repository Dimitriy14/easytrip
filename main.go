package main

import (
	"os"
	"strconv"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/oreuta/easytrip/routers"
	"github.com/oreuta/easytrip/sql1"
)

//var Db *sql.DB

func main() {
	var err error
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	beego.BConfig.Listen.HTTPPort, err = strconv.Atoi(port)
	if err != nil {
		panic(err)
	}
	Con := sql1.CreateConnect(beego.AppConfig.String("connect"))
	go sql1.Update(Con)
	defer Con.Close()
	beego.Run()
}
