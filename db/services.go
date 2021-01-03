package db

import (
	"go-pgdb/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var CreateDAOInstance = GetInstance

func (db *PGDB) GetAllForums() []models.Forum {
	dao := CreateDAOInstance(db)
	var forums []models.Forum
	dao.Preload("Threads.Posts.Post").Find(&forums)
	return forums
}

func (db *PGDB) CreateForum(forum models.Forum) models.Forum {
	dao := CreateDAOInstance(db)
	dao.Create(&forum)
	return forum
}

func (db *PGDB) UpdateForum(forum models.Forum) (models.Forum, error) {
	dao := CreateDAOInstance(db)
	result := dao.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&forum)
	if result.Error() != nil {
		return forum, result.Error()
	}
	return forum, nil
}

func (db *PGDB) GetForumById(id int) (forum models.Forum) {
	dao := CreateDAOInstance(db)
	dao.Where(&models.Forum{ID: id}).Preload("Threads.Posts.Post").First(&forum)
	return
}

func (db *PGDB) GetThreadsByForumId(id int) []models.Thread {
	dao := CreateDAOInstance(db)
	var threads []models.Thread
	dao.Where(&models.Thread{ForumID: id}).Preload("Posts.Post").Find(&threads)
	return threads
}

func (db *PGDB) GetPostsByThreadId(id int) []models.Post {
	dao := CreateDAOInstance(db)
	var posts []models.Post
	dao.Where(&models.Post{ThreadID: id}).Find(&posts)
	return posts
}

func (db *PGDB) GetPostsByForumId(id int) []models.Post {
	dao := CreateDAOInstance(db)
	var threads []models.Thread
	dao.Where(&models.Thread{ForumID: id}).Find(&threads)
	var posts []models.Post
	var threadIds []int
	if threads != nil && len(threads) > 0 {
		for index := range threads {
			threadIds = append(threadIds, threads[index].ID)
		}
	}
	dao.Clauses(clause.Expr{"thread_id IN ?", []interface{}{threadIds}, false}).Find(&posts)
	return posts
}