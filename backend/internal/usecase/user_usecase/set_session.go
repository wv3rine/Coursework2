package user_usecase

import (
	"context"
	"strconv"
	"texts/internal/repository/redis/redis_repository"
	"time"

	"github.com/google/uuid"
)


type Session struct {
	SessionKey string
	TTL        time.Duration
	UserID int64
	Role string
}

func (u *UserUC) setSession(
	ctx context.Context,
	userAgent string,
	userID int64,
	userRole string,
) (Session, error) {
	response := Session{
		SessionKey: strconv.Itoa(int(userID)) + ":" + uuid.NewString(),
		TTL:        time.Hour * 24,
		UserID: userID,
		Role: userRole,
	}

	err := u.userRedisRepo.Set(
		ctx,
		response.SessionKey,
		redis_repository.UserSession{
			ID: userID,
			Role: userRole,
			UserAgent: userAgent,
		},
		response.TTL,
	)
	if err != nil {
		return Session{}, err
	}

	return response, nil
}