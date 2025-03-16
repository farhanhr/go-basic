package golang_database

import (
	"context"
	"database/sql"
	"fmt"
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

	for rows.Next() {
		var id, name string
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println("ID:", id)
		fmt.Println("Name:", name)	
	}

	defer rows.Close()
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