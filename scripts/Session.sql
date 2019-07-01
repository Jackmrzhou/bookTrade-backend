create table session (
    id int primary key auto_increment,
    session_id char(36),
    account_id int,
    expire_time datetime
);