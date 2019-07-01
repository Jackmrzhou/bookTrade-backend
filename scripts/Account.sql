create table account (
    id int primary key auto_increment,
    created_at datetime,
    updated_at datetime,
    deleted_at datetime,
    status tinyint,
    email varchar(254),
    password char(40)
);