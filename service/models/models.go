package models

import (
	"time"

	bun "github.com/uptrace/bun"
)

//TODO verify relation syntaxis
type Order struct {

	bun.BaseModel `bun:"table:wb.order"`

	ID string `bun:"id,pk" json:"order_uid"`
	TrackNumber *string `bun:"track_number" json:"track_number"`
	Entry *string `bun:"entry" json:"entry"`
	Delivery *Delivery `bun:"delivery,rel:has-one,join:delivery=id" json:"delivery"`
	Payment *Payment `bun:"payment,rel:has-one,join:payment=id" json:"payment"`
	Items []*Item `bun:"items,m2m:wb.order_item,join:Order=Item" json:"items"`
	Locale *string `bun:"locale" json:"locale"`
	InternalSignature *string `bun:"internal_signature" json:"internal_signature"`
	CustomerId *string `bun:"customer_id" json:"customer_id"`
	DeliveryService *string `bun:"delivery_service" json:"delivery_service"`
	Shardkey *string `bun:"shardkey" json:"shardkey"`
	SmId *int `bun:"sm_id" json:"sm_id"`
	DateCreated time.Time `bun:"date_created" json:"date_created"`
	OofShard *string `bun:"oof_shard" json:"oof_shard"`
}

type Delivery struct {
	
	bun.BaseModel `bun:"table:wb.delivery"`

	ID int64 `bun:"id,pk,autoincrement" json:"-"`
	Name *string `bun:"name" json:"name"`
	Phone *string `bun:"phone" json:"phone"`
	Zip *string `bun:"zip" json:"zip"`
	City *string `bun:"city" json:"city"`
	Address *string `bun:"address" json:"address"`
	Region *string `bun:"region" json:"region"`
	Email *string `bun:"email" json:"email"`
}

type Payment struct {
	
	bun.BaseModel `bun:"table:wb.payment"`

	ID int64 `bun:"id,pk,autoincrement" json:"-"`
	Transaction *string `bun:"transaction" json:"transaction"`
	RequestId *string `bun:"request_id" json:"request_id"`
	Currency *string `bun:"currency" json:"currency"`
	Provider *string `bun:"provider" json:"provider"`
	Amount *int `bun:"amount" json:"amount"`
	PaymentDate *string `bun:"payment_dt" json:"payment_dt"`
	Bank *string `bun:"bank" json:"bank"`
	DeliveryCost *int `bun:"delivery_cost" json:"delivery_cost"`
	GoodsTotal *int `bun:"goods_total" json:"goods_total"`
	CustmFee *int `bun:"custom_fee" json:"custom_fee"`
}

type Item struct {
	
	bun.BaseModel `bun:"table:wb.item"`

	ChrtId int64 `bun:"id,pk,autoincrement" json:"chrt_id"`
	TrackNumber *string  `bun:"track_number" json:"track_number"`
	Price *int `bun:"price" json:"price"`
	RId *string `bun:"rid" json:"rid"`
	Name *string `bun:"name" json:"name"`
	Sale *int `bun:"sale" json:"sale"`
	Size *string `bun:"size" json:"size"`
	TotalPrice *int `bun:"total_price" json:"total_price"`
	NmId *int `bun:"nm_id" json:"nm_id"`
	Brand *string `bun:"brand" json:"brand"`
	Status *int `bun:"status" json:"status"`
}

type OrderToItem struct {

	bun.BaseModel `bun:"table:wb.order_item"`

	OrderID string  `bun:"order_id,pk"`
	Order   *Order `bun:"rel:belongs-to,join:order_id=id"`
	ItemID  int64  `bun:"item_id,pk"`
	Item    *Item  `bun:"rel:belongs-to,join:item_id=id"`
}