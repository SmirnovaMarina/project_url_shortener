package store

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"os"
)

// Define the struct wrapper around Postgres connection
type StorageService struct {
	psqlConnectionPool *pgxpool.Pool
}

var (
	storeService = &StorageService{}
	ctx = context.Background()
)

func InitializeStore() *StorageService {

	host := "localhost"
	port := 5432
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")

	url := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", user, password, host, port, dbName)
	connectionPool, err := pgxpool.Connect(context.Background(), url)
	if err != nil {
		panic(fmt.Sprintf("Unable to connect to database: %v", err))
	}

	storeService.psqlConnectionPool = connectionPool
	return storeService
}

func IsUrlUnique(id int64) bool {
	var exists bool

	query := `select exists (select 1 from long_urls where id=$1)`
	result := storeService.psqlConnectionPool.QueryRow(ctx, query, id).Scan(&exists)
	if result != nil {
		panic(result)
	}

	return exists
}

func SaveUrlMapping(id int64, originalUrl string){
	query := `insert into long_urls values ($1, $2);`
	_, err := storeService.psqlConnectionPool.Exec(ctx, query, id, originalUrl)
	if err != nil {
		panic(err)
	}
}

func RetrieveOriginalUrl(id int64) string {
	var originalUrl string

	query := `select original_url from long_urls where id=$1`
	row := storeService.psqlConnectionPool.QueryRow(ctx, query, id)
	switch err := row.Scan(&originalUrl); err {
	case pgx.ErrNoRows:
		return "No rows were returned!"
	case nil:
		return originalUrl
	default:
		panic(err)
	}
}
