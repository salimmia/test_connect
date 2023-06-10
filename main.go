package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main(){
	// connect to a database
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/test")
	if err != nil {
		log.Fatal(err)
	}
	
	defer db.Close()

	log.Println("Connected to database")

	// test my connection
	err = db.Ping()

	if err != nil{
		log.Fatal(err)
	}
	log.Println("Pinged Database")


	// get rows from my table
	err = getAllRows(db)
	if err != nil{
		log.Fatal(err)
	}

	// insert a row
	insertToUsersTable(db) /// successfully inserted
	// get rows from my table again
	err = getAllRows(db)
	if err != nil{
		log.Fatal(err)
	}

	// update a row
	updateToUsersTable(db) /// Successfully Updated function
	// get rows from my table again
	err = getAllRows(db)
	if err != nil{
		log.Fatal(err)
	}

	// get one row by id
	queryForSingleRow(db)

	// delete a row
	deleteRowToUsersTable(db)

	// get rows again
	err = getAllRows(db)
	if err != nil{
		log.Fatal(err)
	}
}

/// Verified
func getAllRows(db *sql.DB) error{
	rows, err := db.Query("select * from users")
	
	if err != nil{
		log.Fatal(err)
		return err 
	}
	defer rows.Close()

	var first_name, last_name, email, phone, gender string
	var id, age int

	for rows.Next(){
		err := rows.Scan(&id, &first_name, &last_name, &age, &gender, &email, &phone)
		if err != nil{
			log.Println(err)
			return err
		}
		fmt.Println("Record is:", id, first_name, last_name, age, gender, email, phone)
	}

	if err = rows.Err(); err != nil{
		log.Fatal("Error scanning rows", err)
	}

	fmt.Println("-------------------------------------------------------------------------")

	return nil
}

/// SussessFully inserted
func insertToUsersTable(db *sql.DB){
	query := `INSERT INTO users (first_name, last_name, age, gender, email, phone) VALUES ("Muhammad", "Tomal", "20", "Male", "halimmia313@gmail.com", "+8801314554546");`

	_, err := db.Exec(query)

	if err != nil{
		log.Fatal("Data inserting to table users failed ", err)
	}
}

/// Successfully Updated function
func updateToUsersTable(db *sql.DB){
	stmt := `update users set email = 'tomalmia3135@gmail.com', phone = '+8801733913528' where first_name = 'Muhammad' and last_name = 'Tomal'`

	_, err := db.Exec(stmt)

	if err != nil{
		log.Fatal("Data Updating to table users failed ", err)
	}
}

/// Verified
func queryForSingleRow(db *sql.DB){
	var firstName, lastName, email, phone, gender string
	var id, age int

	query := ("select * from users where id = '1'")

	row := db.QueryRow(query)

	err := row.Scan(&id, &firstName, &lastName, &age, &gender, &email, &phone)

	if err != nil{
		log.Fatal("One row doesn't return ", err)
	}

	log.Println("QueryRow returns", id, firstName, lastName, age, gender, email, phone)
}

/// Verified
func deleteRowToUsersTable(db *sql.DB){
	query := ("delete from users where id > '3'")

	_, err := db.Exec(query)

	if err != nil{
		log.Fatal("Deleted doesn't work ", err)
	}

	log.Println("Successfully Deleted")
}