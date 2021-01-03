package handler

import (
	"go-pgdb/config"
	"log"
	"os"
)

func (env *Env) SetupEnvironment()  {
	log.Printf("Set Up environment.!")
	env.Config = &config.Configuration{
		DbName: 	os.Getenv("DB_NAME"),
		DbHostName: os.Getenv("DB_HOST"),
		DbPort: 	os.Getenv("DB_PORT"),
		DbUserName: os.Getenv("DB_USERNAME"),
		DbPassword: os.Getenv("DB_PASSWORD"),
	}
}

// this is added here just to make the application work. Please remove and add it to your local environment/vault in realtime
func DefaultEnvironment() {
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_USERNAME", "postgres")
	os.Setenv("DB_PASSWORD", "secret")
	os.Setenv("DB_NAME", "forum_threads")
}