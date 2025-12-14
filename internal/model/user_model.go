package model

type UserModel struct {
	ID       []byte
	Login    string
	Email    string
	Password string
	RoleName string
}

type UserPublic struct {
	ID       string
	Login    string
	Email    string
	RoleName string
}

func (u *UserModel) Public() *UserPublic {
	return &UserPublic{
		ID:       string(u.ID),
		Login:    u.Login,
		Email:    u.Email,
		RoleName: u.RoleName,
	}
}
