package main

//

// TODO imports ////////////////////////////////////////////////////////////////////////////////////////////////////////

import (
	"database/sql"
	"fmt"
	"github.com/xuri/excelize/v2"
	"log"
	"time"

	_ "github.com/lib/pq"
)

// TODO imports ////////////////////////////////////////////////////////////////////////////////////////////////////////

//

// TODO actions ////////////////////////////////////////////////////////////////////////////////////////////////////////

func getFromDb() ([]UserDb, error) {
	db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 user=gin_user password=12345Qwerty! dbname=gin_database sslmode=disable")
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

	rows, err := db.Query("select * from users_new")
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err = rows.Close()
		if err != nil {
			log.Fatal(err)
			return
		}
	}(rows)

	users := make([]UserDb, 0)
	for rows.Next() {
		u := UserDb{}
		err = rows.Scan(&u.Id, &u.Username, &u.DateRegister, &u.ActiveStatus)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func writeToExcel(users []UserDb) error {
	file := excelize.NewFile()
	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}()
	index, err := file.NewSheet("Sheet1")
	if err != nil {
		return err
	}
	file.SetActiveSheet(index)

	for indexRow, user := range users {
		for indexCol, value := range user.getMatrix() {
			chars, err := excelize.ColumnNumberToName(indexCol + 1)
			if err != nil {
				return err
			}

			err = file.SetCellValue("Sheet1", fmt.Sprintf("%s%d", chars, indexRow+1), value)
			if err != nil {
				return err
			}
		}
	}
	
	err = file.SaveAs("3_complex/3_get_db_write_excel/to/new.xlsx")
	if err != nil {
		return err
	}

	return nil
}

func run() {
	users, err := getFromDb()
	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(users)

	err = writeToExcel(users)
	if err != nil {
		log.Fatal(err)
		return
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

func (u UserDb) getMatrix() []any {
	return []any{u.Id, u.Username, u.DateRegister, u.ActiveStatus}
}

// TODO methods ////////////////////////////////////////////////////////////////////////////////////////////////////////

//

// TODO main ///////////////////////////////////////////////////////////////////////////////////////////////////////////
func main() {
	fmt.Println("start")
	run()
	fmt.Println("stop")
}

//sudo -i -u postgres
//createuser gin_user
//createdb gin_database -O gin_user
//psql gin_database
//alter user gin_user with password '12345Qwerty!';
//GRANT ALL PRIVILEGES ON DATABASE gin_database TO gin_user;
// GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public to gin_user;
// GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public to gin_user;
// create table users_new (id serial not null unique, username varchar(255) not null, date_register timestamp not null default now(), active_status bool default 'false');
// alter table users_new drop column registered_at;
// alter table users_new add column registered_at timestamp not null default now();
// \l
// \d
// insert into users_new (username, active_status) values ('Alice', 'false');
// select * from users_new;
// \q
// exit

// go get -u github.com/lib/pq
// go get -u github.com/xuri/excelize/v2
// TODO main ///////////////////////////////////////////////////////////////////////////////////////////////////////////
