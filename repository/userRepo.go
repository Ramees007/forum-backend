package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"github.com/rameesThattarath/qaForum/entity"
)

//UserRepo - UserRepo
type UserRepo interface {
	Register(entity.User) (uint, error)
	Login(entity.User) (uint, *entity.User, error)
	GetProfile(uint) (*entity.User, error)
}

//SQLUserRepo - SQLUserRepo
type SQLUserRepo struct {
	db *sql.DB
}

//NewSQLUserRepo - NewMySqlUserRepo
func NewSQLUserRepo(db *sql.DB) *SQLUserRepo {
	return &SQLUserRepo{
		db,
	}
}

//Register - user register
func (r SQLUserRepo) Register(u entity.User) (uint, error) {
	hp, err := hashPassword(u.Password)
	if err != nil {
		return 0, err
	}

	query := fmt.Sprintf("INSERT INTO users (email,name,password) VALUES ('%s', '%s','%s')",
		u.Email, u.Name, hp)

	insert, err := r.db.Query(query)
	if err != nil {
		return 0, err
	}

	q2 := fmt.Sprintf("SELECT id FROM users WHERE email='%s'", u.Email)

	lastID, err := r.db.Query(q2)
	if err != nil {
		return 0, err
	}
	var id uint = 0
	for lastID.Next() {
		lastID.Scan(&id)
	}

	defer insert.Close()
	defer lastID.Close()
	return id, nil
}

//Login - user Login
func (r SQLUserRepo) Login(u entity.User) (uint, *entity.User, error) {
	var id uint = 0

	row := r.db.QueryRow("SELECT id,email,name,password FROM users WHERE email =?", u.Email)

	var user entity.User
	var hashedPass string
	err := row.Scan(&id, &user.Email, &user.Name, &hashedPass)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil, errors.New("Incorrect email ID or password")
		}
		return 0, nil, err
	}

	if !checkPasswordHash(u.Password, hashedPass) {
		return 0, nil, errors.New("Incorrect email ID or password")
	}

	return id, &user, nil
}

//GetProfile - GetProfile
func (r SQLUserRepo) GetProfile(id uint) (*entity.User, error) {
	row := r.db.QueryRow("SELECT id,email,name FROM users WHERE id =?", id)
	var user entity.User
	err := row.Scan(&id, &user.Email, &user.Name)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("Incorrect email ID or password")
		}
		return nil, err
	}

	return &user, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
