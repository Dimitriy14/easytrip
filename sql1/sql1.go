package sql1

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/astaxie/beego/logs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/oreuta/easytrip/clients"
	"github.com/oreuta/easytrip/models"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "bbdc4d9aa08941:ca3c3019@tcp(us-cdbr-iron-east-01.cleardb.net:3306)/heroku_d1f744b3705e71b")
	if err != nil {
		logs.Info("Connection open error: %v", err)
	}
	if err = db.Ping(); err != nil {
		logs.Info("Ping failed  %v", err)
	}
}

//Update updates information from BankUAclient to database
func Update() {
	a := clients.New()
	var err error
	var res []models.CurrencyBank
	for {
		res, err = a.GetCurrBank()
		if err != nil {
			logs.Info("err fo update sql is : %v", err)
		}

		for i := range res {
			result, err := db.Exec("update bankslist set RateBuy=?, RateSale=? where BankName=? and CodeAlpha=?", res[i].RateBuy, res[i].RateSale, res[i].BankName, res[i].CodeAlpha)
			if err != nil && result == nil {
				logs.Info("Number[%v] cant update to database: %v", res[i], err)
			}
		}
		logs.Info("Update DataBase Succesful!")
		time.Sleep(12 * time.Hour)
	}

}

//JsnChanger creates list of banks from database
func JsnChanger() (res []models.CurrencyBank, err error) {
	rows, err := db.Query("select * from bankslist")
	if err != nil {
		return nil, fmt.Errorf("Select query failed:%v", err)
	}
	var a models.CurrencyBank
	for rows.Next() {
		err := rows.Scan(&a.BankName, &a.CodeAlpha, &a.RateBuy, &a.RateSale)
		if err != nil {
			return nil, fmt.Errorf("Scan error:%v", err)
		}
		res = append(res, a)
	}
	return
}

func CheckUser(data models.User) (name string, ok bool) {
	rows, err := db.Query("SELECT count(Id),users.name FROM users where users.login=? and users.pass=?", data.Login, data.Password)
	if err != nil {
		return "", false
	}
	defer rows.Close()
	var a int
	err = rows.Scan(&a, &name)
	if err != nil || a != 1 {
		return "", false
	}
	return name, true

}

func InsertInto(data models.User) (res bool) {

	result, err := db.Exec("insert into users(name, login, pass) values(?,?,?)", data.Name, data.Login, data.Password)
	if err != nil {
		logs.Info("db.exec(Insert) err trouble : %v", err)
		return false
	}
	if result == nil {
		logs.Info("db.exec(Insert) result trouble: %v", result)
		return false
	}
	return true
}
