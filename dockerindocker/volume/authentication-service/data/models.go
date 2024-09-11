package data

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Email     string    `json:"email"`
	FirstName string    `json:"first_name,omitempty"`
	LastName  string    `json:"last_name,omitempty"`
	Password  string    `json:"-"`
	Active    int       `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) GetAll() ([]User, error) {

	var users []User

	result := DB.ORM.Find(&users).Order("created_at desc")
	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

func (u *User) GetByEmail(email string) (*User, error) {
	DB.ORM.Find(u, "email = ?", email)

	if u.ID == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return u, nil
}

func (u *User) getOne(id int) (*User, error) {
	DB.ORM.Find(u, "id = ?", id)
	if u.ID == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return u, nil
}

func (u *User) Update(options ...func(*User) *User) error {
	var user *User
	for _, option := range options {
		user = option(u)
	}

	// update user
	if user != nil {
		if result := DB.ORM.Model(u).Updates(*user); result.Error != nil {
			return result.Error
		}
	}

	return nil
}

func (u *User) Delete() error {
	if result := DB.ORM.Delete(u, u.ID); result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *User) DeleteByID(id int) error {
	result := DB.ORM.Delete(u, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (u *User) Create(options ...func(*User)) error {
	for _, option := range options {
		option(u)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), 12)
	if err != nil {
		return nil
	}

	u.Password = string(hashedPassword)

	if result := DB.ORM.Create(u); result.Error != nil {
		return result.Error
	}

	return nil
}

func (u *User) ResetPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return nil
	}

	u.Password = string(hashedPassword)
	if result := DB.ORM.Model(u).Update("password", string(hashedPassword)); result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *User) PasswordMatches(plantText string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plantText))
	if err != nil {
		switch {
		case errors.Is(err, bcrypt.ErrMismatchedHashAndPassword):
			// invalid password
			return false, nil
		default:
			return false, err
		}
	}
	return true, nil
}
