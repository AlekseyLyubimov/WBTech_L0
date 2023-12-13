package models

import (
	"time"

	bun "github.com/uptrace/bun"
)

type Order struct {

	bun.BaseModel `bun:"table:wb.order"`

	ID *string `bun:"id"`
	TrackNumber *string `bun:"track_number"`
	Entry *string `bun:"entry"`
	Delivery *Delivery `bun:"delivery"`
	Payment *Payment `bun:"payment"`
	Items []*Item `bun:"items"`
	Locale *string `bun:"locale"`
	InternalSignature *string `bun:"internal_signature"`
	CustomerId *string `bun:"customer_id"`
	DeliveryService *string `bun:"delivery_service"`
	Shardkey *string `bun:"shardkey"`
	SmId *int `bun:"sm_id"`
	DateCreated time.Time `bun:"date_created"`
	OofShard *string `bun:"oof_shard"`
}
/*
{
	"order_uid": "b563feb7b2b84b6test",
	"track_number": "WBILMTESTTRACK",
	"entry": "WBIL",
	"delivery": FK_id,
	"payment": FK_id,
	"items": list of FK_ids,
	"locale": "en",
	"internal_signature": "",
	"customer_id": "test",
	"delivery_service": "meest",
	"shardkey": "9",
	"sm_id": 99,
	"date_created": "2021-11-26T06:22:19Z",
	"oof_shard": "1"
  }
*/

type Delivery struct {
	
	bun.BaseModel `bun:"table:wb.delivery"`

	ID int64 `bun:"id,pk,autoincrement" json:"-"`
	Name *string `bun:"name"`
	Phone *string `bun:"phone"`
	Zip *string `bun:"zip"`
	City *string `bun:"city"`
	Address *string `bun:"address"`
	Region *string `bun:"region"`
	Email *string `bun:"email"`
}
/*
"delivery": {
	"name": "Test Testov",
	"phone": "+9720000000",
	"zip": "2639809",
	"city": "Kiryat Mozkin",
	"address": "Ploshad Mira 15",
	"region": "Kraiot",
	"email": "test@gmail.com"
  },
*/

type Payment struct {
	
	bun.BaseModel `bun:"table:wb.payment"`

	ID int64 `bun:"id,pk,autoincrement" json:"-"`
	Transaction *string `bun:"transaction"`
	RequestId *string `bun:"request_id"`
	Currency *string `bun:"currency"`
	Provider *string `bun:"provider"`
	Amount *int `bun:"amount"`
	PaymentDate *time.Time `bun:"payment_dt"`
	Bank *string `bun:"bank"`
	DeliveryCost *int `bun:"delivery_cost"`
	GoodsTotal *int `bun:"goods_total"`
	CustmFee *int `bun:"custom_fee"`
}
/*
	"payment": {
	  "transaction": "b563feb7b2b84b6test",
	  "request_id": "",
	  "currency": "USD",
	  "provider": "wbpay",
	  "amount": 1817,
	  "payment_dt": 1637907727,
	  "bank": "alpha",
	  "delivery_cost": 1500,
	  "goods_total": 317,
	  "custom_fee": 0
	},
*/

type Item struct {
	
	bun.BaseModel `bun:"table:wb.item"`

	ChrtId int64 `bun:"chrt_id,pk,autoincrement" json:"-"`
	TrackNumber *string  `bun:"track_number"`
	Price *int `bun:"price"`
	RId *string `bun:"rid"`
	Name *string `bun:"name"`
	Sale *int `bun:"sale"`
	Size *string `bun:"size"`
	TotalPrice *int `bun:"total_price"`
	NmId *int `bun:"nm_id"`
	Brand *string `bun:"brand"`
	Status *int `bun:"status"`
}
/*
	"items": [
	  {
		"chrt_id": 9934930,
		"track_number": "WBILMTESTTRACK",
		"price": 453,
		"rid": "ab4219087a764ae0btest",
		"name": "Mascaras",
		"sale": 30,
		"size": "0",
		"total_price": 317,
		"nm_id": 2389212,
		"brand": "Vivienne Sabo",
		"status": 202
	  }
	],
*/
