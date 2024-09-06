package data

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"time"
)

const dbTimeout = time.Second * 3

var db *gorm.DB

func Connect() {
	var err error
	db, err = gorm.Open(sqlite.Open("webservice.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	if err := db.AutoMigrate(&User{}); err != nil {
		panic("failed to migrate database")
	}
}

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
	result := db.Find(&users).Order("created_at desc")
	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

func (u *User) GetByEmail(email string) (*User, error) {
	db.Find(u, "email = ?", email)

	if u.ID == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return u, nil
}

func (u *User) getOne(id int) (*User, error) {
	db.Find(u, "id = ?", id)
	if u.ID == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return u, nil
}

//func (u *User) Delete() error {
//	return nil
//}

func (u *User) DeleteByID(id int) error {
	result := db.Delete(u, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

type CreateUser struct {
	email     string
	firstName string
	lastName  string
	password  string
}

func withEmail(email string) func(*CreateUser) {
	return func(user *CreateUser) {
		user.email = email
	}
}

func (u *User) Create(options ...func(*CreateUser)) {
	user := &CreateUser{}

	for _, option := range options {
		option(user)
	}
}

func abc() {
	db.Create()
}
