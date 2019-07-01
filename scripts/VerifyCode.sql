create table verify_code (
    id int primary key auto_increment,
    code char(6),
    email varchar(254),
    expire_time datetime
);