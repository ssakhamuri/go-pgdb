package db

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

//This interface is created to make life easier by mocking the db calls in unit tests
type CRUDCalls interface {
	Preload(string, ...interface{}) CRUDCalls
	Where(interface{}, ...interface{}) CRUDCalls
	First(interface{}, ...interface{}) CRUDCalls
	Clauses(...clause.Expression) CRUDCalls
	Session(config *gorm.Session) CRUDCalls
	Updates(values interface{}) CRUDCalls
	Find(interface{}, ...interface{})
	Create(interface{})
	Save(interface{}) CRUDCalls
	Error() error
}

type CRUDFactory struct {
	CRUDCalls
	PGDB *PGDB
}

func (db *CRUDFactory) Session(config *gorm.Session) CRUDCalls {
	resultSet := db.PGDB.Session(config)
	return &CRUDFactory{PGDB: &PGDB{resultSet}}
}

func (db *CRUDFactory) Where(query interface{}, args ...interface{}) CRUDCalls {
	resultSet := db.PGDB.Where(query, args)
	return &CRUDFactory{PGDB: &PGDB{resultSet}}
}

func (db *CRUDFactory) Preload(query string, args ...interface{}) CRUDCalls {
	resultSet := db.PGDB.Preload(query, args)
	return &CRUDFactory{PGDB: &PGDB{resultSet}}
}

func (db *CRUDFactory) First(dest interface{}, conditions ...interface{}) CRUDCalls {
	resultSet := db.PGDB.First(dest, conditions)
	return &CRUDFactory{PGDB: &PGDB{resultSet}}
}

func (db *CRUDFactory) Find(dest interface{}, conditions ...interface{}) {
	db.PGDB.Find(dest)
}

func (db *CRUDFactory) Create(value interface{}) {
	db.PGDB.Create(value)
}

func (db *CRUDFactory) Save(value interface{}) CRUDCalls {
	resultSet := db.PGDB.Save(value)
	return &CRUDFactory{PGDB: &PGDB{resultSet}}
}

func (db *CRUDFactory) Clauses(conds ...clause.Expression) CRUDCalls {
	resultSet := db.PGDB.Clauses(conds...)
	return &CRUDFactory{PGDB: &PGDB{resultSet}}
}

func (db *CRUDFactory) Updates(values interface{}) CRUDCalls {
	resultSet := db.PGDB.Updates(values)
	return &CRUDFactory{PGDB: &PGDB{resultSet}}
}

func (db *CRUDFactory) Error() error {
	return db.PGDB.Error
}

func GetInstance (pgdb *PGDB) CRUDCalls {
	return &CRUDFactory{PGDB: pgdb}
}