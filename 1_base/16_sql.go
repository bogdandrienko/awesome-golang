package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type User7 struct {
	ID         int64
	Name       string
	Email      string
	Password   string
	RegisterAt time.Time
}

func query1() {
	// sudo docker ps
	// sudo docker run -d --name postgres-docker -e POSTGRES_PASSWORD=31284bogdan -v ${HOME}/pgdata/:/var/lib/postgresql/data -p 5432:5432 postgres
	// sudo docker -it ninja-db bash

	//sudo -i -u postgres
	//createuser gin_user
	//createdb gin_database -O gin_user
	//psql gin_database
	//alter users gin_user with password '12345Qwerty!';
	//GRANT ALL PRIVILEGES ON DATABASE gin_database TO gin_user;
	// GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public to gin_user;
	// GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public to gin_user;
	// create table users (id serial not null unique, name varchar(255) not null, email varchar(255) not null, password varchar(255) not null, registered_at timestamp not null default now());
	// alter table users drop column registered_at;
	// alter table users add column registered_at timestamp not null default now();
	// \l
	// \d
	// insert into users (name, email, password) values ('Alice', 'alice@gmail.com', '12345Qwerty!');
	// select * from users;
	// \q
	// exit

	// go get -u github.com/lib/pq

	// mysql
	db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 users=gin_user password=12345Qwerty! dbname=gin_database sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	fmt.Println(db)

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("successfully connected")

	rows, err := db.Query("select * from users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	users := make([]User7, 0)
	for rows.Next() {
		u := User7{}
		err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.RegisterAt)
		if err != nil {
			log.Fatal(err)
		}

		users = append(users, u)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(users)
}

func query2() {
	db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 users=gin_user password=12345Qwerty! dbname=gin_database sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	user := User7{}
	// err = db.QueryRow(fmt.Sprintf("select * from users where id = %d", 1)) // TODO SQL injection
	err = db.QueryRow("select * from users where id = $1", 1).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.RegisterAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			fmt.Println("not have a data")
			return
		}
		log.Fatal(err)
	}
	fmt.Println(user)
}

func query3() {
	db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 users=gin_user password=12345Qwerty! dbname=gin_database sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(getUsers1(db))

	//fmt.Println(getUser1(db, 1))

	//err = insertUser1(db, User7{
	//	Name:     "Petya",
	//	Email:    "petya@gmail.com",
	//	Password: "12345Qwerty!",
	//})
	//if err != nil {
	//	log.Fatal(err)
	//}

	//fmt.Println(deleteUser1(db, 2))

	//fmt.Println(updateUser1(db, 1, User7{Name: "Polina", Email: "polina@gmail.com"}))
}

func getUsers1(db *sql.DB) ([]User7, error) {
	rows, err := db.Query("select * from users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	users := make([]User7, 0)
	for rows.Next() {
		u := User7{}
		err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.RegisterAt)
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

func getUser1(db *sql.DB, id int) (User7, error) {
	user := User7{}
	// err = db.QueryRow(fmt.Sprintf("select * from users where id = %d", 1)) // TODO SQL injection
	err := db.QueryRow("select * from users where id = $1", id).Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.RegisterAt)
	return user, err
}

func insertUser1(db *sql.DB, u User7) error {
	// TODO transactions
	// begin;
	// delete from users where id = 1;
	// rollback;

	// commit;

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec("insert into users (name, email, password) values ($1, $2, $3);", u.Name, u.Email, u.Password)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func deleteUser1(db *sql.DB, id int) error {
	_, err := db.Exec("delete from users where id = $1;", id)
	return err
}

func updateUser1(db *sql.DB, id int, newUser User7) error {
	_, err := db.Exec("update users set name=$1, email=$2 where id = $3;", newUser.Name, newUser.Email, id)
	return err
}

func main() {
	//query1()
	//query2()
	query3()
}
