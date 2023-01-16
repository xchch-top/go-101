drop table if exists users;
create table users
(
    id         bigint not null primary key auto_increment,
    name       varchar(20),
    age        bigint,
    birthday   datetime,
    created_at datetime,
    updated_at datetime,
    deleted_at datetime,
    index      idx_deleted_at(deleted_at)
);