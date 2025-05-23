package commentary_delivery

import (
	"context"
	"texts/config"
	"texts/internal/usecase/commentary_usecase"

	"github.com/gofiber/fiber/v2"
)

type (
	CommentaryUC interface {
		CreateCommentary(ctx context.Context, createCommentaryReq commentary_usecase.CreateCommentaryReq) (int64, error)
		GetCommentaries(ctx context.Context, getCommentarysReq commentary_usecase.GetCommentarysReq) ([]commentary_usecase.GetCommentaryResp, error)
	}

	Handlers interface {
		GetCommentaries() fiber.Handler
		CreateCommentary() fiber.Handler
	}
)

type CommentaryHandlers struct {
	commentaryUC CommentaryUC
	cfg          *config.Config
}

func NewCommentaryHandler(commentaryUC CommentaryUC, cfg *config.Config) *CommentaryHandlers {
	return &CommentaryHandlers{
		cfg:          cfg,
		commentaryUC: commentaryUC,
	}
}
