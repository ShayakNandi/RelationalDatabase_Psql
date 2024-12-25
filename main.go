package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/url"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	//Connect to PostgreSQL
	dataSourceName := url.URL{
		Scheme: "postgres",
		Host:   "localhost:5432",
		User:   url.UserPassword("user", "password"),
		Path:   "dbname",
	}
	//Indicate to PostgreSQL is we are using SLL mode
	query := dataSourceName.Query()
	query.Add("sslmode", "disable")

	//This creates some url that will be passed to the SQL open function (part of net package)
	dataSourceName.RawQuery = query.Encode()

	//Creates a pool of connections. Sql.Open doesn't actually connect to any db; it's a pool of connections. It checks the dsn and makes sure the value is correct
	db, err := sql.Open("pgx", dataSourceName.String())
	if err != nil {
		fmt.Println("sql.Open", err)
		return
	}
	//Close the connections after main runs
	defer func() {
		_ = db.Close()
		fmt.Println("closed")
	}()

	//Check if connection is alive
	if err := db.PingContext(context.Background()); err != nil {
		fmt.Println("db.PingContext", err)
		return
	}

	//Get Row from DB
	row := db.QueryRowContext(context.Background(), "SELECT bday_year FROM users WHERE name = 'shayak'")
	if err := row.Err(); err != nil {
		fmt.Println("db.QueryRowContext", err)
		return
	}

	//Pass row through scanner. row.Scan is given the address of bday_year.
	var bday_year int64
	if err := row.Scan(&bday_year); err != nil {
		fmt.Println("row.Scan", err)
		return
	}

	fmt.Println("bday_year", bday_year)

	/*Insert Records
	name := "karl"
	bday_year = 1903
	_, err = db.ExecContext(context.Background(), "INSERT INTO users(name, bday_year) VALUES($1, $2)", name, bday_year)

	if err != nil {
		fmt.Println("db.ExecContext", err)
		return
	}

	*/

	// Get Rows
	rows, err := db.QueryContext(context.Background(), "SELECT name, bday_year FROM users")
	if err != nil {
		fmt.Println("db.QueryContext", err)
		return
	}
	//Rows is like an iterator; needs to be closed at the end
	defer func() {
		_ = rows.Close()
	}()
	if rows.Err() != nil {
		fmt.Println("rows.Err()", err)
		return
	}

	//Iterate through rows
	for rows.Next() {
		var name string
		var bday_year int64

		if err := rows.Scan(&name, &bday_year); err != nil {
			fmt.Println("rows.Scan", err)
			return
		}

		fmt.Println("name", name, "bday_year", bday_year)
	}

}
