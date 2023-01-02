drop table if exists users;
create table users
(
    id         bigint not null primary key auto_increment,
    name       varchar(20),
    age        bigint,
    created_at datetime,
    updated_at datetime,
    deleted_at datetime,
    index      idx_deleted_at(deleted_at)
);

insert into users (id, name, age, created_at, updated_at, deleted_at)
values (1, 'zhang3', 3, '2022-12-31 16:42:04', '2022-12-31 16:42:06', null),
       (2, 'li4', 4, '2022-12-31 16:42:28', '2022-12-31 16:42:30', null);