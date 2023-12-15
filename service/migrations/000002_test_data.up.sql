INSERT INTO wb.delivery ("id", name, phone, zip, city, address, region, email) 
    VALUES (1, 'Test Testov', '+9720000000', '2639809', 'Kiryat Mozkin', 'Ploshad Mira 15', 'Kraiot', 'test@gmail.com');

INSERT INTO wb.payment ("id", transaction, request_id, currency, provider, amount, payment_dt, bank, delivery_cost, goods_total, custom_fee) 
    VALUES (1, 'b563feb7b2b84b6test', '', 'USD', 'wbpay', 1817, 1637907727, 'alpha', 1500, 317, 0);

INSERT INTO wb.item ("id", track_number, price, rid, name, sale, size, total_price, nm_id, brand, status) 
    VALUES (1, 'WBILMTESTTRACK', 453, 'ab4219087a764ae0btest', 'Mascaras', 30, '0', 317, 2389212, 'Vivienne Sabo', 202);

INSERT INTO wb.order ("id", track_number, entry, delivery, payment, locale, internal_signature, customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard) 
    VALUES ('b563feb7b2b84b6test', 'WBILMTESTTRACK', 'WBIL', 1, 1, 'en', '', 'test', 'meest', '9', 99, '2021-11-26T06:22:19Z', '1');

INSERT INTO wb.order_item (order_id, item_id) VALUES ('b563feb7b2b84b6test', 1);
