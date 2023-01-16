package main

//

// TODO imports ////////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"database/sql"
	"fmt"
	"github.com/xuri/excelize/v2"
	"log"
	"strconv"
	"time"

	_ "github.com/lib/pq"
)

// TODO imports ////////////////////////////////////////////////////////////////////////////////////////////////////////

//

// TODO actions ////////////////////////////////////////////////////////////////////////////////////////////////////////

func getFromExcel() ([]UserDb, error) {
	f, err := excelize.OpenFile("3_complex/6_read_excel_write_sql/to/data.xlsx")
	if err != nil {
		return nil, err
	}

	//rows, err := f.GetRows("Sheet1")
	matrix, err := f.GetRows("Sheet1", excelize.Options{RawCellValue: true})
	if err != nil {
		return nil, err
	}

	rows := matrix[1:]
	users := make([]UserDb, len(rows))
	for rowIndex, row := range rows {
		var id int
		var username string
		var dateRegister time.Time
		var activeStatus bool

		id = rowIndex
		for columnIndex, cell := range row {
			//fmt.Println(cell, reflect.TypeOf(cell))
			switch columnIndex {
			case 0:
				username = cell
			case 1:
				dateRegisterF, err := strconv.ParseFloat(row[1], 64)
				if err != nil {
					return nil, err
				}
				dateRegister, err = excelize.ExcelDateToTime(dateRegisterF, false)
				if err != nil {
					return nil, err
				}
			case 2:
				if row[2] == "TRUE" {
					activeStatus = true
				} else {
					activeStatus = false
				}
			}
		}

		users[id] = UserDb{Id: id, Username: username, DateRegister: dateRegister, ActiveStatus: activeStatus}
	}

	return users, nil
}

func writeToDb(users []UserDb) error {
	db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 users=gin_user password=12345Qwerty! dbname=gin_database sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err = db.Close()
		if err != nil {
			log.Fatal(err)
			return
		}
	}(db)
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	for _, user := range users {
		_, err = tx.Exec("insert into users_new (username, date_register, active_status) values ($1, $2, $3);",
			user.Username, user.DateRegister, user.ActiveStatus)
		if err != nil {
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func run() {
	users, err := getFromExcel()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(users)

	err = writeToDb(users)
	if err != nil {
		log.Fatal(err)
	}
}

// TODO actions ////////////////////////////////////////////////////////////////////////////////////////////////////////

//

// TODO structs ////////////////////////////////////////////////////////////////////////////////////////////////////////

type UserDb struct {
	Id           int
	Username     string
	DateRegister time.Time
	ActiveStatus bool
}

// TODO structs ////////////////////////////////////////////////////////////////////////////////////////////////////////

//

// TODO methods ////////////////////////////////////////////////////////////////////////////////////////////////////////

// TODO methods ////////////////////////////////////////////////////////////////////////////////////////////////////////

//

// TODO main ///////////////////////////////////////////////////////////////////////////////////////////////////////////

func main() {
	fmt.Println("start")
	run()
	fmt.Println("stop")
}

// TODO main ///////////////////////////////////////////////////////////////////////////////////////////////////////////
