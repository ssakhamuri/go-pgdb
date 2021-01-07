package db

import (
	"go-pgdb/models"
	"reflect"
	"testing"
)

func TestPGDB_CreateForum(t *testing.T) {
	mockDB := &MockCRUDCalls{}
	CreateDAOInstance = mockDB.CreateMockInstance
	type args struct {
		forum models.Forum
	}
	tests := []struct {
		name   string
		args   args
		want   models.Forum
	}{
		{"Test Create Forum", args{models.Forum{}}, models.Forum{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &PGDB{}
			if got := db.CreateForum(tt.args.forum); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateForum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPGDB_GetAllForums(t *testing.T) {
	mockDB := &MockCRUDCalls{}
	CreateDAOInstance = mockDB.CreateMockInstance
	tests := []struct {
		name   string
		want   []models.Forum
	}{
		{"Get All Forums", nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &PGDB{}
			if got := db.GetAllForums(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAllForums() = %v, want %v", got, tt.want)
			}
		})
	}
}