package main

import (
	"database/sql" // 官方mysql操作函数库
	"fmt"
	_ "github.com/go-sql-driver/mysql" // 数据库驱动引入
)

var db *sql.DB // 数据库连接池

//获取db连接池值
func initDb() (err error) {
	// 数据库信息
	dsn := "root:123456@tcp(127.0.0.1:3306)/GoComponent"
	//db, err := sql.Open("mysql", dsn) // 不会校验用户名和密码， := 会赋值重新生成
	db, err = sql.Open("mysql", dsn) //使用全局变量
	if err != nil {
		return
	}
	err = db.Ping() // 校验用户名密码
	if err != nil {
		return
	}
	db.SetMaxOpenConns(10) // 设置数据库最大连接数
	db.SetMaxIdleConns(3)  // 设置最大空闲连接数
	return
}

// 表结构
//CREATE TABLE `user` (
//`id` int NOT NULL AUTO_INCREMENT,
//`name` varchar(255) NOT NULL,
//`age` int NOT NULL,
//`phone` varchar(255) NOT NULL,
//`email` varchar(255) NOT NULL,
//`comment` varchar(255) NOT NULL,
//PRIMARY KEY (`id`)
//) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
type user struct {
	id      int
	name    string
	age     int
	phone   string
	email   string
	comment string
}

// Go 连接Mysql
func main() {
	err := initDb()
	if err != nil {
		fmt.Printf("init Db failed,err:%v", err)
	}
	fmt.Println("连接成功")

	var u user
	//sql 语句
	sqlStr := `select * from user;`
	// 执行
	rawObj := db.QueryRow(sqlStr)
	// 拿到结果
	rawObj.Scan(&u.id, &u.name, &u.age, &u.phone, &u.email, &u.comment)
	fmt.Printf("u:%#v", u)

}
