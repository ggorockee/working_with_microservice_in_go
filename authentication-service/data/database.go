package data

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

type DBInstance struct {
	ORM *gorm.DB
}

var DB DBInstance

type dbConfig struct {
	host     string
	user     string
	password string
	dbname   string
	port     string
	sslmode  string
	timeZone string
}

func (d dbConfig) String() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		d.host,
		d.user,
		d.password,
		d.dbname,
		d.port,
		d.sslmode,
		d.timeZone,
	)
}

func ConnectDB() {
	postgresConfig := dbConfig{
		host:     "postgres",
		user:     "postgres",
		password: "password",
		dbname:   "users",
		port:     "5432",
		sslmode:  "disable",
		timeZone: "Asia/Seoul",
	}
	dsn := fmt.Sprintf("%s", postgresConfig)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	log.Println("Successfully connected to database")
	db.Logger = logger.Default.LogMode(logger.Info)
	if err := db.AutoMigrate(&User{}); err != nil {
		panic("failed to migrate database")
	}

	DB = DBInstance{ORM: db}
}
