package store

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/lib/pq"
	"github.com/parfenovvs/urlshortener/pkg/models/postgresql"
)

type StorageService struct {
	// redisClient   *redis.Client
	PostgresModel *postgresql.LinkModel
}

var (
	storeService = &StorageService{}
)

const CacheDuration = 6 * time.Hour

func InitializeStore() *StorageService {
	// redisClient := redis.NewClient(&redis.Options{
	// 	Addr:     "redis:6379",
	// 	Password: "",
	// 	DB:       0,
	// })

	// pong, err := redisClient.Ping().Result()
	// if err != nil {
	// 	panic(fmt.Sprintf("Error init Redis: %v", err))
	// }

	// fmt.Printf("\nRedis started successfully: pong message = {%s}", pong)
	// storeService.redisClient = redisClient

	postgresHost := getEnv("POSTGRES_HOST")
	postgresUser := getEnv("POSTGRES_USER")
	postgresPassword := getEnv("POSTGRES_PASSWORD")
	postgresPort := getEnv("POSTGRES_PORT")
	postgresDbName := getEnv("POSTGRES_DBNAME")
	psqlconn := fmt.Sprintf("host=%s user=%s password=%s port=%s sslmode=disable dbname=%s", postgresHost, postgresUser, postgresPassword, postgresPort, postgresDbName)
	db, _ := sql.Open("postgres", psqlconn)
	err := db.Ping()
	if err != nil {
		panic(fmt.Sprintf("\nError init PostgreSQL: %v", err))
	}

	fmt.Printf("\nPostgreSQL started successfully")
	storeService.PostgresModel = &postgresql.LinkModel{
		DB: db,
	}

	return storeService
}

func SaveUrlMapping(shortUrl string, originalUrl string, userId string) {
	// err := storeService.redisClient.Set(shortUrl, originalUrl, CacheDuration).Err()
	// if err != nil {
	// 	panic(fmt.Sprintf("Failed saving to Redis | Error: %v - shortUrl: %s - originalUrl: %s\n", err, shortUrl, originalUrl))
	// }

	err := storeService.PostgresModel.Insert(shortUrl, originalUrl)
	if err != nil {
		panic(fmt.Sprintf("Failed saving to PostgreSQL | Error: %v - shortUrl: %s - originalUrl: %s\n", err, shortUrl, originalUrl))
	}
}

func GetInitialUrl(shortUrl string) string {
	// result, err := storeService.redisClient.Get(shortUrl).Result()
	result, err := storeService.PostgresModel.Get(shortUrl)
	if err != nil {
		panic(fmt.Sprintf("Failed GetInitialUrl url | Error: %v - shortUrl: %s\n", err, shortUrl))
	}
	// return result
	return result.OriginalUrl
}

func getEnv(key string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		panic(fmt.Sprintf("Set %s environment variable", key))
	}
	return value
}
