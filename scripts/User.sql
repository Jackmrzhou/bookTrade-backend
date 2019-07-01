create table user (
    id int primary key auto_increment,
    account_id int,
    username varchar(100),
    avatar_key char(36),
    introduction varchar(200),
    phone_number varchar(20),
    address varchar(100)
);