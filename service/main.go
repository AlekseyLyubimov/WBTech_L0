package main

import (
	models "WBTech_L0/service/models"
	"context"
	"fmt"
	"net/http"
	"time"
	"encoding/json"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/nats-io/stan.go"
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
	
	err = initDbConnection()
	if err != nil {
		fmt.Println("DB connection failed: " + err.Error())
	}

	ten_seconds, _ := time.ParseDuration("10s")
	fifteen_minutes, _ := time.ParseDuration("15m")

	server := &http.Server{
		Addr:           ":8080",
		Handler:        buildRouter(),
		ReadTimeout:    ten_seconds ,
		WriteTimeout:   ten_seconds,
		IdleTimeout:    fifteen_minutes,
		MaxHeaderBytes: 2048,
	}

	err = natsSubsription()
	if err != nil {
		fmt.Println("NATS subscription failed: " + err.Error())
	}

	err = server.ListenAndServe()
	if err!= nil {
		fmt.Println("Server failed: " + err.Error())
	}
}

func initDbConnection() error {
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

func buildRouter() http.Handler {

	router := echo.New()

	router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://127.0.0.1*"},
		AllowCredentials: true,
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods:     []string{http.MethodGet, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete, http.MethodOptions}}))

	router.GET("/order/:id", getOrder)

	return router
}

func getOrder(c echo.Context) error {

	order := new(models.Order)
	order.ID = c.Param("id")

	if err := db.NewSelect().
		Model(order).
		Relation("Delivery").
		Relation("Payment").
		Relation("Items").
		WherePK().
		Scan(context.Background()); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, order)
}

func natsSubsription() (err error) {
	sc, err := stan.Connect("test-cluster", "client-123")
	if (err != nil) {return err}

	_, err = sc.Subscribe("foo", func(m *stan.Msg) {
		
		var order models.Order
		err := json.Unmarshal(m.Data, &order)

		if err != nil {
			//should be replaced with actual logging of unparsable input.
			fmt.Printf("mock logging bad input for message %s", m.Data)
			m.Ack()
		} else {
			_, err := db.NewInsert().Model(order).Exec(context.Background())
			if err == nil {
				m.Ack()
			}
		}
		
	}, stan.DurableName("test_subscription"), stan.SetManualAckMode())

	return err
}