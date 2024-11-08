package repository

import (
	"context"
	"database/sql"
	"fmt"

	"greendeco-be/app/models"
	"greendeco-be/platform/database"

	"github.com/google/uuid"
)

type UserRepository interface {
	Create(u *models.CreateUser, ctx context.Context) error
	CreateForStaff(*models.CreateUser) error
	GetUserByIdentifier(identifier string) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	GetUserByPhoneNumber(phoneNumber string) (*models.User, error)
	GetUserById(uuid.UUID) (*models.User, error)
	UpdatePasswordById(password string, id uuid.UUID) error
	UpdateUserInfor(userId uuid.UUID, user *models.UpdateUser) error
}

type UserRepo struct {
	db *database.DB
}

const (
	UserTable = "users"
)

var _ UserRepository = (*UserRepo)(nil)

func NewUserRepo(db *database.DB) UserRepository {
	return &UserRepo{db}
}

func (repo *UserRepo) Create(u *models.CreateUser, ctx context.Context) error {
	query := fmt.Sprintf(`INSERT INTO "%s" (email,identifier,password,first_name,last_name, phone_number) VALUES ($1,$2,$3,$4,$5,$6)`, UserTable)
	_, err := repo.db.Exec(query, u.Email, u.Identifier, u.Password, u.FirstName, u.LastName, u.PhoneNumber)
	if err != nil {
		return err
	}
	return nil
}

func (repo *UserRepo) GetUserByEmail(email string) (*models.User, error) {
	user := models.NewUser()
	query := fmt.Sprintf(`SELECT * FROM "%s" WHERE email = $1`, UserTable)
	err := repo.db.Get(user, query, email)
	if err == sql.ErrNoRows {
		return nil, models.ErrNotFound
	} else if err != nil {
		return nil, err
	}

	return user, nil
}

func (repo *UserRepo) GetUserByIdentifier(identifier string) (*models.User, error) {
	user := models.NewUser()
	query := fmt.Sprintf(`SELECT * FROM "%s" WHERE identifier = $1`, UserTable)
	err := repo.db.Get(user, query, identifier)
	if err == sql.ErrNoRows {
		return nil, models.ErrNotFound
	} else if err != nil {
		return nil, err
	}

	return user, nil
}

func (repo *UserRepo) GetUserByPhoneNumber(phoneNumber string) (*models.User, error) {
	user := models.NewUser()
	query := fmt.Sprintf(`SELECT * FROM "%s" WHERE phone_number = $1`, UserTable)
	err := repo.db.Get(user, query, phoneNumber)
	if err == sql.ErrNoRows {
		return nil, models.ErrNotFound
	} else if err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepo) GetUserById(uId uuid.UUID) (*models.User, error) {
	user := models.NewUser()
	query := fmt.Sprintf(`SELECT * FROM "%s" WHERE id = $1`, UserTable)
	err := repo.db.Get(user, query, uId)

	if err == sql.ErrNoRows {
		return nil, models.ErrNotFound
	} else if err != nil {
		return nil, err
	}

	return user, err
}

func (repo *UserRepo) UpdatePasswordById(password string, id uuid.UUID) error {
	query := fmt.Sprintf(`UPDATE "%s" SET password = $1 WHERE id = $2`, UserTable)
	_, err := repo.db.Exec(query, password, id)
	if err != nil {
		return err
	}

	return nil
}

func (repo *UserRepo) UpdateUserInfor(userId uuid.UUID, user *models.UpdateUser) error {
	query := fmt.Sprintf(`UPDATE "%s" SET first_name = $2, last_name = $3, avatar = $4, phone_number = $5  WHERE id = $1`, UserTable)
	_, err := repo.db.Exec(query, userId, user.FirstName, user.LastName, user.Avatar, user.PhoneNumber)
	if err != nil {
		return err
	}

	return nil
}

func (repo *UserRepo) UpdateRules() error {
	return nil
}

func (repo *UserRepo) CreateForStaff(u *models.CreateUser) error {
	query := fmt.Sprintf(`INSERT INTO "%s" (email,identifier,password,first_name,last_name, phone_number, admin) VALUES ($1,$2,$3,$4,$5,$6,$7)`, UserTable)
	_, err := repo.db.Exec(query, u.Email, u.Identifier, u.Password, u.FirstName, u.LastName, u.PhoneNumber, true)
	if err != nil {
		return err
	}
	return nil
}
