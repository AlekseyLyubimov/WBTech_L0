CREATE SCHEMA IF NOT EXISTS wb;

CREATE TABLE IF NOT EXISTS wb.delivery(
    "id"                 bigserial NOT NULL,
    name                 character varying(250),
    phone                character varying(12),
    zip                  character varying(25),
    city                 character varying(250),
    address              character varying(250),
    region               character varying(250),
    email                character varying(100),

    CONSTRAINT PK_delivery PRIMARY KEY ( "id" )
);

CREATE TABLE IF NOT EXISTS wb.payment(
    "id"                 bigserial NOT NULL,
    transaction          character varying(250),
    request_id           character varying(250),
    currency             character varying(250),
    provider             character varying(250),
    amount               integer,
    payment_dt           character varying(250),
    bank                 character varying(250),
    delivery_cost        integer,
    goods_total          integer,
    custom_fee           integer,

    CONSTRAINT PK_payment PRIMARY KEY ( "id" )
);

CREATE TABLE IF NOT EXISTS wb.item(
    "id"            bigserial NOT NULL,
    track_number    character varying(250),
    price           integer,
    rid             character varying(250),
    name            character varying(250),
    sale            integer,
    size            character varying(250),
    total_price     integer,
    nm_id           integer,
    brand           character varying(250),
    status          integer,

    CONSTRAINT PK_item PRIMARY KEY ( "id" )
);

CREATE TABLE IF NOT EXISTS wb.order (
    "id"             character varying(250) NOT NULL,
    track_number     character varying(250),
    entry            character varying(250),
    delivery         bigint,
    payment          bigint,
    locale           character varying(250),
    internal_signature     character varying(250),
    customer_id      character varying(250),
    delivery_service   character varying(250),
    shardkey         character varying(250),
    sm_id            integer,
    date_created     date,
    oof_shard        character varying(250),

    CONSTRAINT PK_order PRIMARY KEY ( "id" ),
    CONSTRAINT FK_order_delivery FOREIGN KEY ( delivery ) REFERENCES wb.delivery ( "id" ),
    CONSTRAINT FK_order_payment FOREIGN KEY ( payment ) REFERENCES wb.payment ( "id" )
);

CREATE TABLE wb.order_item (
    "id"            bigserial NOT NULL,
    order_id        character varying(250) NOT NULL,
    item_id         bigint NOT NULL,
    CONSTRAINT PK_oi PRIMARY KEY ( "id" ),
    CONSTRAINT FK_oi_order_id FOREIGN KEY ( order_id ) REFERENCES wb.order ( "id" ),
    CONSTRAINT FK_oi_item_id FOREIGN KEY ( item_id ) REFERENCES wb.item ( "id" )
);