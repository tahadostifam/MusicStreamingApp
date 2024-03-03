package auth

import (
	"errors"

	"github.com/tahadostifam/MusicStreamingApp/api/models"
	"gorm.io/gorm"
)

var ErrUserNotFound = errors.New("user not found")
var ErrEmailAlreadyExist = errors.New("email already exist")

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) Create(name, email, password string) (*models.User, error) {
	obj := models.User{Name: name, Email: email, Password: password}
	result := r.db.Create(&obj)

	// var pErr *pgconn.PgError
	// errors.As(result.Error, &pErr)

	// if pErr.Code == "23505" { // check for duplicate key error
	// 	return nil, ErrEmailAlreadyExist
	// }
	if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
		return nil, ErrEmailAlreadyExist
	} else if result.Error != nil {
		return nil, result.Error
	}

	return &obj, nil
}

func (r *Repository) FindByEmail(email string) (*models.User, error) {
	user := &models.User{}
	obj := r.db.Where("email = ?", email).Find(&user)
	if obj.Error != nil {
		return nil, ErrUserNotFound
	}

	return user, nil
}
func (r *Repository) FindByUserID(userID string) (*models.User, error) {
	user := &models.User{}
	obj := r.db.Where("user_id = ?", userID).Find(&user)
	if obj.Error != nil {
		return nil, ErrUserNotFound
	}

	return user, nil
}

func (r *Repository) Delete(email string) error {
	user := &models.User{}
	tx := r.db.Where("email = ?", email).Delete(&user)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (r *Repository) Update(email string, newName, newPassword string) (*models.User, error) {
	user, err := r.FindByEmail(email)
	if err != nil {
		return nil, ErrUserNotFound
	}

	user.Name = newName
	user.Password = newPassword

	tx := r.db.Save(&user)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return user, nil
}
