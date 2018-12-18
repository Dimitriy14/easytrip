package sql1

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/oreuta/easytrip/clients"
	"github.com/oreuta/easytrip/models"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("mysql", "bbdc4d9aa08941:ca3c3019@tcp(us-cdbr-iron-east-01.cleardb.net:3306)/heroku_d1f744b3705e71b")
	if err != nil {
		fmt.Errorf("Connection open error: %v", err)
	}

	if err = Db.Ping(); err != nil {
		log.Fatal(err)
	}

}

//CreateConnect open connection to MySQL database
func CreateConnect(connection string) (Db *sql.DB) {
	var err error
	Db, err = sql.Open("mysql", connection)
	if err != nil {
		fmt.Errorf("Connection open error: %v", err)
	}
	return
}

//Update updates information from BankUAclient to database
func Update() error {
	a := clients.New()
	var err error
	var res []models.CurrencyBank
	res, err = a.GetCurrBank()
	if err != nil {
		fmt.Printf("err fo update sql is : %v", err)
	}
	for {
		for i := range res {
			_, err := Db.Query("update BanksList set RateBuy=? RateSale=? where BankName=? and CodeAlpha=?", res[i].RateBuy, res[i].RateSale, res[i].BankName, res[i].CodeAlpha)
			if err != nil {
				return fmt.Errorf("Number[%v] cant update to database: %v", res[i], err)

			}
		}
		time.Sleep(3 * time.Second)
	}

}

//JsnChanger creates list of banks from database
func JsnChanger() (res []models.CurrencyBank, err error) {
	rows, err := Db.Query("select * from BanksList")
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

func CheckUser(data models.User) bool {
	rows, err := Db.Query("SELECT * FROM users where login=? and pass=?", data.Login, data.Password)
	if err != nil {
		return false
	}
	defer rows.Close()
	var a, b string
	err = rows.Scan(&a, &b)
	if err != nil {
		return false
	}
	if a == "" && b == "" {
		return false
	}
	return true
}

func InsertInto(data models.User) (res bool) {
	_, err := Db.Exec("insert into users values(?,?,?)", data.Name, data.Login, data.Password)
	if err != nil {
		return false
	}
	return true
}
