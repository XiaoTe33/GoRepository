package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// 定义一个初始化数据库的函数
func initDB() (err error) {
	// DSN:Data Source Name
	dsn := "root:root@tcp(localhost:3306)/mydatabase"
	// 不会校验账号密码是否正确
	// 注意！！！这里不要使用:=，我们是给全局变量赋值，然后在main函数中使用全局变量db
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	// 尝试与数据库建立连接（校验dsn是否正确）
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	err := initDB() // 调用输出化数据库的函数
	if err != nil {
		fmt.Printf("init db failed,err:%v\n", err)
		return
	} else {
		fmt.Println("connected")
	}
	QueryRow()
}

type stu struct {
	name string
	id   int
	age  int
}

func QueryRow() {
	sqlStr := "select * from mytable where id=1"
	row := db.QueryRow(sqlStr)
	var stu1 stu
	err := row.Scan(&stu1.name, &stu1.id, &stu1.age)
	if err != nil {
		fmt.Println("scan err")
	}
	fmt.Println(stu1)

}
