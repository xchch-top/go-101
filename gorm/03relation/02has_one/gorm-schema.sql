drop table if exists users;
create table users
(
    id         bigint not null primary key auto_increment,
    name       varchar(20),
    created_at datetime,
    updated_at datetime,
    deleted_at datetime,
    index      idx_deleted_at(deleted_at)
);

drop table if exists credit_cards;
create table credit_cards
(
    id         bigint not null primary key auto_increment,
    number     varchar(20),
    user_id    bigint not null,
    created_at datetime,
    updated_at datetime,
    deleted_at datetime,
    index      idx_deleted_at(deleted_at),
    constraint fk_user foreign key (user_id) references users (id)
);