package server

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"texts/config"
	"texts/pkg/connectiondatabase"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/redis/go-redis/v9"
)

type Server struct {
	cfg   *config.Config
	fiber *fiber.App
	pgDB  *connectiondatabase.Database
	redis *redis.Client
}

func NewServer(
	cfg *config.Config,
	pgDB *connectiondatabase.Database,
	redis *redis.Client,
) *Server {
	return &Server{
		cfg:   cfg,
		fiber: fiber.New(),
		pgDB:  pgDB,
		redis: redis,
	}
}

func (s *Server) Run() error {
	s.fiber.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173",
		AllowCredentials: true,
	}))
	s.MapHandlers()
	go func() {
		log.Println("Starting server...")

		err := s.fiber.Listen(s.cfg.Server.Host)
		if err != nil {
			log.Fatalf("Couldn't start server on %s, err=%v", s.cfg.Server.Host, err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down HTTP server...")
	err := s.fiber.Shutdown()
	if err != nil {
		log.Printf("Couldn't shut down server, err=%v", err)
	} else {
		log.Printf("HTTP server closed properly")
	}

	return nil
}
