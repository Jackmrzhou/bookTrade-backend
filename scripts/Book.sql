create table book (
    id int primary key auto_increment,
    user_id int,
    name varchar(100),
    author varchar(100),
    isbn char(13),
    price float,
    cover_key char(36),
    introduction varchar(1000),
    type tinyint,
    out_link varchar(1000),
    catalog_id int
);