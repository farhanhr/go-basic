package golang_database

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestExectSql(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	qry := "INSERT INTO customers(id, name) VALUES('ID001', 'Farhan')"

	_, err := db.ExecContext(ctx, qry)

	if err != nil {
		panic(err)
	}

	fmt.Println("Success Insert New Customer")
}

func TestQuerySql(t *testing.T)  {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	qry := "SELECT id, name FROM customers"

	rows, err := db.QueryContext(ctx, qry)

	if err != nil {
		panic(err)
	}

	defer rows.Close()
	for rows.Next() {
		var id, name string
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println("ID:", id)
		fmt.Println("Name:", name)	
	}

}

func TestSqlComplex(t *testing.T) {
	
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	qry := "SELECT id, name, email, balance, rating, birth_date,  created_at, married FROM customers"

	rows, err := db.QueryContext(ctx, qry)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var id, name string
		var email sql.NullString
		var balance int32
		var rating float64
		var married bool
		var created_at time.Time
		var birthDate sql.NullTime
		err := rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &created_at, &married)
		if err != nil {
			panic(err)
		}
		fmt.Println("===============================")
		fmt.Println("ID:", id)
		fmt.Println("Name:", name)	
		if email.Valid {
			fmt.Println("Email:", email.String)	
		}
		fmt.Println("Balance:", balance)	
		fmt.Println("rating:", rating)	
		if birthDate.Valid{
			fmt.Println("birth date:", birthDate.Time)	
		}
		fmt.Println("created_at:", created_at)	
		fmt.Println("married:", married)	
	}
}

func TestSqlInjection(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin'; #"
	password := "wrong"

	qry := "SELECT username FROM users WHERE username = '"+ username +
	"' AND password = '"+ password +"' LIMIT 1"

	rows, err := db.QueryContext(ctx, qry)

	if err != nil {
		panic(err)
	}

	defer rows.Close()
	
	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Selamat datang", username)
	} else {
		fmt.Println("Login Gagal")
	}

}

func TestSqlInjectionSafe(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()

	username := "admin'; #"
	password := "wrong"

	qry := "SELECT username FROM users WHERE username = ? AND password = ?  LIMIT 1"

	rows, err := db.QueryContext(ctx, qry, username, password)

	if err != nil {
		panic(err)
	}

	defer rows.Close()
	
	if rows.Next() {
		var username string
		err := rows.Scan(&username)
		if err != nil {
			panic(err)
		}
		fmt.Println("Selamat datang", username)
	} else {
		fmt.Println("Login Gagal")
	}

}

func TestTransaction(t *testing.T) {
	db := GetConnection()
	defer db.Close()

	ctx := context.Background()
	tx, err := db.Begin()

	if err != nil {
		panic(err)
	}

	script := "INSERT INTO comments(email, comment) VALUES (?,?)"
	// do transaction
	for i := 0; i<10; i++ {
		email := "Farhan" + strconv.Itoa(i) + "@mail.com"
		comment := "Komentar" + strconv.Itoa(i)

		result, err := tx.ExecContext(ctx, script, email, comment)
		if err != nil {
			panic(err)
		}

		id, err := result.LastInsertId()
		if err != nil {
			panic(err)
		}

		fmt.Println("comment id", id)
	}
	
	
	err = tx.Commit()
	
	if err != nil {
		panic(err)
	}

}