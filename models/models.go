package models

type Forum struct {
	ID   		int         `json:"id" gorm:"primaryKey"`
	Name 		string		`json:"name"`
	Threads 	[]Thread 	`json:"threads" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Thread struct {
	ID   	int     `json:"id" gorm:"primaryKey"`
	Title 	string	`json:"title"`
	ForumID int		`json:"forumId"`
	Posts   []Post	`json:"posts" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Post struct {
	ID   		int     `json:"id" gorm:"primaryKey"`
	Title 		string	`json:"title"`
	Body 		string	`json:"body"`
	ThreadID 	int		`json:"threadId"`
}