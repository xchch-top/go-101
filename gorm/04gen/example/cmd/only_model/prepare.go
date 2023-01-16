package main

import "gorm.io/gorm"

// prepare table for test

const mytableSQL = `create table users (
id         bigint not null primary key auto_increment,
name       varchar(20),
age        bigint,
birthday   datetime,
created_at datetime,
updated_at datetime,
deleted_at datetime,
index      idx_deleted_at(deleted_at)
);`

func prepare(db *gorm.DB) {
	db.Exec(mytableSQL)
}
