package repository

import (
	"database/sql"

	"github.com/exange_rate_service/repository/model"

	_ "github.com/mattn/go-sqlite3"
)

type ExchangeRateRepository struct {
	driverName     string
	dataSourceName string
}

func (e *ExchangeRateRepository) Init() {
	e.driverName = "sqlite3"
	e.dataSourceName = "./foo.db"

	db, err := sql.Open(e.driverName, e.dataSourceName)
	defer db.Close()

	checkErr(err)

	rows, err := db.Query(`SELECT name FROM sqlite_master WHERE type='table' AND name='exchange_rate';`)
	rows.Close()
	if !rows.Next() {
		sql_table := `
		CREATE TABLE IF NOT EXISTS exchange_rate(
			uid INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL COLLATE NOCASE,
			cashBuyingRate REAL NOT NULL,
			cashSellingRate REAL NOT NULL,
			signBuyingRate REAL NOT NULL,
			signSellingRate REAL NOT NULL,
			timestamp DATETIME DEFAULT CURRENT_TIMESTAMP
		);
		`
		db.Exec(sql_table)
	}
}

func (e *ExchangeRateRepository) InsertOrUpdate(entity model.ExchangeRateDao) int64 {
	var result int64
	db, err := sql.Open(e.driverName, e.dataSourceName)
	defer db.Close()

	checkErr(err)

	row := e.Get(entity.Name)
	if row.Uid > 0 {
		stmt, err := db.Prepare("UPDATE exchange_rate SET cashBuyingRate = ?, cashSellingRate = ?, signBuyingRate = ?, signSellingRate = ? where name = ?")
		checkErr(err)

		res, err := stmt.Exec(entity.CashBuyingRate, entity.CashSellingRate, entity.SignBuyingRate, entity.SignSellingRate, entity.Name)
		checkErr(err)
		result, err = res.LastInsertId()
	} else {
		stmt, err := db.Prepare("INSERT INTO exchange_rate(name, cashBuyingRate, cashSellingRate, signBuyingRate, signSellingRate) values(?,?,?,?,?)")
		checkErr(err)

		res, err := stmt.Exec(entity.Name, entity.CashBuyingRate, entity.CashSellingRate, entity.SignBuyingRate, entity.SignSellingRate)
		checkErr(err)
		result, err = res.LastInsertId()
	}

	return result
}

func (e *ExchangeRateRepository) Get(name string) model.ExchangeRateDao {
	db, err := sql.Open(e.driverName, e.dataSourceName)
	defer db.Close()
	rows, err := db.Query(`SELECT 
	Name, CashBuyingRate, CashSellingRate, SignBuyingRate, SignSellingRate, datetime(TimeStamp, 'localtime')
	FROM exchange_rate where name=?`, name)
	defer rows.Close()
	checkErr(err)

	var result model.ExchangeRateDao
	for rows.Next() {
		var name string
		var cashBuyingRate float32
		var cashSellingRate float32
		var signBuyingRate float32
		var signSellingRate float32
		var timestamp string

		err = rows.Scan(&name, &cashBuyingRate, &cashSellingRate, &signBuyingRate, &signSellingRate, &timestamp)
		checkErr(err)

		result.Name = name
		result.CashBuyingRate = cashBuyingRate
		result.CashSellingRate = cashSellingRate
		result.SignBuyingRate = signBuyingRate
		result.SignSellingRate = signSellingRate
		result.TimeStamp = timestamp
	}

	return result
}

func (e *ExchangeRateRepository) List() []model.ExchangeRateDao {
	db, err := sql.Open(e.driverName, e.dataSourceName)
	defer db.Close()
	rows, err := db.Query("SELECT * FROM exchange_rate")
	rows.Close()
	checkErr(err)
	var result []model.ExchangeRateDao
	for rows.Next() {

		var row model.ExchangeRateDao
		err = rows.Scan(&row.Uid, &row.Name, &row.CashBuyingRate, &row.CashSellingRate, &row.SignBuyingRate, &row.SignSellingRate)
		checkErr(err)

		result = append(result, row)
	}

	return result
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
