package db

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
)

type MockCRUDCalls struct {
	CRUDCalls
	CalledWithQuery []interface{}
	CalledWithArgs  []interface{}
}

var mockDB = &MockCRUDCalls{}
var mockError error

func (db *MockCRUDCalls) CreateMockInstance(pgdb *PGDB) CRUDCalls {
	return mockDB
}

func (db *MockCRUDCalls) Preload(query string, args ...interface{}) CRUDCalls {
	log.Println("Preload Mock Called")
	mockDB.CalledWithQuery = append(mockDB.CalledWithQuery, query)
	mockDB.CalledWithArgs = append(mockDB.CalledWithArgs, args)
	return mockDB
}


func (db *MockCRUDCalls) Where(query interface{}, args ...interface{}) CRUDCalls {
	log.Println("Where Mock Called")
	mockDB.CalledWithQuery = append(mockDB.CalledWithQuery, query)
	mockDB.CalledWithArgs = append(mockDB.CalledWithArgs, args)
	return mockDB
}

func (db *MockCRUDCalls) First(dest interface{}, conditions ...interface{}) CRUDCalls {
	log.Println("First Mock Called")
	return mockDB
}

func (db *MockCRUDCalls) Clauses(conds ...clause.Expression) CRUDCalls {
	log.Println("Clauses Mock Called")
	return mockDB
}

func (db *MockCRUDCalls) Session(config *gorm.Session) CRUDCalls {
	log.Println("Session Mock Called")
	return mockDB
}

func (db *MockCRUDCalls) Find(dest interface{}, conditions ...interface{}) {
	log.Println("Find Mock Called")
}

func (db *MockCRUDCalls) Updates(values interface{}) CRUDCalls {
	log.Println("Updates Mock Called")
	return mockDB
}

func (db *MockCRUDCalls) Create(value interface{}) {
	log.Println("Create Mock Called")
}

func (db *MockCRUDCalls) Save(value interface{}) CRUDCalls {
	log.Println("Save Mock Called")
	return mockDB
}

func (db *MockCRUDCalls) Error() error {
	log.Println("Error Mock Called")
	return mockError
}
