package golang_database

import (
	"context"
	"fmt"
	"testing"
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