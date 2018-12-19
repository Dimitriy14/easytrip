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

/*func init() {
	orm.RegisterDriver("mysql", orm.DR_MySQL)
	mysqlReg := beego.AppConfig.String("mysqluser") + ":" +
		beego.AppConfig.String("mysqlpass") + "@tcp(127.0.0.1:3306)/" +
		beego.AppConfig.String("mysqldb")
	orm.RegisterDataBase("default", "mysql", mysqlReg)
}*/

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

	go sql1.Update(sql1.Db)
	beego.Run()
}
