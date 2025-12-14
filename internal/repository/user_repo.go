package repository

import (
	"database/sql"
	"example/shop-progect/internal/model"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) GetRoleByName(name string) (*model.RoleModel, error) {
	row := r.DB.QueryRow(`
        SELECT r.ID, r.NAME
		FROM  SHOP.ROLES r
		WHERE r.NAME = :1
    `, name)

	var roleModel model.RoleModel
	err := row.Scan(&roleModel.ID, &roleModel.Name)
	if err != nil {
		return nil, err
	}

	return &roleModel, nil
}

func (r *UserRepository) GetUserByEmail(email string) (*model.UserModel, error) {
	var userModel model.UserModel

	err := r.DB.QueryRow(`
        SELECT U.ID, U.LOGIN, U.EMAIL, U.PASSWORD, R.NAME AS ROLE_NAME
		FROM  SHOP.USERS U
		LEFT JOIN SHOP.ROLES R on R.ID = U.ROLE_ID
		WHERE U.EMAIL = :1
    `, email).Scan(&userModel.ID, &userModel.Login, &userModel.Email, &userModel.Password, &userModel.RoleName)

	if err != nil {
		return nil, err
	}

	return &userModel, nil
}

func (r *UserRepository) CreateUser(login string, email string, password string, roleId int) error {
	_, err := r.DB.Exec(`INSERT INTO SHOP.USERS (LOGIN, EMAIL, PASSWORD, ROLE_ID) VALUES (:1, :2, :3, :4)`, login, email, password, roleId)
	return err
}
