package model

import "github.com/google/uuid"

type UserModel struct {
	ID       uuid.UUID
	Login    string
	Email    string
	Password string
	RoleName string
}

type UserSess struct {
	ID       string
	Login    string
	Email    string
	RoleName string
}

func (u *UserModel) Public() *UserSess {
	return &UserSess{
		ID:       u.ID.String(),
		Login:    u.Login,
		Email:    u.Email,
		RoleName: u.RoleName,
	}
}
