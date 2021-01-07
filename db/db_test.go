package db

import (
	"go-pgdb/config"
	"testing"
)

func Test_buildDatabaseConnectString(t *testing.T) {
	conf := &config.Configuration{
		DbName: 	"DBNAME",
		DbHostName: "DBHOSTNAME",
		DbPort: 	"DBPORT",
		DbUserName: "DBUSERNAME",
		DbPassword: "DBPASSWORD",
	}
	type args struct {
		config *config.Configuration
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"test database connection string", args{config: conf}, "host=DBHOSTNAME port=DBPORT user=DBUSERNAME dbname=DBNAME password=DBPASSWORD"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildDatabaseConnectString(tt.args.config); got != tt.want {
				t.Errorf("buildDatabaseConnectString() = %v, want %v", got, tt.want)
			}
		})
	}
}
