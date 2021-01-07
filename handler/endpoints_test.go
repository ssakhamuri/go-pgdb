package handler

import (
	"go-pgdb/db"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/stretchr/testify/assert"
)

var mockDB = &db.MockCRUDCalls{}

func TestEnv_GetAllForums(t *testing.T) {
	assert := New(t)
	db.CreateDAOInstance = mockDB.CreateMockInstance
	tests := []struct {
		name  string
		urlString  string
		expectedBody string
		expectedResponseCode int
	}{
		{"GetAllForums", "/api/forum/all", "null\n", 200},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", tt.urlString, nil)
			env := &Env{nil, nil, &db.PGDB{}}
			recorder := httptest.NewRecorder()
			handler := http.HandlerFunc(env.GetAllForums)
			handler.ServeHTTP(recorder, req)
			responseBody, err := ioutil.ReadAll(recorder.Body)
			assert.NoError(err, "Error is not nil")
			assert.Equal(tt.expectedResponseCode, recorder.Code, tt.name)
			assert.Equal(tt.expectedBody, string(responseBody), tt.name)
		})
	}
}
