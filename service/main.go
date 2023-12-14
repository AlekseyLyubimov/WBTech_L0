package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/jackc/pgx/v4/pgxpool"
)

var dbPool *pgxpool.Pool

func main() {

	//requires golang_migrate schema to exist in db
	m, err := migrate.New(
		"file://../migrations",
		"postgres://postgres:qwerty123@localhost:5432/postgres?sslmode=disable&search_path=golang_migrate")
	if err == nil {
		if err := m.Up(); err != nil {
			if err.Error() == "no change" {
				fmt.Println("migrations up to date")
			} else {
				fmt.Println("migration failure: " + err.Error())
			}
		}
	} else {
		fmt.Println("Migrations failed: " + err.Error())
	}
	
	err = InitConnection("host=localhost port=5432 user=postgres password=qwerty123 dbname=postgres sslmode=disable")
	if err != nil {
		fmt.Println("DB connection failed: " + err.Error())
	}

	ten_seconds, _ := time.ParseDuration("10s")
	fifteen_minutes, _ := time.ParseDuration("15m")

	server := &http.Server{
		Addr:           ":8080",
		Handler:        BuildRouter(),
		ReadTimeout:    ten_seconds ,
		WriteTimeout:   ten_seconds,
		IdleTimeout:    fifteen_minutes,
		MaxHeaderBytes: 2048,
	}

	err = server.ListenAndServe()
	if err!= nil {
		fmt.Println("Server failed: " + err.Error())
	}
}

func InitConnection(config string) error {
	var err error
	ctx := context.Background()
	dbPool, err = pgxpool.Connect(ctx, config)
	if err != nil {
		fmt.Println("DB connection failed: " + err.Error())
	}

	asd, err := dbPool.Acquire(ctx)
	asd.Conn().PgConn().Conn()

	return err
}

func BuildRouter() http.Handler {

	router := echo.New()

	router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://127.0.0.1*"},
		AllowCredentials: true,
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete, http.MethodOptions}}))

	router.GET("/order", TestHandler)

	return router
}

func TestHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, Responce{"test ok"})
}

type Responce struct {
	Message string `json:"message"`
}