drop table if exists users;
create table users
(
    id         bigint primary key auto_increment,
    name       varchar(20),
    age        bigint,
    birthday   datetime,
    created_at datetime,
    updated_at datetime,
    deleted_at datetime
);