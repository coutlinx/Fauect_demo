package Config

import (
	"database/sql"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
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

func SelectAllAddress() ([]common.Address, error) {
	var Address []common.Address
	var sqlStr = "select address from transfer  where time >= date(now()) and time < DATE_ADD(date(now()),INTERVAL 1 DAY) group by address"
	row, err := db.Query(sqlStr)
	if err != nil {
		return nil, err
	}
	for row.Next() {
		var acc common.Address
		err = row.Scan(acc)
		if err != nil {
			return nil, err
		}
		Address = append(Address, acc)
	}
	return Address, nil
}
