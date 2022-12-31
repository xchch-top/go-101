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

insert into users (id, name, age, birthday, created_at, updated_at, deleted_at)
    values  (1, 'zhang3', 3, '2022-12-31 16:42:04', '2022-12-31 16:42:04', '2022-12-31 16:42:04', null);