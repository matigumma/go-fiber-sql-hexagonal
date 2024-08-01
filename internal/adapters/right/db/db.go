package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	// sq "github.com/Masterminds/squirrel"
	// "github.com/matigumma/go-fiber-sql-hexagonal/internal/ports"
	"github.com/joho/godotenv"
	_ "github.com/microsoft/go-mssqldb"
)

type DB struct {
	*sql.DB
}

var (
	USER          string
	PASSWORD      string
	DATABASE_URL  string
	PORT          string
	DATABASE_NAME string
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	} else {
		USER = os.Getenv("DB_USER")
		PASSWORD = os.Getenv("DB_PASSWORD")
		DATABASE_URL = os.Getenv("DB_SERVER")
		PORT = os.Getenv("DB_PORT")
		DATABASE_NAME = os.Getenv("DB_DATABASE")
		fmt.Println(".env file loaded successfully")
	}
}
func New() (*DB, error) {
	connection_string := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s&encrypt=true&TrustServerCertificate=true&tlsmin=1.0", USER, PASSWORD, DATABASE_URL, PORT, DATABASE_NAME)
	// fmt.Println("Connection String:", connection_string)

	db, err := sql.Open("sqlserver", connection_string)
	if err != nil {
		log.Fatal(err)
	}
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}
	return &DB{db}, nil
}

func (this DB) CloseDbConnection() {
	err := this.DB.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func fetchData(db *sql.DB, query string, args ...interface{}) ([]map[string]interface{}, error) {
	// Debugging: Print the type and value of flatArgs
	// fmt.Printf("FlatArgs type: %T, FlatArgs value: %v\n", args, args)

	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, fmt.Errorf("error fetching data: %w", err)
	}
	defer rows.Close()

	var results []map[string]interface{}

	columns, err := rows.Columns()
	if err != nil {
		return nil, fmt.Errorf("error fetching columns: %w", err)
	}

	values := make([]interface{}, len(columns))
	valuePtrs := make([]interface{}, len(columns))
	for i := range columns {
		valuePtrs[i] = &values[i]
	}

	for rows.Next() {
		if err := rows.Scan(valuePtrs...); err != nil {
			return nil, fmt.Errorf("error scanning data: %w", err)
		}

		rowMap := make(map[string]interface{})
		for i, col := range columns {
			val := values[i]
			if b, ok := val.([]byte); ok {
				rowMap[col] = string(b)
			} else {
				rowMap[col] = val
			}
		}
		results = append(results, rowMap)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error with rows: %w", err)
	}

	return results, nil
}

func getStringValue(row map[string]interface{}, key string) string {
	if val, ok := row[key].(string); ok {
		return val
	}
	return ""
}

func getInt64Value(row map[string]interface{}, key string) int64 {
	if val, ok := row[key].(int64); ok {
		return val
	}
	return 0
}
