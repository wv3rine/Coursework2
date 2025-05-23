package user_usecase

import (
	"context"
)


func (u *UserUC) CheckSession(
	ctx context.Context,
	sessionKey string,
) (Session, error) {
	session, err := u.userRedisRepo.Get(
		ctx,
		sessionKey,
	)
	if err != nil {
		return Session{}, err
	}

	return Session{
		SessionKey: sessionKey,
		UserID: session.ID,
		Role: session.Role,
	}, nil
}