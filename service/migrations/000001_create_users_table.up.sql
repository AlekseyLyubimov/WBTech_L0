CREATE SCHEMA IF NOT EXISTS wb;

CREATE TABLE IF NOT EXISTS wb.delivery(
    "id"                 bigserial NOT NULL,
    name                 character varying(250) NOT NULL,
    phone                character varying(12) NOT NULL,
    zip                  character varying(6) NULL,
    city                 character varying(250) NOT NULL,
    address              character varying(250) NOT NULL,
    region               character varying(250) NOT NULL,
    email                character varying(100) NOT NULL,

    CONSTRAINT PK_delivery PRIMARY KEY ( "id" )
);

CREATE TABLE IF NOT EXISTS wb.payment(
    "id"                 bigserial NOT NULL,
    transaction          character varying(250) NOT NULL,
    request_id           character varying(250) NOT NULL,
    currency             character varying(250) NOT NULL,
    provider             character varying(250) NOT NULL,
    amount               integer NOT NULL,
    payment_dt           date NOT NULL,
    Bank                 character varying(250) NOT NULL,
    delivery_cost        integer NOT NULL,
    goods_total          integer NOT NULL,
    custom_fee           integer NOT NULL,

    CONSTRAINT PK_payment PRIMARY KEY ( "id" )
);

CREATE TABLE IF NOT EXISTS wb.item(
    "id"            bigserial NOT NULL,
    track_number    character varying(250) NOT NULL,
    price           integer NOT NULL,
    rid             character varying(250) NOT NULL,
    name            character varying(250) NOT NULL,
    sale            integer NOT NULL,
    size            character varying(250) NOT NULL,
    total_price     integer NOT NULL,
    nm_id           integer NOT NULL,
    brand           character varying(250) NOT NULL,
    status          integer NOT NULL,

    CONSTRAINT PK_item PRIMARY KEY ( "id" )
);

CREATE TABLE IF NOT EXISTS wb.order (
    "id"             character varying(250) NOT NULL,
    track_number     character varying(250) NOT NULL,
    entry            character varying(250) NOT NULL,
    delivery         bigint NOT NULL,
    payment          bigint NOT NULL,
    items            bigint[] NOT NULL,
    locale           character varying(250) NOT NULL,
    internal_signature     character varying(250) NOT NULL,
    customer_id      character varying(250) NOT NULL,
    delivery_service   character varying(250) NOT NULL,
    shardkey         character varying(250) NOT NULL,
    sm_id            integer NOT NULL,
    date_created     date NOT NULL,
    oof_shard        character varying(250) NOT NULL,

    CONSTRAINT PK_order PRIMARY KEY ( "id" ),
    CONSTRAINT FK_order_delivery FOREIGN KEY ( delivery ) REFERENCES wb.delivery ( "id" ),
    CONSTRAINT FK_order_payment FOREIGN KEY ( payment ) REFERENCES wb.payment ( "id" )
);

CREATE TABLE order_item (
    "id"            bigserial NOT NULL,
    order_id        character varying(250) NOT NULL,
    item_id         bigint NOT NULL,
    CONSTRAINT PK_oi PRIMARY KEY ( "id" ),
    CONSTRAINT FK_oi_order_id FOREIGN KEY ( order_id ) REFERENCES wb.order ( "id" ),
    CONSTRAINT FK_oi_item_id FOREIGN KEY ( item_id ) REFERENCES wb.item ( "id" )
);