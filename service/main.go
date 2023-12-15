package main

import (
	models "WBTech_L0/service/models"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

var postgresConfig = "postgres://postgres:qwerty123@localhost:5432/postgres?sslmode=disable&search_path=golang_migrate"
var db *bun.DB

func main() {

	//requires golang_migrate schema to exist in db
	m, err := migrate.New(
		"file://migrations",
		postgresConfig)
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
	
	err = InitDbConnection()
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

func InitDbConnection() error {
	config, err := pgx.ParseConfig(postgresConfig)
	if err != nil {
		panic(err)
	}
	config.PreferSimpleProtocol = true

	sqldb := stdlib.OpenDB(*config)
	db = bun.NewDB(sqldb, pgdialect.New())

	db.RegisterModel((*models.OrderToItem)(nil))

	return err
}

func BuildRouter() http.Handler {

	router := echo.New()

	router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://127.0.0.1*"},
		AllowCredentials: true,
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete, http.MethodOptions}}))

	router.GET("/order/:id", GetOrder)

	return router
}

func TestHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, Responce{"test ok"})
}

type Responce struct {
	Message string `json:"message"`
}

func GetOrder(c echo.Context) error {

	//TODO unsafe
	id := c.Param("id")

	order := new(models.Order)
	order.ID = &id

	if err := db.NewSelect().
		Model(order).
		Relation("Delivery").
		Relation("Payment").
		//Relation("Items").
		//Where("wb.order.id = ?", c.Param("id")).
		WherePK().
		Scan(context.Background()); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, order)
}