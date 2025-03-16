package golang_database

import (
	"context"
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

	for rows.Next() {
		var id, name, email string
		var balance int32
		var rating float64
		var married bool
		var birthDate, created_at time.Time
		err := rows.Scan(&id, &name, &email, &balance, &rating, &birthDate, &created_at, &married)
		if err != nil {
			panic(err)
		}
		fmt.Println("===============================")
		fmt.Println("ID:", id)
		fmt.Println("Name:", name)	
		fmt.Println("Email:", email)	
		fmt.Println("Balance:", balance)	
		fmt.Println("rating:", rating)	
		fmt.Println("birth date:", birthDate)	
		fmt.Println("created_at:", created_at)	
		fmt.Println("married:", married)	
	}

	defer rows.Close()
}