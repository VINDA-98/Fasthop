package db

import (
	"database/sql"
	"fmt"
	"time"
)

// @Title  db
// @Description  MyGO
// @Author  WeiDa  2023/6/29 15:14
// @Update  WeiDa  2023/6/29 15:14

var db *sql.DB

// ConnectToDb 连接到数据库
func ConnectToDb() *sql.DB {
	db, _ := sql.Open("mysql", "root:123456@/FastShop?charset=utf8")
	// 设置可重用链接的最长时间（0为不限制）
	db.SetConnMaxLifetime(time.Hour * 1)
	// 设置连接到数据库的最大数量（默认值为0，即不限制）
	db.SetMaxOpenConns(5)
	// 设置空闲连接的最大数量（默认值为2）
	db.SetMaxIdleConns(5)
	fmt.Println("连接成功！！")
	return db
}

// CreateTable 创建数据表
func CreateTable() {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS `notebook` (" +
		"`id` bigint(20) NOT NULL AUTO_INCREMENT," +
		"`title` varchar(45) DEFAULT ''," +
		"`content` varchar(45) DEFAULT ''," +
		"`dateTime` varchar(45) DEFAULT ''," +
		"PRIMARY KEY (`id`)" +
		") ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;")
	if err != nil {
		return
	}
	fmt.Println("notebook表存在或成功创建")
}
