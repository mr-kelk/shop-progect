package service

import (
	"errors"
	"example/shop-progect/internal/enum"
	"example/shop-progect/internal/model"
	"example/shop-progect/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	user *repository.UserRepository
}

func NewAuthService(repo *repository.UserRepository) *AuthService {
	return &AuthService{user: repo}
}

func isPasswordVerification(hashedPassword []byte, password []byte) (bool, error) {
	err := bcrypt.CompareHashAndPassword(hashedPassword, password)
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return false, errors.New("wrong password")
		}
		return false, err
	}
	return true, nil

}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func (s *AuthService) Login(email string, password string) (*model.UserSess, error) {
	user, err := s.user.GetUserByEmail(email)

	if err != nil || user == nil {
		return nil, errors.New("user not found")
	}

	ok, err := isPasswordVerification([]byte(user.Password), []byte(password))

	if err != nil || !ok {
		return nil, errors.New("invalid credentials")
	}

	return user.Public(), nil
}

func (s *AuthService) Register(login string, email string, password string) (*model.UserSess, error) {
	user, err := s.user.GetUserByEmail(email)

	if user != nil {
		return nil, errors.New("user already exists")
	}

	if err == nil {
		return nil, errors.New("user with email already exists")
	}

	role, _ := s.user.GetRoleByName(enum.USER)

	hashed, _ := hashPassword(password)

	err = s.user.CreateUser(login, email, hashed, role.ID)

	if err != nil {
		return nil, errors.New("user not created")
	}

	user, _ = s.user.GetUserByEmail(email)

	return user.Public(), nil
}
