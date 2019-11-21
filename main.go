package main

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/ilamazuliaf/desaku/delivery"
	"github.com/ilamazuliaf/desaku/middlewares"
	"github.com/ilamazuliaf/desaku/models"
	"github.com/ilamazuliaf/desaku/repository"
	"github.com/ilamazuliaf/desaku/usecase"

	"gopkg.in/go-playground/validator.v9"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile("config.json")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

func connect() *sql.DB {
	var dbHost, dbUser, dbPass, dbName string
	dbHost = viper.GetString("database.host")
	dbUser = viper.GetString("database.user")
	dbPass = viper.GetString("database.pass")
	dbName = viper.GetString("database.db")
	conn := fmt.Sprintf("%s:%s@tcp(%s)/%s", dbUser, dbPass, dbHost, dbName)
	db, err := sql.Open("mysql", conn)
	if err != nil {
		panic(err)
	}
	return db
}

func main() {
	db := connect()
	if err := db.Ping(); err != nil {
		panic(err)
	}
	defer db.Close()
	// defer profile.Start().Stop()
	// defer profile.Start(profile.MemProfile).Stop()

	db.SetConnMaxLifetime(time.Duration(viper.GetInt("database.maxLifeTime")) * time.Minute)
	db.SetMaxIdleConns(viper.GetInt("maxIdleCons"))
	db.SetMaxOpenConns(viper.GetInt("maxOpenCons"))

	e := echo.New()
	e.Use(middleware.Gzip())
	e.Use(echo.WrapMiddleware(middlewares.MiddlewareJWTAuthorization))
	e.Validator = &models.CustomValidator{Validator: validator.New()}

	timeOut := time.Duration(viper.GetInt("context.timeout")) * time.Second
	repo := repository.NewRepositoryConfig(db)
	usec := usecase.NewUsecaseConfig(repo, timeOut)
	delivery.NewHandler(e, usec)

	port := viper.GetString("server.port")

	if viper.GetBool("running.debug") {
		fmt.Println("Service RUN on DEBUG mode")
	}
	fmt.Println("Service RUN on ", viper.GetString("running.status"), " mode")
	e.Logger.Fatal(e.Start(port))
}
