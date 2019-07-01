create table book_order (
    id int primary key auto_increment,
    book_id int,
    transport_type tinyint,
    order_type tinyint,
    seller_id int,
    buyer_id int,
    status tinyint,
    create_time datetime
);