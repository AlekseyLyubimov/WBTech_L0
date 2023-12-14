package models

import (
	"time"

	bun "github.com/uptrace/bun"
)

//TODO verify relation syntaxis
type Order struct {

	bun.BaseModel `bun:"table:wb.order"`

	ID *string `bun:"id"`
	TrackNumber *string `bun:"track_number"`
	Entry *string `bun:"entry"`
	Delivery *Delivery `bun:"delivery, rel:has-one,join:delivery=id"`
	Payment *Payment `bun:"payment, rel:has-one,join:payment=id"`
	Items []*Item `bun:"items,m2m:order_to_items,join:Order=Item"`
	Locale *string `bun:"locale"`
	InternalSignature *string `bun:"internal_signature"`
	CustomerId *string `bun:"customer_id"`
	DeliveryService *string `bun:"delivery_service"`
	Shardkey *string `bun:"shardkey"`
	SmId *int `bun:"sm_id"`
	DateCreated time.Time `bun:"date_created"`
	OofShard *string `bun:"oof_shard"`
}

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

type Item struct {
	
	bun.BaseModel `bun:"table:wb.item"`

	ChrtId int64 `bun:"chrt_id,pk,autoincrement"`
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

//TODO register relation https://bun.uptrace.dev/guide/relations.html#many-to-many-relation
type OrderToItem struct {

	bun.BaseModel `bun:"table:wb.order_item"`

	OrderID int64  `bun:"order_id,pk"`
	Order   *Order `bun:"rel:belongs-to,join:order_id=id"`
	ItemID  int64  `bun:"item_id,pk"`
	Item    *Item  `bun:"rel:belongs-to,join:item_id=id"`
}