package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"net/http"
)

var db *sql.DB
var err error

func main() {
	db, err = sql.Open("mysql", "ID:PASS@tcp(URL:3306)/test?charset=utf8")
	check(err)
	defer db.Close()

	err = db.Ping()
	check(err)

	http.HandleFunc("/", index)
	http.HandleFunc("/amigos", amigos)
	http.HandleFunc("/create", create)
	http.HandleFunc("/insert", insert)
	http.HandleFunc("/read", read)
	http.HandleFunc("/update", update)
	http.HandleFunc("/delete", del)
	http.HandleFunc("/drop", drop)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err := http.ListenAndServe(":8080", nil)
	check(err)
}

func index (w http.ResponseWriter, r *http.Request) {
	_, err = io.WriteString(w, "Successfully completed")
	check(err)
}

func amigos(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query(`SELECT aName FROM amigos`)
	check(err)

	var s, name string
	s = "RETRIEVED RECORDS:\n"

	for rows.Next() {
		err = rows.Scan(&name)
		check(err)
		s += name + "\n"
	}
	fmt.Fprintln(w, s)
}

func create(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Prepare(`CREATE TABLE customer (name VARCHAR(20));`)
	check(err)

	ro, err := stmt.Exec()
	check(err)

	n, err := ro.RowsAffected()
	check(err)

	fmt.Fprintln(w, "Created Table customer", n)
}

func insert(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Prepare(`INSERT INTO customer VALUES ("James");`)
	check(err)

	ro, err := stmt.Exec()
	check(err)

	n, err := ro.RowsAffected()
	check(err)

	fmt.Fprintln(w, "Inserted Record", n)
}

func read(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query(`SELECT * FROM customer`)
	check(err)

	var name string
	for rows.Next() {
		err = rows.Scan(&name)
		check(err)
		fmt.Println(name)

		fmt.Fprintln(w, "Retreived Record:", name)
	}
}

func update(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Prepare(`UPDATE customer SET name="jimmy"`)
	check(err)

	ro, err := stmt.Exec()
	check(err)

	n, err := ro.RowsAffected()
	check(err)

	fmt.Fprintln(w, "Updated Record", n)
}

func del(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Prepare(`DELETE FROM customer WHERE name="jimmy";`)
	check(err)

	ro, err := stmt.Exec()
	check(err)

	n, err := ro.RowsAffected()
	check(err)

	fmt.Fprintln(w, "Deleted Record", n)
}

func drop(w http.ResponseWriter, r *http.Request) {
	stmt, err := db.Prepare(`DROP TABLE customer;`)
	check(err)

	_, err = stmt.Exec()
	check(err)

	fmt.Fprintln(w, "Dropped Table customer")
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}