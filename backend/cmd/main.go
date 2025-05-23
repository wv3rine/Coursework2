package main

import (
	"texts/config"
	"texts/internal/server"
	"texts/pkg/connectiondatabase"
	"texts/pkg/postgresConnector"
	"texts/pkg/redisConnector"

	"github.com/gofiber/fiber/v2/log"
)

func main() {
	cfg := config.Config{
		Server: struct{ Host string }{"localhost:8000"},
	}
	cfgPostgres := postgresConnector.Config{
		Host:     "localhost",
		Port:     "5432",
		User:     "postgres",
		Password: "12345",
		DBName:   "texts",
		// PGDriver: "postgres",
		SSLMode: "disable",
		Settings: postgresConnector.Settings{
			MaxOpenConns:    60,
			ConnMaxLifetime: 120,
			MaxIdleConns:    30,
			ConnMaxIdleTime: 20,
		},
	}

	cfgRedis := redisConnector.Config{
		Host: "localhost",
		Port: "6379",
	}

	log.Info("Setting up postgres...")
	sqlDB, err := postgresConnector.GetConnection(cfgPostgres)
	if err != nil {
		log.Fatalf("PostgreSQL init error: %s", err)
		return
	}
	log.Info("Success!")

	log.Info("Setting up DB connection...")
	psqlDB := connectiondatabase.NewDB(sqlDB)

	defer func(psqlDB *connectiondatabase.Database) {
		if err := psqlDB.Close(); err != nil {
			log.Errorf(err.Error())
		} else {
			log.Errorf("PostgresSQL closed properly")
		}
	}(psqlDB)

	log.Infof("Trying to connect to redis, host=%s", cfgRedis.Host)
	redis, err := redisConnector.GetConnection(cfgRedis)
	if err != nil {
		log.Fatalf("Failed to connect to redis: %s", err.Error())
	}
	log.Info("Success!")

	s := server.NewServer(
		&cfg,
		psqlDB,
		redis,
	)

	if err = s.Run(); err != nil {
		log.Errorf("Cannot start server: %v", err)
		return
	}
}
