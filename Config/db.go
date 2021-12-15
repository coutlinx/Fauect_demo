package Config

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "root:102400@tcp(127.0.0.1:3306)/faucet?parseTime=true&loc=Local")
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("已经连接MYSQL")
}

func InsertAdd(address, ip string, amount int64) error {
	now := time.Now()
	var sqlStr = "insert into transfer (address,amount,ip,time) values (?,?,?,?)"
	_, err := db.Exec(sqlStr, address, amount, ip, now)
	if err != nil {
		return err
	}
	return nil
}

func InsertCon(amount int64) error {
	now := time.Now()
	var sqlStr = "insert into amount (time,amount) values (?,?)"
	_, err := db.Exec(sqlStr, now, amount)
	if err != nil {
		return err
	}
	return nil
}
