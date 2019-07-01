create table message (
    id int primary key auto_increment,
    contact_id int,
    from_id int,
    to_id int,
    content varchar(200),
    create_time datetime,
    is_read tinyint
);