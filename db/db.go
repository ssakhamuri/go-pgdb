package db

import (
	"fmt"
	"go-pgdb/config"
	"go-pgdb/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

type PGDB struct {
	*gorm.DB
}

type CRUDOperations interface {
	GetAllForums() []models.Forum
	CreateForum(forum models.Forum) models.Forum
	UpdateForum(forum models.Forum) (models.Forum, error)
	GetForumById(id int) models.Forum
	GetThreadsByForumId(id int) []models.Thread
	GetPostsByThreadId(id int) []models.Post
	GetPostsByForumId(id int) []models.Post
}

func Init(config *config.Configuration) (pgdb *PGDB, err error) {
	dsn := buildDatabaseConnectString(config)
	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Print(err)
		time.Sleep(10 * time.Second)
		connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		conn = connection
		fmt.Print(err)
	}
	pgdb = &PGDB{conn}
	pgdb.Debug().AutoMigrate(&models.Forum{})
	pgdb.Debug().AutoMigrate(&models.Thread{})
	pgdb.Debug().AutoMigrate(&models.Post{})
	return
}

func buildDatabaseConnectString(config *config.Configuration) string {
	return "host=" + config.DbHostName + " port=" + config.DbPort + " user=" + config.DbUserName + " dbname=" + config.DbName + " password=" + config.DbPassword
}
