package main

import (
	"context"
	"database/sql"
	"fmt"
	"math"
	"os"

	"github.com/DATA-DOG/go-txdb"
	_ "github.com/go-sql-driver/mysql"
)

var dsn = os.Getenv("SQL_DSN")

func init() {
	txdb.Register("txdb", "mysql", dsn)
}

func main() {
	ctx := context.Background()
	mDB, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		os.Exit(1)
	}

	tDB, err := sql.Open("txdb", "dummy")
	if err != nil {
		fmt.Printf("err: %v\n", err)
		os.Exit(1)
	}

	dbs := map[string]*sql.DB{
		"mysql": mDB,
		"txdb":  tDB,
	}
	for name, db := range dbs {
		var v uint64
		row := db.QueryRowContext(ctx, "SELECT ?", uint64(math.MaxUint64))
		err := row.Scan(&v)
		if err != nil {
			fmt.Printf("err: %v\n", err)
		}
		fmt.Printf("name: %s, math.MaxUint64: %d, v: %d\n", name, uint64(math.MaxUint64), v)
	}
}
