drop table if exists companies;
create table companies
(
    id         bigint not null primary key auto_increment,
    name       varchar(20),
    created_at datetime,
    updated_at datetime,
    deleted_at datetime,
    index      idx_deleted_at(deleted_at)
);

drop table if exists users;
create table users
(
    id         bigint not null primary key auto_increment,
    name       varchar(20),
    company_id bigint not null,
    created_at datetime,
    updated_at datetime,
    deleted_at datetime,
    index      idx_deleted_at(deleted_at),
    constraint fk_company foreign key (company_id) references companies (id)
);