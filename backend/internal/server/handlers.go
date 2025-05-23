package server

import (
	"texts/internal/delivery/commentary_delivery"
	"texts/internal/delivery/post_delivery"
	"texts/internal/delivery/user_delivery"
	"texts/internal/middleware"
	"texts/internal/repository/postgres/commentary_repository"
	"texts/internal/repository/postgres/post_repository"
	"texts/internal/repository/postgres/user_repository"
	"texts/internal/repository/redis/redis_repository"
	"texts/internal/usecase/commentary_usecase"
	"texts/internal/usecase/post_usecase"
	"texts/internal/usecase/user_usecase"

	"github.com/avito-tech/go-transaction-manager/sqlx"
	"github.com/avito-tech/go-transaction-manager/trm/manager"
)

func (s *Server) MapHandlers() {
	userPGRepo := user_repository.NewUserPGRepo(s.cfg, s.pgDB, sqlx.DefaultCtxGetter)
	userRedisRepo := redis_repository.NewUserRedisRepo(s.redis, s.cfg)

	postPGRepo := post_repository.NewPostPGRepo(s.cfg, s.pgDB, sqlx.DefaultCtxGetter)

	commentaryPGRepo := commentary_repository.NewCommentaryPGRepo(s.cfg, s.pgDB, sqlx.DefaultCtxGetter)

	trManager := manager.Must(sqlx.NewDefaultFactory(s.pgDB.GetPool()))

	userUC := user_usecase.NewUserUC(s.cfg, userPGRepo, userRedisRepo, trManager)
	postUC := post_usecase.NewPostUC(s.cfg, postPGRepo, trManager)
	commentaryUC := commentary_usecase.NewCommentaryUC(s.cfg, commentaryPGRepo, trManager)

	userDelivery := user_delivery.NewUserHandler(userUC, s.cfg)
	postDelivery := post_delivery.NewPostHandler(postUC, s.cfg)
	commentaryDelivery := commentary_delivery.NewCommentaryHandler(commentaryUC, s.cfg)

	mw := middleware.NewMiddlewareManager(s.cfg, userRedisRepo)

	userGroup := s.fiber.Group("user")
	postGroup := s.fiber.Group("post")
	commentaryGroup := s.fiber.Group("commentary")

	user_delivery.MapUserRoutes(userGroup, userDelivery, mw)
	post_delivery.MapPostRoutes(postGroup, postDelivery, mw)
	commentary_delivery.MapCommentaryRoutes(commentaryGroup, commentaryDelivery, mw)
}
