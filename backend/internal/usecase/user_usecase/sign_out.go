package user_usecase

import (
	"context"

	"github.com/pkg/errors"
)


type SignOutUserReq struct {
	SessionKey string
}

func (u *UserUC) SignOut(ctx context.Context, signOutUserReq SignOutUserReq) (error) {
	spanName := "UserUC.SignOut"

	err := u.userRedisRepo.Delete(ctx, signOutUserReq.SessionKey)
	if err != nil {
		return errors.Wrap(err, spanName)
	}
 
	return  nil
}
 