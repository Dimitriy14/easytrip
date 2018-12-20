package repository

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
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

func CheckUser(data models.User) (user models.User, err error) {
	data.Password = getMD5Hash(data.Password)
	rows, err1 := db.Query("SELECT users.name,users.login,users.pass FROM users where users.login=? and users.pass=?", data.Login, data.Password)
	if err != nil {
		return user, fmt.Errorf("query select failed err:%v\n", err1)
	}
	if rows.Next() {
		err = rows.Scan(&user.Name, &user.Login, &user.Password)
		if err != nil {
			return user, fmt.Errorf("scan err:%v\n", err)
		}
	}
	return user, err
}

func InsertInto(data models.User) (err error) {
	data.Password = getMD5Hash(data.Password)
	_, err = db.Exec("insert into users(name, login, pass) values(?,?,?)", data.Name, data.Login, data.Password)
	if err != nil {
		return fmt.Errorf("db.exec(Insert) err trouble : %v", err)
	}

	return err
}

func InsertHist(data models.User, req models.MainRequest, cont string) error {
	var id int

	link := requestConvert(req, cont)

	rows, err := db.Query("select user.id from users where users.login=? and users.pass=?", data.Login, data.Password)
	if err != nil {
		return fmt.Errorf("InsertHist: Query error:%v", err)
	}
	if rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			return fmt.Errorf("InsertHist: rows.Scan error:%v", err)
		}
	}

	_, err = db.Exec("Insert into history values(?,?)", id, link)
	if err != nil {
		return fmt.Errorf("InsertHist: Exec error:%v", err)
	}
	return nil
}

func getMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func requestConvert(req models.MainRequest, opt string) (res string) {

	res += "/" + opt + "?"
	for i := range req.Bank {
		res += "bank=" + req.Bank[i] + "&"
	}
	for i := range req.Currency {
		res += "currency=" + req.Bank[i] + "&"
	}
	res += "option=" + req.Option
	return
}
