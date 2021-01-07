package models

import (
	"gorm.io/gorm"
	"time"
)

type Forum struct {
	ID   			int         `json:"id" gorm:"primaryKey"`
	Name 			string		`json:"name"`
	Threads 		[]Thread 	`json:"threads" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreateDate  	int64   	`json:"createDate"`
	LastUpdateTime	int64       `json:"lastUpdateTime"`

}

type Thread struct {
	ID   			int     `json:"id" gorm:"primaryKey"`
	Title 			string	`json:"title"`
	ForumID 		int		`json:"forumId"`
	Posts   		[]Post	`json:"posts" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Post struct {
	ID   			int     `json:"id" gorm:"primaryKey"`
	Title 			string	`json:"title"`
	Body 			string	`json:"body"`
	ThreadID 		int		`json:"threadId"`
}

func GetCurrentSystemTime() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func (forum *Forum) BeforeCreate(tx *gorm.DB) error {
	var systemTime = GetCurrentSystemTime()
	forum.LastUpdateTime = systemTime
	forum.CreateDate = systemTime
	return nil
}

func (forum *Forum) BeforeUpdate(tx *gorm.DB) error {
	forum.LastUpdateTime = GetCurrentSystemTime()
	return nil
}