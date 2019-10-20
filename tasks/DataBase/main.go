package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func userList(db *sql.DB)  {
	rows, err := db.Query("select id, fname, lname from user_list")

	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var lname string

		err = rows.Scan(&id, &name, &lname)

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(id, name, lname)
	}
}

func userInsert(db *sql.DB, fname, lname string)  {
	sqlStatement := `insert into user_list (fname, lname) values ($1, $2)`

	_, err := db.Exec(sqlStatement, fname, lname)

	if err != nil {
		panic(err)
	}
}

func userUpdate(db * sql.DB, fname, lname string, userId int)  {
	sqlStatemaent := `update user_list set fname=$1, lname=$2 where id=$3`

	_, err := db.Exec(sqlStatemaent, fname, lname, userId)

	if err != nil {
		panic(err)
	}
}

func userDelete(db *sql.DB, userId int)  {
	sqlStatement := `delete from user_list where id=$1`

	_, err := db.Exec(sqlStatement, userId)

	if err != nil {
		panic(err)
	}
}

func contactList(db *sql.DB)  {
	rows, err := db.Query("select id, name, phone_number, (select fname || ' ' || lname from user_list where id = created_by) from contact_list")

	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var phone_number string
		var created_by string

		err = rows.Scan(&id, &name, &phone_number, &created_by)

		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(id, name, phone_number, created_by)
	}
}

func contactInsert(db *sql.DB, contactName, phoneNumber string, createdBy int)  {
	sqlStatement := `insert into contact_list (name, phone_number, created_by) values ($1, $2, $3)`

	_, err := db.Exec(sqlStatement, contactName, phoneNumber, createdBy)

	if err != nil {
		panic(err)
	}
}

func contactUpdate(db *sql.DB, contactName, phoneNumber string, createdBy int, contactId int)  {
	sqlStatement := `update contact_list set name=$1, phone_number=$2, created_by=$3 where id=$4`

	_, err := db.Exec(sqlStatement, contactName, phoneNumber, createdBy, contactId)

	if err != nil {
		panic(err)
	}
}

func contactDelete(db * sql.DB, contactId int)  {
	sqlStatement := `delete from contact_list where id=$1`

	_, err := db.Exec(sqlStatement, contactId)

	if err != nil {
		panic(err)
	}
}

func main() {
	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "1"
		dbname   = "postgres"
	)

	dbinfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", dbinfo)

	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()

	//insertUser(db, "user", "user")
	//userUpdate(db, "userUpdate", "userUpdate", 1)
	//userDelete(db, 10)
	userList(db)

	//contactInsert(db, "contact", "+998998023940", 3)
	//contactUpdate(db, "contactUpdate", "+998988023941", 11, 8)
	//contactDelete(db, 8)
	contactList(db)
}

