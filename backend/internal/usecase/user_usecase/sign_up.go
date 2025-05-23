package user_usecase

import (
	"context"
	"texts/internal/repository/postgres/user_repository"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type SignUpUserReq struct {
	Login    string
	Password string
	Role     string
	UserAgent string
}


func (u *UserUC) SignUp(ctx context.Context, user SignUpUserReq) (Session, error) {
 spanName := "UserUC.SignUp"

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return Session{}, errors.Wrap(err, spanName)
	}
	user.Password = string(hashedPassword)

	var session Session

	if err := u.trManager.Do(ctx, func(ctx context.Context) error {
		uid, err := u.userPGRepo.InsertUser(ctx, user_repository.InsertUserReq{
			Login: user.Login,
			Password: user.Password,
			Role: user.Role,
		})
		if err != nil {
			return errors.Wrap(err, spanName)
		}
		
		session, err = u.setSession(ctx, user.UserAgent, uid, user.Role)
		if err != nil {
			return errors.Wrap(err, spanName)
		}

		return nil
	}); err != nil {
		return Session{}, errors.Wrap(err, spanName)
	}

	return session, nil
}
