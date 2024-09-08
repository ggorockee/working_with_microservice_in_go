package data

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

const dbTimeout = time.Second * 3

type DBInstance struct {
	DB *gorm.DB
}

var Database DBInstance

type dbConfig struct {
	host            string
	user            string
	password        string
	dbname          string
	port            string
	sslmode         string
	timeZone        string
	connect_timeout string
}

func (d dbConfig) String() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s connect_timeout=%s",
		d.host,
		d.user,
		d.password,
		d.dbname,
		d.port,
		d.sslmode,
		d.timeZone,
		d.connect_timeout,
	)
}

func ConnectDB() {
	postgresConfig := dbConfig{
		host:            "postgres",
		user:            "postgres",
		password:        "password",
		dbname:          "users",
		port:            "5432",
		sslmode:         "disable",
		timeZone:        "Asia/Seoul",
		connect_timeout: "5",
	}
	dsn := fmt.Sprintf("%s", postgresConfig)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	if err := db.AutoMigrate(&User{}); err != nil {
		panic("failed to migrate database")
	}

	Database = DBInstance{DB: db}
}

type Models struct {
	User User
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

	result := Database.DB.Find(&users).Order("created_at desc")
	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

func (u *User) GetByEmail(email string) (*User, error) {
	Database.DB.Find(u, "email = ?", email)

	if u.ID == 0 {
		return nil, gorm.ErrRecordNotFound
	}

	return u, nil
}

func (u *User) getOne(id int) (*User, error) {
	Database.DB.Find(u, "id = ?", id)
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
	if result := Database.DB.Model(u).Updates(*user); result.Error != nil {
		return result.Error
	}

	return nil
}

func (u *User) Delete() error {
	if result := Database.DB.Delete(u, u.ID); result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *User) DeleteByID(id int) error {
	result := Database.DB.Delete(u, id)
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

	if result := Database.DB.Create(u); result.Error != nil {
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
	if result := Database.DB.Model(u).Update("password", string(hashedPassword)); result.Error != nil {
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
