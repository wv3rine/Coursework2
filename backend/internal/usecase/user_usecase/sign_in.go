package user_usecase

import (
	"context"
	"texts/internal/repository/postgres/user_repository"

	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

type SignInUserReq struct {
	Login    string
	Password string
	UserAgent string
}

func (u *UserUC) SignIn(ctx context.Context, signInUserReq SignInUserReq) (Session, error) {
	spanName := "UserUC.SignIn"

	users, err := u.userPGRepo.SelectUsers(ctx, user_repository.SelectUserReq{
		Logins: []string{signInUserReq.Login},
	})
	if err != nil {
		return Session{}, errors.Wrap(err, spanName)
	}
	if len(users) != 1 {
		return Session{}, errors.New("user's login is not unique")
	}
	user := users[0]

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(signInUserReq.Password), bcrypt.DefaultCost)
	if err != nil {
		return Session{}, errors.Wrap(err, spanName)
	}
	user.Password = string(hashedPassword)
 
	if user.Password != string(hashedPassword) {
		return Session{}, errors.New("wrong password")
	}

	session, err := u.setSession(ctx, signInUserReq.UserAgent, user.UserID, user.Role)
	if err != nil {
		return Session{}, errors.Wrap(err, spanName)
	}
 
	return session, nil
}
 